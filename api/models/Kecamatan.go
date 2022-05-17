package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type TablerF interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Cr_Kecamatan) TableName() string {
	return "Cr_Kecamatan"
}

type Cr_Kecamatan struct {
	Kec_Id       uint32 `gorm:"primary_key;auto_increment" json:"Kec_Id"`
	Kec_Nama     string `gorm:"size:255;not null;unique" json:"Kec_Nama"`
	Kec_Kab_Kota int    `gorm:"size:255;not null;" json:"Kec_Kab_Kota"`
	Kec_Propinsi string `gorm:"size:100;not null;" json:"Kec_Propinsi"`
	Kec_Kodepos  string `gorm:"size:100;not null;" json:"Kec_Kodepos"`
	Kec_Lat      string `gorm:"size:100;not null;" json:"Kec_Lat"`
	Kec_Long     string `gorm:"size:100;not null;" json:"Kec_Long"`
	Kec_Long_Lat string `gorm:"size:100;not null;" json:"Kec_Long_Lat"`
}

func (u *Cr_Kecamatan) FindAllKecamatans(db *gorm.DB) (*[]Cr_Kecamatan, error) {
	var err error
	city := []Cr_Kecamatan{}
	err = db.Debug().Model(&Cr_Kecamatan{}).Limit(100).Find(&city).Error
	if err != nil {
		return &[]Cr_Kecamatan{}, err
	}
	return &city, err
}

func (u *Cr_Kecamatan) FindKecamatanByID(db *gorm.DB, uid uint32) (*Cr_Kecamatan, error) {
	var err error
	err = db.Debug().Model(Cr_Kecamatan{}).Where("kec_id = ?", uid).Take(&u).Error
	if err != nil {
		return &Cr_Kecamatan{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cr_Kecamatan{}, errors.New("Kecamatan Not Found")
	}
	return u, err
}
