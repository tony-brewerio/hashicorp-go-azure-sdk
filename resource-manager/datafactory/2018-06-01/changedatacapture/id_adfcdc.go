package changedatacapture

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AdfcdcId{}

// AdfcdcId is a struct representing the Resource ID for a Adfcdc
type AdfcdcId struct {
	SubscriptionId    string
	ResourceGroupName string
	FactoryName       string
	AdfcdcName        string
}

// NewAdfcdcID returns a new AdfcdcId struct
func NewAdfcdcID(subscriptionId string, resourceGroupName string, factoryName string, adfcdcName string) AdfcdcId {
	return AdfcdcId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		FactoryName:       factoryName,
		AdfcdcName:        adfcdcName,
	}
}

// ParseAdfcdcID parses 'input' into a AdfcdcId
func ParseAdfcdcID(input string) (*AdfcdcId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdfcdcId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdfcdcId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAdfcdcIDInsensitively parses 'input' case-insensitively into a AdfcdcId
// note: this method should only be used for API response data and not user input
func ParseAdfcdcIDInsensitively(input string) (*AdfcdcId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdfcdcId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdfcdcId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AdfcdcId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AdfcdcName, ok = input.Parsed["adfcdcName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "adfcdcName", input)
	}

	return nil
}

// ValidateAdfcdcID checks that 'input' can be parsed as a Adfcdc ID
func ValidateAdfcdcID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdfcdcID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Adfcdc ID
func (id AdfcdcId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/adfcdcs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.AdfcdcName)
}

// Segments returns a slice of Resource ID Segments which comprise this Adfcdc ID
func (id AdfcdcId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticAdfcdcs", "adfcdcs", "adfcdcs"),
		resourceids.UserSpecifiedSegment("adfcdcName", "adfcdcValue"),
	}
}

// String returns a human-readable description of this Adfcdc ID
func (id AdfcdcId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Adfcdc Name: %q", id.AdfcdcName),
	}
	return fmt.Sprintf("Adfcdc (%s)", strings.Join(components, "\n"))
}
