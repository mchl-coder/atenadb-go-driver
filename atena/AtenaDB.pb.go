// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: AtenaDB.proto

package atenadb

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

type AuthLookupModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Db       string `protobuf:"bytes,3,opt,name=db,proto3" json:"db,omitempty"`
}

func (x *AuthLookupModel) Reset() {
	*x = AuthLookupModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLookupModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLookupModel) ProtoMessage() {}

func (x *AuthLookupModel) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLookupModel.ProtoReflect.Descriptor instead.
func (*AuthLookupModel) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{0}
}

func (x *AuthLookupModel) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuthLookupModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthLookupModel) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

type AuthUserLookupModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthUserLookupModel) Reset() {
	*x = AuthUserLookupModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserLookupModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserLookupModel) ProtoMessage() {}

func (x *AuthUserLookupModel) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserLookupModel.ProtoReflect.Descriptor instead.
func (*AuthUserLookupModel) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{1}
}

func (x *AuthUserLookupModel) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuthUserLookupModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserModel) Reset() {
	*x = CreateUserModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserModel) ProtoMessage() {}

func (x *CreateUserModel) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserModel.ProtoReflect.Descriptor instead.
func (*CreateUserModel) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserModel) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateUserModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type NewPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	NewPsw string `protobuf:"bytes,2,opt,name=NewPsw,proto3" json:"NewPsw,omitempty"`
}

func (x *NewPassword) Reset() {
	*x = NewPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPassword) ProtoMessage() {}

func (x *NewPassword) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPassword.ProtoReflect.Descriptor instead.
func (*NewPassword) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{3}
}

func (x *NewPassword) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *NewPassword) GetNewPsw() string {
	if x != nil {
		return x.NewPsw
	}
	return ""
}

type DBModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LFU   bool   `protobuf:"varint,3,opt,name=LFU,proto3" json:"LFU,omitempty"`
	Save  bool   `protobuf:"varint,4,opt,name=save,proto3" json:"save,omitempty"`
}

func (x *DBModel) Reset() {
	*x = DBModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBModel) ProtoMessage() {}

func (x *DBModel) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBModel.ProtoReflect.Descriptor instead.
func (*DBModel) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{4}
}

func (x *DBModel) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DBModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DBModel) GetLFU() bool {
	if x != nil {
		return x.LFU
	}
	return false
}

func (x *DBModel) GetSave() bool {
	if x != nil {
		return x.Save
	}
	return false
}

type DBInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DBInfo) Reset() {
	*x = DBInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBInfo) ProtoMessage() {}

func (x *DBInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBInfo.ProtoReflect.Descriptor instead.
func (*DBInfo) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{5}
}

func (x *DBInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DBInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AtenaSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AtenaSet) Reset() {
	*x = AtenaSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtenaSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtenaSet) ProtoMessage() {}

func (x *AtenaSet) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtenaSet.ProtoReflect.Descriptor instead.
func (*AtenaSet) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{6}
}

func (x *AtenaSet) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AtenaSet) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AtenaSet) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AtenaGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *AtenaGet) Reset() {
	*x = AtenaGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtenaGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtenaGet) ProtoMessage() {}

func (x *AtenaGet) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtenaGet.ProtoReflect.Descriptor instead.
func (*AtenaGet) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{7}
}

func (x *AtenaGet) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AtenaGet) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type AtenaDel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *AtenaDel) Reset() {
	*x = AtenaDel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtenaDel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtenaDel) ProtoMessage() {}

func (x *AtenaDel) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtenaDel.ProtoReflect.Descriptor instead.
func (*AtenaDel) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{8}
}

func (x *AtenaDel) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AtenaDel) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type AtenaIncr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Inc   int32  `protobuf:"varint,3,opt,name=inc,proto3" json:"inc,omitempty"`
}

func (x *AtenaIncr) Reset() {
	*x = AtenaIncr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtenaIncr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtenaIncr) ProtoMessage() {}

func (x *AtenaIncr) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtenaIncr.ProtoReflect.Descriptor instead.
func (*AtenaIncr) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{9}
}

func (x *AtenaIncr) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AtenaIncr) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AtenaIncr) GetInc() int32 {
	if x != nil {
		return x.Inc
	}
	return 0
}

type RemoveAllRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RemoveAllRecords) Reset() {
	*x = RemoveAllRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAllRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAllRecords) ProtoMessage() {}

func (x *RemoveAllRecords) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAllRecords.ProtoReflect.Descriptor instead.
func (*RemoveAllRecords) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveAllRecords) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogoutUser) Reset() {
	*x = LogoutUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutUser) ProtoMessage() {}

func (x *LogoutUser) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutUser.ProtoReflect.Descriptor instead.
func (*LogoutUser) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{11}
}

func (x *LogoutUser) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AtenaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AtenaResponse) Reset() {
	*x = AtenaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtenaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtenaResponse) ProtoMessage() {}

func (x *AtenaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtenaResponse.ProtoReflect.Descriptor instead.
func (*AtenaResponse) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{12}
}

func (x *AtenaResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AtenaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
}

func (x *AtenaReply) Reset() {
	*x = AtenaReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtenaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtenaReply) ProtoMessage() {}

func (x *AtenaReply) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtenaReply.ProtoReflect.Descriptor instead.
func (*AtenaReply) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{13}
}

func (x *AtenaReply) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

type AtenaAuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool   `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AtenaAuthReply) Reset() {
	*x = AtenaAuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AtenaDB_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtenaAuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtenaAuthReply) ProtoMessage() {}

func (x *AtenaAuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_AtenaDB_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtenaAuthReply.ProtoReflect.Descriptor instead.
func (*AtenaAuthReply) Descriptor() ([]byte, []int) {
	return file_AtenaDB_proto_rawDescGZIP(), []int{14}
}

func (x *AtenaAuthReply) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

func (x *AtenaAuthReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_AtenaDB_proto protoreflect.FileDescriptor

var file_AtenaDB_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x44, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x51, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x64, 0x62, 0x22, 0x45, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x57, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x3b, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x50, 0x73,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x65, 0x77, 0x50, 0x73, 0x77, 0x22,
	0x59, 0x0a, 0x07, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x46, 0x55, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x4c, 0x46, 0x55, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x76, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x61, 0x76, 0x65, 0x22, 0x32, 0x0a, 0x06, 0x44, 0x42,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48,
	0x0a, 0x08, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x36, 0x0a, 0x08, 0x41, 0x74, 0x65, 0x6e,
	0x61, 0x47, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x22, 0x32, 0x0a, 0x08, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x44, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x45, 0x0a, 0x09, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x49, 0x6e, 0x63,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x6e, 0x63, 0x22, 0x28, 0x0a, 0x10, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x25, 0x0a, 0x0d, 0x41, 0x74, 0x65,
	0x6e, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x2c, 0x0a, 0x0a, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x22, 0x46,
	0x0a, 0x0e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc0, 0x04, 0x0a, 0x07, 0x41, 0x74, 0x65, 0x6e, 0x61,
	0x44, 0x42, 0x12, 0x29, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x10, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x0f, 0x2e, 0x41,
	0x74, 0x65, 0x6e, 0x61, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a,
	0x0f, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x0b, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a,
	0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x0c, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x0b, 0x2e,
	0x41, 0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x08, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x42, 0x12, 0x07, 0x2e, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x0b, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x42, 0x48, 0x54, 0x12, 0x08, 0x2e, 0x44, 0x42, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x0b, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x24, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x42, 0x52, 0x42, 0x54,
	0x12, 0x08, 0x2e, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x0b, 0x2e, 0x41, 0x74, 0x65,
	0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x42, 0x12, 0x07, 0x2e, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0b, 0x2e, 0x41,
	0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x09, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x09, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x53, 0x65,
	0x74, 0x1a, 0x0b, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x09, 0x2e, 0x41, 0x74,
	0x65, 0x6e, 0x61, 0x47, 0x65, 0x74, 0x1a, 0x0e, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x09, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x44, 0x65,
	0x6c, 0x1a, 0x0b, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28,
	0x0a, 0x0a, 0x49, 0x6e, 0x63, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0a, 0x2e, 0x41,
	0x74, 0x65, 0x6e, 0x61, 0x49, 0x6e, 0x63, 0x72, 0x1a, 0x0e, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x0b, 0x2e, 0x41, 0x74, 0x65, 0x6e, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12,
	0x0b, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x41,
	0x74, 0x65, 0x6e, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x74, 0x65,
	0x6e, 0x61, 0x2f, 0x61, 0x74, 0x65, 0x6e, 0x61, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_AtenaDB_proto_rawDescOnce sync.Once
	file_AtenaDB_proto_rawDescData = file_AtenaDB_proto_rawDesc
)

func file_AtenaDB_proto_rawDescGZIP() []byte {
	file_AtenaDB_proto_rawDescOnce.Do(func() {
		file_AtenaDB_proto_rawDescData = protoimpl.X.CompressGZIP(file_AtenaDB_proto_rawDescData)
	})
	return file_AtenaDB_proto_rawDescData
}

var file_AtenaDB_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_AtenaDB_proto_goTypes = []interface{}{
	(*AuthLookupModel)(nil),     // 0: AuthLookupModel
	(*AuthUserLookupModel)(nil), // 1: AuthUserLookupModel
	(*CreateUserModel)(nil),     // 2: CreateUserModel
	(*NewPassword)(nil),         // 3: NewPassword
	(*DBModel)(nil),             // 4: DBModel
	(*DBInfo)(nil),              // 5: DBInfo
	(*AtenaSet)(nil),            // 6: AtenaSet
	(*AtenaGet)(nil),            // 7: AtenaGet
	(*AtenaDel)(nil),            // 8: AtenaDel
	(*AtenaIncr)(nil),           // 9: AtenaIncr
	(*RemoveAllRecords)(nil),    // 10: RemoveAllRecords
	(*LogoutUser)(nil),          // 11: LogoutUser
	(*AtenaResponse)(nil),       // 12: AtenaResponse
	(*AtenaReply)(nil),          // 13: AtenaReply
	(*AtenaAuthReply)(nil),      // 14: AtenaAuthReply
}
var file_AtenaDB_proto_depIdxs = []int32{
	0,  // 0: AtenaDB.Auth:input_type -> AuthLookupModel
	1,  // 1: AtenaDB.AuthUser:input_type -> AuthUserLookupModel
	2,  // 2: AtenaDB.CreateUser:input_type -> CreateUserModel
	3,  // 3: AtenaDB.ChangePassword:input_type -> NewPassword
	5,  // 4: AtenaDB.CreateDB:input_type -> DBInfo
	4,  // 5: AtenaDB.CreateDBHT:input_type -> DBModel
	4,  // 6: AtenaDB.CreateDBRBT:input_type -> DBModel
	5,  // 7: AtenaDB.DeleteDB:input_type -> DBInfo
	6,  // 8: AtenaDB.SetRecord:input_type -> AtenaSet
	7,  // 9: AtenaDB.GetRecord:input_type -> AtenaGet
	8,  // 10: AtenaDB.DeleteRecord:input_type -> AtenaDel
	9,  // 11: AtenaDB.IncrRecord:input_type -> AtenaIncr
	10, // 12: AtenaDB.RemoveAll:input_type -> RemoveAllRecords
	11, // 13: AtenaDB.Logout:input_type -> LogoutUser
	14, // 14: AtenaDB.Auth:output_type -> AtenaAuthReply
	14, // 15: AtenaDB.AuthUser:output_type -> AtenaAuthReply
	13, // 16: AtenaDB.CreateUser:output_type -> AtenaReply
	13, // 17: AtenaDB.ChangePassword:output_type -> AtenaReply
	13, // 18: AtenaDB.CreateDB:output_type -> AtenaReply
	13, // 19: AtenaDB.CreateDBHT:output_type -> AtenaReply
	13, // 20: AtenaDB.CreateDBRBT:output_type -> AtenaReply
	13, // 21: AtenaDB.DeleteDB:output_type -> AtenaReply
	13, // 22: AtenaDB.SetRecord:output_type -> AtenaReply
	12, // 23: AtenaDB.GetRecord:output_type -> AtenaResponse
	13, // 24: AtenaDB.DeleteRecord:output_type -> AtenaReply
	12, // 25: AtenaDB.IncrRecord:output_type -> AtenaResponse
	13, // 26: AtenaDB.RemoveAll:output_type -> AtenaReply
	13, // 27: AtenaDB.Logout:output_type -> AtenaReply
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_AtenaDB_proto_init() }
func file_AtenaDB_proto_init() {
	if File_AtenaDB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AtenaDB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLookupModel); i {
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
		file_AtenaDB_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserLookupModel); i {
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
		file_AtenaDB_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserModel); i {
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
		file_AtenaDB_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPassword); i {
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
		file_AtenaDB_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBModel); i {
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
		file_AtenaDB_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBInfo); i {
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
		file_AtenaDB_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtenaSet); i {
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
		file_AtenaDB_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtenaGet); i {
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
		file_AtenaDB_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtenaDel); i {
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
		file_AtenaDB_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtenaIncr); i {
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
		file_AtenaDB_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAllRecords); i {
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
		file_AtenaDB_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutUser); i {
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
		file_AtenaDB_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtenaResponse); i {
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
		file_AtenaDB_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtenaReply); i {
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
		file_AtenaDB_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtenaAuthReply); i {
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
			RawDescriptor: file_AtenaDB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_AtenaDB_proto_goTypes,
		DependencyIndexes: file_AtenaDB_proto_depIdxs,
		MessageInfos:      file_AtenaDB_proto_msgTypes,
	}.Build()
	File_AtenaDB_proto = out.File
	file_AtenaDB_proto_rawDesc = nil
	file_AtenaDB_proto_goTypes = nil
	file_AtenaDB_proto_depIdxs = nil
}
