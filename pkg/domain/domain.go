package domain

type Patient struct {
	Id            uint   `json:"id" gorm:"uniquekey; not null"`
	Fullname      string `json:"fullname" gorm:"validate:required"`
	Email         string `json:"email" gorm:"validate:required"`
	Password      string `json:"password" gorm:"validate:required"`
	Gender        string `json:"gender" gorm:"validate:required"`
	Contactnumber string `json:"contactnumber" gorm:"validate:required"`
}
func (Patient) TableName() string {
    return "patient"
}