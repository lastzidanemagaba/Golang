package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type TablerT interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Cr_Negara) TableName() string {
	return "Cr_Negara"
}

type Cr_Negara struct {
	Id     uint32 `gorm:"primary_key;auto_increment" json:"Id"`
	Nama   string `gorm:"size:255;not null;unique" json:"Nama"`
	Alpha3 string `gorm:"size:255;not null;" json:"Alpha3"`
	Ket    string `gorm:"size:100;not null;" json:"Ket"`
}

func (u *Cr_Negara) FindAllNegaras(db *gorm.DB) (*[]Cr_Negara, error) {
	var err error
	city := []Cr_Negara{}
	err = db.Debug().Model(&Cr_Negara{}).Limit(100).Find(&city).Error
	if err != nil {
		return &[]Cr_Negara{}, err
	}
	return &city, err
}

func (u *Cr_Negara) FindNegaraByID(db *gorm.DB, uid uint32) (*Cr_Negara, error) {
	var err error
	err = db.Debug().Model(Cr_Negara{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Cr_Negara{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cr_Negara{}, errors.New("Negara Not Found")
	}
	return u, err
}
