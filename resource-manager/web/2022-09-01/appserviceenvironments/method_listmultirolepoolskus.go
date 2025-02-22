package appserviceenvironments

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

type ListMultiRolePoolSkusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SkuInfo
}

type ListMultiRolePoolSkusCompleteResult struct {
	Items []SkuInfo
}

// ListMultiRolePoolSkus ...
func (c AppServiceEnvironmentsClient) ListMultiRolePoolSkus(ctx context.Context, id commonids.AppServiceEnvironmentId) (result ListMultiRolePoolSkusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/multiRolePools/default/skus", id.ID()),
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

// ListMultiRolePoolSkusComplete retrieves all the results into a single object
func (c AppServiceEnvironmentsClient) ListMultiRolePoolSkusComplete(ctx context.Context, id commonids.AppServiceEnvironmentId) (ListMultiRolePoolSkusCompleteResult, error) {
	return c.ListMultiRolePoolSkusCompleteMatchingPredicate(ctx, id, SkuInfoOperationPredicate{})
}

// ListMultiRolePoolSkusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentsClient) ListMultiRolePoolSkusCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, predicate SkuInfoOperationPredicate) (result ListMultiRolePoolSkusCompleteResult, err error) {
	items := make([]SkuInfo, 0)

	resp, err := c.ListMultiRolePoolSkus(ctx, id)
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

	result = ListMultiRolePoolSkusCompleteResult{
		Items: items,
	}
	return
}
