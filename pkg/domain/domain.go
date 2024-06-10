package domain

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Id            uint   `json:"id" gorm:"uniquekey; not null"`
	GoogleId string  `json:"googleid" gorm:"uniquekey; not null"`
	Fullname      string `json:"fullname" gorm:"validate:required"`
	Email         string `json:"email" gorm:"validate:required"`
	Password      string `json:"password" gorm:"validate:required"`
	Gender        string `json:"gender" gorm:"validate:required"`
	Contactnumber string `json:"contactnumber" gorm:"validate:required"`
}
type Prescription struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	PatientID  uint      `json:"patient_id" gorm:"not null"`
	DoctorID   uint      `json:"doctor_id" gorm:"not null"`
	DoctorName string    `json:"doctor_name" gorm:"not null"`
	Medicine   string    `json:"medicine" gorm:"not null"`
	Dosage     string    `json:"dosage" gorm:"not null"`
	Notes      string    `json:"notes"`
	CreatedAt  time.Time `json:"created_at"`
}
