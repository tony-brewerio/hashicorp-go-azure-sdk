package productgrouplink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupLinkId{}

// GroupLinkId is a struct representing the Resource ID for a Group Link
type GroupLinkId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ProductId         string
	GroupLinkId       string
}

// NewGroupLinkID returns a new GroupLinkId struct
func NewGroupLinkID(subscriptionId string, resourceGroupName string, serviceName string, productId string, groupLinkId string) GroupLinkId {
	return GroupLinkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ProductId:         productId,
		GroupLinkId:       groupLinkId,
	}
}

// ParseGroupLinkID parses 'input' into a GroupLinkId
func ParseGroupLinkID(input string) (*GroupLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupLinkId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupLinkIDInsensitively parses 'input' case-insensitively into a GroupLinkId
// note: this method should only be used for API response data and not user input
func ParseGroupLinkIDInsensitively(input string) (*GroupLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupLinkId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupLinkId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServiceName, ok = input.Parsed["serviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceName", input)
	}

	if id.ProductId, ok = input.Parsed["productId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "productId", input)
	}

	if id.GroupLinkId, ok = input.Parsed["groupLinkId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupLinkId", input)
	}

	return nil
}

// ValidateGroupLinkID checks that 'input' can be parsed as a Group Link ID
func ValidateGroupLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Link ID
func (id GroupLinkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/products/%s/groupLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ProductId, id.GroupLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Link ID
func (id GroupLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticProducts", "products", "products"),
		resourceids.UserSpecifiedSegment("productId", "productIdValue"),
		resourceids.StaticSegment("staticGroupLinks", "groupLinks", "groupLinks"),
		resourceids.UserSpecifiedSegment("groupLinkId", "groupLinkIdValue"),
	}
}

// String returns a human-readable description of this Group Link ID
func (id GroupLinkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Product: %q", id.ProductId),
		fmt.Sprintf("Group Link: %q", id.GroupLinkId),
	}
	return fmt.Sprintf("Group Link (%s)", strings.Join(components, "\n"))
}
