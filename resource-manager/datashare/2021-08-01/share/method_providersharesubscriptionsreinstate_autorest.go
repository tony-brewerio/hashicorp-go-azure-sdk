package share

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderShareSubscriptionsReinstateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProviderShareSubscription
}

// ProviderShareSubscriptionsReinstate ...
func (c ShareClient) ProviderShareSubscriptionsReinstate(ctx context.Context, id ProviderShareSubscriptionId, input ProviderShareSubscription) (result ProviderShareSubscriptionsReinstateOperationResponse, err error) {
	req, err := c.preparerForProviderShareSubscriptionsReinstate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "share.ShareClient", "ProviderShareSubscriptionsReinstate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "share.ShareClient", "ProviderShareSubscriptionsReinstate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProviderShareSubscriptionsReinstate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "share.ShareClient", "ProviderShareSubscriptionsReinstate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProviderShareSubscriptionsReinstate prepares the ProviderShareSubscriptionsReinstate request.
func (c ShareClient) preparerForProviderShareSubscriptionsReinstate(ctx context.Context, id ProviderShareSubscriptionId, input ProviderShareSubscription) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reinstate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProviderShareSubscriptionsReinstate handles the response to the ProviderShareSubscriptionsReinstate request. The method always
// closes the http.Response Body.
func (c ShareClient) responderForProviderShareSubscriptionsReinstate(resp *http.Response) (result ProviderShareSubscriptionsReinstateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
