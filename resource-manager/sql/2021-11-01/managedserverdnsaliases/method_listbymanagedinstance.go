package managedserverdnsaliases

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

type ListByManagedInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedServerDnsAlias
}

type ListByManagedInstanceCompleteResult struct {
	Items []ManagedServerDnsAlias
}

// ListByManagedInstance ...
func (c ManagedServerDnsAliasesClient) ListByManagedInstance(ctx context.Context, id commonids.SqlManagedInstanceId) (result ListByManagedInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/dnsAliases", id.ID()),
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
		Values *[]ManagedServerDnsAlias `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByManagedInstanceComplete retrieves all the results into a single object
func (c ManagedServerDnsAliasesClient) ListByManagedInstanceComplete(ctx context.Context, id commonids.SqlManagedInstanceId) (ListByManagedInstanceCompleteResult, error) {
	return c.ListByManagedInstanceCompleteMatchingPredicate(ctx, id, ManagedServerDnsAliasOperationPredicate{})
}

// ListByManagedInstanceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedServerDnsAliasesClient) ListByManagedInstanceCompleteMatchingPredicate(ctx context.Context, id commonids.SqlManagedInstanceId, predicate ManagedServerDnsAliasOperationPredicate) (result ListByManagedInstanceCompleteResult, err error) {
	items := make([]ManagedServerDnsAlias, 0)

	resp, err := c.ListByManagedInstance(ctx, id)
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

	result = ListByManagedInstanceCompleteResult{
		Items: items,
	}
	return
}
