package webapps

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

type ListInstanceIdentifiersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebSiteInstanceStatus
}

type ListInstanceIdentifiersCompleteResult struct {
	Items []WebSiteInstanceStatus
}

// ListInstanceIdentifiers ...
func (c WebAppsClient) ListInstanceIdentifiers(ctx context.Context, id commonids.AppServiceId) (result ListInstanceIdentifiersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/instances", id.ID()),
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
		Values *[]WebSiteInstanceStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInstanceIdentifiersComplete retrieves all the results into a single object
func (c WebAppsClient) ListInstanceIdentifiersComplete(ctx context.Context, id commonids.AppServiceId) (ListInstanceIdentifiersCompleteResult, error) {
	return c.ListInstanceIdentifiersCompleteMatchingPredicate(ctx, id, WebSiteInstanceStatusOperationPredicate{})
}

// ListInstanceIdentifiersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListInstanceIdentifiersCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate WebSiteInstanceStatusOperationPredicate) (result ListInstanceIdentifiersCompleteResult, err error) {
	items := make([]WebSiteInstanceStatus, 0)

	resp, err := c.ListInstanceIdentifiers(ctx, id)
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

	result = ListInstanceIdentifiersCompleteResult{
		Items: items,
	}
	return
}
