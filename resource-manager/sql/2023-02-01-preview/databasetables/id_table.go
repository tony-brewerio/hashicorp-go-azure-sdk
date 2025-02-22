package databasetables

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TableId{}

// TableId is a struct representing the Resource ID for a Table
type TableId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	DatabaseName      string
	SchemaName        string
	TableName         string
}

// NewTableID returns a new TableId struct
func NewTableID(subscriptionId string, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string) TableId {
	return TableId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		DatabaseName:      databaseName,
		SchemaName:        schemaName,
		TableName:         tableName,
	}
}

// ParseTableID parses 'input' into a TableId
func ParseTableID(input string) (*TableId, error) {
	parser := resourceids.NewParserFromResourceIdType(TableId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TableId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTableIDInsensitively parses 'input' case-insensitively into a TableId
// note: this method should only be used for API response data and not user input
func ParseTableIDInsensitively(input string) (*TableId, error) {
	parser := resourceids.NewParserFromResourceIdType(TableId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TableId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TableId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServerName, ok = input.Parsed["serverName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serverName", input)
	}

	if id.DatabaseName, ok = input.Parsed["databaseName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseName", input)
	}

	if id.SchemaName, ok = input.Parsed["schemaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schemaName", input)
	}

	if id.TableName, ok = input.Parsed["tableName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tableName", input)
	}

	return nil
}

// ValidateTableID checks that 'input' can be parsed as a Table ID
func ValidateTableID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTableID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Table ID
func (id TableId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/databases/%s/schemas/%s/tables/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseName, id.SchemaName, id.TableName)
}

// Segments returns a slice of Resource ID Segments which comprise this Table ID
func (id TableId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
		resourceids.StaticSegment("staticSchemas", "schemas", "schemas"),
		resourceids.UserSpecifiedSegment("schemaName", "schemaValue"),
		resourceids.StaticSegment("staticTables", "tables", "tables"),
		resourceids.UserSpecifiedSegment("tableName", "tableValue"),
	}
}

// String returns a human-readable description of this Table ID
func (id TableId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Schema Name: %q", id.SchemaName),
		fmt.Sprintf("Table Name: %q", id.TableName),
	}
	return fmt.Sprintf("Table (%s)", strings.Join(components, "\n"))
}
