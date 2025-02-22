package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotNetworkFeatureId{}

// SlotNetworkFeatureId is a struct representing the Resource ID for a Slot Network Feature
type SlotNetworkFeatureId struct {
	SubscriptionId     string
	ResourceGroupName  string
	SiteName           string
	SlotName           string
	NetworkFeatureName string
}

// NewSlotNetworkFeatureID returns a new SlotNetworkFeatureId struct
func NewSlotNetworkFeatureID(subscriptionId string, resourceGroupName string, siteName string, slotName string, networkFeatureName string) SlotNetworkFeatureId {
	return SlotNetworkFeatureId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		SiteName:           siteName,
		SlotName:           slotName,
		NetworkFeatureName: networkFeatureName,
	}
}

// ParseSlotNetworkFeatureID parses 'input' into a SlotNetworkFeatureId
func ParseSlotNetworkFeatureID(input string) (*SlotNetworkFeatureId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotNetworkFeatureId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SlotNetworkFeatureId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSlotNetworkFeatureIDInsensitively parses 'input' case-insensitively into a SlotNetworkFeatureId
// note: this method should only be used for API response data and not user input
func ParseSlotNetworkFeatureIDInsensitively(input string) (*SlotNetworkFeatureId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotNetworkFeatureId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SlotNetworkFeatureId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SlotNetworkFeatureId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SlotName, ok = input.Parsed["slotName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "slotName", input)
	}

	if id.NetworkFeatureName, ok = input.Parsed["networkFeatureName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "networkFeatureName", input)
	}

	return nil
}

// ValidateSlotNetworkFeatureID checks that 'input' can be parsed as a Slot Network Feature ID
func ValidateSlotNetworkFeatureID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotNetworkFeatureID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Network Feature ID
func (id SlotNetworkFeatureId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/networkFeatures/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.NetworkFeatureName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Network Feature ID
func (id SlotNetworkFeatureId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticSlots", "slots", "slots"),
		resourceids.UserSpecifiedSegment("slotName", "slotValue"),
		resourceids.StaticSegment("staticNetworkFeatures", "networkFeatures", "networkFeatures"),
		resourceids.UserSpecifiedSegment("networkFeatureName", "networkFeatureValue"),
	}
}

// String returns a human-readable description of this Slot Network Feature ID
func (id SlotNetworkFeatureId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Network Feature Name: %q", id.NetworkFeatureName),
	}
	return fmt.Sprintf("Slot Network Feature (%s)", strings.Join(components, "\n"))
}
