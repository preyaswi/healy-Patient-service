// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: pkg/pb/patient.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PatientSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname        string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Confirmpassword string `protobuf:"bytes,4,opt,name=confirmpassword,proto3" json:"confirmpassword,omitempty"`
	Gender          string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Contactnumber   string `protobuf:"bytes,6,opt,name=contactnumber,proto3" json:"contactnumber,omitempty"`
}

func (x *PatientSignUpRequest) Reset() {
	*x = PatientSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientSignUpRequest) ProtoMessage() {}

func (x *PatientSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientSignUpRequest.ProtoReflect.Descriptor instead.
func (*PatientSignUpRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{0}
}

func (x *PatientSignUpRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *PatientSignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientSignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PatientSignUpRequest) GetConfirmpassword() string {
	if x != nil {
		return x.Confirmpassword
	}
	return ""
}

func (x *PatientSignUpRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PatientSignUpRequest) GetContactnumber() string {
	if x != nil {
		return x.Contactnumber
	}
	return ""
}

type PatientDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fullname      string `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Gender        string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Contactnumber string `protobuf:"bytes,5,opt,name=contactnumber,proto3" json:"contactnumber,omitempty"`
}

func (x *PatientDetails) Reset() {
	*x = PatientDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientDetails) ProtoMessage() {}

func (x *PatientDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientDetails.ProtoReflect.Descriptor instead.
func (*PatientDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{1}
}

func (x *PatientDetails) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatientDetails) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *PatientDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientDetails) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PatientDetails) GetContactnumber() string {
	if x != nil {
		return x.Contactnumber
	}
	return ""
}

type PatientSignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientDetails *PatientDetails `protobuf:"bytes,1,opt,name=PatientDetails,proto3" json:"PatientDetails,omitempty"`
	AccessToken    string          `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken   string          `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *PatientSignUpResponse) Reset() {
	*x = PatientSignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientSignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientSignUpResponse) ProtoMessage() {}

func (x *PatientSignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientSignUpResponse.ProtoReflect.Descriptor instead.
func (*PatientSignUpResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{2}
}

func (x *PatientSignUpResponse) GetPatientDetails() *PatientDetails {
	if x != nil {
		return x.PatientDetails
	}
	return nil
}

func (x *PatientSignUpResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *PatientSignUpResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type PatientLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PatientLoginRequest) Reset() {
	*x = PatientLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientLoginRequest) ProtoMessage() {}

func (x *PatientLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientLoginRequest.ProtoReflect.Descriptor instead.
func (*PatientLoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{3}
}

func (x *PatientLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PatientLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientDetails *PatientDetails `protobuf:"bytes,1,opt,name=PatientDetails,proto3" json:"PatientDetails,omitempty"`
	AccessToken    string          `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken   string          `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *PatientLoginResponse) Reset() {
	*x = PatientLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientLoginResponse) ProtoMessage() {}

func (x *PatientLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientLoginResponse.ProtoReflect.Descriptor instead.
func (*PatientLoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{4}
}

func (x *PatientLoginResponse) GetPatientDetails() *PatientDetails {
	if x != nil {
		return x.PatientDetails
	}
	return nil
}

func (x *PatientLoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *PatientLoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type Idreq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Idreq) Reset() {
	*x = Idreq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Idreq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Idreq) ProtoMessage() {}

func (x *Idreq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Idreq.ProtoReflect.Descriptor instead.
func (*Idreq) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{5}
}

func (x *Idreq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type InPatientDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname      string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email         string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Gender        string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Contactnumber string `protobuf:"bytes,4,opt,name=contactnumber,proto3" json:"contactnumber,omitempty"`
}

func (x *InPatientDetails) Reset() {
	*x = InPatientDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InPatientDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InPatientDetails) ProtoMessage() {}

func (x *InPatientDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InPatientDetails.ProtoReflect.Descriptor instead.
func (*InPatientDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{6}
}

func (x *InPatientDetails) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *InPatientDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *InPatientDetails) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *InPatientDetails) GetContactnumber() string {
	if x != nil {
		return x.Contactnumber
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId        uint64            `protobuf:"varint,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	InPatientDetails *InPatientDetails `protobuf:"bytes,2,opt,name=InPatientDetails,proto3" json:"InPatientDetails,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRequest) GetPatientId() uint64 {
	if x != nil {
		return x.PatientId
	}
	return 0
}

func (x *UpdateRequest) GetInPatientDetails() *InPatientDetails {
	if x != nil {
		return x.InPatientDetails
	}
	return nil
}

type UpdatePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId          uint64 `protobuf:"varint,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	OldPassword        string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword        string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	ConfirmNewPassword string `protobuf:"bytes,4,opt,name=confirm_new_password,json=confirmNewPassword,proto3" json:"confirm_new_password,omitempty"`
}

func (x *UpdatePasswordRequest) Reset() {
	*x = UpdatePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordRequest) ProtoMessage() {}

func (x *UpdatePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePasswordRequest) GetPatientId() uint64 {
	if x != nil {
		return x.PatientId
	}
	return 0
}

func (x *UpdatePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *UpdatePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *UpdatePasswordRequest) GetConfirmNewPassword() string {
	if x != nil {
		return x.ConfirmNewPassword
	}
	return ""
}

type UpdatePasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePasswordResponse) Reset() {
	*x = UpdatePasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordResponse) ProtoMessage() {}

func (x *UpdatePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordResponse.ProtoReflect.Descriptor instead.
func (*UpdatePasswordResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{9}
}

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{10}
}

type Listpares struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pali []*PatientDetails `protobuf:"bytes,1,rep,name=pali,proto3" json:"pali,omitempty"`
}

func (x *Listpares) Reset() {
	*x = Listpares{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listpares) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listpares) ProtoMessage() {}

func (x *Listpares) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listpares.ProtoReflect.Descriptor instead.
func (*Listpares) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_proto_rawDescGZIP(), []int{11}
}

func (x *Listpares) GetPali() []*PatientDetails {
	if x != nil {
		return x.Pali
	}
	return nil
}

var File_pkg_pb_patient_proto protoreflect.FileDescriptor

var file_pkg_pb_patient_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0xcc, 0x01, 0x0a, 0x14, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x55,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x90,
	0x01, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x9e, 0x01, 0x0a, 0x15, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x47, 0x0a, 0x13, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x14,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20, 0x0a, 0x05, 0x69,
	0x64, 0x72, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x82, 0x01,
	0x0a, 0x10, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x75, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x45, 0x0a, 0x10, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x10, 0x49, 0x6e, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x05, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x22, 0x38, 0x0a, 0x09, 0x6c,
	0x69, 0x73, 0x74, 0x70, 0x61, 0x72, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x6c, 0x69,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x04, 0x70, 0x61, 0x6c, 0x69, 0x32, 0xc0, 0x03, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x50, 0x0a, 0x0d, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11, 0x49, 0x6e, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x69, 0x64, 0x72, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12,
	0x53, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x70, 0x61, 0x72, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_patient_proto_rawDescOnce sync.Once
	file_pkg_pb_patient_proto_rawDescData = file_pkg_pb_patient_proto_rawDesc
)

func file_pkg_pb_patient_proto_rawDescGZIP() []byte {
	file_pkg_pb_patient_proto_rawDescOnce.Do(func() {
		file_pkg_pb_patient_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_patient_proto_rawDescData)
	})
	return file_pkg_pb_patient_proto_rawDescData
}

var file_pkg_pb_patient_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pkg_pb_patient_proto_goTypes = []interface{}{
	(*PatientSignUpRequest)(nil),   // 0: patient.PatientSignUpRequest
	(*PatientDetails)(nil),         // 1: patient.PatientDetails
	(*PatientSignUpResponse)(nil),  // 2: patient.PatientSignUpResponse
	(*PatientLoginRequest)(nil),    // 3: patient.PatientLoginRequest
	(*PatientLoginResponse)(nil),   // 4: patient.PatientLoginResponse
	(*Idreq)(nil),                  // 5: patient.idreq
	(*InPatientDetails)(nil),       // 6: patient.InPatientDetails
	(*UpdateRequest)(nil),          // 7: patient.UpdateRequest
	(*UpdatePasswordRequest)(nil),  // 8: patient.UpdatePasswordRequest
	(*UpdatePasswordResponse)(nil), // 9: patient.UpdatePasswordResponse
	(*Req)(nil),                    // 10: patient.req
	(*Listpares)(nil),              // 11: patient.listpares
}
var file_pkg_pb_patient_proto_depIdxs = []int32{
	1,  // 0: patient.PatientSignUpResponse.PatientDetails:type_name -> patient.PatientDetails
	1,  // 1: patient.PatientLoginResponse.PatientDetails:type_name -> patient.PatientDetails
	6,  // 2: patient.UpdateRequest.InPatientDetails:type_name -> patient.InPatientDetails
	1,  // 3: patient.listpares.pali:type_name -> patient.PatientDetails
	0,  // 4: patient.Patient.PatientSignUp:input_type -> patient.PatientSignUpRequest
	3,  // 5: patient.Patient.PatientLogin:input_type -> patient.PatientLoginRequest
	5,  // 6: patient.Patient.IndPatientDetails:input_type -> patient.idreq
	7,  // 7: patient.Patient.UpdatePatientDetails:input_type -> patient.UpdateRequest
	8,  // 8: patient.Patient.UpdatePassword:input_type -> patient.UpdatePasswordRequest
	10, // 9: patient.Patient.ListPatients:input_type -> patient.req
	2,  // 10: patient.Patient.PatientSignUp:output_type -> patient.PatientSignUpResponse
	4,  // 11: patient.Patient.PatientLogin:output_type -> patient.PatientLoginResponse
	1,  // 12: patient.Patient.IndPatientDetails:output_type -> patient.PatientDetails
	6,  // 13: patient.Patient.UpdatePatientDetails:output_type -> patient.InPatientDetails
	9,  // 14: patient.Patient.UpdatePassword:output_type -> patient.UpdatePasswordResponse
	11, // 15: patient.Patient.ListPatients:output_type -> patient.listpares
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_patient_proto_init() }
func file_pkg_pb_patient_proto_init() {
	if File_pkg_pb_patient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_patient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientSignUpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientSignUpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientLoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Idreq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InPatientDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listpares); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pb_patient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_patient_proto_goTypes,
		DependencyIndexes: file_pkg_pb_patient_proto_depIdxs,
		MessageInfos:      file_pkg_pb_patient_proto_msgTypes,
	}.Build()
	File_pkg_pb_patient_proto = out.File
	file_pkg_pb_patient_proto_rawDesc = nil
	file_pkg_pb_patient_proto_goTypes = nil
	file_pkg_pb_patient_proto_depIdxs = nil
}
