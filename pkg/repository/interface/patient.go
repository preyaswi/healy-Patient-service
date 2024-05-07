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
}
