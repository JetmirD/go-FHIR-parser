package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"hl7-fhir-parser/models"
	"hl7-fhir-parser/utils"

	"github.com/gin-gonic/gin"
)

func ParseHL7(c *gin.Context) {
	var data models.HL7Wrapper

	body, err := c.GetRawData()
	if err != nil {
		log.Println("Error reading request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	log.Println("Request body:", string(body))

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("Error parsing HL7 data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input structure"})
		return
	}

	log.Println("Parsed HL7 Patient:", data.Patient)
	log.Println("Parsed HL7 Observation:", data.Observation)

	patient := utils.ConvertHL7ToPatient(data.Patient)
	observation := utils.ConvertHL7ToObservation(data.Observation)

	c.JSON(http.StatusOK, gin.H{
		"patient":     patient,
		"observation": observation,
	})
}
