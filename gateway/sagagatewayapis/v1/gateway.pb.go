// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: sagagatewayapis/v1/gateway.proto

package v1

import (
	v1 "github.com/awe76/saga/api/sagatransactionapis/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start   string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sagagatewayapis_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sagagatewayapis_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_sagagatewayapis_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *InitRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *v1.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sagagatewayapis_v1_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sagagatewayapis_v1_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_sagagatewayapis_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetState() *v1.State {
	if x != nil {
		return x.State
	}
	return nil
}

type CompleteOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsRollback bool          `protobuf:"varint,1,opt,name=is_rollback,json=isRollback,proto3" json:"is_rollback,omitempty"`
	Id         string        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Payload    string        `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Operation  *v1.Operation `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *CompleteOperationRequest) Reset() {
	*x = CompleteOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sagagatewayapis_v1_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOperationRequest) ProtoMessage() {}

func (x *CompleteOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sagagatewayapis_v1_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOperationRequest.ProtoReflect.Descriptor instead.
func (*CompleteOperationRequest) Descriptor() ([]byte, []int) {
	return file_sagagatewayapis_v1_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteOperationRequest) GetIsRollback() bool {
	if x != nil {
		return x.IsRollback
	}
	return false
}

func (x *CompleteOperationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompleteOperationRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *CompleteOperationRequest) GetOperation() *v1.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type CompleteOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *v1.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *CompleteOperationResponse) Reset() {
	*x = CompleteOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sagagatewayapis_v1_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOperationResponse) ProtoMessage() {}

func (x *CompleteOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sagagatewayapis_v1_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOperationResponse.ProtoReflect.Descriptor instead.
func (*CompleteOperationResponse) Descriptor() ([]byte, []int) {
	return file_sagagatewayapis_v1_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *CompleteOperationResponse) GetState() *v1.State {
	if x != nil {
		return x.State
	}
	return nil
}

var File_sagagatewayapis_v1_gateway_proto protoreflect.FileDescriptor

var file_sagagatewayapis_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x61, 0x67, 0x61, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x73, 0x61, 0x67, 0x61, 0x73, 0x74, 0x61, 0x74, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x73, 0x61, 0x67, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x67,
	0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3d, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x43, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x61, 0x67, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3f, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a,
	0x19, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x61, 0x67, 0x61,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32,
	0x8c, 0x02, 0x0a, 0x10, 0x53, 0x61, 0x67, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x1d, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x73, 0x74, 0x61, 0x74, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x61,
	0x67, 0x61, 0x73, 0x74, 0x61, 0x74, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x69,
	0x6e, 0x69, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x95, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x73,
	0x61, 0x67, 0x61, 0x73, 0x74, 0x61, 0x74, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x77, 0x65,
	0x37, 0x36, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x61,
	0x67, 0x61, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sagagatewayapis_v1_gateway_proto_rawDescOnce sync.Once
	file_sagagatewayapis_v1_gateway_proto_rawDescData = file_sagagatewayapis_v1_gateway_proto_rawDesc
)

func file_sagagatewayapis_v1_gateway_proto_rawDescGZIP() []byte {
	file_sagagatewayapis_v1_gateway_proto_rawDescOnce.Do(func() {
		file_sagagatewayapis_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_sagagatewayapis_v1_gateway_proto_rawDescData)
	})
	return file_sagagatewayapis_v1_gateway_proto_rawDescData
}

var file_sagagatewayapis_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sagagatewayapis_v1_gateway_proto_goTypes = []interface{}{
	(*InitRequest)(nil),               // 0: sagastateapis.v1.InitRequest
	(*InitResponse)(nil),              // 1: sagastateapis.v1.InitResponse
	(*CompleteOperationRequest)(nil),  // 2: sagastateapis.v1.CompleteOperationRequest
	(*CompleteOperationResponse)(nil), // 3: sagastateapis.v1.CompleteOperationResponse
	(*v1.State)(nil),                  // 4: sagatransactionapis.v1.State
	(*v1.Operation)(nil),              // 5: sagatransactionapis.v1.Operation
}
var file_sagagatewayapis_v1_gateway_proto_depIdxs = []int32{
	4, // 0: sagastateapis.v1.InitResponse.state:type_name -> sagatransactionapis.v1.State
	5, // 1: sagastateapis.v1.CompleteOperationRequest.operation:type_name -> sagatransactionapis.v1.Operation
	4, // 2: sagastateapis.v1.CompleteOperationResponse.state:type_name -> sagatransactionapis.v1.State
	0, // 3: sagastateapis.v1.SagaStateService.Init:input_type -> sagastateapis.v1.InitRequest
	2, // 4: sagastateapis.v1.SagaStateService.CompleteOperation:input_type -> sagastateapis.v1.CompleteOperationRequest
	1, // 5: sagastateapis.v1.SagaStateService.Init:output_type -> sagastateapis.v1.InitResponse
	3, // 6: sagastateapis.v1.SagaStateService.CompleteOperation:output_type -> sagastateapis.v1.CompleteOperationResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sagagatewayapis_v1_gateway_proto_init() }
func file_sagagatewayapis_v1_gateway_proto_init() {
	if File_sagagatewayapis_v1_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sagagatewayapis_v1_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_sagagatewayapis_v1_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
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
		file_sagagatewayapis_v1_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOperationRequest); i {
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
		file_sagagatewayapis_v1_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOperationResponse); i {
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
			RawDescriptor: file_sagagatewayapis_v1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sagagatewayapis_v1_gateway_proto_goTypes,
		DependencyIndexes: file_sagagatewayapis_v1_gateway_proto_depIdxs,
		MessageInfos:      file_sagagatewayapis_v1_gateway_proto_msgTypes,
	}.Build()
	File_sagagatewayapis_v1_gateway_proto = out.File
	file_sagagatewayapis_v1_gateway_proto_rawDesc = nil
	file_sagagatewayapis_v1_gateway_proto_goTypes = nil
	file_sagagatewayapis_v1_gateway_proto_depIdxs = nil
}
