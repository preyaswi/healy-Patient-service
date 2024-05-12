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
func (p *PatientServer) PatientSignUp(ctx context.Context, patientsighupRequest *pb.PatientSignUpRequest) (*pb.PatientSignUpResponse, error) {
	patientCreateDetails := models.PatientSignUp{
		Fullname:        patientsighupRequest.Fullname,
		Email:           patientsighupRequest.Email,
		Password:        patientsighupRequest.Password,
		Confirmpassword: patientsighupRequest.Confirmpassword,
		Gender:          patientsighupRequest.Gender,
		Contactnumber:   patientsighupRequest.Contactnumber,
	}
	data, err := p.patientUseCase.PatientsSignUp(patientCreateDetails)
	if err != nil {
		return &pb.PatientSignUpResponse{}, err
	}
	userDetails := &pb.PatientDetails{
		Id:            uint64(data.Patient.Id),
		Fullname:      data.Patient.Fullname,
		Email:         data.Patient.Email,
		Gender:        data.Patient.Gender,
		Contactnumber: data.Patient.Contactnumber}
	return &pb.PatientSignUpResponse{
		PatientDetails: userDetails,
		AccessToken:    data.AccessToken,
		RefreshToken:   data.RefreshToken,
	},nil
}
func (p *PatientServer) PatientLogin(ctx context.Context, PatientLoginRequest *pb.PatientLoginRequest) (*pb.PatientLoginResponse, error) {
	patient := models.PatientLogin{
		Email:    PatientLoginRequest.Email,
		Password: PatientLoginRequest.Password,
	}
	data, err := p.patientUseCase.PatientLogin(patient)

	if err != nil {
		return &pb.PatientLoginResponse{}, err
	}
	patientdetail := &pb.PatientDetails{
		Id:            uint64(data.Patient.Id),
		Fullname:      data.Patient.Fullname,
		Email:         patient.Email,
		Gender:        data.Patient.Gender,
		Contactnumber: data.Patient.Contactnumber,
	}
	return &pb.PatientLoginResponse{
		PatientDetails: patientdetail,
		AccessToken:    data.AccessToken,
		RefreshToken:   data.RefreshToken,
	}, nil

}
func (p *PatientServer)IndPatientDetails(ctx context.Context,req *pb.Idreq) (*pb.PatientDetails, error)  {
	doctor,err:=p.patientUseCase.IndPatientDetails(req.UserId)
	if err!=nil{
		return &pb.PatientDetails{},err
	}
	return &pb.PatientDetails{
		Id: uint64(doctor.Id),
		Fullname: doctor.Fullname,
		Email: doctor.Email,
		Gender: doctor.Gender,
		Contactnumber: doctor.Contactnumber,
	},nil
}