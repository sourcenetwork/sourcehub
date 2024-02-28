// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sourcehub/acp/tx.proto

package acp

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

const (
	Msg_UpdateParams_FullMethodName       = "/sourcehub.acp.Msg/UpdateParams"
	Msg_CreatePolicy_FullMethodName       = "/sourcehub.acp.Msg/CreatePolicy"
	Msg_SetRelationship_FullMethodName    = "/sourcehub.acp.Msg/SetRelationship"
	Msg_DeleteRelationship_FullMethodName = "/sourcehub.acp.Msg/DeleteRelationship"
	Msg_RegisterObject_FullMethodName     = "/sourcehub.acp.Msg/RegisterObject"
	Msg_UnregisterObject_FullMethodName   = "/sourcehub.acp.Msg/UnregisterObject"
	Msg_CheckAccess_FullMethodName        = "/sourcehub.acp.Msg/CheckAccess"
	Msg_PolicyCmd_FullMethodName          = "/sourcehub.acp.Msg/PolicyCmd"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// CreatePolicy adds a new Policy to SourceHub.
	// The Policy models an aplication's high level access control rules.
	CreatePolicy(ctx context.Context, in *MsgCreatePolicy, opts ...grpc.CallOption) (*MsgCreatePolicyResponse, error)
	// SetRelationship creates or updates a Relationship within a Policy
	// A Relationship is a statement which ties together an object and a subjecto with a "relation",
	// which means the set of high level rules defined in the Policy will apply to these entities.
	SetRelationship(ctx context.Context, in *MsgSetRelationship, opts ...grpc.CallOption) (*MsgSetRelationshipResponse, error)
	// DelereRelationship removes a Relationship from a Policy.
	// If the Relationship was not found in a Policy, this Msg is a no-op.
	DeleteRelationship(ctx context.Context, in *MsgDeleteRelationship, opts ...grpc.CallOption) (*MsgDeleteRelationshipResponse, error)
	// Attempting to register a previously registered Object is an error,
	// Object IDs are therefore assumed to be unique within a Policy.
	RegisterObject(ctx context.Context, in *MsgRegisterObject, opts ...grpc.CallOption) (*MsgRegisterObjectResponse, error)
	// Suppose Bob owns object Foo, which is shared with Bob but not Eve.
	// Eve wants to access Foo but was not given permission to, they could "hijack" Bob's object by waiting for Bob to Unregister Foo,
	// then submitting a RegisterObject Msg, effectively becoming Foo's new owner.
	// If Charlie has a copy of the object, Eve could convince Charlie to share his copy, granting Eve access to Foo.
	// The previous scenario where an unauthorized user is able to claim ownership to data previously unaccessible to them
	// is an "ownership hijack".
	UnregisterObject(ctx context.Context, in *MsgUnregisterObject, opts ...grpc.CallOption) (*MsgUnregisterObjectResponse, error)
	// The resulting evaluation is used to generate a cryptographic proof that the given Access Request
	// was valid at a particular block height.
	CheckAccess(ctx context.Context, in *MsgCheckAccess, opts ...grpc.CallOption) (*MsgCheckAccessResponse, error)
	PolicyCmd(ctx context.Context, in *MsgPolicyCmd, opts ...grpc.CallOption) (*MsgPolicyCmdResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePolicy(ctx context.Context, in *MsgCreatePolicy, opts ...grpc.CallOption) (*MsgCreatePolicyResponse, error) {
	out := new(MsgCreatePolicyResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetRelationship(ctx context.Context, in *MsgSetRelationship, opts ...grpc.CallOption) (*MsgSetRelationshipResponse, error) {
	out := new(MsgSetRelationshipResponse)
	err := c.cc.Invoke(ctx, Msg_SetRelationship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteRelationship(ctx context.Context, in *MsgDeleteRelationship, opts ...grpc.CallOption) (*MsgDeleteRelationshipResponse, error) {
	out := new(MsgDeleteRelationshipResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteRelationship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterObject(ctx context.Context, in *MsgRegisterObject, opts ...grpc.CallOption) (*MsgRegisterObjectResponse, error) {
	out := new(MsgRegisterObjectResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnregisterObject(ctx context.Context, in *MsgUnregisterObject, opts ...grpc.CallOption) (*MsgUnregisterObjectResponse, error) {
	out := new(MsgUnregisterObjectResponse)
	err := c.cc.Invoke(ctx, Msg_UnregisterObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CheckAccess(ctx context.Context, in *MsgCheckAccess, opts ...grpc.CallOption) (*MsgCheckAccessResponse, error) {
	out := new(MsgCheckAccessResponse)
	err := c.cc.Invoke(ctx, Msg_CheckAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PolicyCmd(ctx context.Context, in *MsgPolicyCmd, opts ...grpc.CallOption) (*MsgPolicyCmdResponse, error) {
	out := new(MsgPolicyCmdResponse)
	err := c.cc.Invoke(ctx, Msg_PolicyCmd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// CreatePolicy adds a new Policy to SourceHub.
	// The Policy models an aplication's high level access control rules.
	CreatePolicy(context.Context, *MsgCreatePolicy) (*MsgCreatePolicyResponse, error)
	// SetRelationship creates or updates a Relationship within a Policy
	// A Relationship is a statement which ties together an object and a subjecto with a "relation",
	// which means the set of high level rules defined in the Policy will apply to these entities.
	SetRelationship(context.Context, *MsgSetRelationship) (*MsgSetRelationshipResponse, error)
	// DelereRelationship removes a Relationship from a Policy.
	// If the Relationship was not found in a Policy, this Msg is a no-op.
	DeleteRelationship(context.Context, *MsgDeleteRelationship) (*MsgDeleteRelationshipResponse, error)
	// Attempting to register a previously registered Object is an error,
	// Object IDs are therefore assumed to be unique within a Policy.
	RegisterObject(context.Context, *MsgRegisterObject) (*MsgRegisterObjectResponse, error)
	// Suppose Bob owns object Foo, which is shared with Bob but not Eve.
	// Eve wants to access Foo but was not given permission to, they could "hijack" Bob's object by waiting for Bob to Unregister Foo,
	// then submitting a RegisterObject Msg, effectively becoming Foo's new owner.
	// If Charlie has a copy of the object, Eve could convince Charlie to share his copy, granting Eve access to Foo.
	// The previous scenario where an unauthorized user is able to claim ownership to data previously unaccessible to them
	// is an "ownership hijack".
	UnregisterObject(context.Context, *MsgUnregisterObject) (*MsgUnregisterObjectResponse, error)
	// The resulting evaluation is used to generate a cryptographic proof that the given Access Request
	// was valid at a particular block height.
	CheckAccess(context.Context, *MsgCheckAccess) (*MsgCheckAccessResponse, error)
	PolicyCmd(context.Context, *MsgPolicyCmd) (*MsgPolicyCmdResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreatePolicy(context.Context, *MsgCreatePolicy) (*MsgCreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedMsgServer) SetRelationship(context.Context, *MsgSetRelationship) (*MsgSetRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRelationship not implemented")
}
func (UnimplementedMsgServer) DeleteRelationship(context.Context, *MsgDeleteRelationship) (*MsgDeleteRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelationship not implemented")
}
func (UnimplementedMsgServer) RegisterObject(context.Context, *MsgRegisterObject) (*MsgRegisterObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterObject not implemented")
}
func (UnimplementedMsgServer) UnregisterObject(context.Context, *MsgUnregisterObject) (*MsgUnregisterObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterObject not implemented")
}
func (UnimplementedMsgServer) CheckAccess(context.Context, *MsgCheckAccess) (*MsgCheckAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccess not implemented")
}
func (UnimplementedMsgServer) PolicyCmd(context.Context, *MsgPolicyCmd) (*MsgPolicyCmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyCmd not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePolicy(ctx, req.(*MsgCreatePolicy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SetRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetRelationship(ctx, req.(*MsgSetRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteRelationship(ctx, req.(*MsgDeleteRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterObject(ctx, req.(*MsgRegisterObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnregisterObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnregisterObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnregisterObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UnregisterObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnregisterObject(ctx, req.(*MsgUnregisterObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CheckAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCheckAccess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CheckAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CheckAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CheckAccess(ctx, req.(*MsgCheckAccess))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PolicyCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPolicyCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PolicyCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_PolicyCmd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PolicyCmd(ctx, req.(*MsgPolicyCmd))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sourcehub.acp.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _Msg_CreatePolicy_Handler,
		},
		{
			MethodName: "SetRelationship",
			Handler:    _Msg_SetRelationship_Handler,
		},
		{
			MethodName: "DeleteRelationship",
			Handler:    _Msg_DeleteRelationship_Handler,
		},
		{
			MethodName: "RegisterObject",
			Handler:    _Msg_RegisterObject_Handler,
		},
		{
			MethodName: "UnregisterObject",
			Handler:    _Msg_UnregisterObject_Handler,
		},
		{
			MethodName: "CheckAccess",
			Handler:    _Msg_CheckAccess_Handler,
		},
		{
			MethodName: "PolicyCmd",
			Handler:    _Msg_PolicyCmd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sourcehub/acp/tx.proto",
}
