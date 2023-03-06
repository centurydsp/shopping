// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/cart/cart.proto

/*
Package go_micro_service_cart is a generated protocol buffer package.

It is generated from these files:
	proto/cart/cart.proto

It has these top-level messages:
	CartInfo
	ResponseAdd
	Clean
	Response
	Item
	CartID
	CartFindAll
	CartAll
*/
package cart

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CartInfo struct {
	Id        int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,3,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	SizeId    int64 `protobuf:"varint,4,opt,name=size_id,json=sizeId" json:"size_id,omitempty"`
	Num       int64 `protobuf:"varint,5,opt,name=num" json:"num,omitempty"`
}

func (m *CartInfo) Reset()                    { *m = CartInfo{} }
func (m *CartInfo) String() string            { return proto.CompactTextString(m) }
func (*CartInfo) ProtoMessage()               {}
func (*CartInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CartInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CartInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CartInfo) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *CartInfo) GetSizeId() int64 {
	if m != nil {
		return m.SizeId
	}
	return 0
}

func (m *CartInfo) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ResponseAdd struct {
	CartId int64  `protobuf:"varint,1,opt,name=cart_id,json=cartId" json:"cart_id,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *ResponseAdd) Reset()                    { *m = ResponseAdd{} }
func (m *ResponseAdd) String() string            { return proto.CompactTextString(m) }
func (*ResponseAdd) ProtoMessage()               {}
func (*ResponseAdd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseAdd) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *ResponseAdd) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Clean struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Clean) Reset()                    { *m = Clean{} }
func (m *Clean) String() string            { return proto.CompactTextString(m) }
func (*Clean) ProtoMessage()               {}
func (*Clean) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Clean) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Response struct {
	Meg string `protobuf:"bytes,1,opt,name=meg" json:"meg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetMeg() string {
	if m != nil {
		return m.Meg
	}
	return ""
}

type Item struct {
	Id        int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ChangeNum int64 `protobuf:"varint,2,opt,name=change_num,json=changeNum" json:"change_num,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetChangeNum() int64 {
	if m != nil {
		return m.ChangeNum
	}
	return 0
}

type CartID struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CartID) Reset()                    { *m = CartID{} }
func (m *CartID) String() string            { return proto.CompactTextString(m) }
func (*CartID) ProtoMessage()               {}
func (*CartID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CartID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CartFindAll struct {
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *CartFindAll) Reset()                    { *m = CartFindAll{} }
func (m *CartFindAll) String() string            { return proto.CompactTextString(m) }
func (*CartFindAll) ProtoMessage()               {}
func (*CartFindAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CartFindAll) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CartAll struct {
	CartInfo []*CartInfo `protobuf:"bytes,1,rep,name=cart_info,json=cartInfo" json:"cart_info,omitempty"`
}

func (m *CartAll) Reset()                    { *m = CartAll{} }
func (m *CartAll) String() string            { return proto.CompactTextString(m) }
func (*CartAll) ProtoMessage()               {}
func (*CartAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CartAll) GetCartInfo() []*CartInfo {
	if m != nil {
		return m.CartInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CartInfo)(nil), "go.micro.service.cart.CartInfo")
	proto.RegisterType((*ResponseAdd)(nil), "go.micro.service.cart.ResponseAdd")
	proto.RegisterType((*Clean)(nil), "go.micro.service.cart.Clean")
	proto.RegisterType((*Response)(nil), "go.micro.service.cart.Response")
	proto.RegisterType((*Item)(nil), "go.micro.service.cart.Item")
	proto.RegisterType((*CartID)(nil), "go.micro.service.cart.CartID")
	proto.RegisterType((*CartFindAll)(nil), "go.micro.service.cart.CartFindAll")
	proto.RegisterType((*CartAll)(nil), "go.micro.service.cart.CartAll")
}

func init() { proto.RegisterFile("proto/cart/cart.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x5d, 0x8b, 0xd3, 0x50,
	0x10, 0x6d, 0x9a, 0x6c, 0xda, 0x4c, 0x61, 0x91, 0x0b, 0x8b, 0xa1, 0x6e, 0xdd, 0x72, 0x1f, 0x64,
	0x9f, 0x22, 0xac, 0x08, 0x3e, 0xf8, 0x12, 0x37, 0xec, 0x12, 0x90, 0x45, 0xf2, 0x07, 0x4a, 0xcc,
	0x9d, 0x8d, 0x81, 0xe4, 0xde, 0x72, 0x93, 0x08, 0x0a, 0xfe, 0x51, 0x7f, 0x8d, 0xcc, 0xa4, 0x2b,
	0x45, 0x4d, 0xdb, 0x07, 0x5f, 0xc2, 0x7c, 0x9c, 0x73, 0x38, 0x33, 0x93, 0x0b, 0x17, 0x5b, 0x6b,
	0x3a, 0xf3, 0xba, 0xc8, 0x6d, 0xc7, 0x9f, 0x88, 0x73, 0x71, 0x51, 0x9a, 0xa8, 0xa9, 0x0a, 0x6b,
	0xa2, 0x16, 0xed, 0xd7, 0xaa, 0xc0, 0x88, 0x9a, 0xf2, 0x07, 0xcc, 0x6f, 0x73, 0xdb, 0xa5, 0xfa,
	0xd1, 0x88, 0x73, 0x98, 0x56, 0x2a, 0x74, 0xd6, 0xce, 0xb5, 0x9b, 0x4d, 0x2b, 0x25, 0x9e, 0xc3,
	0xac, 0x6f, 0xd1, 0x6e, 0x2a, 0x15, 0x4e, 0xb9, 0xe8, 0x53, 0x9a, 0x2a, 0xb1, 0x02, 0xd8, 0x5a,
	0xa3, 0xfa, 0xa2, 0xa3, 0x9e, 0xcb, 0xbd, 0x60, 0x57, 0x49, 0x99, 0xd7, 0x56, 0xdf, 0x91, 0x7a,
	0xde, 0xc0, 0xa3, 0x34, 0x55, 0xe2, 0x19, 0xb8, 0xba, 0x6f, 0xc2, 0x33, 0x2e, 0x52, 0x28, 0xdf,
	0xc1, 0x22, 0xc3, 0x76, 0x6b, 0x74, 0x8b, 0xb1, 0x62, 0x26, 0xb9, 0xda, 0xfc, 0xb6, 0xe1, 0x53,
	0x3a, 0x30, 0x9b, 0xb6, 0x64, 0x1b, 0x41, 0x46, 0xa1, 0x5c, 0xc3, 0xd9, 0x6d, 0x8d, 0xb9, 0xde,
	0x77, 0xe9, 0xec, 0xbb, 0x94, 0x97, 0x30, 0x7f, 0xd2, 0x66, 0x3e, 0x96, 0x0c, 0x20, 0x3e, 0x96,
	0xf2, 0x2d, 0x78, 0x69, 0x87, 0xcd, 0x5f, 0x43, 0xaf, 0x00, 0x8a, 0x2f, 0xb9, 0x2e, 0x71, 0x43,
	0x56, 0x87, 0xb9, 0x83, 0xa1, 0xf2, 0xd0, 0x37, 0x32, 0x04, 0x9f, 0xf7, 0x95, 0xfc, 0x49, 0x94,
	0xaf, 0x60, 0x41, 0x9d, 0xbb, 0x4a, 0xab, 0xb8, 0xae, 0xc7, 0x6d, 0xdd, 0xc3, 0x8c, 0x70, 0x84,
	0x79, 0x0f, 0xc1, 0x30, 0xae, 0x7e, 0x34, 0xa1, 0xb3, 0x76, 0xaf, 0x17, 0x37, 0x57, 0xd1, 0x3f,
	0xef, 0x14, 0x3d, 0x1d, 0x29, 0x9b, 0x17, 0xbb, 0xe8, 0xe6, 0xa7, 0x0b, 0x1e, 0x95, 0xc5, 0x27,
	0x98, 0xc5, 0x4a, 0x71, 0x78, 0x8c, 0xbe, 0x94, 0x23, 0x80, 0xbd, 0x2b, 0xc8, 0x89, 0xf8, 0x08,
	0x01, 0x2f, 0x97, 0x35, 0x2f, 0xc7, 0x34, 0x09, 0xb1, 0xbc, 0x3a, 0x22, 0x28, 0x27, 0xe2, 0x0e,
	0xbc, 0x54, 0x17, 0x56, 0xbc, 0x18, 0x81, 0xd2, 0x1d, 0x4e, 0xd4, 0x49, 0xf0, 0x3f, 0xe8, 0x64,
	0x70, 0x9e, 0x60, 0x8d, 0x1d, 0x12, 0xe1, 0xc3, 0xb7, 0x34, 0x11, 0xab, 0x43, 0x6b, 0x4b, 0x4e,
	0xd1, 0x7c, 0x00, 0xff, 0x1e, 0xf9, 0xa8, 0xf2, 0x80, 0xd6, 0xee, 0xe7, 0x58, 0xbe, 0x3c, 0x80,
	0x89, 0xeb, 0x5a, 0x4e, 0x3e, 0xfb, 0xfc, 0x6a, 0xdf, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0xf4, 0x79, 0x58, 0xce, 0x03, 0x00, 0x00,
}