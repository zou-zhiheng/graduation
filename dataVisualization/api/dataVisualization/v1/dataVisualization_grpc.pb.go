// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.0
// source: api/dataVisualization/v1/dataVisualization.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataVisualizationClient is the client API for DataVisualization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataVisualizationClient interface {
	//设备数据查询
	DeviceDataGet(ctx context.Context, in *DeviceDataGetReq, opts ...grpc.CallOption) (*DeviceDataGetRes, error)
	//折线图，当天，七天，近一个月
	DeviceDataCurve(ctx context.Context, in *DeviceDataCurveReq, opts ...grpc.CallOption) (*DeviceDataCurveRes, error)
}

type dataVisualizationClient struct {
	cc grpc.ClientConnInterface
}

func NewDataVisualizationClient(cc grpc.ClientConnInterface) DataVisualizationClient {
	return &dataVisualizationClient{cc}
}

func (c *dataVisualizationClient) DeviceDataGet(ctx context.Context, in *DeviceDataGetReq, opts ...grpc.CallOption) (*DeviceDataGetRes, error) {
	out := new(DeviceDataGetRes)
	err := c.cc.Invoke(ctx, "/api.dataVisualization.v1.DataVisualization/DeviceDataGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataVisualizationClient) DeviceDataCurve(ctx context.Context, in *DeviceDataCurveReq, opts ...grpc.CallOption) (*DeviceDataCurveRes, error) {
	out := new(DeviceDataCurveRes)
	err := c.cc.Invoke(ctx, "/api.dataVisualization.v1.DataVisualization/DeviceDataCurve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataVisualizationServer is the server API for DataVisualization service.
// All implementations must embed UnimplementedDataVisualizationServer
// for forward compatibility
type DataVisualizationServer interface {
	//设备数据查询
	DeviceDataGet(context.Context, *DeviceDataGetReq) (*DeviceDataGetRes, error)
	//折线图，当天，七天，近一个月
	DeviceDataCurve(context.Context, *DeviceDataCurveReq) (*DeviceDataCurveRes, error)
	mustEmbedUnimplementedDataVisualizationServer()
}

// UnimplementedDataVisualizationServer must be embedded to have forward compatible implementations.
type UnimplementedDataVisualizationServer struct {
}

func (UnimplementedDataVisualizationServer) DeviceDataGet(context.Context, *DeviceDataGetReq) (*DeviceDataGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceDataGet not implemented")
}
func (UnimplementedDataVisualizationServer) DeviceDataCurve(context.Context, *DeviceDataCurveReq) (*DeviceDataCurveRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceDataCurve not implemented")
}
func (UnimplementedDataVisualizationServer) mustEmbedUnimplementedDataVisualizationServer() {}

// UnsafeDataVisualizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataVisualizationServer will
// result in compilation errors.
type UnsafeDataVisualizationServer interface {
	mustEmbedUnimplementedDataVisualizationServer()
}

func RegisterDataVisualizationServer(s grpc.ServiceRegistrar, srv DataVisualizationServer) {
	s.RegisterService(&DataVisualization_ServiceDesc, srv)
}

func _DataVisualization_DeviceDataGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceDataGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataVisualizationServer).DeviceDataGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.dataVisualization.v1.DataVisualization/DeviceDataGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataVisualizationServer).DeviceDataGet(ctx, req.(*DeviceDataGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataVisualization_DeviceDataCurve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceDataCurveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataVisualizationServer).DeviceDataCurve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.dataVisualization.v1.DataVisualization/DeviceDataCurve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataVisualizationServer).DeviceDataCurve(ctx, req.(*DeviceDataCurveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DataVisualization_ServiceDesc is the grpc.ServiceDesc for DataVisualization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataVisualization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.dataVisualization.v1.DataVisualization",
	HandlerType: (*DataVisualizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceDataGet",
			Handler:    _DataVisualization_DeviceDataGet_Handler,
		},
		{
			MethodName: "DeviceDataCurve",
			Handler:    _DataVisualization_DeviceDataCurve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/dataVisualization/v1/dataVisualization.proto",
}
