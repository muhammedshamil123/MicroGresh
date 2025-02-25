// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: proto/cart-pb/cart.proto

package cart_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cart struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	OfferAmount   float64                `protobuf:"fixed64,5,opt,name=offer_amount,json=offerAmount,proto3" json:"offer_amount,omitempty"`
	StockLeft     uint32                 `protobuf:"varint,6,opt,name=stock_left,json=stockLeft,proto3" json:"stock_left,omitempty"`
	Quantity      uint32                 `protobuf:"varint,7,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cart) Reset() {
	*x = Cart{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{0}
}

func (x *Cart) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cart) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cart) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cart) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Cart) GetOfferAmount() float64 {
	if x != nil {
		return x.OfferAmount
	}
	return 0
}

func (x *Cart) GetStockLeft() uint32 {
	if x != nil {
		return x.StockLeft
	}
	return 0
}

func (x *Cart) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddToCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint32                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	UserId        uint32                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Quantity      uint32                 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddToCartRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *AddToCartRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddToCartRequest) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddToCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{2}
}

func (x *AddToCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddToCartResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RemoveFromCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId     uint32                 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveFromCartRequest) Reset() {
	*x = RemoveFromCartRequest{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartRequest) ProtoMessage() {}

func (x *RemoveFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFromCartRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RemoveFromCartRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type RemoveFromCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveFromCartResponse) Reset() {
	*x = RemoveFromCartResponse{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartResponse) ProtoMessage() {}

func (x *RemoveFromCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartResponse.ProtoReflect.Descriptor instead.
func (*RemoveFromCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveFromCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RemoveFromCartResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ViewCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ViewCartRequest) Reset() {
	*x = ViewCartRequest{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ViewCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartRequest) ProtoMessage() {}

func (x *ViewCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartRequest.ProtoReflect.Descriptor instead.
func (*ViewCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{5}
}

func (x *ViewCartRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ViewCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	CartsItems    []*Cart                `protobuf:"bytes,3,rep,name=carts_items,json=cartsItems,proto3" json:"carts_items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ViewCartResponse) Reset() {
	*x = ViewCartResponse{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ViewCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartResponse) ProtoMessage() {}

func (x *ViewCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartResponse.ProtoReflect.Descriptor instead.
func (*ViewCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{6}
}

func (x *ViewCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ViewCartResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ViewCartResponse) GetCartsItems() []*Cart {
	if x != nil {
		return x.CartsItems
	}
	return nil
}

type ClearCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearCartRequest) Reset() {
	*x = ClearCartRequest{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartRequest) ProtoMessage() {}

func (x *ClearCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartRequest.ProtoReflect.Descriptor instead.
func (*ClearCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{7}
}

func (x *ClearCartRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ClearCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearCartResponse) Reset() {
	*x = ClearCartResponse{}
	mi := &file_proto_cart_pb_cart_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartResponse) ProtoMessage() {}

func (x *ClearCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_pb_cart_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartResponse.ProtoReflect.Descriptor instead.
func (*ClearCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_pb_cart_proto_rawDescGZIP(), []int{8}
}

func (x *ClearCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ClearCartResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_cart_pb_cart_proto protoreflect.FileDescriptor

var file_proto_cart_pb_cart_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2d, 0x70, 0x62, 0x2f,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x22, 0xc0, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x66, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x47, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73, 0x0a, 0x10, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2b,
	0x0a, 0x0b, 0x63, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x0a, 0x63, 0x61, 0x72, 0x74, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x47, 0x0a, 0x11, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x91, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x69,
	0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2d, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_cart_pb_cart_proto_rawDescOnce sync.Once
	file_proto_cart_pb_cart_proto_rawDescData []byte
)

func file_proto_cart_pb_cart_proto_rawDescGZIP() []byte {
	file_proto_cart_pb_cart_proto_rawDescOnce.Do(func() {
		file_proto_cart_pb_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_cart_pb_cart_proto_rawDesc), len(file_proto_cart_pb_cart_proto_rawDesc)))
	})
	return file_proto_cart_pb_cart_proto_rawDescData
}

var file_proto_cart_pb_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_cart_pb_cart_proto_goTypes = []any{
	(*Cart)(nil),                   // 0: auth.Cart
	(*AddToCartRequest)(nil),       // 1: auth.AddToCartRequest
	(*AddToCartResponse)(nil),      // 2: auth.AddToCartResponse
	(*RemoveFromCartRequest)(nil),  // 3: auth.RemoveFromCartRequest
	(*RemoveFromCartResponse)(nil), // 4: auth.RemoveFromCartResponse
	(*ViewCartRequest)(nil),        // 5: auth.ViewCartRequest
	(*ViewCartResponse)(nil),       // 6: auth.ViewCartResponse
	(*ClearCartRequest)(nil),       // 7: auth.ClearCartRequest
	(*ClearCartResponse)(nil),      // 8: auth.ClearCartResponse
}
var file_proto_cart_pb_cart_proto_depIdxs = []int32{
	0, // 0: auth.ViewCartResponse.carts_items:type_name -> auth.Cart
	1, // 1: auth.CartService.AddToCart:input_type -> auth.AddToCartRequest
	3, // 2: auth.CartService.RemoveFromCart:input_type -> auth.RemoveFromCartRequest
	5, // 3: auth.CartService.ViewCart:input_type -> auth.ViewCartRequest
	7, // 4: auth.CartService.ClearCart:input_type -> auth.ClearCartRequest
	2, // 5: auth.CartService.AddToCart:output_type -> auth.AddToCartResponse
	4, // 6: auth.CartService.RemoveFromCart:output_type -> auth.RemoveFromCartResponse
	6, // 7: auth.CartService.ViewCart:output_type -> auth.ViewCartResponse
	8, // 8: auth.CartService.ClearCart:output_type -> auth.ClearCartResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cart_pb_cart_proto_init() }
func file_proto_cart_pb_cart_proto_init() {
	if File_proto_cart_pb_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_cart_pb_cart_proto_rawDesc), len(file_proto_cart_pb_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cart_pb_cart_proto_goTypes,
		DependencyIndexes: file_proto_cart_pb_cart_proto_depIdxs,
		MessageInfos:      file_proto_cart_pb_cart_proto_msgTypes,
	}.Build()
	File_proto_cart_pb_cart_proto = out.File
	file_proto_cart_pb_cart_proto_goTypes = nil
	file_proto_cart_pb_cart_proto_depIdxs = nil
}
