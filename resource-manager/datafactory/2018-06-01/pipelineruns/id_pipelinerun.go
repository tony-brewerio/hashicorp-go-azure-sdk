package pipelineruns

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PipelineRunId{}

// PipelineRunId is a struct representing the Resource ID for a Pipeline Run
type PipelineRunId struct {
	SubscriptionId    string
	ResourceGroupName string
	FactoryName       string
	RunId             string
}

// NewPipelineRunID returns a new PipelineRunId struct
func NewPipelineRunID(subscriptionId string, resourceGroupName string, factoryName string, runId string) PipelineRunId {
	return PipelineRunId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		FactoryName:       factoryName,
		RunId:             runId,
	}
}

// ParsePipelineRunID parses 'input' into a PipelineRunId
func ParsePipelineRunID(input string) (*PipelineRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(PipelineRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PipelineRunId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePipelineRunIDInsensitively parses 'input' case-insensitively into a PipelineRunId
// note: this method should only be used for API response data and not user input
func ParsePipelineRunIDInsensitively(input string) (*PipelineRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(PipelineRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PipelineRunId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PipelineRunId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.FactoryName, ok = input.Parsed["factoryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "factoryName", input)
	}

	if id.RunId, ok = input.Parsed["runId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runId", input)
	}

	return nil
}

// ValidatePipelineRunID checks that 'input' can be parsed as a Pipeline Run ID
func ValidatePipelineRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePipelineRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pipeline Run ID
func (id PipelineRunId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/pipelineRuns/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.RunId)
}

// Segments returns a slice of Resource ID Segments which comprise this Pipeline Run ID
func (id PipelineRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticPipelineRuns", "pipelineRuns", "pipelineRuns"),
		resourceids.UserSpecifiedSegment("runId", "runIdValue"),
	}
}

// String returns a human-readable description of this Pipeline Run ID
func (id PipelineRunId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Run: %q", id.RunId),
	}
	return fmt.Sprintf("Pipeline Run (%s)", strings.Join(components, "\n"))
}
