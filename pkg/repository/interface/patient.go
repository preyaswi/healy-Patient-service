package interfaces

import (
	"patient-service/pkg/domain"
	"patient-service/pkg/models"
)

type PatientRepository interface {
	CheckPatientExistsByEmail(email string) (*domain.Patient, error)
	CheckPatientExistsByPhone(phone string) (*domain.Patient, error)
	PatientSignUp(patient models.PatientSignUp) (models.SignupdetailResponse, error)
	FindPatientByEmail(email string)(models.PatientDetails,error)
	IndPatientDetails(patient_id uint64)(models.SignupdetailResponse,error)
	CheckPatientAvailability(email string) bool 
	UpdatePatientEmail(email string, PatientID uint) error
	UpdatePatientPhone(phone string, PatientID uint) error
	UpdateName(name string, PatientID uint) error
	UserDetails(userID int) (models.PatientProfile, error)
	PatientPassword(userID int) (string, error)
	UpdatePatientPassword(password string, userID int) error
}
