package contenttype

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ContentTypeId{}

// ContentTypeId is a struct representing the Resource ID for a Content Type
type ContentTypeId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ContentTypeId     string
}

// NewContentTypeID returns a new ContentTypeId struct
func NewContentTypeID(subscriptionId string, resourceGroupName string, serviceName string, contentTypeId string) ContentTypeId {
	return ContentTypeId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ContentTypeId:     contentTypeId,
	}
}

// ParseContentTypeID parses 'input' into a ContentTypeId
func ParseContentTypeID(input string) (*ContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ContentTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentTypeId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseContentTypeIDInsensitively parses 'input' case-insensitively into a ContentTypeId
// note: this method should only be used for API response data and not user input
func ParseContentTypeIDInsensitively(input string) (*ContentTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ContentTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentTypeId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ContentTypeId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ContentTypeId, ok = input.Parsed["contentTypeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentTypeId", input)
	}

	return nil
}

// ValidateContentTypeID checks that 'input' can be parsed as a Content Type ID
func ValidateContentTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseContentTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Content Type ID
func (id ContentTypeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/contentTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ContentTypeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Content Type ID
func (id ContentTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
	}
}

// String returns a human-readable description of this Content Type ID
func (id ContentTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
	}
	return fmt.Sprintf("Content Type (%s)", strings.Join(components, "\n"))
}
