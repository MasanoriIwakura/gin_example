package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"app/database/models"
	"fmt"
)

type Sample = models.Sample

func SampleIndex(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	var sample Sample
	err := db.First(&sample).Error

	fmt.Println(err)
	if err != nil {
		db.Create(&Sample{Name: "TEST"})
		db.First(&sample)
	}

	c.JSON(200, gin.H{
		"sample": &sample.Name,
	})
}
