package utils

import (
	"fmt"
	"hl7-fhir-parser/models"
)

func ConvertHL7ToPatient(hl7 models.HL7Patient) models.Patient {
	return models.Patient{
		ID:        hl7.PID,
		Name:      hl7.Name,
		Gender:    hl7.Sex,
		BirthDate: hl7.DOB,
	}
}

func ConvertHL7ToObservation(hl7 models.HL7Observation) models.Observation {
	return models.Observation{
		ID:              hl7.OBR,
		PatientID:       hl7.PatientID,
		Code:            hl7.OBXCode,
		Value:           hl7.OBXValue,
		ObservationDate: hl7.OBXDate,
	}
}

func ConvertFHIRToPatient(fhir models.FHIRPatient) models.Patient {
	return models.Patient{
		ID:        fhir.ID,
		Name:      fhir.Name[0].Text,
		Gender:    fhir.Gender,
		BirthDate: fhir.BirthDate,
	}
}

func ConvertFHIRToObservation(fhir models.FHIRObservation) models.Observation {
	return models.Observation{
		ID:              fhir.ID,
		PatientID:       fhir.Subject.Reference,
		Code:            fhir.Code.Coding[0].Code,
		Value:           fmt.Sprintf("%f", fhir.ValueQuantity.Value),
		ObservationDate: fhir.EffectiveDateTime,
	}
}
