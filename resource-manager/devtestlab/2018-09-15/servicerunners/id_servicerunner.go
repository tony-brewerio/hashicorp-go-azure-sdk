package servicerunners

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServiceRunnerId{}

// ServiceRunnerId is a struct representing the Resource ID for a Service Runner
type ServiceRunnerId struct {
	SubscriptionId    string
	ResourceGroupName string
	LabName           string
	ServiceRunnerName string
}

// NewServiceRunnerID returns a new ServiceRunnerId struct
func NewServiceRunnerID(subscriptionId string, resourceGroupName string, labName string, serviceRunnerName string) ServiceRunnerId {
	return ServiceRunnerId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LabName:           labName,
		ServiceRunnerName: serviceRunnerName,
	}
}

// ParseServiceRunnerID parses 'input' into a ServiceRunnerId
func ParseServiceRunnerID(input string) (*ServiceRunnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServiceRunnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServiceRunnerId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServiceRunnerIDInsensitively parses 'input' case-insensitively into a ServiceRunnerId
// note: this method should only be used for API response data and not user input
func ParseServiceRunnerIDInsensitively(input string) (*ServiceRunnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServiceRunnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServiceRunnerId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServiceRunnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LabName, ok = input.Parsed["labName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "labName", input)
	}

	if id.ServiceRunnerName, ok = input.Parsed["serviceRunnerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceRunnerName", input)
	}

	return nil
}

// ValidateServiceRunnerID checks that 'input' can be parsed as a Service Runner ID
func ValidateServiceRunnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServiceRunnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Runner ID
func (id ServiceRunnerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/serviceRunners/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.ServiceRunnerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Runner ID
func (id ServiceRunnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticServiceRunners", "serviceRunners", "serviceRunners"),
		resourceids.UserSpecifiedSegment("serviceRunnerName", "serviceRunnerValue"),
	}
}

// String returns a human-readable description of this Service Runner ID
func (id ServiceRunnerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Service Runner Name: %q", id.ServiceRunnerName),
	}
	return fmt.Sprintf("Service Runner (%s)", strings.Join(components, "\n"))
}
