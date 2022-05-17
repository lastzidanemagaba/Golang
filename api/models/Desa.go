package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type TablerU interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Cr_Desa) TableName() string {
	return "Cr_Desa"
}

type Cr_Desa struct {
	Desa_Id        uint32 `gorm:"primary_key;auto_increment" json:"Desa_Id"`
	Desa_Nama      string `gorm:"size:255;not null;unique" json:"Desa_Nama"`
	Desa_Kecamatan int    `gorm:"size:255;not null;" json:"Desa_Kecamatan"`
	Desa_Kab_Kota  string `gorm:"size:100;not null;" json:"Desa_Kab_Kota"`
	Desa_Propinsi  string `gorm:"size:100;not null;" json:"Desa_Propinsi"`
	Desa_Kodepos   string `gorm:"size:100;not null;" json:"Desa_Kodepos"`
	Desa_Lat       string `gorm:"size:100;not null;" json:"Desa_Lat"`
	Desa_Long      string `gorm:"size:100;not null;" json:"Desa_Long"`
	Desa_Long_Lat  string `gorm:"size:100;not null;" json:"Desa_Long_Lat"`
}

func (u *Cr_Desa) FindAllDesas(db *gorm.DB) (*[]Cr_Desa, error) {
	var err error
	city := []Cr_Desa{}
	err = db.Debug().Model(&Cr_Desa{}).Limit(100).Find(&city).Error
	if err != nil {
		return &[]Cr_Desa{}, err
	}
	return &city, err
}

func (u *Cr_Desa) FindDesaByID(db *gorm.DB, uid uint32) (*Cr_Desa, error) {
	var err error
	err = db.Debug().Model(Cr_Desa{}).Where("kota_id = ?", uid).Take(&u).Error
	if err != nil {
		return &Cr_Desa{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cr_Desa{}, errors.New("Desa Not Found")
	}
	return u, err
}
