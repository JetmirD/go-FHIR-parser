package models

type HL7Patient struct {
	PID  string `json:"pid"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
	DOB  string `json:"dob"`
}

type HL7Observation struct {
	OBR       string `json:"obr"`
	PatientID string `json:"patientId"`
	OBXCode   string `json:"obxCode"`
	OBXValue  string `json:"obxValue"`
	OBXDate   string `json:"obxDate"`
}
