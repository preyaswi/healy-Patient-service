package repository

import (
	"errors"
	"patient-service/pkg/domain"
	"patient-service/pkg/models"
	interfaces "patient-service/pkg/repository/interface"

	"gorm.io/gorm"
)

type patientRepository struct {
	DB *gorm.DB
}

func NewPatientRepository(DB *gorm.DB) interfaces.PatientRepository {
	return &patientRepository{
		DB: DB,
	}
}
func (ur *patientRepository) CheckPatientExistsByEmail(email string) (*domain.Patient, error) {
	var patient domain.Patient
	res := ur.DB.Where(&domain.Patient{Email: email}).First(&patient)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Patient{}, res.Error
	}
	return &patient, nil
}
func (ur *patientRepository) CheckPatientExistsByPhone(phone string) (*domain.Patient, error) {
	var patient domain.Patient
	res := ur.DB.Where(&domain.Patient{Contactnumber: phone}).First(&patient)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Patient{}, res.Error
	}
	return &patient, nil
}
func (ur *patientRepository) PatientSignUp(patient models.PatientSignUp) (models.SignupdetailResponse, error) {
	var signupDetail models.SignupdetailResponse
	err := ur.DB.Raw(`
		INSERT INTO Patient(fullname,email,password,gender,contactnumber)
		VALUES(?, ?, ?, ?, ?)
		RETURNING id,fullname,email,gender,contactnumber
	`, patient.Fullname, patient.Email, patient.Password, patient.Gender, patient.Contactnumber).
		Scan(&signupDetail).Error

	if err != nil {
		return models.SignupdetailResponse{}, err
	}
	return signupDetail, nil
}
func (ur *patientRepository)FindPatientByEmail(email string)(models.PatientDetails,error)  {
	var patientdetail models.PatientDetails
	err := ur.DB.Raw("SELECT * FROM patient WHERE email=?", email).Scan(&patientdetail).Error
	if err != nil {
		return models.PatientDetails{}, errors.New("error checking user details")
	}
	return patientdetail, nil
}