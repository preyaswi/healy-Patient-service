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
    Id            uint  `json:"id"`
    Fullname      string `json:"fullname"`
    Email         string `json:"email"`
    Gender        string `json:"gender"`
    Contactnumber string `json:"contactnumber"`
}

type TokenPatient struct {
	Patient         SignupdetailResponse
	AccessToken  string
	RefreshToken string
}

