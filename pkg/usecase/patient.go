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

func NewPatientUseCase(repository interfaces.PatientRepository) usecaseint.PatientUseCase{
	return &patientUseCase{
		patientRepository: repository,
	}
}

func (pr *patientUseCase) PatientsSignUp(patient models.PatientSignUp) (models.TokenPatient, error) {
	email, err := pr.patientRepository.CheckPatientExistsByEmail(patient.Email)
	if err != nil {
		return models.TokenPatient{}, errors.New("error with server")
	}
	if email != nil {
		return models.TokenPatient{}, errors.New("user with this email is already exists")
	}

	phone, err := pr.patientRepository.CheckPatientExistsByPhone(patient.Contactnumber)
	if err != nil {
		return models.TokenPatient{}, errors.New("error with server")
	}
	if phone != nil {
		return models.TokenPatient{}, errors.New("user with this phone is already exists")
	}
	if patient.Password!=patient.Confirmpassword{
		return models.TokenPatient{},errors.New("confirm password is not matching")
	}

	hashPassword, err := helper.PasswordHash(patient.Password)
	if err != nil {
		return models.TokenPatient{}, errors.New("error in hashing password")
	}
	patient.Password = hashPassword
	PatientData, err := pr.patientRepository.PatientSignUp(patient)
	if err != nil {
		return models.TokenPatient{}, errors.New("could not add the user")
	}
	accessToken, err := helper.GenerateAccessToken(PatientData)
	if err != nil {
		return models.TokenPatient{}, errors.New("couldn't create access token due to error")
	}
	refreshToken, err := helper.GenerateRefreshToken(PatientData)
	if err != nil {
		return models.TokenPatient{}, errors.New("couldn't create refresh token due to error")
	}
	return models.TokenPatient{
		Patient:         PatientData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
