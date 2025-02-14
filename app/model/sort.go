package model

type SortRequestPayload struct {
	SortKeys []string               `json:"sortKeys"`
	Payload  map[string]interface{} `json:"payload"`
}
