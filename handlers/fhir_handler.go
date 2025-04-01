package handlers

import (
	"log"
	"net/http"

	"hl7-fhir-parser/models"
	"hl7-fhir-parser/utils"

	"github.com/gin-gonic/gin"
)

func ParseFHIR(c *gin.Context) {
	var fhirPatient models.FHIRPatient
	var fhirObservation models.FHIRObservation

	if err := c.ShouldBindJSON(&fhirPatient); err != nil {
		log.Println("Error parsing FHIR patient data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if err := c.ShouldBindJSON(&fhirObservation); err != nil {
		log.Println("Error parsing FHIR observation data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	patient := utils.ConvertFHIRToPatient(fhirPatient)
	observation := utils.ConvertFHIRToObservation(fhirObservation)

	c.JSON(http.StatusOK, gin.H{"patient": patient, "observation": observation})
}
