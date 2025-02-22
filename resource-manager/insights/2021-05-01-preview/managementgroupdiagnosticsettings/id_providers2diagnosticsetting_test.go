package managementgroupdiagnosticsettings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2DiagnosticSettingId{}

func TestNewProviders2DiagnosticSettingID(t *testing.T) {
	id := NewProviders2DiagnosticSettingID("managementGroupIdValue", "diagnosticSettingValue")

	if id.ManagementGroupId != "managementGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupId'", id.ManagementGroupId, "managementGroupIdValue")
	}

	if id.DiagnosticSettingName != "diagnosticSettingValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DiagnosticSettingName'", id.DiagnosticSettingName, "diagnosticSettingValue")
	}
}

func TestFormatProviders2DiagnosticSettingID(t *testing.T) {
	actual := NewProviders2DiagnosticSettingID("managementGroupIdValue", "diagnosticSettingValue").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights/diagnosticSettings/diagnosticSettingValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2DiagnosticSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2DiagnosticSettingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights/diagnosticSettings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights/diagnosticSettings/diagnosticSettingValue",
			Expected: &Providers2DiagnosticSettingId{
				ManagementGroupId:     "managementGroupIdValue",
				DiagnosticSettingName: "diagnosticSettingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights/diagnosticSettings/diagnosticSettingValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2DiagnosticSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.DiagnosticSettingName != v.Expected.DiagnosticSettingName {
			t.Fatalf("Expected %q but got %q for DiagnosticSettingName", v.Expected.DiagnosticSettingName, actual.DiagnosticSettingName)
		}

	}
}

func TestParseProviders2DiagnosticSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2DiagnosticSettingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.iNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights/diagnosticSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dIaGnOsTiCsEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights/diagnosticSettings/diagnosticSettingValue",
			Expected: &Providers2DiagnosticSettingId{
				ManagementGroupId:     "managementGroupIdValue",
				DiagnosticSettingName: "diagnosticSettingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Insights/diagnosticSettings/diagnosticSettingValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dIaGnOsTiCsEtTiNgS/dIaGnOsTiCsEtTiNgVaLuE",
			Expected: &Providers2DiagnosticSettingId{
				ManagementGroupId:     "mAnAgEmEnTgRoUpIdVaLuE",
				DiagnosticSettingName: "dIaGnOsTiCsEtTiNgVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.iNsIgHtS/dIaGnOsTiCsEtTiNgS/dIaGnOsTiCsEtTiNgVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2DiagnosticSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.DiagnosticSettingName != v.Expected.DiagnosticSettingName {
			t.Fatalf("Expected %q but got %q for DiagnosticSettingName", v.Expected.DiagnosticSettingName, actual.DiagnosticSettingName)
		}

	}
}

func TestSegmentsForProviders2DiagnosticSettingId(t *testing.T) {
	segments := Providers2DiagnosticSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2DiagnosticSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
