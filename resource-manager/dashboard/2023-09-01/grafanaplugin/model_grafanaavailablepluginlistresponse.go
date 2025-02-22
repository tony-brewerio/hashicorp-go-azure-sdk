package grafanaplugin

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GrafanaAvailablePluginListResponse struct {
	NextLink *string                   `json:"nextLink,omitempty"`
	Value    *[]GrafanaAvailablePlugin `json:"value,omitempty"`
}
