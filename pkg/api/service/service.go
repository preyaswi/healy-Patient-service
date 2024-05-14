package service

import (
	"context"
	"fmt"
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
func (p *PatientServer) UpdatePatientDetails(ctx context.Context,req *pb.UpdateRequest) (*pb.InPatientDetails, error){
	patient:=models.SignupdetailResponse{
		Id: uint(req.PatientId),
		Fullname: req.InPatientDetails.Fullname,
		Email: req.InPatientDetails.Email,
		Gender: req.InPatientDetails.Gender,
		Contactnumber: req.InPatientDetails.Contactnumber,
	}
	res,err:=p.patientUseCase.UpdatePatientDetails(patient)
	if err!=nil{
		return &pb.InPatientDetails{},err
	}
	return &pb.InPatientDetails{
		Fullname: res.Fullname,
		Email: res.Email,
		Gender: res.Gender,
		Contactnumber: res.Contactnumber,
	},nil
}
func (p *PatientServer) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	patient_id:=req.PatientId
	body:=models.UpdatePassword{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
		ConfirmNewPassword: req.ConfirmNewPassword,
	}
	fmt.Println(body,"body")
	if err:=p.patientUseCase.UpdatePassword(ctx,patient_id, body);err!=nil{
		return &pb.UpdatePasswordResponse{},err
	}
	return&pb.UpdatePasswordResponse{},nil
}
func (p *PatientServer)ListPatients(ctx context.Context,req *pb.Req) (*pb.Listpares, error)  {
	listed,err:=p.patientUseCase.ListPatients()
	if err!=nil{
		return &pb.Listpares{},err
	}
	patientlist:=make([]*pb.PatientDetails,len(listed))
	for i,patient:=range listed{
		patientlist[i]=&pb.PatientDetails{
			Id: uint64(patient.Id),
			Fullname: patient.Fullname,
			Email: patient.Email,
			Gender: patient.Gender,
			Contactnumber: patient.Contactnumber,
		}
	}
	return &pb.Listpares{
		Pali: patientlist,
	},nil

}