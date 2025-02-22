package componentcontinuousexportapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ExportConfigurationId{}

// ExportConfigurationId is a struct representing the Resource ID for a Export Configuration
type ExportConfigurationId struct {
	SubscriptionId    string
	ResourceGroupName string
	ComponentName     string
	ExportId          string
}

// NewExportConfigurationID returns a new ExportConfigurationId struct
func NewExportConfigurationID(subscriptionId string, resourceGroupName string, componentName string, exportId string) ExportConfigurationId {
	return ExportConfigurationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ComponentName:     componentName,
		ExportId:          exportId,
	}
}

// ParseExportConfigurationID parses 'input' into a ExportConfigurationId
func ParseExportConfigurationID(input string) (*ExportConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExportConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExportConfigurationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseExportConfigurationIDInsensitively parses 'input' case-insensitively into a ExportConfigurationId
// note: this method should only be used for API response data and not user input
func ParseExportConfigurationIDInsensitively(input string) (*ExportConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExportConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExportConfigurationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ExportConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ComponentName, ok = input.Parsed["componentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "componentName", input)
	}

	if id.ExportId, ok = input.Parsed["exportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "exportId", input)
	}

	return nil
}

// ValidateExportConfigurationID checks that 'input' can be parsed as a Export Configuration ID
func ValidateExportConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExportConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Export Configuration ID
func (id ExportConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/components/%s/exportConfiguration/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ComponentName, id.ExportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Export Configuration ID
func (id ExportConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentValue"),
		resourceids.StaticSegment("staticExportConfiguration", "exportConfiguration", "exportConfiguration"),
		resourceids.UserSpecifiedSegment("exportId", "exportIdValue"),
	}
}

// String returns a human-readable description of this Export Configuration ID
func (id ExportConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Export: %q", id.ExportId),
	}
	return fmt.Sprintf("Export Configuration (%s)", strings.Join(components, "\n"))
}
