package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type category struct {
	Id           int        `json:"id"`
	CategoryName string     `json:"categoryName"`
	CreateAt     *time.Time `json:"createAt"`
}

type toCreateCategory struct {
	Id           int    `json:"-" gorm:"id"`
	CategoryName string `json:"categoryName" gorm:"column:categoryName"`
}

func (toCreateCategory) TableName() string { return "categories" }

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	// now := time.Now().UTC()

	// item := category{
	// 	Id:           1,
	// 	CategoryName: "Duoc Ly",
	// 	CreateAt:     &now,
	// }

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		categories := v1.Group("/categories")
		{
			categories.POST("", createCategory(db))
		}
	}

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": item,
	// 	})
	// })

	r.Run(":8080")
}

func createCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data toCreateCategory

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
