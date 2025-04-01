package main

import (
	"hl7-fhir-parser/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/parse-hl7", handlers.ParseHL7)
	r.POST("/parse-fhir", handlers.ParseFHIR)

	r.Run(":8080")
}
