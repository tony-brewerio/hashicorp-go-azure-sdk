package jobtargetgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAgentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]JobTargetGroup
}

type ListByAgentCompleteResult struct {
	Items []JobTargetGroup
}

// ListByAgent ...
func (c JobTargetGroupsClient) ListByAgent(ctx context.Context, id JobAgentId) (result ListByAgentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/targetGroups", id.ID()),
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
		Values *[]JobTargetGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByAgentComplete retrieves all the results into a single object
func (c JobTargetGroupsClient) ListByAgentComplete(ctx context.Context, id JobAgentId) (ListByAgentCompleteResult, error) {
	return c.ListByAgentCompleteMatchingPredicate(ctx, id, JobTargetGroupOperationPredicate{})
}

// ListByAgentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JobTargetGroupsClient) ListByAgentCompleteMatchingPredicate(ctx context.Context, id JobAgentId, predicate JobTargetGroupOperationPredicate) (result ListByAgentCompleteResult, err error) {
	items := make([]JobTargetGroup, 0)

	resp, err := c.ListByAgent(ctx, id)
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

	result = ListByAgentCompleteResult{
		Items: items,
	}
	return
}
