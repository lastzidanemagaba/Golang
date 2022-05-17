package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type TablerD interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Cr_Kab_Kota) TableName() string {
	return "Cr_Kab_Kota"
}

type Cr_Kab_Kota struct {
	Kota_Id       uint32 `gorm:"primary_key;auto_increment" json:"Kota_Id"`
	Kota_Nama     string `gorm:"size:255;not null;unique" json:"Kota_Nama"`
	Kota_Propinsi int    `gorm:"size:255;not null;" json:"Kota_Propinsi"`
	Kota_Lat      string `gorm:"size:100;not null;" json:"Kota_Lat"`
	Kota_Long     string `gorm:"size:100;not null;" json:"Kota_Long"`
	Kota_Long_Lat string `gorm:"size:100;not null;" json:"Kota_Long_Lat"`
}

func (u *Cr_Kab_Kota) FindAllCities(db *gorm.DB) (*[]Cr_Kab_Kota, error) {
	var err error
	city := []Cr_Kab_Kota{}
	err = db.Debug().Model(&Cr_Kab_Kota{}).Limit(100).Find(&city).Error
	if err != nil {
		return &[]Cr_Kab_Kota{}, err
	}
	return &city, err
}

func (u *Cr_Kab_Kota) FindCityByID(db *gorm.DB, uid uint32) (*Cr_Kab_Kota, error) {
	var err error
	err = db.Debug().Model(Cr_Kab_Kota{}).Where("kota_id = ?", uid).Take(&u).Error
	if err != nil {
		return &Cr_Kab_Kota{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cr_Kab_Kota{}, errors.New("Kota Not Found")
	}
	return u, err
}
