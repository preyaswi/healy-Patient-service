package models

type PatientSignUp struct {
	Fullname        string
	Email           string
	Password        string
	Confirmpassword string
	Gender          string
	Contactnumber   string
}
type SignupdetailResponse struct {
	Id            uint   `json:"id"`
	Fullname      string `json:"fullname"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	Contactnumber string `json:"contactnumber"`
}

type TokenPatient struct {
	Patient      SignupdetailResponse
	AccessToken  string
	RefreshToken string
}

type PatientLogin struct {
	Email    string
	Password string
}
type PatientDetails struct {
	Id            uint
	Fullname      string
	Email         string
	Password  string
	Gender        string
	Contactnumber string
}
type PatientProfile struct{
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
type UpdatePassword struct{
	OldPassword        string `json:"old_password" binding:"required"`
    NewPassword        string `json:"new_password" binding:"required"`
    ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}