package models

type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthDate"`
}

type Observation struct {
	ID              string `json:"id"`
	PatientID       string `json:"patientId"`
	Code            string `json:"code"`
	Value           string `json:"value"`
	ObservationDate string `json:"observationDate"`
}
