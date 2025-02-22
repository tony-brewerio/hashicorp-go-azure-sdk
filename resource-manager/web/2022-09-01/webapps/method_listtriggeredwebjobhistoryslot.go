package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTriggeredWebJobHistorySlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TriggeredJobHistory
}

type ListTriggeredWebJobHistorySlotCompleteResult struct {
	Items []TriggeredJobHistory
}

// ListTriggeredWebJobHistorySlot ...
func (c WebAppsClient) ListTriggeredWebJobHistorySlot(ctx context.Context, id SlotTriggeredWebJobId) (result ListTriggeredWebJobHistorySlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/history", id.ID()),
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
		Values *[]TriggeredJobHistory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTriggeredWebJobHistorySlotComplete retrieves all the results into a single object
func (c WebAppsClient) ListTriggeredWebJobHistorySlotComplete(ctx context.Context, id SlotTriggeredWebJobId) (ListTriggeredWebJobHistorySlotCompleteResult, error) {
	return c.ListTriggeredWebJobHistorySlotCompleteMatchingPredicate(ctx, id, TriggeredJobHistoryOperationPredicate{})
}

// ListTriggeredWebJobHistorySlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListTriggeredWebJobHistorySlotCompleteMatchingPredicate(ctx context.Context, id SlotTriggeredWebJobId, predicate TriggeredJobHistoryOperationPredicate) (result ListTriggeredWebJobHistorySlotCompleteResult, err error) {
	items := make([]TriggeredJobHistory, 0)

	resp, err := c.ListTriggeredWebJobHistorySlot(ctx, id)
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

	result = ListTriggeredWebJobHistorySlotCompleteResult{
		Items: items,
	}
	return
}
