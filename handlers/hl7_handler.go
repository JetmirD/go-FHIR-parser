package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ParseHL7(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Println("Error reading request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	hl7Message := string(body)

	segments := strings.Split(hl7Message, "\n")

	var pid, firstName, lastName, sex, dob, obr, patientID, obxCode, obxValue, obxDate string

	for _, segment := range segments {
		fields := strings.Split(segment, "|")
		if len(fields) > 0 {
			switch fields[0] {
			case "PID":
				pid = fields[1]
				names := strings.Split(fields[4], "^")
				if len(names) > 1 {
					lastName = names[2]
					firstName = names[1]
				}
				sex = fields[7]
				dob = fields[6]
				dobParsed, err := time.Parse("20060102", dob)
				if err == nil {
					dob = dobParsed.Format("01/02/2006")
				} else {
					log.Println("Error parsing DOB:", err)
				}

			case "OBR":
				obr = fields[1]
				patientID = fields[3]
			case "OBX":
				obxCode = fields[3]
				obxValue = fields[5]
				if len(fields) > 14 {
					obxDate = fields[14]
				}
			}
		}
	}

	log.Println("✅ Parsed HL7 Patient:")
	log.Println("  - PID:     ", pid)
	log.Println("  - First Name:", firstName)
	log.Println("  - Last Name: ", lastName)
	log.Println("  - Sex:     ", sex)
	log.Println("  - DOB:     ", dob)

	log.Println("✅ Parsed HL7 Observation:")
	log.Println("  - OBR:      ", obr)
	log.Println("  - Patient ID:", patientID)
	log.Println("  - OBX Code: ", obxCode)
	log.Println("  - OBX Value:", obxValue)
	log.Println("  - OBX Date: ", obxDate)

	patient := map[string]string{
		"PID":       pid,
		"FirstName": firstName,
		"LastName":  lastName,
		"Sex":       sex,
		"DOB":       dob,
	}
	observation := map[string]string{
		"OBR":       obr,
		"PatientID": patientID,
		"OBXCode":   obxCode,
		"OBXValue":  obxValue,
		"OBXDate":   obxDate,
	}

	c.JSON(http.StatusOK, gin.H{
		"patient":     patient,
		"observation": observation,
	})
}
