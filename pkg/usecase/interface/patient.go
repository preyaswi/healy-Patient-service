package usecaseint

import (
	"patient-service/pkg/models"
)

type PatientUseCase interface {
	PatientsSignUp(patient models.PatientSignUp) (models.TokenPatient, error)
}
