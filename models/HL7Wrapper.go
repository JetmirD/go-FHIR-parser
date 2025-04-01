package models

type HL7Wrapper struct {
	Patient     HL7Patient     `json:"patient"`
	Observation HL7Observation `json:"observation"`
}
