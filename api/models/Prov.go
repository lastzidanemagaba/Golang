package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Cr_Propinsi) TableName() string {
	return "Cr_Propinsi"
}

type Cr_Propinsi struct {
	Prop_Id           uint32 `gorm:"primary_key;auto_increment" json:"Prop_Id"`
	Prop_Nama         string `gorm:"size:255;not null;unique" json:"Prop_Nama"`
	Prop_Nama_Singkat string `gorm:"size:255;not null;" json:"Prop_Nama_Singkat"`
	Prop_Kode         string `gorm:"size:255;not null;" json:"Prop_Kode"`
	Prop_Negara_Id    uint32 `gorm:"size:255;not null;" json:"Prop_Negara_Id"`
	Prop_Lat          string `gorm:"size:100;not null;" json:"Prop_Lat"`
	Prop_Long         string `gorm:"size:100;not null;" json:"Prop_Long"`
	Prop_Long_Lat     string `gorm:"size:100;not null;" json:"Prop_Long_Lat"`
}

func (u *Cr_Propinsi) FindAllProvs(db *gorm.DB) (*[]Cr_Propinsi, error) {
	var err error
	prov := []Cr_Propinsi{}
	err = db.Debug().Model(&Cr_Propinsi{}).Limit(100).Find(&prov).Error
	if err != nil {
		return &[]Cr_Propinsi{}, err
	}
	return &prov, err
}

func (u *Cr_Propinsi) FindProvByID(db *gorm.DB, uid uint32) (*Cr_Propinsi, error) {
	var err error
	err = db.Debug().Model(Cr_Propinsi{}).Where("prop_id  = ?", uid).Take(&u).Error
	if err != nil {
		return &Cr_Propinsi{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cr_Propinsi{}, errors.New("Prov Not Found")
	}
	return u, err
}
