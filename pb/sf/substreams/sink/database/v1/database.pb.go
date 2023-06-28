// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sf/substreams/sink/database/v1/database.proto

package pbdatabase

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

type TableChange_Operation int32

const (
	TableChange_OPERATION_UNSPECIFIED TableChange_Operation = 0 // Protobuf default should not be used, this is used so that the consume can ensure that the value was actually specified
	TableChange_OPERATION_CREATE      TableChange_Operation = 1
	TableChange_OPERATION_UPDATE      TableChange_Operation = 2
	TableChange_OPERATION_DELETE      TableChange_Operation = 3
)

// Enum value maps for TableChange_Operation.
var (
	TableChange_Operation_name = map[int32]string{
		0: "OPERATION_UNSPECIFIED",
		1: "OPERATION_CREATE",
		2: "OPERATION_UPDATE",
		3: "OPERATION_DELETE",
	}
	TableChange_Operation_value = map[string]int32{
		"OPERATION_UNSPECIFIED": 0,
		"OPERATION_CREATE":      1,
		"OPERATION_UPDATE":      2,
		"OPERATION_DELETE":      3,
	}
)

func (x TableChange_Operation) Enum() *TableChange_Operation {
	p := new(TableChange_Operation)
	*p = x
	return p
}

func (x TableChange_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TableChange_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_substreams_sink_database_v1_database_proto_enumTypes[0].Descriptor()
}

func (TableChange_Operation) Type() protoreflect.EnumType {
	return &file_sf_substreams_sink_database_v1_database_proto_enumTypes[0]
}

func (x TableChange_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TableChange_Operation.Descriptor instead.
func (TableChange_Operation) EnumDescriptor() ([]byte, []int) {
	return file_sf_substreams_sink_database_v1_database_proto_rawDescGZIP(), []int{1, 0}
}

type DatabaseChanges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableChanges []*TableChange `protobuf:"bytes,1,rep,name=table_changes,json=tableChanges,proto3" json:"table_changes,omitempty"`
}

func (x *DatabaseChanges) Reset() {
	*x = DatabaseChanges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_sink_database_v1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseChanges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseChanges) ProtoMessage() {}

func (x *DatabaseChanges) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_sink_database_v1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseChanges.ProtoReflect.Descriptor instead.
func (*DatabaseChanges) Descriptor() ([]byte, []int) {
	return file_sf_substreams_sink_database_v1_database_proto_rawDescGZIP(), []int{0}
}

func (x *DatabaseChanges) GetTableChanges() []*TableChange {
	if x != nil {
		return x.TableChanges
	}
	return nil
}

type TableChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	// Types that are assignable to PrimaryKey:
	//	*TableChange_Pk
	//	*TableChange_CompositePk
	PrimaryKey isTableChange_PrimaryKey `protobuf_oneof:"primary_key"`
	Ordinal    uint64                   `protobuf:"varint,3,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	Operation  TableChange_Operation    `protobuf:"varint,4,opt,name=operation,proto3,enum=sf.substreams.sink.database.v1.TableChange_Operation" json:"operation,omitempty"`
	Fields     []*Field                 `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *TableChange) Reset() {
	*x = TableChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_sink_database_v1_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableChange) ProtoMessage() {}

func (x *TableChange) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_sink_database_v1_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableChange.ProtoReflect.Descriptor instead.
func (*TableChange) Descriptor() ([]byte, []int) {
	return file_sf_substreams_sink_database_v1_database_proto_rawDescGZIP(), []int{1}
}

func (x *TableChange) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (m *TableChange) GetPrimaryKey() isTableChange_PrimaryKey {
	if m != nil {
		return m.PrimaryKey
	}
	return nil
}

func (x *TableChange) GetPk() string {
	if x, ok := x.GetPrimaryKey().(*TableChange_Pk); ok {
		return x.Pk
	}
	return ""
}

func (x *TableChange) GetCompositePk() *CompositePrimaryKey {
	if x, ok := x.GetPrimaryKey().(*TableChange_CompositePk); ok {
		return x.CompositePk
	}
	return nil
}

func (x *TableChange) GetOrdinal() uint64 {
	if x != nil {
		return x.Ordinal
	}
	return 0
}

func (x *TableChange) GetOperation() TableChange_Operation {
	if x != nil {
		return x.Operation
	}
	return TableChange_OPERATION_UNSPECIFIED
}

func (x *TableChange) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type isTableChange_PrimaryKey interface {
	isTableChange_PrimaryKey()
}

type TableChange_Pk struct {
	Pk string `protobuf:"bytes,2,opt,name=pk,proto3,oneof"`
}

type TableChange_CompositePk struct {
	CompositePk *CompositePrimaryKey `protobuf:"bytes,6,opt,name=composite_pk,json=compositePk,proto3,oneof"`
}

func (*TableChange_Pk) isTableChange_PrimaryKey() {}

func (*TableChange_CompositePk) isTableChange_PrimaryKey() {}

type CompositePrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys map[string]string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CompositePrimaryKey) Reset() {
	*x = CompositePrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_sink_database_v1_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositePrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositePrimaryKey) ProtoMessage() {}

func (x *CompositePrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_sink_database_v1_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompositePrimaryKey.ProtoReflect.Descriptor instead.
func (*CompositePrimaryKey) Descriptor() ([]byte, []int) {
	return file_sf_substreams_sink_database_v1_database_proto_rawDescGZIP(), []int{2}
}

func (x *CompositePrimaryKey) GetKeys() map[string]string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NewValue string `protobuf:"bytes,2,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	OldValue string `protobuf:"bytes,3,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_sink_database_v1_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_sink_database_v1_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_sf_substreams_sink_database_v1_database_proto_rawDescGZIP(), []int{3}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

func (x *Field) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

var File_sf_substreams_sink_database_v1_database_proto protoreflect.FileDescriptor

var file_sf_substreams_sink_database_v1_database_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x73, 0x69, 0x6e, 0x6b, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73,
	0x69, 0x6e, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0x63, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x12, 0x50, 0x0a, 0x0d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x66, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x22, 0xb6, 0x03, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x02, 0x70, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x58, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x70, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x65, 0x50, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c,
	0x12, 0x53, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x22, 0x68, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x42, 0x0d,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0xa1, 0x01,
	0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x51, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x55, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xb8, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73,
	0x69, 0x6e, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x66, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2d, 0x73, 0x69, 0x6e, 0x6b, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x73,
	0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6e,
	0x6b, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0xa2, 0x02, 0x04, 0x53, 0x53, 0x53, 0x44, 0xaa,
	0x02, 0x1e, 0x53, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x53, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x1e, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x5c, 0x53, 0x69, 0x6e, 0x6b, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x2a, 0x53, 0x66, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x5c, 0x53, 0x69, 0x6e, 0x6b, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x22, 0x53, 0x66, 0x3a, 0x3a, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x3a,
	0x3a, 0x53, 0x69, 0x6e, 0x6b, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_substreams_sink_database_v1_database_proto_rawDescOnce sync.Once
	file_sf_substreams_sink_database_v1_database_proto_rawDescData = file_sf_substreams_sink_database_v1_database_proto_rawDesc
)

func file_sf_substreams_sink_database_v1_database_proto_rawDescGZIP() []byte {
	file_sf_substreams_sink_database_v1_database_proto_rawDescOnce.Do(func() {
		file_sf_substreams_sink_database_v1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_sink_database_v1_database_proto_rawDescData)
	})
	return file_sf_substreams_sink_database_v1_database_proto_rawDescData
}

var file_sf_substreams_sink_database_v1_database_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sf_substreams_sink_database_v1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sf_substreams_sink_database_v1_database_proto_goTypes = []interface{}{
	(TableChange_Operation)(0),  // 0: sf.substreams.sink.database.v1.TableChange.Operation
	(*DatabaseChanges)(nil),     // 1: sf.substreams.sink.database.v1.DatabaseChanges
	(*TableChange)(nil),         // 2: sf.substreams.sink.database.v1.TableChange
	(*CompositePrimaryKey)(nil), // 3: sf.substreams.sink.database.v1.CompositePrimaryKey
	(*Field)(nil),               // 4: sf.substreams.sink.database.v1.Field
	nil,                         // 5: sf.substreams.sink.database.v1.CompositePrimaryKey.KeysEntry
}
var file_sf_substreams_sink_database_v1_database_proto_depIdxs = []int32{
	2, // 0: sf.substreams.sink.database.v1.DatabaseChanges.table_changes:type_name -> sf.substreams.sink.database.v1.TableChange
	3, // 1: sf.substreams.sink.database.v1.TableChange.composite_pk:type_name -> sf.substreams.sink.database.v1.CompositePrimaryKey
	0, // 2: sf.substreams.sink.database.v1.TableChange.operation:type_name -> sf.substreams.sink.database.v1.TableChange.Operation
	4, // 3: sf.substreams.sink.database.v1.TableChange.fields:type_name -> sf.substreams.sink.database.v1.Field
	5, // 4: sf.substreams.sink.database.v1.CompositePrimaryKey.keys:type_name -> sf.substreams.sink.database.v1.CompositePrimaryKey.KeysEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sf_substreams_sink_database_v1_database_proto_init() }
func file_sf_substreams_sink_database_v1_database_proto_init() {
	if File_sf_substreams_sink_database_v1_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_sink_database_v1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseChanges); i {
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
		file_sf_substreams_sink_database_v1_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableChange); i {
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
		file_sf_substreams_sink_database_v1_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompositePrimaryKey); i {
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
		file_sf_substreams_sink_database_v1_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
	file_sf_substreams_sink_database_v1_database_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TableChange_Pk)(nil),
		(*TableChange_CompositePk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sf_substreams_sink_database_v1_database_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_sink_database_v1_database_proto_goTypes,
		DependencyIndexes: file_sf_substreams_sink_database_v1_database_proto_depIdxs,
		EnumInfos:         file_sf_substreams_sink_database_v1_database_proto_enumTypes,
		MessageInfos:      file_sf_substreams_sink_database_v1_database_proto_msgTypes,
	}.Build()
	File_sf_substreams_sink_database_v1_database_proto = out.File
	file_sf_substreams_sink_database_v1_database_proto_rawDesc = nil
	file_sf_substreams_sink_database_v1_database_proto_goTypes = nil
	file_sf_substreams_sink_database_v1_database_proto_depIdxs = nil
}
