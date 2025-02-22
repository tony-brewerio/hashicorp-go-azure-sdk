package sharesubscription

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSourceShareSynchronizationSettingsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SourceShareSynchronizationSetting

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSourceShareSynchronizationSettingsOperationResponse, error)
}

type ListSourceShareSynchronizationSettingsCompleteResult struct {
	Items []SourceShareSynchronizationSetting
}

func (r ListSourceShareSynchronizationSettingsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSourceShareSynchronizationSettingsOperationResponse) LoadMore(ctx context.Context) (resp ListSourceShareSynchronizationSettingsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSourceShareSynchronizationSettings ...
func (c ShareSubscriptionClient) ListSourceShareSynchronizationSettings(ctx context.Context, id ShareSubscriptionId) (resp ListSourceShareSynchronizationSettingsOperationResponse, err error) {
	req, err := c.preparerForListSourceShareSynchronizationSettings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ListSourceShareSynchronizationSettings", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ListSourceShareSynchronizationSettings", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSourceShareSynchronizationSettings(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ListSourceShareSynchronizationSettings", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSourceShareSynchronizationSettings prepares the ListSourceShareSynchronizationSettings request.
func (c ShareSubscriptionClient) preparerForListSourceShareSynchronizationSettings(ctx context.Context, id ShareSubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listSourceShareSynchronizationSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSourceShareSynchronizationSettingsWithNextLink prepares the ListSourceShareSynchronizationSettings request with the given nextLink token.
func (c ShareSubscriptionClient) preparerForListSourceShareSynchronizationSettingsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSourceShareSynchronizationSettings handles the response to the ListSourceShareSynchronizationSettings request. The method always
// closes the http.Response Body.
func (c ShareSubscriptionClient) responderForListSourceShareSynchronizationSettings(resp *http.Response) (result ListSourceShareSynchronizationSettingsOperationResponse, err error) {
	type page struct {
		Values   []json.RawMessage `json:"value"`
		NextLink *string           `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	temp := make([]SourceShareSynchronizationSetting, 0)
	for i, v := range respObj.Values {
		val, err := unmarshalSourceShareSynchronizationSettingImplementation(v)
		if err != nil {
			err = fmt.Errorf("unmarshalling item %d for SourceShareSynchronizationSetting (%q): %+v", i, v, err)
			return result, err
		}
		temp = append(temp, val)
	}
	result.Model = &temp
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSourceShareSynchronizationSettingsOperationResponse, err error) {
			req, err := c.preparerForListSourceShareSynchronizationSettingsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ListSourceShareSynchronizationSettings", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ListSourceShareSynchronizationSettings", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSourceShareSynchronizationSettings(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ListSourceShareSynchronizationSettings", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSourceShareSynchronizationSettingsComplete retrieves all of the results into a single object
func (c ShareSubscriptionClient) ListSourceShareSynchronizationSettingsComplete(ctx context.Context, id ShareSubscriptionId) (ListSourceShareSynchronizationSettingsCompleteResult, error) {
	return c.ListSourceShareSynchronizationSettingsCompleteMatchingPredicate(ctx, id, SourceShareSynchronizationSettingOperationPredicate{})
}

// ListSourceShareSynchronizationSettingsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ShareSubscriptionClient) ListSourceShareSynchronizationSettingsCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, predicate SourceShareSynchronizationSettingOperationPredicate) (resp ListSourceShareSynchronizationSettingsCompleteResult, err error) {
	items := make([]SourceShareSynchronizationSetting, 0)

	page, err := c.ListSourceShareSynchronizationSettings(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListSourceShareSynchronizationSettingsCompleteResult{
		Items: items,
	}
	return out, nil
}
