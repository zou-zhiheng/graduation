// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.9.0
// source: api/gateway/v1/dataVisualization.proto

package v1

import (
	common "gateway/api/common"
	
	
	
	
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

type IntervalType int32

const (
	// 1. 空
	IntervalType_Null IntervalType = 0
	// 2. 天
	IntervalType_Day IntervalType = 1
	// 3. 星期
	IntervalType_Week IntervalType = 2
	// 4 . 月
	IntervalType_Month IntervalType = 3
)

// Enum value maps for IntervalType.
var (
	IntervalType_name = map[int32]string{
		0: "Null",
		1: "Day",
		2: "Week",
		3: "Month",
	}
	IntervalType_value = map[string]int32{
		"Null":  0,
		"Day":   1,
		"Week":  2,
		"Month": 3,
	}
)

func (x IntervalType) Enum() *IntervalType {
	p := new(IntervalType)
	*p = x
	return p
}

func (x IntervalType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IntervalType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_gateway_v1_dataVisualization_proto_enumTypes[0].Descriptor()
}

func (IntervalType) Type() protoreflect.EnumType {
	return &file_api_gateway_v1_dataVisualization_proto_enumTypes[0]
}

func (x IntervalType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IntervalType.Descriptor instead.
func (IntervalType) EnumDescriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{0}
}

type ChartType int32

const (
	// 1. 空
	ChartType_ChartTypeNull ChartType = 0
	// 1.饼图
	ChartType_Pip ChartType = 1
	// 2.折线图
	ChartType_Line ChartType = 2
)

// Enum value maps for ChartType.
var (
	ChartType_name = map[int32]string{
		0: "ChartTypeNull",
		1: "Pip",
		2: "Line",
	}
	ChartType_value = map[string]int32{
		"ChartTypeNull": 0,
		"Pip":           1,
		"Line":          2,
	}
)

func (x ChartType) Enum() *ChartType {
	p := new(ChartType)
	*p = x
	return p
}

func (x ChartType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChartType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_gateway_v1_dataVisualization_proto_enumTypes[1].Descriptor()
}

func (ChartType) Type() protoreflect.EnumType {
	return &file_api_gateway_v1_dataVisualization_proto_enumTypes[1]
}

func (x ChartType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChartType.Descriptor instead.
func (ChartType) EnumDescriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{1}
}

type DeviceDataGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1. 设备名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// 2. 设备编号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 3. 页大小
	PageSize uint64 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize"`
	// 4. 开始时间
	StartTime string `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime"`
	// 5. 结束时间
	EndTime string `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime"`
	// 6. 页码
	CurrPage uint64 `protobuf:"varint,6,opt,name=currPage,proto3" json:"currPage"`
	// 7. 用户ID
	UserId uint64 `protobuf:"varint,7,opt,name=userId,proto3" json:"userId"`
}

func (x *DeviceDataGetReq) Reset() {
	*x = DeviceDataGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDataGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDataGetReq) ProtoMessage() {}

func (x *DeviceDataGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDataGetReq.ProtoReflect.Descriptor instead.
func (*DeviceDataGetReq) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceDataGetReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceDataGetReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeviceDataGetReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DeviceDataGetReq) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *DeviceDataGetReq) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *DeviceDataGetReq) GetCurrPage() uint64 {
	if x != nil {
		return x.CurrPage
	}
	return 0
}

func (x *DeviceDataGetReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeviceDataGetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1. 设备名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// 2. 设备编号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 3. 设备数据
	Data []*common.DeviceData `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	// 4. 数据量
	Count uint64 `protobuf:"varint,4,opt,name=count,proto3" json:"count"`
}

func (x *DeviceDataGetRes) Reset() {
	*x = DeviceDataGetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDataGetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDataGetRes) ProtoMessage() {}

func (x *DeviceDataGetRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDataGetRes.ProtoReflect.Descriptor instead.
func (*DeviceDataGetRes) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceDataGetRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceDataGetRes) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeviceDataGetRes) GetData() []*common.DeviceData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DeviceDataGetRes) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeviceDataPushReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1. 设备编号
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`
}

func (x *DeviceDataPushReq) Reset() {
	*x = DeviceDataPushReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDataPushReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDataPushReq) ProtoMessage() {}

func (x *DeviceDataPushReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDataPushReq.ProtoReflect.Descriptor instead.
func (*DeviceDataPushReq) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceDataPushReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type DeviceDataPushRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1. 设备名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// 2. 设备编号
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	// 3. 设备数据
	Data []*common.DataDetail `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *DeviceDataPushRes) Reset() {
	*x = DeviceDataPushRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDataPushRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDataPushRes) ProtoMessage() {}

func (x *DeviceDataPushRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDataPushRes.ProtoReflect.Descriptor instead.
func (*DeviceDataPushRes) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceDataPushRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceDataPushRes) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeviceDataPushRes) GetData() []*common.DataDetail {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeviceDataCurveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1. 用户ID
	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	// 2. 设备编号
	DeviceCode string `protobuf:"bytes,2,opt,name=deviceCode,proto3" json:"deviceCode"`
	// 3. 时间间隔
	Interval IntervalType `protobuf:"varint,3,opt,name=interval,proto3,enum=api.gateway.v1.IntervalType" json:"interval"`
	// 4. 图类型
	ChartType []ChartType `protobuf:"varint,4,rep,packed,name=chartType,proto3,enum=api.gateway.v1.ChartType" json:"chartType"`
}

func (x *DeviceDataCurveReq) Reset() {
	*x = DeviceDataCurveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDataCurveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDataCurveReq) ProtoMessage() {}

func (x *DeviceDataCurveReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDataCurveReq.ProtoReflect.Descriptor instead.
func (*DeviceDataCurveReq) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceDataCurveReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeviceDataCurveReq) GetDeviceCode() string {
	if x != nil {
		return x.DeviceCode
	}
	return ""
}

func (x *DeviceDataCurveReq) GetInterval() IntervalType {
	if x != nil {
		return x.Interval
	}
	return IntervalType_Null
}

func (x *DeviceDataCurveReq) GetChartType() []ChartType {
	if x != nil {
		return x.ChartType
	}
	return nil
}

type DataLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//1. 名称
	Key []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key"`
	//2. 值
	Value []float32 `protobuf:"fixed32,2,rep,packed,name=value,proto3" json:"value"`
}

func (x *DataLine) Reset() {
	*x = DataLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataLine) ProtoMessage() {}

func (x *DataLine) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataLine.ProtoReflect.Descriptor instead.
func (*DataLine) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{5}
}

func (x *DataLine) GetKey() []string {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *DataLine) GetValue() []float32 {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeviceDataCurveRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//1. 近一个月数据量
	Month *DataLine `protobuf:"bytes,1,opt,name=month,proto3" json:"month"`
	//2. 饼图
	Pip *DataLine `protobuf:"bytes,2,opt,name=pip,proto3" json:"pip"`
	//3. 折线图
	Line *DataLine `protobuf:"bytes,3,opt,name=line,proto3" json:"line"`
	//4. 电流
	Elect *DataLine `protobuf:"bytes,4,opt,name=elect,proto3" json:"elect"`
	//5. 电压
	Volt *DataLine `protobuf:"bytes,5,opt,name=volt,proto3" json:"volt"`
}

func (x *DeviceDataCurveRes) Reset() {
	*x = DeviceDataCurveRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceDataCurveRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceDataCurveRes) ProtoMessage() {}

func (x *DeviceDataCurveRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_dataVisualization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceDataCurveRes.ProtoReflect.Descriptor instead.
func (*DeviceDataCurveRes) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_dataVisualization_proto_rawDescGZIP(), []int{6}
}

func (x *DeviceDataCurveRes) GetMonth() *DataLine {
	if x != nil {
		return x.Month
	}
	return nil
}

func (x *DeviceDataCurveRes) GetPip() *DataLine {
	if x != nil {
		return x.Pip
	}
	return nil
}

func (x *DeviceDataCurveRes) GetLine() *DataLine {
	if x != nil {
		return x.Line
	}
	return nil
}

func (x *DeviceDataCurveRes) GetElect() *DataLine {
	if x != nil {
		return x.Elect
	}
	return nil
}

func (x *DeviceDataCurveRes) GetVolt() *DataLine {
	if x != nil {
		return x.Volt
	}
	return nil
}

var File_api_gateway_v1_dataVisualization_proto protoreflect.FileDescriptor

var file_api_gateway_v1_dataVisualization_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01,
	0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x78, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x11,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x63, 0x0a, 0x11, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbf, 0x01, 0x0a, 0x12, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x75, 0x72, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x32, 0x0a, 0x08,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xfc, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43,
	0x75, 0x72, 0x76, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x6e, 0x65,
	0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x03, 0x70, 0x69, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x03,
	0x70, 0x69, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x05, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x12, 0x2c, 0x0a, 0x04, 0x76, 0x6f, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x76, 0x6f, 0x6c, 0x74, 0x2a,
	0x36, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x61, 0x79,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x65, 0x65, 0x6b, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x10, 0x03, 0x2a, 0x31, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x69, 0x70, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x65, 0x10, 0x02, 0x32, 0xd9, 0x04, 0x0a, 0x11, 0x44,
	0x61, 0x74, 0x61, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0xc4, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x47,
	0x65, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x22, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x22, 0x27,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x56, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3a, 0x12, 0x12, 0xe6,
	0x9f, 0xa5, 0xe8, 0xaf, 0xa2, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0xe6, 0x95, 0xb0, 0xe6, 0x8d,
	0xae, 0x1a, 0x24, 0xe6, 0x9c, 0xac, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0xe4, 0xb8, 0xba, 0xe6,
	0x9f, 0xa5, 0xe8, 0xaf, 0xa2, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0xe6, 0x95, 0xb0, 0xe6, 0x8d,
	0xae, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x12, 0xb4, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x75, 0x73, 0x68, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x22, 0x5c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x50, 0x75, 0x73,
	0x68, 0x92, 0x41, 0x3a, 0x12, 0x12, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0xe6, 0x95, 0xb0, 0xe6,
	0x8d, 0xae, 0xe6, 0x8e, 0xa8, 0xe9, 0x80, 0x81, 0x1a, 0x24, 0xe6, 0x9c, 0xac, 0xe6, 0x8e, 0xa5,
	0xe5, 0x8f, 0xa3, 0xe4, 0xb8, 0xba, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0xe6, 0x95, 0xb0, 0xe6,
	0x8d, 0xae, 0xe6, 0x8e, 0xa8, 0xe9, 0x80, 0x81, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x12, 0xc5,
	0x01, 0x0a, 0x0f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x75, 0x72,
	0x76, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x75,
	0x72, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x75, 0x72, 0x76, 0x65, 0x52, 0x65, 0x73, 0x22, 0x6a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x21, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x92, 0x41, 0x40, 0x12, 0x15, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0xe6, 0x95,
	0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0x8f, 0xaf, 0xe8, 0xa7, 0x86, 0xe5, 0x8c, 0x96, 0x1a, 0x27, 0xe6,
	0x9c, 0xac, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0xe4, 0xb8, 0xba, 0xe8, 0xae, 0xbe, 0xe5, 0xa4,
	0x87, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0x8f, 0xaf, 0xe8, 0xa7, 0x86, 0xe5, 0x8c, 0x96,
	0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x42, 0xc2, 0x01, 0x5a, 0x19, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x92, 0x41, 0xa3, 0x01, 0x12, 0x1a, 0x0a, 0x13, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x70, 0x69, 0x20, 0x64, 0x6f, 0x63, 0x73, 0x32,
	0x03, 0x32, 0x2e, 0x30, 0x1a, 0x0e, 0x31, 0x32, 0x37, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x3a,
	0x38, 0x30, 0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x4d, 0x0a, 0x4b,
	0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x3d, 0x08, 0x02,
	0x12, 0x28, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x61, 0x20, 0x22, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x2d, 0x6a, 0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x20, 0x74, 0x6f, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_gateway_v1_dataVisualization_proto_rawDescOnce sync.Once
	file_api_gateway_v1_dataVisualization_proto_rawDescData = file_api_gateway_v1_dataVisualization_proto_rawDesc
)

func file_api_gateway_v1_dataVisualization_proto_rawDescGZIP() []byte {
	file_api_gateway_v1_dataVisualization_proto_rawDescOnce.Do(func() {
		file_api_gateway_v1_dataVisualization_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gateway_v1_dataVisualization_proto_rawDescData)
	})
	return file_api_gateway_v1_dataVisualization_proto_rawDescData
}

var file_api_gateway_v1_dataVisualization_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_gateway_v1_dataVisualization_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_gateway_v1_dataVisualization_proto_goTypes = []interface{}{
	(IntervalType)(0),          // 0: api.gateway.v1.IntervalType
	(ChartType)(0),             // 1: api.gateway.v1.ChartType
	(*DeviceDataGetReq)(nil),   // 2: api.gateway.v1.DeviceDataGetReq
	(*DeviceDataGetRes)(nil),   // 3: api.gateway.v1.DeviceDataGetRes
	(*DeviceDataPushReq)(nil),  // 4: api.gateway.v1.DeviceDataPushReq
	(*DeviceDataPushRes)(nil),  // 5: api.gateway.v1.DeviceDataPushRes
	(*DeviceDataCurveReq)(nil), // 6: api.gateway.v1.DeviceDataCurveReq
	(*DataLine)(nil),           // 7: api.gateway.v1.DataLine
	(*DeviceDataCurveRes)(nil), // 8: api.gateway.v1.DeviceDataCurveRes
	(*common.DeviceData)(nil),  // 9: common.DeviceData
	(*common.DataDetail)(nil),  // 10: common.DataDetail
}
var file_api_gateway_v1_dataVisualization_proto_depIdxs = []int32{
	9,  // 0: api.gateway.v1.DeviceDataGetRes.data:type_name -> common.DeviceData
	10, // 1: api.gateway.v1.DeviceDataPushRes.data:type_name -> common.DataDetail
	0,  // 2: api.gateway.v1.DeviceDataCurveReq.interval:type_name -> api.gateway.v1.IntervalType
	1,  // 3: api.gateway.v1.DeviceDataCurveReq.chartType:type_name -> api.gateway.v1.ChartType
	7,  // 4: api.gateway.v1.DeviceDataCurveRes.month:type_name -> api.gateway.v1.DataLine
	7,  // 5: api.gateway.v1.DeviceDataCurveRes.pip:type_name -> api.gateway.v1.DataLine
	7,  // 6: api.gateway.v1.DeviceDataCurveRes.line:type_name -> api.gateway.v1.DataLine
	7,  // 7: api.gateway.v1.DeviceDataCurveRes.elect:type_name -> api.gateway.v1.DataLine
	7,  // 8: api.gateway.v1.DeviceDataCurveRes.volt:type_name -> api.gateway.v1.DataLine
	2,  // 9: api.gateway.v1.DataVisualization.DeviceDataGet:input_type -> api.gateway.v1.DeviceDataGetReq
	4,  // 10: api.gateway.v1.DataVisualization.DeviceDataPush:input_type -> api.gateway.v1.DeviceDataPushReq
	6,  // 11: api.gateway.v1.DataVisualization.DeviceDataCurve:input_type -> api.gateway.v1.DeviceDataCurveReq
	3,  // 12: api.gateway.v1.DataVisualization.DeviceDataGet:output_type -> api.gateway.v1.DeviceDataGetRes
	5,  // 13: api.gateway.v1.DataVisualization.DeviceDataPush:output_type -> api.gateway.v1.DeviceDataPushRes
	8,  // 14: api.gateway.v1.DataVisualization.DeviceDataCurve:output_type -> api.gateway.v1.DeviceDataCurveRes
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_gateway_v1_dataVisualization_proto_init() }
func file_api_gateway_v1_dataVisualization_proto_init() {
	if File_api_gateway_v1_dataVisualization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gateway_v1_dataVisualization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDataGetReq); i {
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
		file_api_gateway_v1_dataVisualization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDataGetRes); i {
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
		file_api_gateway_v1_dataVisualization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDataPushReq); i {
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
		file_api_gateway_v1_dataVisualization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDataPushRes); i {
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
		file_api_gateway_v1_dataVisualization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDataCurveReq); i {
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
		file_api_gateway_v1_dataVisualization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataLine); i {
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
		file_api_gateway_v1_dataVisualization_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceDataCurveRes); i {
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
			RawDescriptor: file_api_gateway_v1_dataVisualization_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_gateway_v1_dataVisualization_proto_goTypes,
		DependencyIndexes: file_api_gateway_v1_dataVisualization_proto_depIdxs,
		EnumInfos:         file_api_gateway_v1_dataVisualization_proto_enumTypes,
		MessageInfos:      file_api_gateway_v1_dataVisualization_proto_msgTypes,
	}.Build()
	File_api_gateway_v1_dataVisualization_proto = out.File
	file_api_gateway_v1_dataVisualization_proto_rawDesc = nil
	file_api_gateway_v1_dataVisualization_proto_goTypes = nil
	file_api_gateway_v1_dataVisualization_proto_depIdxs = nil
}