package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:category_products"`
	gorm.Model
}

type Product struct {
	ID            int `gorm:"primaryKey"`
	Name          string
	Price         float64
	Categories    []Category     `gorm:"many2many:category_products"`
	SerialNumbers []SerialNumber `gorm:"many2many:products_serial_numbers"`
	gorm.Model
}

type SerialNumber struct {
	ID       int `gorm:"primaryKey"`
	Number   string
	Products []Product `gorm:"many2many:products_serial_numbers"`
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
		Categories: []Category{category, category2},
	}
	db.Create(&product)

	product2 := Product{
		Name:       "Cooking pot",
		Price:      150.00,
		Categories: []Category{category, category2},
	}
	db.Create(&product2)

	//Create serial number with product
	db.Create([]SerialNumber{
		{Number: "987654321", Products: []Product{product}},
		{Number: "123456789", Products: []Product{product}},
		{Number: "321321321", Products: []Product{product2}},
		{Number: "456456456", Products: []Product{product2}},
	})

	var categories []Category
	err = db.Model(&Category{}).
		Preload("Products").
		Preload("Products.SerialNumbers").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println("  - ", product.Name)
			for _, serialNumber := range product.SerialNumbers {
				fmt.Println("    * ", serialNumber.Number)
			}
		}
	}
}
