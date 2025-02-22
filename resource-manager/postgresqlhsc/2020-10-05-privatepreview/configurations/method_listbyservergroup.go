package configurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServerGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServerGroupConfiguration
}

type ListByServerGroupCompleteResult struct {
	Items []ServerGroupConfiguration
}

// ListByServerGroup ...
func (c ConfigurationsClient) ListByServerGroup(ctx context.Context, id ServerGroupsv2Id) (result ListByServerGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/configurations", id.ID()),
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
		Values *[]ServerGroupConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByServerGroupComplete retrieves all the results into a single object
func (c ConfigurationsClient) ListByServerGroupComplete(ctx context.Context, id ServerGroupsv2Id) (ListByServerGroupCompleteResult, error) {
	return c.ListByServerGroupCompleteMatchingPredicate(ctx, id, ServerGroupConfigurationOperationPredicate{})
}

// ListByServerGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationsClient) ListByServerGroupCompleteMatchingPredicate(ctx context.Context, id ServerGroupsv2Id, predicate ServerGroupConfigurationOperationPredicate) (result ListByServerGroupCompleteResult, err error) {
	items := make([]ServerGroupConfiguration, 0)

	resp, err := c.ListByServerGroup(ctx, id)
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

	result = ListByServerGroupCompleteResult{
		Items: items,
	}
	return
}
