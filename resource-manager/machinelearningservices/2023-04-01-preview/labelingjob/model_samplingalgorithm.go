package labelingjob

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamplingAlgorithm interface {
}

// RawSamplingAlgorithmImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSamplingAlgorithmImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalSamplingAlgorithmImplementation(input []byte) (SamplingAlgorithm, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SamplingAlgorithm into map[string]interface: %+v", err)
	}

	value, ok := temp["samplingAlgorithmType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Bayesian") {
		var out BayesianSamplingAlgorithm
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BayesianSamplingAlgorithm: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Grid") {
		var out GridSamplingAlgorithm
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GridSamplingAlgorithm: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Random") {
		var out RandomSamplingAlgorithm
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RandomSamplingAlgorithm: %+v", err)
		}
		return out, nil
	}

	out := RawSamplingAlgorithmImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
