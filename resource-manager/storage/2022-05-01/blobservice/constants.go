package blobservice

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedMethods string

const (
	AllowedMethodsDELETE  AllowedMethods = "DELETE"
	AllowedMethodsGET     AllowedMethods = "GET"
	AllowedMethodsHEAD    AllowedMethods = "HEAD"
	AllowedMethodsMERGE   AllowedMethods = "MERGE"
	AllowedMethodsOPTIONS AllowedMethods = "OPTIONS"
	AllowedMethodsPATCH   AllowedMethods = "PATCH"
	AllowedMethodsPOST    AllowedMethods = "POST"
	AllowedMethodsPUT     AllowedMethods = "PUT"
)

func PossibleValuesForAllowedMethods() []string {
	return []string{
		string(AllowedMethodsDELETE),
		string(AllowedMethodsGET),
		string(AllowedMethodsHEAD),
		string(AllowedMethodsMERGE),
		string(AllowedMethodsOPTIONS),
		string(AllowedMethodsPATCH),
		string(AllowedMethodsPOST),
		string(AllowedMethodsPUT),
	}
}

func (s *AllowedMethods) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowedMethods(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowedMethods(input string) (*AllowedMethods, error) {
	vals := map[string]AllowedMethods{
		"delete":  AllowedMethodsDELETE,
		"get":     AllowedMethodsGET,
		"head":    AllowedMethodsHEAD,
		"merge":   AllowedMethodsMERGE,
		"options": AllowedMethodsOPTIONS,
		"patch":   AllowedMethodsPATCH,
		"post":    AllowedMethodsPOST,
		"put":     AllowedMethodsPUT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowedMethods(input)
	return &out, nil
}

type Name string

const (
	NameAccessTimeTracking Name = "AccessTimeTracking"
)

func PossibleValuesForName() []string {
	return []string{
		string(NameAccessTimeTracking),
	}
}

func (s *Name) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseName(input string) (*Name, error) {
	vals := map[string]Name{
		"accesstimetracking": NameAccessTimeTracking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Name(input)
	return &out, nil
}

type SkuName string

const (
	SkuNamePremiumLRS     SkuName = "Premium_LRS"
	SkuNamePremiumZRS     SkuName = "Premium_ZRS"
	SkuNameStandardGRS    SkuName = "Standard_GRS"
	SkuNameStandardGZRS   SkuName = "Standard_GZRS"
	SkuNameStandardLRS    SkuName = "Standard_LRS"
	SkuNameStandardRAGRS  SkuName = "Standard_RAGRS"
	SkuNameStandardRAGZRS SkuName = "Standard_RAGZRS"
	SkuNameStandardZRS    SkuName = "Standard_ZRS"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNamePremiumLRS),
		string(SkuNamePremiumZRS),
		string(SkuNameStandardGRS),
		string(SkuNameStandardGZRS),
		string(SkuNameStandardLRS),
		string(SkuNameStandardRAGRS),
		string(SkuNameStandardRAGZRS),
		string(SkuNameStandardZRS),
	}
}

func (s *SkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"premium_lrs":     SkuNamePremiumLRS,
		"premium_zrs":     SkuNamePremiumZRS,
		"standard_grs":    SkuNameStandardGRS,
		"standard_gzrs":   SkuNameStandardGZRS,
		"standard_lrs":    SkuNameStandardLRS,
		"standard_ragrs":  SkuNameStandardRAGRS,
		"standard_ragzrs": SkuNameStandardRAGZRS,
		"standard_zrs":    SkuNameStandardZRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}
