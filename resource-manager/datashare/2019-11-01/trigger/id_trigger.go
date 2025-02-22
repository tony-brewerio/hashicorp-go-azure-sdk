package trigger

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TriggerId{}

// TriggerId is a struct representing the Resource ID for a Trigger
type TriggerId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AccountName           string
	ShareSubscriptionName string
	TriggerName           string
}

// NewTriggerID returns a new TriggerId struct
func NewTriggerID(subscriptionId string, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string) TriggerId {
	return TriggerId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AccountName:           accountName,
		ShareSubscriptionName: shareSubscriptionName,
		TriggerName:           triggerName,
	}
}

// ParseTriggerID parses 'input' into a TriggerId
func ParseTriggerID(input string) (*TriggerId, error) {
	parser := resourceids.NewParserFromResourceIdType(TriggerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggerId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTriggerIDInsensitively parses 'input' case-insensitively into a TriggerId
// note: this method should only be used for API response data and not user input
func ParseTriggerIDInsensitively(input string) (*TriggerId, error) {
	parser := resourceids.NewParserFromResourceIdType(TriggerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggerId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TriggerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AccountName, ok = input.Parsed["accountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accountName", input)
	}

	if id.ShareSubscriptionName, ok = input.Parsed["shareSubscriptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "shareSubscriptionName", input)
	}

	if id.TriggerName, ok = input.Parsed["triggerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "triggerName", input)
	}

	return nil
}

// ValidateTriggerID checks that 'input' can be parsed as a Trigger ID
func ValidateTriggerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTriggerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Trigger ID
func (id TriggerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataShare/accounts/%s/shareSubscriptions/%s/triggers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ShareSubscriptionName, id.TriggerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Trigger ID
func (id TriggerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataShare", "Microsoft.DataShare", "Microsoft.DataShare"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticShareSubscriptions", "shareSubscriptions", "shareSubscriptions"),
		resourceids.UserSpecifiedSegment("shareSubscriptionName", "shareSubscriptionValue"),
		resourceids.StaticSegment("staticTriggers", "triggers", "triggers"),
		resourceids.UserSpecifiedSegment("triggerName", "triggerValue"),
	}
}

// String returns a human-readable description of this Trigger ID
func (id TriggerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Share Subscription Name: %q", id.ShareSubscriptionName),
		fmt.Sprintf("Trigger Name: %q", id.TriggerName),
	}
	return fmt.Sprintf("Trigger (%s)", strings.Join(components, "\n"))
}
