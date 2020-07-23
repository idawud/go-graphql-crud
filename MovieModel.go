package main

import "github.com/jinzhu/gorm"

type Movie struct {
	ID uint `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Title string `json:"title" gorm:"unique;not null"`
	Minutes int `json:"minutes"`
}

var(
	DBConn *gorm.DB
)