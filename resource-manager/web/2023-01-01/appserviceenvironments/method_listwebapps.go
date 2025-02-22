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

type ListWebAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Site
}

type ListWebAppsCompleteResult struct {
	Items []Site
}

type ListWebAppsOperationOptions struct {
	PropertiesToInclude *string
}

func DefaultListWebAppsOperationOptions() ListWebAppsOperationOptions {
	return ListWebAppsOperationOptions{}
}

func (o ListWebAppsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListWebAppsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListWebAppsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.PropertiesToInclude != nil {
		out.Append("propertiesToInclude", fmt.Sprintf("%v", *o.PropertiesToInclude))
	}
	return &out
}

// ListWebApps ...
func (c AppServiceEnvironmentsClient) ListWebApps(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListWebAppsOperationOptions) (result ListWebAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/sites", id.ID()),
		OptionsObject: options,
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
		Values *[]Site `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWebAppsComplete retrieves all the results into a single object
func (c AppServiceEnvironmentsClient) ListWebAppsComplete(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListWebAppsOperationOptions) (ListWebAppsCompleteResult, error) {
	return c.ListWebAppsCompleteMatchingPredicate(ctx, id, options, SiteOperationPredicate{})
}

// ListWebAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentsClient) ListWebAppsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListWebAppsOperationOptions, predicate SiteOperationPredicate) (result ListWebAppsCompleteResult, err error) {
	items := make([]Site, 0)

	resp, err := c.ListWebApps(ctx, id, options)
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

	result = ListWebAppsCompleteResult{
		Items: items,
	}
	return
}
