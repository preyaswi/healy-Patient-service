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
func (p *PatientServer)GoogleSignIn(ctx context.Context,req *pb.GoogleSignInRequest) (*pb.PatientSignUpResponse, error)  {

	res,err:=p.patientUseCase.GoogleSignIn(
		req.GoogleId,req.Name,req.Email,
	)
	if err!=nil{
		return &pb.PatientSignUpResponse{},err
	}
	return &pb.PatientSignUpResponse{
		PatientDetails: &pb.GoogleSignInResponse{
			Id: uint32(res.Patient.Id),
			GoogleId: res.Patient.GoogleId,
			Fullname: res.Patient.FullName,
			Email: res.Patient.Email,
		},
		AccessToken: res.AccessToken,
		RefreshToken: res.RefreshToken,
	},nil
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
