package proxy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteRenderingAccountsListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RemoteRenderingAccount
}

type RemoteRenderingAccountsListBySubscriptionCompleteResult struct {
	Items []RemoteRenderingAccount
}

// RemoteRenderingAccountsListBySubscription ...
func (c ProxyClient) RemoteRenderingAccountsListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result RemoteRenderingAccountsListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.MixedReality/remoteRenderingAccounts", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]RemoteRenderingAccount `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RemoteRenderingAccountsListBySubscriptionComplete retrieves all the results into a single object
func (c ProxyClient) RemoteRenderingAccountsListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (RemoteRenderingAccountsListBySubscriptionCompleteResult, error) {
	return c.RemoteRenderingAccountsListBySubscriptionCompleteMatchingPredicate(ctx, id, RemoteRenderingAccountOperationPredicate{})
}

// RemoteRenderingAccountsListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProxyClient) RemoteRenderingAccountsListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate RemoteRenderingAccountOperationPredicate) (result RemoteRenderingAccountsListBySubscriptionCompleteResult, err error) {
	items := make([]RemoteRenderingAccount, 0)

	resp, err := c.RemoteRenderingAccountsListBySubscription(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = RemoteRenderingAccountsListBySubscriptionCompleteResult{
		Items: items,
	}
	return
}
