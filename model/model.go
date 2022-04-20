package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Imto struct {
	ID       uint64 `json:"id" gorm:"primarykey;autoincrement:true"`
	Redirect string `json:"redirect"`
	Imto     string `json:"imto" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

type Update struct {
	Redirect string `json:"redirect"`
	Imto     string `json:"imto" gorm:"unique;not null"`
	Random   bool   `json:"random"`
}

func Setup() {
	dsn := "host=172.18.0.2 user=root password=toor dbname=imto port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db
	DB.AutoMigrate(&Imto{})
	if err != nil {
		log.Fatal(err)
	}

}

func GetAll() ([]Imto, error) {
	var imtos []Imto

	tx := DB.Find(&imtos)
	if tx.Error != nil {
		return imtos, tx.Error
	}

	return imtos, nil
}

func GetOne(id int) (Imto, error) {
	var imto Imto

	tx := DB.Where("id = ?", id).First(&imto)
	if tx.Error != nil {
		return imto, tx.Error
	}

	return imto, tx.Error
}

func CreateOne(imto Imto) (Imto, error) {
	tx := DB.Create(&imto)
	if tx.Error != nil {
		return Imto{}, tx.Error
	}

	return imto, nil
}

func UpdateOne(imto Imto) (Imto, error) {
	tx := DB.Save(&imto)
	if tx.Error != nil {
		return Imto{}, tx.Error
	}

	return imto, nil
}

func DeleteOne(id uint64) error {
	tx := DB.Unscoped().Delete(&Imto{}, id)
	return tx.Error
}

func FindByUniqueUrl(url string) (Imto, error) {
	var imto Imto
	tx := DB.Where("imto = ? ", url).First(&imto)
	return imto, tx.Error
}
