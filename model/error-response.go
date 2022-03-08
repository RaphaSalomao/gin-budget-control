package model

import (
	"github.com/dlclark/regexp2"
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Error   string    `json:"error,omitempty"`
	Message string    `json:"message,omitempty"`
	Id      uuid.UUID `json:"id,omitempty"`
}

var ValidationErrorResponse = func(err error) map[string]string {
	resp := make(map[string]string)
	keyRegex := regexp2.MustCompile(`[\w.]+(?='\sError)`, 0)
	valueRegex := regexp2.MustCompile(`(?<=Error:).+?tag`, 0)

	for keyMatches, _ := keyRegex.FindStringMatch(err.Error()); keyMatches != nil; keyMatches, _ = keyRegex.FindNextMatch(keyMatches) {
		valueMatches, _ := valueRegex.FindStringMatchStartingAt(err.Error(), keyMatches.Index)
		resp[keyMatches.String()] = valueMatches.String()
	}

	return resp
}
