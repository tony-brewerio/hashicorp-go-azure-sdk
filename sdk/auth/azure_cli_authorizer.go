// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MPL-2.0 License. See NOTICE.txt in the project root for license information.

package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/azurecli"
	"golang.org/x/oauth2"
)

type AzureCliAuthorizerOptions struct {
	// Api describes the Azure API being used
	Api environments.Api

	// TenantId is the tenant to authenticate against
	TenantId string

	// AuxTenantIds lists additional tenants to authenticate against, currently only
	// used for Resource Manager when auxiliary tenants are needed.
	// e.g. https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/authenticate-multi-tenant
	AuxTenantIds []string

	// SubscriptionId is the subscription to authenticate against
	// If not provided, Azure CLI will default to the current subscription
	// This will be the subscription from the last az login call, or the one assigned with az account set
	SubscriptionId string
}

// NewAzureCliAuthorizer returns an Authorizer which authenticates using the Azure CLI.
func NewAzureCliAuthorizer(ctx context.Context, options AzureCliAuthorizerOptions) (Authorizer, error) {
	conf, err := newAzureCliConfig(options.Api, options.TenantId, options.AuxTenantIds, options.SubscriptionId)
	if err != nil {
		return nil, err
	}
	return conf.TokenSource(ctx)
}

var _ Authorizer = &AzureCliAuthorizer{}

// AzureCliAuthorizer is an Authorizer which supports the Azure CLI.
type AzureCliAuthorizer struct {
	// TenantID is the specified tenant ID, or the auto-detected tenant ID if none was specified
	TenantID string

	// DefaultSubscriptionID is the default subscription, when detected
	DefaultSubscriptionID string

	// SubscriptionId is the specified subscription ID, or empty value if none was specified
	SubscriptionId string

	conf *azureCliConfig
}

// Token returns an access token using the Azure CLI as an authentication mechanism.
func (a *AzureCliAuthorizer) Token(_ context.Context, _ *http.Request) (*oauth2.Token, error) {
	if a.conf == nil {
		return nil, fmt.Errorf("could not request token: conf is nil")
	}

	azArgs := []string{"account", "get-access-token"}

	scope, err := environments.Scope(a.conf.Api)
	if err != nil {
		return nil, fmt.Errorf("determining scope for %q: %+v", a.conf.Api.Name(), err)
	}
	azArgs = append(azArgs, "--scope", *scope)

	// If user provided an explicit subscription ID, add it to the command arguments
	// This overrides tenant setting, since those two are not compatible
	if a.conf.SubscriptionId != "" {
		azArgs = append(azArgs, "--subscription", a.conf.SubscriptionId)
	} else {
		// Try to detect if we're running in Cloud Shell
		if cloudShell := os.Getenv("AZUREPS_HOST_ENVIRONMENT"); !strings.HasPrefix(cloudShell, "cloud-shell/") {
			// Seemingly not, so we'll append the tenant ID to the az args
			azArgs = append(azArgs, "--tenant", a.conf.TenantID)
		}
	}

	var token azureCliToken
	if err = azurecli.JSONUnmarshalAzCmd(false, &token, azArgs...); err != nil {
		return nil, err
	}

	var expiry time.Time
	if token.ExpiresOn != "" {
		if expiry, err = time.ParseInLocation("2006-01-02 15:04:05.999999", token.ExpiresOn, time.Local); err != nil {
			return nil, fmt.Errorf("internal-error: parsing expiresOn value for az-cli token")
		}
	}

	return &oauth2.Token{
		AccessToken: token.AccessToken,
		Expiry:      expiry,
		TokenType:   token.TokenType,
	}, nil
}

// AuxiliaryTokens returns additional tokens for auxiliary tenant IDs, for use in multi-tenant scenarios
func (a *AzureCliAuthorizer) AuxiliaryTokens(_ context.Context, _ *http.Request) ([]*oauth2.Token, error) {
	if a.conf == nil {
		return nil, fmt.Errorf("could not request token: conf is nil")
	}

	// Return early if no auxiliary tenants are configured
	if len(a.conf.AuxiliaryTenantIDs) == 0 {
		return []*oauth2.Token{}, nil
	}

	// Try to detect if we're running in Cloud Shell
	if cloudShell := os.Getenv("AZUREPS_HOST_ENVIRONMENT"); strings.HasPrefix(cloudShell, "cloud-shell/") {
		return nil, fmt.Errorf("auxiliary tokens not supported in Cloud Shell")
	}

	azArgs := []string{"account", "get-access-token"}

	scope, err := environments.Scope(a.conf.Api)
	if err != nil {
		return nil, fmt.Errorf("determining scope for %q: %+v", a.conf.Api.Name(), err)
	}
	azArgs = append(azArgs, "--scope", *scope)

	tokens := make([]*oauth2.Token, 0)
	for _, tenantId := range a.conf.AuxiliaryTenantIDs {
		argsWithTenant := append(azArgs, "--tenant", tenantId)

		var token azureCliToken
		if err = azurecli.JSONUnmarshalAzCmd(false, &token, argsWithTenant...); err != nil {
			return nil, err
		}

		tokens = append(tokens, &oauth2.Token{
			AccessToken: token.AccessToken,
			TokenType:   token.TokenType,
		})
	}

	return tokens, nil
}

// azureCliConfig configures an AzureCliAuthorizer.
type azureCliConfig struct {
	Api environments.Api

	// TenantID is the required tenant ID for the primary token
	TenantID string

	// AuxiliaryTenantIDs is an optional list of tenant IDs for which to obtain additional tokens
	AuxiliaryTenantIDs []string

	// SubscriptionId is the optional explicit subscription ID
	SubscriptionId string

	// DefaultSubscriptionID is the optional default subscription ID
	DefaultSubscriptionID string
}

// newAzureCliConfig validates the supplied tenant ID and returns a new azureCliConfig.
func newAzureCliConfig(api environments.Api, tenantId string, auxiliaryTenantIds []string, subscriptionId string) (*azureCliConfig, error) {
	// check az-cli version, ensure that MSAL is supported
	if err := azurecli.CheckAzVersion(); err != nil {
		return nil, err
	}

	if strings.TrimSpace(subscriptionId) != "" {
		if strings.TrimSpace(tenantId) != "" || auxiliaryTenantIds != nil || len(auxiliaryTenantIds) != 0 {
			return nil, fmt.Errorf("only one of subscription ID and tenant ID can be specified, not both")
		}
	}

	// obtain default tenant ID if no tenant ID was provided
	if strings.TrimSpace(tenantId) == "" {
		if defaultTenantId, err := azurecli.GetDefaultTenantID(); err != nil {
			return nil, fmt.Errorf("tenant ID was not specified and the default tenant ID could not be determined: %v", err)
		} else if defaultTenantId == nil {
			return nil, fmt.Errorf("tenant ID was not specified and the default tenant ID could not be determined")
		} else {
			tenantId = *defaultTenantId
		}
	}

	// validate tenant ID
	if valid, err := azurecli.ValidateTenantID(tenantId); err != nil {
		return nil, err
	} else if !valid {
		return nil, fmt.Errorf("invalid tenant ID was provided")
	}

	// get the default subscription ID
	var defaultSubscriptionId string
	if azureCliDefaultSubscriptionId, err := azurecli.GetDefaultSubscriptionID(); err != nil {
		return nil, err
	} else if azureCliDefaultSubscriptionId != nil {
		defaultSubscriptionId = *azureCliDefaultSubscriptionId
	}

	return &azureCliConfig{
		Api:                   api,
		TenantID:              tenantId,
		AuxiliaryTenantIDs:    auxiliaryTenantIds,
		DefaultSubscriptionID: defaultSubscriptionId,
		SubscriptionId:        subscriptionId,
	}, nil
}

// TokenSource provides a source for obtaining access tokens using AzureCliAuthorizer.
func (c *azureCliConfig) TokenSource(ctx context.Context) (Authorizer, error) {
	// Cache access tokens internally to avoid unnecessary `az` invocations
	return NewCachedAuthorizer(&AzureCliAuthorizer{
		TenantID:              c.TenantID,
		DefaultSubscriptionID: c.DefaultSubscriptionID,
		SubscriptionId:        c.SubscriptionId,
		conf:                  c,
	})
}

type azureCliToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresOn   string `json:"expiresOn"`
	Tenant      string `json:"tenant"`
	TokenType   string `json:"tokenType"`
}
