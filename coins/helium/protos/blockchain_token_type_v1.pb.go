// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/blockchain_token_type_v1.proto

package protos

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockchainTokenTypeV1 int32

const (
	BlockchainTokenTypeV1_hnt    BlockchainTokenTypeV1 = 0
	BlockchainTokenTypeV1_hst    BlockchainTokenTypeV1 = 1
	BlockchainTokenTypeV1_mobile BlockchainTokenTypeV1 = 2
	BlockchainTokenTypeV1_iot    BlockchainTokenTypeV1 = 3
)

// Enum value maps for BlockchainTokenTypeV1.
var (
	BlockchainTokenTypeV1_name = map[int32]string{
		0: "hnt",
		1: "hst",
		2: "mobile",
		3: "iot",
	}
	BlockchainTokenTypeV1_value = map[string]int32{
		"hnt":    0,
		"hst":    1,
		"mobile": 2,
		"iot":    3,
	}
)

func (x BlockchainTokenTypeV1) Enum() *BlockchainTokenTypeV1 {
	p := new(BlockchainTokenTypeV1)
	*p = x
	return p
}

func (x BlockchainTokenTypeV1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockchainTokenTypeV1) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_blockchain_token_type_v1_proto_enumTypes[0].Descriptor()
}

func (BlockchainTokenTypeV1) Type() protoreflect.EnumType {
	return &file_proto_blockchain_token_type_v1_proto_enumTypes[0]
}

func (x BlockchainTokenTypeV1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockchainTokenTypeV1.Descriptor instead.
func (BlockchainTokenTypeV1) EnumDescriptor() ([]byte, []int) {
	return file_proto_blockchain_token_type_v1_proto_rawDescGZIP(), []int{0}
}

var File_proto_blockchain_token_type_v1_proto protoreflect.FileDescriptor

var file_proto_blockchain_token_type_v1_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2a, 0x41,
	0x0a, 0x18, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x31, 0x12, 0x07, 0x0a, 0x03, 0x68, 0x6e,
	0x74, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x68, 0x73, 0x74, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x69, 0x6f, 0x74, 0x10,
	0x03, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_blockchain_token_type_v1_proto_rawDescOnce sync.Once
	file_proto_blockchain_token_type_v1_proto_rawDescData = file_proto_blockchain_token_type_v1_proto_rawDesc
)

func file_proto_blockchain_token_type_v1_proto_rawDescGZIP() []byte {
	file_proto_blockchain_token_type_v1_proto_rawDescOnce.Do(func() {
		file_proto_blockchain_token_type_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_blockchain_token_type_v1_proto_rawDescData)
	})
	return file_proto_blockchain_token_type_v1_proto_rawDescData
}

var file_proto_blockchain_token_type_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_blockchain_token_type_v1_proto_goTypes = []interface{}{
	(BlockchainTokenTypeV1)(0), // 0: protos.blockchain_token_type_v1
}
var file_proto_blockchain_token_type_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_blockchain_token_type_v1_proto_init() }
func file_proto_blockchain_token_type_v1_proto_init() {
	if File_proto_blockchain_token_type_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_blockchain_token_type_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_blockchain_token_type_v1_proto_goTypes,
		DependencyIndexes: file_proto_blockchain_token_type_v1_proto_depIdxs,
		EnumInfos:         file_proto_blockchain_token_type_v1_proto_enumTypes,
	}.Build()
	File_proto_blockchain_token_type_v1_proto = out.File
	file_proto_blockchain_token_type_v1_proto_rawDesc = nil
	file_proto_blockchain_token_type_v1_proto_goTypes = nil
	file_proto_blockchain_token_type_v1_proto_depIdxs = nil
}
