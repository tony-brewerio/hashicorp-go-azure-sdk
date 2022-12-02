package protectioncontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ErrorDetail struct {
	Code            *string   `json:"code,omitempty"`
	Message         *string   `json:"message,omitempty"`
	Recommendations *[]string `json:"recommendations,omitempty"`
}
