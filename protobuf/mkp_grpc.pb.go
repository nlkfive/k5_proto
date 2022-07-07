// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: proto/mkp.proto

package protobufpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MkpSvcClient is the client API for MkpSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MkpSvcClient interface {
	TransferErc721(ctx context.Context, in *TransferErc721Req, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TransferErc20(ctx context.Context, in *TransferErc20Req, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateAuction(ctx context.Context, in *CreateAuctionReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateAuction(ctx context.Context, in *UpdateAuctionReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelAuction(ctx context.Context, in *CancelAuctionReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateBidding(ctx context.Context, in *CreateBiddingReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevealBidding(ctx context.Context, in *RevealBiddingReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateOrUpdateNft(ctx context.Context, in *CreateOrUpdateNftReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mkpSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewMkpSvcClient(cc grpc.ClientConnInterface) MkpSvcClient {
	return &mkpSvcClient{cc}
}

func (c *mkpSvcClient) TransferErc721(ctx context.Context, in *TransferErc721Req, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/TransferErc721", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) TransferErc20(ctx context.Context, in *TransferErc20Req, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/TransferErc20", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) CreateAuction(ctx context.Context, in *CreateAuctionReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/CreateAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) UpdateAuction(ctx context.Context, in *UpdateAuctionReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/UpdateAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) CancelAuction(ctx context.Context, in *CancelAuctionReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/CancelAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) CreateBidding(ctx context.Context, in *CreateBiddingReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/CreateBidding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) RevealBidding(ctx context.Context, in *RevealBiddingReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/RevealBidding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) CreateTransaction(ctx context.Context, in *CreateTransactionReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mkpSvcClient) CreateOrUpdateNft(ctx context.Context, in *CreateOrUpdateNftReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.MkpSvc/CreateOrUpdateNft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MkpSvcServer is the server API for MkpSvc service.
// All implementations must embed UnimplementedMkpSvcServer
// for forward compatibility
type MkpSvcServer interface {
	TransferErc721(context.Context, *TransferErc721Req) (*emptypb.Empty, error)
	TransferErc20(context.Context, *TransferErc20Req) (*emptypb.Empty, error)
	CreateAuction(context.Context, *CreateAuctionReq) (*emptypb.Empty, error)
	UpdateAuction(context.Context, *UpdateAuctionReq) (*emptypb.Empty, error)
	CancelAuction(context.Context, *CancelAuctionReq) (*emptypb.Empty, error)
	CreateBidding(context.Context, *CreateBiddingReq) (*emptypb.Empty, error)
	RevealBidding(context.Context, *RevealBiddingReq) (*emptypb.Empty, error)
	CreateOrder(context.Context, *CreateOrderReq) (*emptypb.Empty, error)
	UpdateOrder(context.Context, *UpdateOrderReq) (*emptypb.Empty, error)
	CancelOrder(context.Context, *CancelOrderReq) (*emptypb.Empty, error)
	CreateTransaction(context.Context, *CreateTransactionReq) (*emptypb.Empty, error)
	CreateOrUpdateNft(context.Context, *CreateOrUpdateNftReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedMkpSvcServer()
}

// UnimplementedMkpSvcServer must be embedded to have forward compatible implementations.
type UnimplementedMkpSvcServer struct {
}

func (UnimplementedMkpSvcServer) TransferErc721(context.Context, *TransferErc721Req) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferErc721 not implemented")
}
func (UnimplementedMkpSvcServer) TransferErc20(context.Context, *TransferErc20Req) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferErc20 not implemented")
}
func (UnimplementedMkpSvcServer) CreateAuction(context.Context, *CreateAuctionReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuction not implemented")
}
func (UnimplementedMkpSvcServer) UpdateAuction(context.Context, *UpdateAuctionReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuction not implemented")
}
func (UnimplementedMkpSvcServer) CancelAuction(context.Context, *CancelAuctionReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAuction not implemented")
}
func (UnimplementedMkpSvcServer) CreateBidding(context.Context, *CreateBiddingReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBidding not implemented")
}
func (UnimplementedMkpSvcServer) RevealBidding(context.Context, *RevealBiddingReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevealBidding not implemented")
}
func (UnimplementedMkpSvcServer) CreateOrder(context.Context, *CreateOrderReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedMkpSvcServer) UpdateOrder(context.Context, *UpdateOrderReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedMkpSvcServer) CancelOrder(context.Context, *CancelOrderReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedMkpSvcServer) CreateTransaction(context.Context, *CreateTransactionReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedMkpSvcServer) CreateOrUpdateNft(context.Context, *CreateOrUpdateNftReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateNft not implemented")
}
func (UnimplementedMkpSvcServer) mustEmbedUnimplementedMkpSvcServer() {}

// UnsafeMkpSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MkpSvcServer will
// result in compilation errors.
type UnsafeMkpSvcServer interface {
	mustEmbedUnimplementedMkpSvcServer()
}

func RegisterMkpSvcServer(s grpc.ServiceRegistrar, srv MkpSvcServer) {
	s.RegisterService(&MkpSvc_ServiceDesc, srv)
}

func _MkpSvc_TransferErc721_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferErc721Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).TransferErc721(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/TransferErc721",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).TransferErc721(ctx, req.(*TransferErc721Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_TransferErc20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferErc20Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).TransferErc20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/TransferErc20",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).TransferErc20(ctx, req.(*TransferErc20Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_CreateAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuctionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).CreateAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/CreateAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).CreateAuction(ctx, req.(*CreateAuctionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_UpdateAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuctionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).UpdateAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/UpdateAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).UpdateAuction(ctx, req.(*UpdateAuctionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_CancelAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelAuctionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).CancelAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/CancelAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).CancelAuction(ctx, req.(*CancelAuctionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_CreateBidding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBiddingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).CreateBidding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/CreateBidding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).CreateBidding(ctx, req.(*CreateBiddingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_RevealBidding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevealBiddingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).RevealBidding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/RevealBidding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).RevealBidding(ctx, req.(*RevealBiddingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).UpdateOrder(ctx, req.(*UpdateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).CancelOrder(ctx, req.(*CancelOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).CreateTransaction(ctx, req.(*CreateTransactionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MkpSvc_CreateOrUpdateNft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateNftReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MkpSvcServer).CreateOrUpdateNft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MkpSvc/CreateOrUpdateNft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MkpSvcServer).CreateOrUpdateNft(ctx, req.(*CreateOrUpdateNftReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MkpSvc_ServiceDesc is the grpc.ServiceDesc for MkpSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MkpSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.MkpSvc",
	HandlerType: (*MkpSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransferErc721",
			Handler:    _MkpSvc_TransferErc721_Handler,
		},
		{
			MethodName: "TransferErc20",
			Handler:    _MkpSvc_TransferErc20_Handler,
		},
		{
			MethodName: "CreateAuction",
			Handler:    _MkpSvc_CreateAuction_Handler,
		},
		{
			MethodName: "UpdateAuction",
			Handler:    _MkpSvc_UpdateAuction_Handler,
		},
		{
			MethodName: "CancelAuction",
			Handler:    _MkpSvc_CancelAuction_Handler,
		},
		{
			MethodName: "CreateBidding",
			Handler:    _MkpSvc_CreateBidding_Handler,
		},
		{
			MethodName: "RevealBidding",
			Handler:    _MkpSvc_RevealBidding_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _MkpSvc_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _MkpSvc_UpdateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _MkpSvc_CancelOrder_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _MkpSvc_CreateTransaction_Handler,
		},
		{
			MethodName: "CreateOrUpdateNft",
			Handler:    _MkpSvc_CreateOrUpdateNft_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mkp.proto",
}
