package diagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AnalysisId{}

// AnalysisId is a struct representing the Resource ID for a Analysis
type AnalysisId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	DiagnosticName    string
	AnalysisName      string
}

// NewAnalysisID returns a new AnalysisId struct
func NewAnalysisID(subscriptionId string, resourceGroupName string, siteName string, diagnosticName string, analysisName string) AnalysisId {
	return AnalysisId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		DiagnosticName:    diagnosticName,
		AnalysisName:      analysisName,
	}
}

// ParseAnalysisID parses 'input' into a AnalysisId
func ParseAnalysisID(input string) (*AnalysisId, error) {
	parser := resourceids.NewParserFromResourceIdType(AnalysisId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AnalysisId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAnalysisIDInsensitively parses 'input' case-insensitively into a AnalysisId
// note: this method should only be used for API response data and not user input
func ParseAnalysisIDInsensitively(input string) (*AnalysisId, error) {
	parser := resourceids.NewParserFromResourceIdType(AnalysisId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AnalysisId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AnalysisId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SiteName, ok = input.Parsed["siteName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteName", input)
	}

	if id.DiagnosticName, ok = input.Parsed["diagnosticName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "diagnosticName", input)
	}

	if id.AnalysisName, ok = input.Parsed["analysisName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "analysisName", input)
	}

	return nil
}

// ValidateAnalysisID checks that 'input' can be parsed as a Analysis ID
func ValidateAnalysisID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAnalysisID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Analysis ID
func (id AnalysisId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/diagnostics/%s/analyses/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.DiagnosticName, id.AnalysisName)
}

// Segments returns a slice of Resource ID Segments which comprise this Analysis ID
func (id AnalysisId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticDiagnostics", "diagnostics", "diagnostics"),
		resourceids.UserSpecifiedSegment("diagnosticName", "diagnosticValue"),
		resourceids.StaticSegment("staticAnalyses", "analyses", "analyses"),
		resourceids.UserSpecifiedSegment("analysisName", "analysisValue"),
	}
}

// String returns a human-readable description of this Analysis ID
func (id AnalysisId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Diagnostic Name: %q", id.DiagnosticName),
		fmt.Sprintf("Analysis Name: %q", id.AnalysisName),
	}
	return fmt.Sprintf("Analysis (%s)", strings.Join(components, "\n"))
}
