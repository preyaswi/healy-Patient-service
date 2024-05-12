package usecaseint

import (
	"patient-service/pkg/models"
)

type PatientUseCase interface {
	PatientsSignUp(patient models.PatientSignUp) (models.TokenPatient, error)
	PatientLogin(patient models.PatientLogin)(models.TokenPatient,error)
	IndPatientDetails(patient_id uint64)(models.SignupdetailResponse,error)
}
