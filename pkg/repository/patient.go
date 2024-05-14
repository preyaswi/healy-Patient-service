package repository

import (
	"errors"
	"fmt"
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
	fmt.Println(phone, "contact number")
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
func (ur *patientRepository) FindPatientByEmail(email string) (models.PatientDetails, error) {
	var patientdetail models.PatientDetails
	err := ur.DB.Raw("SELECT * FROM patient WHERE email=?", email).Scan(&patientdetail).Error
	if err != nil {
		return models.PatientDetails{}, errors.New("error checking user details")
	}
	return patientdetail, nil
}
func (ur *patientRepository) IndPatientDetails(patient_id uint64) (models.SignupdetailResponse, error) {
	var patient models.SignupdetailResponse
	err := ur.DB.Raw("Select * from patient where id=?", patient_id).Scan(&patient).Error
	if err != nil {
		return models.SignupdetailResponse{}, err
	}
	return patient, nil
}
func (ur *patientRepository) CheckPatientAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from patient where email='%s'", email)
	if err := ur.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}
func (ur *patientRepository) UpdatePatientEmail(email string, PatientID uint) error {

	err := ur.DB.Exec("update patient set email = ? where id = ?", email, PatientID).Error
	if err != nil {
		return err
	}
	return nil

}

func (ur *patientRepository) UpdatePatientPhone(phone string, PatientID uint) error {

	err := ur.DB.Exec("update patient set contactnumber = ? where id = ?", phone, PatientID).Error
	if err != nil {
		return err
	}
	return nil

}

func (ur *patientRepository) UpdateName(name string, PatientID uint) error {

	err := ur.DB.Exec("update patient set fullname = ? where id = ?", name, PatientID).Error
	if err != nil {
		return err
	}
	return nil

}
func (ur *patientRepository) UserDetails(userID int) (models.PatientProfile, error) {

	var userDetails models.PatientProfile
	err := ur.DB.Raw("select patient.fullname,patient.email,patient.gender,patient.contactnumber from patient  where patient.id = ?", userID).Row().Scan(&userDetails.Fullname, &userDetails.Email, &userDetails.Gender, &userDetails.Contactnumber)
	if err != nil {
		return models.PatientProfile{}, errors.New("could not get user details")
	}
	return userDetails, nil
}
func (ur *patientRepository) PatientPassword(userID int) (string, error) {

	var userPassword string
	err := ur.DB.Raw("select password from patient where id = ?", userID).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}
func (ur *patientRepository) UpdatePatientPassword(password string, userID int) error {
	err := ur.DB.Exec("update patient set password = ? where id = ?", password, userID).Error
	if err != nil {
		return err
	}
	fmt.Println("password Updated succesfully")
	return nil
}
func (ur *patientRepository) ListPatients() ([]models.SignupdetailResponse, error) {
	row, err := ur.DB.Raw("select id,fullname,email,gender,contactnumber from patient").Rows()
	if err != nil {
		return []models.SignupdetailResponse{}, err
	}
	defer row.Close()
	var patientDetails []models.SignupdetailResponse
	for row.Next() {
		var patientdetail models.SignupdetailResponse

		// Scan the row into variables
		if err := row.Scan(&patientdetail.Id, &patientdetail.Fullname, &patientdetail.Email,&patientdetail.Gender, &patientdetail.Contactnumber); err != nil {
			return nil, err
		}

		patientDetails = append(patientDetails, patientdetail)
	}
	return patientDetails, nil
}
