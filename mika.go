package mika

import (
	"encoding/json"

	"github.com/ekokurniadi/micagen-for-dart/generator"
	"github.com/ekokurniadi/micagen-for-dart/schemas"
	"github.com/gin-gonic/gin"
)

type InterfaceHandler interface {
	RunGenerator()
}

type doFunction struct{}

func Initialized() *doFunction {
	return &doFunction{}
}

func (d *doFunction) RunGenerator() {
	router := gin.Default()
	router.POST("/generate", func(c *gin.Context) {
		var project schemas.Project
		c.ShouldBindJSON(&project)

		gen := generator.NewGeneratorHandler(project)

		gen.GenerateFeature()

		p, err := json.Marshal(project)
		if err != nil {
			c.Writer.Write([]byte("Parameter yang anda kirim tidak valid"))
			return
		}
		c.Writer.Write([]byte("{\n"))
		c.Writer.Write([]byte("\"status\":200,\n"))
		c.Writer.Write([]byte("\"message\": \"Generate Berhasil, silahkan cek directory " + project.FeatureName + " pada project anda,\",\n"))
		c.Writer.Write([]byte("\t\"parameter\":" + string(p) + "\n"))
		c.Writer.Write([]byte("}"))
	})
	router.Run()
}
