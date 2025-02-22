package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListWorkerPoolSkusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SkuInfo
}

type ListWorkerPoolSkusCompleteResult struct {
	Items []SkuInfo
}

// ListWorkerPoolSkus ...
func (c AppServiceEnvironmentsClient) ListWorkerPoolSkus(ctx context.Context, id WorkerPoolId) (result ListWorkerPoolSkusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/skus", id.ID()),
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
		Values *[]SkuInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWorkerPoolSkusComplete retrieves all the results into a single object
func (c AppServiceEnvironmentsClient) ListWorkerPoolSkusComplete(ctx context.Context, id WorkerPoolId) (ListWorkerPoolSkusCompleteResult, error) {
	return c.ListWorkerPoolSkusCompleteMatchingPredicate(ctx, id, SkuInfoOperationPredicate{})
}

// ListWorkerPoolSkusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentsClient) ListWorkerPoolSkusCompleteMatchingPredicate(ctx context.Context, id WorkerPoolId, predicate SkuInfoOperationPredicate) (result ListWorkerPoolSkusCompleteResult, err error) {
	items := make([]SkuInfo, 0)

	resp, err := c.ListWorkerPoolSkus(ctx, id)
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

	result = ListWorkerPoolSkusCompleteResult{
		Items: items,
	}
	return
}
