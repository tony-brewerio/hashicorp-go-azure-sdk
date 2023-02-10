package policyassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentUpdateProperties struct {
	Overrides         *[]Override         `json:"overrides,omitempty"`
	ResourceSelectors *[]ResourceSelector `json:"resourceSelectors,omitempty"`
}
