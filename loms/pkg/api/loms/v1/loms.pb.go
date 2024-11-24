// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.3
// source: loms.proto

package loms

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *Item) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type OrderCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64   `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *OrderCreateRequest) Reset() {
	*x = OrderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateRequest) ProtoMessage() {}

func (x *OrderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateRequest.ProtoReflect.Descriptor instead.
func (*OrderCreateRequest) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{1}
}

func (x *OrderCreateRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *OrderCreateRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *OrderCreateResponse) Reset() {
	*x = OrderCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateResponse) ProtoMessage() {}

func (x *OrderCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateResponse.ProtoReflect.Descriptor instead.
func (*OrderCreateResponse) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{2}
}

func (x *OrderCreateResponse) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type OrderInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *OrderInfoRequest) Reset() {
	*x = OrderInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoRequest) ProtoMessage() {}

func (x *OrderInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoRequest.ProtoReflect.Descriptor instead.
func (*OrderInfoRequest) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{3}
}

func (x *OrderInfoRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type OrderInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // (new | awaiting payment | failed | payed | cancelled)
	User   int64   `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Items  []*Item `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *OrderInfoResponse) Reset() {
	*x = OrderInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoResponse) ProtoMessage() {}

func (x *OrderInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoResponse.ProtoReflect.Descriptor instead.
func (*OrderInfoResponse) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{4}
}

func (x *OrderInfoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderInfoResponse) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *OrderInfoResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderPayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *OrderPayRequest) Reset() {
	*x = OrderPayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayRequest) ProtoMessage() {}

func (x *OrderPayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayRequest.ProtoReflect.Descriptor instead.
func (*OrderPayRequest) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{5}
}

func (x *OrderPayRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type OrderPayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderPayResponse) Reset() {
	*x = OrderPayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayResponse) ProtoMessage() {}

func (x *OrderPayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayResponse.ProtoReflect.Descriptor instead.
func (*OrderPayResponse) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{6}
}

type OrderCancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *OrderCancelRequest) Reset() {
	*x = OrderCancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCancelRequest) ProtoMessage() {}

func (x *OrderCancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCancelRequest.ProtoReflect.Descriptor instead.
func (*OrderCancelRequest) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{7}
}

func (x *OrderCancelRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type OrderCancelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderCancelResponse) Reset() {
	*x = OrderCancelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCancelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCancelResponse) ProtoMessage() {}

func (x *OrderCancelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCancelResponse.ProtoReflect.Descriptor instead.
func (*OrderCancelResponse) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{8}
}

type StocksInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *StocksInfoRequest) Reset() {
	*x = StocksInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocksInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocksInfoRequest) ProtoMessage() {}

func (x *StocksInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocksInfoRequest.ProtoReflect.Descriptor instead.
func (*StocksInfoRequest) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{9}
}

func (x *StocksInfoRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

type StocksInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StocksInfoResponse) Reset() {
	*x = StocksInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocksInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocksInfoResponse) ProtoMessage() {}

func (x *StocksInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocksInfoResponse.ProtoReflect.Descriptor instead.
func (*StocksInfoResponse) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{10}
}

func (x *StocksInfoResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_loms_proto protoreflect.FileDescriptor

var file_loms_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x12,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x2f, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x2c, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x5c, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x2b, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x12, 0x0a,
	0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2e, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x15, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22,
	0x2a, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x9d, 0x02, 0x0a, 0x0b,
	0x4c, 0x4f, 0x4d, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x11, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x12, 0x10, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x63,
	0x6d, 0x64, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loms_proto_rawDescOnce sync.Once
	file_loms_proto_rawDescData = file_loms_proto_rawDesc
)

func file_loms_proto_rawDescGZIP() []byte {
	file_loms_proto_rawDescOnce.Do(func() {
		file_loms_proto_rawDescData = protoimpl.X.CompressGZIP(file_loms_proto_rawDescData)
	})
	return file_loms_proto_rawDescData
}

var file_loms_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_loms_proto_goTypes = []interface{}{
	(*Item)(nil),                // 0: Item
	(*OrderCreateRequest)(nil),  // 1: OrderCreateRequest
	(*OrderCreateResponse)(nil), // 2: OrderCreateResponse
	(*OrderInfoRequest)(nil),    // 3: OrderInfoRequest
	(*OrderInfoResponse)(nil),   // 4: OrderInfoResponse
	(*OrderPayRequest)(nil),     // 5: OrderPayRequest
	(*OrderPayResponse)(nil),    // 6: OrderPayResponse
	(*OrderCancelRequest)(nil),  // 7: OrderCancelRequest
	(*OrderCancelResponse)(nil), // 8: OrderCancelResponse
	(*StocksInfoRequest)(nil),   // 9: StocksInfoRequest
	(*StocksInfoResponse)(nil),  // 10: StocksInfoResponse
}
var file_loms_proto_depIdxs = []int32{
	0,  // 0: OrderCreateRequest.items:type_name -> Item
	0,  // 1: OrderInfoResponse.items:type_name -> Item
	1,  // 2: LOMSService.OrderCreate:input_type -> OrderCreateRequest
	3,  // 3: LOMSService.OrderInfo:input_type -> OrderInfoRequest
	5,  // 4: LOMSService.OrderPay:input_type -> OrderPayRequest
	7,  // 5: LOMSService.OrderCancel:input_type -> OrderCancelRequest
	9,  // 6: LOMSService.StocksInfo:input_type -> StocksInfoRequest
	2,  // 7: LOMSService.OrderCreate:output_type -> OrderCreateResponse
	4,  // 8: LOMSService.OrderInfo:output_type -> OrderInfoResponse
	6,  // 9: LOMSService.OrderPay:output_type -> OrderPayResponse
	8,  // 10: LOMSService.OrderCancel:output_type -> OrderCancelResponse
	10, // 11: LOMSService.StocksInfo:output_type -> StocksInfoResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_loms_proto_init() }
func file_loms_proto_init() {
	if File_loms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_loms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateRequest); i {
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
		file_loms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateResponse); i {
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
		file_loms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoRequest); i {
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
		file_loms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoResponse); i {
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
		file_loms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayRequest); i {
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
		file_loms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayResponse); i {
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
		file_loms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCancelRequest); i {
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
		file_loms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCancelResponse); i {
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
		file_loms_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocksInfoRequest); i {
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
		file_loms_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocksInfoResponse); i {
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
			RawDescriptor: file_loms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loms_proto_goTypes,
		DependencyIndexes: file_loms_proto_depIdxs,
		MessageInfos:      file_loms_proto_msgTypes,
	}.Build()
	File_loms_proto = out.File
	file_loms_proto_rawDesc = nil
	file_loms_proto_goTypes = nil
	file_loms_proto_depIdxs = nil
}
