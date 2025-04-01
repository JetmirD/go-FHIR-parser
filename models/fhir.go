package models

type FHIRPatient struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Name         []struct {
		Text string `json:"text"`
	} `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthDate"`
}

type FHIRObservation struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Subject      struct {
		Reference string `json:"reference"`
	} `json:"subject"`
	Code struct {
		Coding []struct {
			Code string `json:"code"`
		} `json:"coding"`
	} `json:"code"`
	ValueQuantity struct {
		Value float64 `json:"value"`
	} `json:"valueQuantity"`
	EffectiveDateTime string `json:"effectiveDateTime"`
}
