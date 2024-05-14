package usecaseint

import (
	"context"
	"patient-service/pkg/models"
)

type PatientUseCase interface {
	PatientsSignUp(patient models.PatientSignUp) (models.TokenPatient, error)
	PatientLogin(patient models.PatientLogin)(models.TokenPatient,error)
	IndPatientDetails(patient_id uint64)(models.SignupdetailResponse,error)
	UpdatePatientDetails(patient models.SignupdetailResponse) (models.PatientProfile,error)
	UpdatePassword(ctx context.Context,patinet_id uint64,nbody models.UpdatePassword)error

	ListPatients() ([]models.SignupdetailResponse, error)
}
