// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: oci.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The following is a proposed custom index for OCI repositories.
// For charts-syncer use case, we are only interested on charts artifacts
//
type OCIIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Charts []*OCIIndex_Charts `protobuf:"bytes,2,rep,name=charts,proto3" json:"charts,omitempty"`
}

func (x *OCIIndex) Reset() {
	*x = OCIIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCIIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCIIndex) ProtoMessage() {}

func (x *OCIIndex) ProtoReflect() protoreflect.Message {
	mi := &file_oci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCIIndex.ProtoReflect.Descriptor instead.
func (*OCIIndex) Descriptor() ([]byte, []int) {
	return file_oci_proto_rawDescGZIP(), []int{0}
}

func (x *OCIIndex) GetCharts() []*OCIIndex_Charts {
	if x != nil {
		return x.Charts
	}
	return nil
}

type OCIIndex_ChartVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Digest  string `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	Uri     string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *OCIIndex_ChartVersions) Reset() {
	*x = OCIIndex_ChartVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCIIndex_ChartVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCIIndex_ChartVersions) ProtoMessage() {}

func (x *OCIIndex_ChartVersions) ProtoReflect() protoreflect.Message {
	mi := &file_oci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCIIndex_ChartVersions.ProtoReflect.Descriptor instead.
func (*OCIIndex_ChartVersions) Descriptor() ([]byte, []int) {
	return file_oci_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OCIIndex_ChartVersions) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *OCIIndex_ChartVersions) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *OCIIndex_ChartVersions) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type OCIIndex_Charts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Versions []*OCIIndex_ChartVersions `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *OCIIndex_Charts) Reset() {
	*x = OCIIndex_Charts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oci_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCIIndex_Charts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCIIndex_Charts) ProtoMessage() {}

func (x *OCIIndex_Charts) ProtoReflect() protoreflect.Message {
	mi := &file_oci_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCIIndex_Charts.ProtoReflect.Descriptor instead.
func (*OCIIndex_Charts) Descriptor() ([]byte, []int) {
	return file_oci_proto_rawDescGZIP(), []int{0, 1}
}

func (x *OCIIndex_Charts) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OCIIndex_Charts) GetVersions() []*OCIIndex_ChartVersions {
	if x != nil {
		return x.Versions
	}
	return nil
}

var File_oci_proto protoreflect.FileDescriptor

var file_oci_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6f, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0xe4, 0x01, 0x0a, 0x08, 0x4f, 0x43, 0x49, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2c, 0x0a,
	0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x43, 0x49, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x73, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x1a, 0x53, 0x0a, 0x0d, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69,
	0x1a, 0x55, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x43, 0x49, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x08, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x69, 0x2d, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x2d, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oci_proto_rawDescOnce sync.Once
	file_oci_proto_rawDescData = file_oci_proto_rawDesc
)

func file_oci_proto_rawDescGZIP() []byte {
	file_oci_proto_rawDescOnce.Do(func() {
		file_oci_proto_rawDescData = protoimpl.X.CompressGZIP(file_oci_proto_rawDescData)
	})
	return file_oci_proto_rawDescData
}

var file_oci_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_oci_proto_goTypes = []interface{}{
	(*OCIIndex)(nil),               // 0: api.OCIIndex
	(*OCIIndex_ChartVersions)(nil), // 1: api.OCIIndex.ChartVersions
	(*OCIIndex_Charts)(nil),        // 2: api.OCIIndex.Charts
}
var file_oci_proto_depIdxs = []int32{
	2, // 0: api.OCIIndex.charts:type_name -> api.OCIIndex.Charts
	1, // 1: api.OCIIndex.Charts.versions:type_name -> api.OCIIndex.ChartVersions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_oci_proto_init() }
func file_oci_proto_init() {
	if File_oci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCIIndex); i {
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
		file_oci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCIIndex_ChartVersions); i {
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
		file_oci_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCIIndex_Charts); i {
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
			RawDescriptor: file_oci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oci_proto_goTypes,
		DependencyIndexes: file_oci_proto_depIdxs,
		MessageInfos:      file_oci_proto_msgTypes,
	}.Build()
	File_oci_proto = out.File
	file_oci_proto_rawDesc = nil
	file_oci_proto_goTypes = nil
	file_oci_proto_depIdxs = nil
}
