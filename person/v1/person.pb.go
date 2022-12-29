// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: person/v1/person.proto

package person

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId        string        `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	AccountId  string        `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Profile    *Profile      `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	Credential []*Credential `protobuf:"bytes,4,rep,name=credential,proto3" json:"credential,omitempty"`
	StaffId    bool          `protobuf:"varint,5,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_v1_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_v1_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_v1_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Person) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Person) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *Person) GetCredential() []*Credential {
	if x != nil {
		return x.Credential
	}
	return nil
}

func (x *Person) GetStaffId() bool {
	if x != nil {
		return x.StaffId
	}
	return false
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastName         string   `protobuf:"bytes,1,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName        string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	PhoneNumber      string   `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	NationalIdentity string   `protobuf:"bytes,4,opt,name=national_identity,json=nationalIdentity,proto3" json:"national_identity,omitempty"`
	Address          *Address `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_v1_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_person_v1_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_person_v1_person_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Profile) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Profile) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Profile) GetNationalIdentity() string {
	if x != nil {
		return x.NationalIdentity
	}
	return ""
}

func (x *Profile) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_v1_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_person_v1_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_person_v1_person_proto_rawDescGZIP(), []int{2}
}

func (x *Credential) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Credential) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Credential) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street  string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	Town    string `protobuf:"bytes,2,opt,name=town,proto3" json:"town,omitempty"`
	Lga     string `protobuf:"bytes,3,opt,name=lga,proto3" json:"lga,omitempty"`
	State   string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Country string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_v1_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_person_v1_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_person_v1_person_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

func (x *Address) GetLga() string {
	if x != nil {
		return x.Lga
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

var File_person_v1_person_proto protoreflect.FileDescriptor

var file_person_v1_person_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0xb8, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0f,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x22, 0xc3,
	0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x5a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x77, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x67, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x67, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x66, 0x65, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x79, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_v1_person_proto_rawDescOnce sync.Once
	file_person_v1_person_proto_rawDescData = file_person_v1_person_proto_rawDesc
)

func file_person_v1_person_proto_rawDescGZIP() []byte {
	file_person_v1_person_proto_rawDescOnce.Do(func() {
		file_person_v1_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_v1_person_proto_rawDescData)
	})
	return file_person_v1_person_proto_rawDescData
}

var file_person_v1_person_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_person_v1_person_proto_goTypes = []interface{}{
	(*Person)(nil),     // 0: person.v1.Person
	(*Profile)(nil),    // 1: person.v1.Profile
	(*Credential)(nil), // 2: person.v1.Credential
	(*Address)(nil),    // 3: person.v1.Address
}
var file_person_v1_person_proto_depIdxs = []int32{
	1, // 0: person.v1.Person.profile:type_name -> person.v1.Profile
	2, // 1: person.v1.Person.credential:type_name -> person.v1.Credential
	3, // 2: person.v1.Profile.address:type_name -> person.v1.Address
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_person_v1_person_proto_init() }
func file_person_v1_person_proto_init() {
	if File_person_v1_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_v1_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_person_v1_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_person_v1_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
		file_person_v1_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_person_v1_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_v1_person_proto_goTypes,
		DependencyIndexes: file_person_v1_person_proto_depIdxs,
		MessageInfos:      file_person_v1_person_proto_msgTypes,
	}.Build()
	File_person_v1_person_proto = out.File
	file_person_v1_person_proto_rawDesc = nil
	file_person_v1_person_proto_goTypes = nil
	file_person_v1_person_proto_depIdxs = nil
}