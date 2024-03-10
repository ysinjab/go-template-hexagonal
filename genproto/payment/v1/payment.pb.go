// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.25.3
// source: payment/v1/payment.proto

package paymentpbv1

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

type CreatePaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string  `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	OrderId    string  `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TotalPrice float64 `protobuf:"fixed64,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *CreatePaymentRequest) Reset() {
	*x = CreatePaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentRequest) ProtoMessage() {}

func (x *CreatePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentRequest.ProtoReflect.Descriptor instead.
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaymentRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreatePaymentRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreatePaymentRequest) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type CreatePaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePaymentResponse) Reset() {
	*x = CreatePaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResponse) ProtoMessage() {}

func (x *CreatePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResponse.ProtoReflect.Descriptor instead.
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePaymentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_payment_v1_payment_proto protoreflect.FileDescriptor

var file_payment_v1_payment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x27,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x6c, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_v1_payment_proto_rawDescOnce sync.Once
	file_payment_v1_payment_proto_rawDescData = file_payment_v1_payment_proto_rawDesc
)

func file_payment_v1_payment_proto_rawDescGZIP() []byte {
	file_payment_v1_payment_proto_rawDescOnce.Do(func() {
		file_payment_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_v1_payment_proto_rawDescData)
	})
	return file_payment_v1_payment_proto_rawDescData
}

var file_payment_v1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payment_v1_payment_proto_goTypes = []interface{}{
	(*CreatePaymentRequest)(nil),  // 0: proto.payment.CreatePaymentRequest
	(*CreatePaymentResponse)(nil), // 1: proto.payment.CreatePaymentResponse
}
var file_payment_v1_payment_proto_depIdxs = []int32{
	0, // 0: proto.payment.PaymentService.CreatePayment:input_type -> proto.payment.CreatePaymentRequest
	1, // 1: proto.payment.PaymentService.CreatePayment:output_type -> proto.payment.CreatePaymentResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_v1_payment_proto_init() }
func file_payment_v1_payment_proto_init() {
	if File_payment_v1_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_v1_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentRequest); i {
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
		file_payment_v1_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentResponse); i {
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
			RawDescriptor: file_payment_v1_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_v1_payment_proto_goTypes,
		DependencyIndexes: file_payment_v1_payment_proto_depIdxs,
		MessageInfos:      file_payment_v1_payment_proto_msgTypes,
	}.Build()
	File_payment_v1_payment_proto = out.File
	file_payment_v1_payment_proto_rawDesc = nil
	file_payment_v1_payment_proto_goTypes = nil
	file_payment_v1_payment_proto_depIdxs = nil
}
