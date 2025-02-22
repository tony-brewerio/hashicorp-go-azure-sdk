package agreements

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PlanId{}

func TestNewPlanID(t *testing.T) {
	id := NewPlanID("12345678-1234-9876-4563-123456789012", "publisherIdValue", "offerIdValue", "planIdValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.PublisherId != "publisherIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PublisherId'", id.PublisherId, "publisherIdValue")
	}

	if id.OfferId != "offerIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OfferId'", id.OfferId, "offerIdValue")
	}

	if id.PlanId != "planIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlanId'", id.PlanId, "planIdValue")
	}
}

func TestFormatPlanID(t *testing.T) {
	actual := NewPlanID("12345678-1234-9876-4563-123456789012", "publisherIdValue", "offerIdValue", "planIdValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue/plans/planIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PlanId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue/plans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue/plans/planIdValue",
			Expected: &PlanId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				PublisherId:    "publisherIdValue",
				OfferId:        "offerIdValue",
				PlanId:         "planIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue/plans/planIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePlanID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.PublisherId != v.Expected.PublisherId {
			t.Fatalf("Expected %q but got %q for PublisherId", v.Expected.PublisherId, actual.PublisherId)
		}

		if actual.OfferId != v.Expected.OfferId {
			t.Fatalf("Expected %q but got %q for OfferId", v.Expected.OfferId, actual.OfferId)
		}

		if actual.PlanId != v.Expected.PlanId {
			t.Fatalf("Expected %q but got %q for PlanId", v.Expected.PlanId, actual.PlanId)
		}

	}
}

func TestParsePlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PlanId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mArKeTpLaCeOrDeRiNg",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mArKeTpLaCeOrDeRiNg/aGrEeMeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mArKeTpLaCeOrDeRiNg/aGrEeMeNtS/pUbLiShErIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mArKeTpLaCeOrDeRiNg/aGrEeMeNtS/pUbLiShErIdVaLuE/oFfErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mArKeTpLaCeOrDeRiNg/aGrEeMeNtS/pUbLiShErIdVaLuE/oFfErS/oFfErIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue/plans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mArKeTpLaCeOrDeRiNg/aGrEeMeNtS/pUbLiShErIdVaLuE/oFfErS/oFfErIdVaLuE/pLaNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue/plans/planIdValue",
			Expected: &PlanId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				PublisherId:    "publisherIdValue",
				OfferId:        "offerIdValue",
				PlanId:         "planIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MarketplaceOrdering/agreements/publisherIdValue/offers/offerIdValue/plans/planIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mArKeTpLaCeOrDeRiNg/aGrEeMeNtS/pUbLiShErIdVaLuE/oFfErS/oFfErIdVaLuE/pLaNs/pLaNiDvAlUe",
			Expected: &PlanId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				PublisherId:    "pUbLiShErIdVaLuE",
				OfferId:        "oFfErIdVaLuE",
				PlanId:         "pLaNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mArKeTpLaCeOrDeRiNg/aGrEeMeNtS/pUbLiShErIdVaLuE/oFfErS/oFfErIdVaLuE/pLaNs/pLaNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePlanIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.PublisherId != v.Expected.PublisherId {
			t.Fatalf("Expected %q but got %q for PublisherId", v.Expected.PublisherId, actual.PublisherId)
		}

		if actual.OfferId != v.Expected.OfferId {
			t.Fatalf("Expected %q but got %q for OfferId", v.Expected.OfferId, actual.OfferId)
		}

		if actual.PlanId != v.Expected.PlanId {
			t.Fatalf("Expected %q but got %q for PlanId", v.Expected.PlanId, actual.PlanId)
		}

	}
}

func TestSegmentsForPlanId(t *testing.T) {
	segments := PlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
