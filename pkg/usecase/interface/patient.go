package usecaseint

import (
	"patient-service/pkg/models"
)

type PatientUseCase interface {
	GoogleSignIn(googleid,googlename,googleEmail string) (models.TokenPatient, error) 
	IndPatientDetails(patient_id uint64)(models.SignupdetailResponse,error)
	UpdatePatientDetails(patient models.SignupdetailResponse) (models.PatientProfile,error)

	ListPatients() ([]models.SignupdetailResponse, error)
}
