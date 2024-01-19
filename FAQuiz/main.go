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
	Id                 int        `json:"id"`
	CategoryName       string     `json:"categoryName" gorm:"column:categoryName"`
	CategoryImage      string     `json:"categoryImage" gorm:"column:categoryImage"`
	CategoryImageColor string     `json:"categoryImageColor" gorm:"column:categoryImageColor"`
	CreateAt           *time.Time `json:"createAt" gorm:"column:createAt"`
}

// type toCreateCategory struct {
// 	Id           int    `json:"-" gorm:"column:id"`
// 	CategoryName string `json:"categoryName" gorm:"column:categoryName"`
// 	CategoryImage string `json:"categoryImage" gorm:"column:categoryImage"`
// 	CategoryImageColor string `json:"categoryImageColor" gorm:"column:categoryImageColor"`
// }

// func (toCreateCategory) TableName() string { return "categories" }

// type categoryTags struct {
// 	Id         int    `json:"id"`
// 	CategoryId string `json:"categoryId" gorm:"column:categoryId"`
// 	TagId      string `json:"tagId" gorm:"column:tagId"`
// }

// type categoryMajors struct {
// 	Id         int    `json:"id"`
// 	CategoryId string `json:"categoryId" gorm:"column:categoryId"`
// 	MajorsId   string `json:"majorsId" gorm:"column:majorId"`
// }

type toGetCategory struct {
	Keyword  string  `json:"keyword"`
	TagId    []int64 `json:"tagId"`
	MajorsId []int64 `json:"majorsId"`
	Offset   int     `json:"offset"`
	Limit    int     `json:"limit"`
}

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		categories := v1.Group("/categories")
		{
			// categories.POST("", createCategory(db))
			categories.POST("", getManyCategory(db))
		}
	}

	r.Run(":8080")
}

// func createCategory(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var data toCreateCategory

// 		if err := c.ShouldBind(&data); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		}

// 		if err := db.Create(&data).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"data": data,
// 		})
// 	}
// }

func getManyCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data toGetCategory

		if err := c.ShouldBind(&data); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		fmt.Printf("%+v", data)

		var categories []category

		if data.Keyword == "" {
			if cap(data.TagId) == 0 && cap(data.MajorsId) == 0 {
				if err := db.Limit(data.Limit).Offset(data.Offset).Find(&categories).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})

					return
				}
			} else if cap(data.TagId) != 0 && cap(data.MajorsId) == 0 {
				if err := db.Table("categories").Select("DISTINCT categories.id", "categories.categoryName",
					"categories.categoryImage", "categories.categoryImageColor", "categories.createAt").Joins("INNER JOIN "+
					"categoryTags ON categoryTags.categoryId= categories.id").Where("categoryTags.tagId IN ?", data.TagId).Limit(data.Limit).Offset(data.Offset).Find(&categories).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})

					return
				}
			} else {
				if err := db.Table("categories").Select("DISTINCT categories.id", "categories.categoryName",
					"categories.categoryImage", "categories.categoryImageColor",
					"categories.createAt").Joins("INNER JOIN categoryTags ON categoryTags.categoryId ="+
					"categories.id").Joins("INNER JOIN categoryMajors ON categoryMajors.categoryId ="+
					"categories.id").Where("categoryTags.tagId IN ? AND categoryMajors.majorsId IN ?", data.TagId, data.MajorsId).Limit(data.Limit).Offset(data.Offset).Find(&categories).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})

					return
				}
			}
		} else {
			if err := db.Where("categoryName LIKE ?", "%"+data.Keyword+"%").Limit(data.Limit).Offset(data.Offset).Find(&categories).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"data": categories,
		})
	}
}

// func convertNumeric(s string) []int64 {
// 	numberStrings := strings.Split(s, ", ")

// 	var numbers []int64

// 	for _, str := range numberStrings {
// 		num, err := strconv.Atoi(str)

// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return []int64{}
// 		}

// 		numbers = append(numbers, int64(num))
// 	}

// 	return numbers
// }
