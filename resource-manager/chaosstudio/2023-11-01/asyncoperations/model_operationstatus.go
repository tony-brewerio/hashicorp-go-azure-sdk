package asyncoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatus struct {
	EndTime   *string      `json:"endTime,omitempty"`
	Error     *ErrorDetail `json:"error,omitempty"`
	Id        *string      `json:"id,omitempty"`
	Name      *string      `json:"name,omitempty"`
	StartTime *string      `json:"startTime,omitempty"`
	Status    *string      `json:"status,omitempty"`
}
