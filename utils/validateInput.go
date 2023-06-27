package utils

import (
	"encoding/json"
)

func IsInputValid(input interface{}) bool {
	if input == nil {
		return false
	}

	auxInput := make(map[string]interface{})
	toJson, err := json.Marshal(input)
	if err != nil {
		return false
	}

	err = json.Unmarshal(toJson, &auxInput)
	if err != nil {
		return false
	}

	for _, value := range auxInput {
		if value == nil || value == "" {
			return false
		}
	}

	return true
}
