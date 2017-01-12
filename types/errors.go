package types

import "encoding/json"

/*
APIRequestError ...
*/
type APIRequestError struct {
	Err apiError `json:"error,omitempty"`
}

type apiError struct {
	Message string `json:"message,omitempty"`
	Sent    string `json:"sent,omitempty"`
}

func (ruleError *APIRequestError) Error() string {
	resp, err := json.Marshal(ruleError)
	if err != nil {
		return err.Error()
	}
	return string(resp[:])
}
