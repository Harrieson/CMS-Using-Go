package model

import "gorm.io/gorm"

var db *gorm.DB

type Goly struct {
	ID       uint64 `json: "id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func SetUp() {
	dsn := "host=172.17.0.2 user=admin password=Harr1350n dbname=admin port=5432 sshmode=disbale"

	gorm.Open()
}
