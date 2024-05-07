package service

import (
	"context"
	"patient-service/pkg/models"
	"patient-service/pkg/pb"
	usecaseint "patient-service/pkg/usecase/interface"
)

type PatientServer struct {
	patientUseCase usecaseint.PatientUseCase
	pb.UnimplementedPatientServer
}

func NewPatientServer(useCase usecaseint.PatientUseCase) pb.PatientServer {
	return &PatientServer{
		patientUseCase: useCase,
	}

}
func (p *PatientServer) PatientSignUp(ctx context.Context,patientsighupRequest *pb.PatientSignUpRequest) (*pb.PatientSignUpResponse, error) {
	patientCreateDetails := models.PatientSignUp{
		Fullname: patientsighupRequest.Fullname,
		Email:     patientsighupRequest.Email,
		Password:  patientsighupRequest.Password,
		Confirmpassword: patientsighupRequest.Confirmpassword,
		Gender: patientsighupRequest.Gender,
		Contactnumber: patientsighupRequest.Contactnumber,
	}
	data, err := p.patientUseCase.PatientsSignUp(patientCreateDetails)
	if err != nil {
		return &pb.PatientSignUpResponse{}, err
	}
	userDetails := &pb.PatientDetails{Id:uint64(data.Patient.Id), Fullname: data.Patient.Fullname, Email: data.Patient.Email, Gender: data.Patient.Gender,Contactnumber: data.Patient.Contactnumber}
	return &pb.PatientSignUpResponse{
		PatientDetails:  userDetails,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}, nil
}
