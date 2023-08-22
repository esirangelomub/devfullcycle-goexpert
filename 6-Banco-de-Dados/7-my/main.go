package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
	gorm.Model
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3309)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	category2 := Category{Name: "Kitchen"}
	db.Create(&category2)

	//Create product with category
	product := Product{
		Name:       "Mouse",
		Price:      87.52,
		CategoryID: category.ID,
	}
	db.Create(&product)

	product2 := Product{
		Name:       "Cooking pot",
		Price:      150.00,
		CategoryID: category2.ID,
	}
	db.Create(&product2)

	//Create serial number with product
	db.Create([]SerialNumber{
		{Number: "987654321", ProductID: product.ID},
		{Number: "123123123", ProductID: product2.ID},
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
