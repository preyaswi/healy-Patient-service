package usecase

import (
	"errors"
	"patient-service/pkg/helper"
	"patient-service/pkg/models"
	interfaces "patient-service/pkg/repository/interface"
	usecaseint "patient-service/pkg/usecase/interface"

)

type patientUseCase struct {
	patientRepository interfaces.PatientRepository
}

func NewPatientUseCase(repository interfaces.PatientRepository) usecaseint.PatientUseCase {
	return &patientUseCase{
		patientRepository: repository,
	}
}

func (pr *patientUseCase) GoogleSignIn(googleid,googlename,googleEmail string) (models.TokenPatient, error)  {
	patient,err:=pr.patientRepository.FindOrCreatePatientByGoogleID(googleid,googleEmail,googlename)
	if err!=nil{
		return models.TokenPatient{},err
	}
	accesstoken,err:=helper.GenerateAccessToken(patient)
	if err!=nil{
		return models.TokenPatient{},err
	}
	refreshToken,err:=helper.GenerateRefreshToken(patient)
	if err!=nil{
		return models.TokenPatient{},err
	}
	return models.TokenPatient{
		Patient: patient,
		AccessToken: accesstoken,
		RefreshToken: refreshToken,
	},nil

}
func (pr *patientUseCase) IndPatientDetails(patient_id uint64) (models.SignupdetailResponse, error) {
	pateint, err := pr.patientRepository.IndPatientDetails(patient_id)
	if err != nil {
		return models.SignupdetailResponse{}, nil
	}
	return pateint, nil
}
func (pr *patientUseCase) UpdatePatientDetails(patient models.SignupdetailResponse) (models.PatientProfile, error) {
	userExist := pr.patientRepository.CheckPatientAvailability(patient.Email)
	// update with email that does not already exist
	if userExist {
		return models.PatientProfile{}, errors.New("user already exist, choose different email")
	}
	if patient.Contactnumber != "" {

		userExistByPhone, err := pr.patientRepository.CheckPatientExistsByPhone(patient.Contactnumber)

		if err != nil {
			return models.PatientProfile{}, errors.New("error with server")
		}
		if userExistByPhone != nil {
			return models.PatientProfile{}, errors.New("user with this phone is already exists")
		}
	}
	// which all field are not empty (which are provided from the front end should be updated)
	if patient.Email != "" {
		if err := pr.patientRepository.UpdatePatientEmail(patient.Email, patient.Id); err != nil {
			return models.PatientProfile{}, err
		}
	}
	if patient.Fullname != "" {
		if err := pr.patientRepository.UpdateName(patient.Fullname, patient.Id); err != nil {
			return models.PatientProfile{}, err
		}
	}
	if patient.Contactnumber != "" {
		if err := pr.patientRepository.UpdatePatientPhone(patient.Contactnumber, patient.Id); err != nil {
			return models.PatientProfile{}, err
		}
	}
	return pr.patientRepository.UserDetails(int(patient.Id))
}

func (pr *patientUseCase) ListPatients() ([]models.SignupdetailResponse, error) {
	patients, err := pr.patientRepository.ListPatients()
	if err != nil {
		return []models.SignupdetailResponse{}, err
	}
	return patients, nil
}
