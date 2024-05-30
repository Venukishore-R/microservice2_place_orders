// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: protos/order.proto

package protos

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

type OrderEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderEmpty) Reset() {
	*x = OrderEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderEmpty) ProtoMessage() {}

func (x *OrderEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderEmpty.ProtoReflect.Descriptor instead.
func (*OrderEmpty) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Quantity string `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price    string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Order) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *Order) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type GetDeleteOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDeleteOrderReq) Reset() {
	*x = GetDeleteOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeleteOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeleteOrderReq) ProtoMessage() {}

func (x *GetDeleteOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeleteOrderReq.ProtoReflect.Descriptor instead.
func (*GetDeleteOrderReq) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeleteOrderReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Order *Order `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *UpdateOrderReq) Reset() {
	*x = UpdateOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderReq) ProtoMessage() {}

func (x *UpdateOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderReq.ProtoReflect.Descriptor instead.
func (*UpdateOrderReq) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOrderReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateOrderReq) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type GetOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetOrderResp) Reset() {
	*x = GetOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResp) ProtoMessage() {}

func (x *GetOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResp.ProtoReflect.Descriptor instead.
func (*GetOrderResp) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderResp) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type CreateUpdateDeleteOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateUpdateDeleteOrderResp) Reset() {
	*x = CreateUpdateDeleteOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUpdateDeleteOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUpdateDeleteOrderResp) ProtoMessage() {}

func (x *CreateUpdateDeleteOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUpdateDeleteOrderResp.ProtoReflect.Descriptor instead.
func (*CreateUpdateDeleteOrderResp) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUpdateDeleteOrderResp) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetOrdersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []*Order `protobuf:"bytes,1,rep,name=Order,proto3" json:"Order,omitempty"`
}

func (x *GetOrdersResp) Reset() {
	*x = GetOrdersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersResp) ProtoMessage() {}

func (x *GetOrdersResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersResp.ProtoReflect.Descriptor instead.
func (*GetOrdersResp) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrdersResp) GetOrder() []*Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_protos_order_proto protoreflect.FileDescriptor

var file_protos_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x4d, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x9b, 0x02, 0x0a, 0x0c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x28, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x65, 0x6e, 0x75, 0x6b, 0x69, 0x73, 0x68, 0x6f,
	0x72, 0x65, 0x2d, 0x52, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x32, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_order_proto_rawDescOnce sync.Once
	file_protos_order_proto_rawDescData = file_protos_order_proto_rawDesc
)

func file_protos_order_proto_rawDescGZIP() []byte {
	file_protos_order_proto_rawDescOnce.Do(func() {
		file_protos_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_order_proto_rawDescData)
	})
	return file_protos_order_proto_rawDescData
}

var file_protos_order_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_order_proto_goTypes = []interface{}{
	(*OrderEmpty)(nil),                  // 0: OrderEmpty
	(*Order)(nil),                       // 1: Order
	(*GetDeleteOrderReq)(nil),           // 2: GetDeleteOrderReq
	(*UpdateOrderReq)(nil),              // 3: UpdateOrderReq
	(*GetOrderResp)(nil),                // 4: GetOrderResp
	(*CreateUpdateDeleteOrderResp)(nil), // 5: CreateUpdateDeleteOrderResp
	(*GetOrdersResp)(nil),               // 6: GetOrdersResp
}
var file_protos_order_proto_depIdxs = []int32{
	1, // 0: UpdateOrderReq.order:type_name -> Order
	1, // 1: GetOrderResp.order:type_name -> Order
	1, // 2: GetOrdersResp.Order:type_name -> Order
	1, // 3: OrderService.CreateOrder:input_type -> Order
	0, // 4: OrderService.GetOrders:input_type -> OrderEmpty
	2, // 5: OrderService.GetOrder:input_type -> GetDeleteOrderReq
	3, // 6: OrderService.UpdateOrder:input_type -> UpdateOrderReq
	2, // 7: OrderService.DeleteOrder:input_type -> GetDeleteOrderReq
	5, // 8: OrderService.CreateOrder:output_type -> CreateUpdateDeleteOrderResp
	6, // 9: OrderService.GetOrders:output_type -> GetOrdersResp
	4, // 10: OrderService.GetOrder:output_type -> GetOrderResp
	5, // 11: OrderService.UpdateOrder:output_type -> CreateUpdateDeleteOrderResp
	5, // 12: OrderService.DeleteOrder:output_type -> CreateUpdateDeleteOrderResp
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_order_proto_init() }
func file_protos_order_proto_init() {
	if File_protos_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderEmpty); i {
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
		file_protos_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_protos_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeleteOrderReq); i {
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
		file_protos_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderReq); i {
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
		file_protos_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderResp); i {
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
		file_protos_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUpdateDeleteOrderResp); i {
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
		file_protos_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersResp); i {
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
			RawDescriptor: file_protos_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_order_proto_goTypes,
		DependencyIndexes: file_protos_order_proto_depIdxs,
		MessageInfos:      file_protos_order_proto_msgTypes,
	}.Build()
	File_protos_order_proto = out.File
	file_protos_order_proto_rawDesc = nil
	file_protos_order_proto_goTypes = nil
	file_protos_order_proto_depIdxs = nil
}
