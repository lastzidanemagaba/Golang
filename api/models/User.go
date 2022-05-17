package models

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                            uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Nickname                      string `gorm:"size:255;not null;unique" json:"nickname"`
	Device_Name                   string `gorm:"size:255;not null;unique" json:"device_name"`
	Email                         string `gorm:"size:100;not null;unique" json:"email"`
	Password                      string `gorm:"size:100;not null;" json:"password"`
	Temp_Password                 string `gorm:"size:100;not null;" json:"temp_password"`
	Name                          string `gorm:"size:100;not null;" json:"name"`
	Photo                         string `gorm:"size:100;not null;" json:"photo"`
	Last_Juz                      string `gorm:"size:100;not null;" json:"last_juz"`
	Last_Chapter                  string `gorm:"size:100;not null;" json:"last_chapter"`
	Last_Verse                    string `gorm:"size:100;not null;" json:"last_verse"`
	Last_Read_Date                string `gorm:"size:100;not null;" json:"last_read_date"`
	Scheduled_Khatam_Date         string `gorm:"size:100;not null;" json:"scheduled_khatam_date"`
	Default_Voice_Path            string `gorm:"size:100;not null;" json:"default_voice_path"`
	Penyakit_Khusus               string `gorm:"size:100;not null;" json:"Penyakit_Khusus"`
	Obat_Khusus                   string `gorm:"size:100;not null;" json:"obat_khusus"`
	Gol_Darah                     string
	Ktp_Image                     string `gorm:"size:100;not null;" json:"ktp_image"`
	Passport_Image                string `gorm:"size:100;not null;" json:"passport_image"`
	Vaccination_Certificate_Image string `gorm:"size:100;not null;" json:"vaccination_certificate_image"`
	Alergi                        string `gorm:"size:100;not null;" json:"alergi"`

	Verification_Token            string `gorm:"size:100;not null;" json:"Verification_Token"`
	Verification_Token_Date       string `gorm:"size:100;not null;" json:"Verification_Token_Date"`
	Email_Verified                string `gorm:"size:100;not null;" json:"Email_Verified"`
	Lat                           string `gorm:"size:100;not null;" json:"Lat"`
	Lng                           string `gorm:"size:100;not null;" json:"Lng"`
	Current_Location              string `gorm:"size:100;not null;" json:"Current_Location"`
	Parking_Lat                   string `gorm:"size:100;not null;" json:"Parking_Lat"`
	Parking_Lng                   string `gorm:"size:100;not null;" json:"Parking_Lng"`
	Driver_Status                 string `gorm:"size:100;not null;" json:"Driver_Status"`
	Google_Uid                    string `gorm:"size:100;not null;" json:"Google_Uid"`
	Facebook_Uid                  string `gorm:"size:100;not null;" json:"Facebook_Uid"`
	Fcm_Key                       string `gorm:"size:100;not null;" json:"Fcm_Key"`
	Pushy_Token                   string `gorm:"size:100;not null;" json:"Pushy_Token"`
	Xmpp_Password                 string `gorm:"size:100;not null;" json:"Xmpp_Password"`
	Lang_Code                     string `gorm:"size:100;not null;" json:"Lang_Code"`
	Premium                       string `gorm:"size:100;not null;" json:"Premium"`
	Premium_Months                string `gorm:"size:100;not null;" json:"Premium_Months"`
	Notification_Channel_Id       string `gorm:"size:100;not null;" json:"Notification_Channel_Id"`
	Group_Notification_Channel_Id string `gorm:"size:100;not null;" json:"Group_Notification_Channel_Id"`
	Last_Premium_Date             string `gorm:"size:100;not null;" json:"Last_Premium_Date"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"nickname":   u.Nickname,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
