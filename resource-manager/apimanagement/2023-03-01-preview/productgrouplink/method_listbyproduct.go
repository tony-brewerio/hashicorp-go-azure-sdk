package productgrouplink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByProductOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProductGroupLinkContract
}

type ListByProductCompleteResult struct {
	Items []ProductGroupLinkContract
}

type ListByProductOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultListByProductOperationOptions() ListByProductOperationOptions {
	return ListByProductOperationOptions{}
}

func (o ListByProductOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByProductOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByProductOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListByProduct ...
func (c ProductGroupLinkClient) ListByProduct(ctx context.Context, id ProductId, options ListByProductOperationOptions) (result ListByProductOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/groupLinks", id.ID()),
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
		Values *[]ProductGroupLinkContract `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByProductComplete retrieves all the results into a single object
func (c ProductGroupLinkClient) ListByProductComplete(ctx context.Context, id ProductId, options ListByProductOperationOptions) (ListByProductCompleteResult, error) {
	return c.ListByProductCompleteMatchingPredicate(ctx, id, options, ProductGroupLinkContractOperationPredicate{})
}

// ListByProductCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProductGroupLinkClient) ListByProductCompleteMatchingPredicate(ctx context.Context, id ProductId, options ListByProductOperationOptions, predicate ProductGroupLinkContractOperationPredicate) (result ListByProductCompleteResult, err error) {
	items := make([]ProductGroupLinkContract, 0)

	resp, err := c.ListByProduct(ctx, id, options)
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

	result = ListByProductCompleteResult{
		Items: items,
	}
	return
}
