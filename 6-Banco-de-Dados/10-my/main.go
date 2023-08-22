package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db.AutoMigrate(&Product{}, &Category{})

	db.Create(&Category{
		Name: "Eletronicos",
		Products: []Product{
			{Name: "TV", Price: 1000},
			{Name: "Radio", Price: 500},
			{Name: "Celular", Price: 1500},
			{Name: "Notebook", Price: 2500},
		},
	})

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	c.Name = "Novo Eletronicos"
	err = tx.Debug().Save(&c).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}
