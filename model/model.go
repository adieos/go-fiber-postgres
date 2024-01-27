package model

import "gorm.io/gorm"

type Books struct {
	ID        uint   `gorm:"primary key;autoIncrement" json:"id"`
	Title     string `gorm:"not null" json:"title"`
	Author    string `gorm:"not null" json:"author"`
	Publisher string `gorm:"not null" json:"publisher"`
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
