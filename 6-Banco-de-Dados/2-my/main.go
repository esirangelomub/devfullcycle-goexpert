package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3309)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{})

	// Create
	//db.Create(&Product{Name: "Laptop", Price: 1987.52})

	// Create batch
	//db.Create([]Product{
	//	{Name: "Laptop", Price: 1000},
	//	{Name: "Mouse", Price: 100},
	//	{Name: "Keyboard", Price: 200},
	//})

	// Select one by conditions
	//var product Product
	////db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "name = ?", "Mouse") // find product with name Mouse
	//fmt.Println(product)

	// Select all
	//var products []Product
	//// db.Find(&products) // find all products
	//db.Where("price <= ?", 200).Find(&products) // find all products with price <= 200
	//for _, p := range products {
	//	fmt.Println(p)
	//}

	//var p Product
	//db.First(&p, 1)
	//p.Name = "New Laptop"
	//db.Save(&p)
	//
	//var p2 Product
	//db.First(&p2, 1)
	//fmt.Println(p2.Name)
	//db.Delete(&p2)
}
