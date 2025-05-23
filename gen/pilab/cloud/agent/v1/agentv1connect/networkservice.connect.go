// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pilab/cloud/agent/v1/networkservice.proto

package agentv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "go.pilab.hu/cloud/virtpb/gen/pilab/cloud/agent/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// NetworkServiceName is the fully-qualified name of the NetworkService service.
	NetworkServiceName = "pilab.cloud.agent.v1.NetworkService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// NetworkServiceOvsBridgeCreateProcedure is the fully-qualified name of the NetworkService's
	// OvsBridgeCreate RPC.
	NetworkServiceOvsBridgeCreateProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsBridgeCreate"
	// NetworkServiceOvsBridgeDeleteProcedure is the fully-qualified name of the NetworkService's
	// OvsBridgeDelete RPC.
	NetworkServiceOvsBridgeDeleteProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsBridgeDelete"
	// NetworkServiceOvsBridgeListProcedure is the fully-qualified name of the NetworkService's
	// OvsBridgeList RPC.
	NetworkServiceOvsBridgeListProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsBridgeList"
	// NetworkServiceOvsBridgeGetProcedure is the fully-qualified name of the NetworkService's
	// OvsBridgeGet RPC.
	NetworkServiceOvsBridgeGetProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsBridgeGet"
	// NetworkServiceOvsPortCreateProcedure is the fully-qualified name of the NetworkService's
	// OvsPortCreate RPC.
	NetworkServiceOvsPortCreateProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsPortCreate"
	// NetworkServiceOvsPortDeleteProcedure is the fully-qualified name of the NetworkService's
	// OvsPortDelete RPC.
	NetworkServiceOvsPortDeleteProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsPortDelete"
	// NetworkServiceOvsPortListProcedure is the fully-qualified name of the NetworkService's
	// OvsPortList RPC.
	NetworkServiceOvsPortListProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsPortList"
	// NetworkServiceOvsPortGetProcedure is the fully-qualified name of the NetworkService's OvsPortGet
	// RPC.
	NetworkServiceOvsPortGetProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsPortGet"
	// NetworkServiceOvsPortUpdateProcedure is the fully-qualified name of the NetworkService's
	// OvsPortUpdate RPC.
	NetworkServiceOvsPortUpdateProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsPortUpdate"
	// NetworkServiceOvsQosRuleCreateProcedure is the fully-qualified name of the NetworkService's
	// OvsQosRuleCreate RPC.
	NetworkServiceOvsQosRuleCreateProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsQosRuleCreate"
	// NetworkServiceOvsQosRuleDeleteProcedure is the fully-qualified name of the NetworkService's
	// OvsQosRuleDelete RPC.
	NetworkServiceOvsQosRuleDeleteProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsQosRuleDelete"
	// NetworkServiceOvsQosRuleListProcedure is the fully-qualified name of the NetworkService's
	// OvsQosRuleList RPC.
	NetworkServiceOvsQosRuleListProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsQosRuleList"
	// NetworkServiceOvsQosRuleGetProcedure is the fully-qualified name of the NetworkService's
	// OvsQosRuleGet RPC.
	NetworkServiceOvsQosRuleGetProcedure = "/pilab.cloud.agent.v1.NetworkService/OvsQosRuleGet"
	// NetworkServiceNetworkCreateProcedure is the fully-qualified name of the NetworkService's
	// NetworkCreate RPC.
	NetworkServiceNetworkCreateProcedure = "/pilab.cloud.agent.v1.NetworkService/NetworkCreate"
	// NetworkServiceNetworkDeleteProcedure is the fully-qualified name of the NetworkService's
	// NetworkDelete RPC.
	NetworkServiceNetworkDeleteProcedure = "/pilab.cloud.agent.v1.NetworkService/NetworkDelete"
	// NetworkServiceNetworkListProcedure is the fully-qualified name of the NetworkService's
	// NetworkList RPC.
	NetworkServiceNetworkListProcedure = "/pilab.cloud.agent.v1.NetworkService/NetworkList"
	// NetworkServiceNetworkGetProcedure is the fully-qualified name of the NetworkService's NetworkGet
	// RPC.
	NetworkServiceNetworkGetProcedure = "/pilab.cloud.agent.v1.NetworkService/NetworkGet"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	networkServiceServiceDescriptor                = v1.File_pilab_cloud_agent_v1_networkservice_proto.Services().ByName("NetworkService")
	networkServiceOvsBridgeCreateMethodDescriptor  = networkServiceServiceDescriptor.Methods().ByName("OvsBridgeCreate")
	networkServiceOvsBridgeDeleteMethodDescriptor  = networkServiceServiceDescriptor.Methods().ByName("OvsBridgeDelete")
	networkServiceOvsBridgeListMethodDescriptor    = networkServiceServiceDescriptor.Methods().ByName("OvsBridgeList")
	networkServiceOvsBridgeGetMethodDescriptor     = networkServiceServiceDescriptor.Methods().ByName("OvsBridgeGet")
	networkServiceOvsPortCreateMethodDescriptor    = networkServiceServiceDescriptor.Methods().ByName("OvsPortCreate")
	networkServiceOvsPortDeleteMethodDescriptor    = networkServiceServiceDescriptor.Methods().ByName("OvsPortDelete")
	networkServiceOvsPortListMethodDescriptor      = networkServiceServiceDescriptor.Methods().ByName("OvsPortList")
	networkServiceOvsPortGetMethodDescriptor       = networkServiceServiceDescriptor.Methods().ByName("OvsPortGet")
	networkServiceOvsPortUpdateMethodDescriptor    = networkServiceServiceDescriptor.Methods().ByName("OvsPortUpdate")
	networkServiceOvsQosRuleCreateMethodDescriptor = networkServiceServiceDescriptor.Methods().ByName("OvsQosRuleCreate")
	networkServiceOvsQosRuleDeleteMethodDescriptor = networkServiceServiceDescriptor.Methods().ByName("OvsQosRuleDelete")
	networkServiceOvsQosRuleListMethodDescriptor   = networkServiceServiceDescriptor.Methods().ByName("OvsQosRuleList")
	networkServiceOvsQosRuleGetMethodDescriptor    = networkServiceServiceDescriptor.Methods().ByName("OvsQosRuleGet")
	networkServiceNetworkCreateMethodDescriptor    = networkServiceServiceDescriptor.Methods().ByName("NetworkCreate")
	networkServiceNetworkDeleteMethodDescriptor    = networkServiceServiceDescriptor.Methods().ByName("NetworkDelete")
	networkServiceNetworkListMethodDescriptor      = networkServiceServiceDescriptor.Methods().ByName("NetworkList")
	networkServiceNetworkGetMethodDescriptor       = networkServiceServiceDescriptor.Methods().ByName("NetworkGet")
)

// NetworkServiceClient is a client for the pilab.cloud.agent.v1.NetworkService service.
type NetworkServiceClient interface {
	// OVS Bridge Management
	OvsBridgeCreate(context.Context, *connect.Request[v1.OvsBridgeCreateRequest]) (*connect.Response[v1.OvsBridgeCreateResponse], error)
	OvsBridgeDelete(context.Context, *connect.Request[v1.OvsBridgeDeleteRequest]) (*connect.Response[v1.OvsBridgeDeleteResponse], error)
	OvsBridgeList(context.Context, *connect.Request[v1.OvsBridgeListRequest]) (*connect.Response[v1.OvsBridgeListResponse], error)
	OvsBridgeGet(context.Context, *connect.Request[v1.OvsBridgeGetRequest]) (*connect.Response[v1.OvsBridgeGetResponse], error)
	// OVS Port Management
	OvsPortCreate(context.Context, *connect.Request[v1.OvsPortCreateRequest]) (*connect.Response[v1.OvsPortCreateResponse], error)
	OvsPortDelete(context.Context, *connect.Request[v1.OvsPortDeleteRequest]) (*connect.Response[v1.OvsPortDeleteResponse], error)
	OvsPortList(context.Context, *connect.Request[v1.OvsPortListRequest]) (*connect.Response[v1.OvsPortListResponse], error)
	OvsPortGet(context.Context, *connect.Request[v1.OvsPortGetRequest]) (*connect.Response[v1.OvsPortGetResponse], error)
	OvsPortUpdate(context.Context, *connect.Request[v1.OvsPortUpdateRequest]) (*connect.Response[v1.OvsPortUpdateResponse], error)
	// OVS QoS Management
	OvsQosRuleCreate(context.Context, *connect.Request[v1.OvsQosRuleCreateRequest]) (*connect.Response[v1.OvsQosRuleCreateResponse], error)
	OvsQosRuleDelete(context.Context, *connect.Request[v1.OvsQosRuleDeleteRequest]) (*connect.Response[v1.OvsQosRuleDeleteResponse], error)
	OvsQosRuleList(context.Context, *connect.Request[v1.OvsQosRuleListRequest]) (*connect.Response[v1.OvsQosRuleListResponse], error)
	OvsQosRuleGet(context.Context, *connect.Request[v1.OvsQosRuleGetRequest]) (*connect.Response[v1.OvsQosRuleGetResponse], error)
	// Libvirt Network Management
	NetworkCreate(context.Context, *connect.Request[v1.NetworkCreateRequest]) (*connect.Response[v1.NetworkCreateResponse], error)
	NetworkDelete(context.Context, *connect.Request[v1.NetworkDeleteRequest]) (*connect.Response[v1.NetworkDeleteResponse], error)
	NetworkList(context.Context, *connect.Request[v1.NetworkListRequest]) (*connect.Response[v1.NetworkListResponse], error)
	NetworkGet(context.Context, *connect.Request[v1.NetworkGetRequest]) (*connect.Response[v1.NetworkGetResponse], error)
}

// NewNetworkServiceClient constructs a client for the pilab.cloud.agent.v1.NetworkService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNetworkServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) NetworkServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &networkServiceClient{
		ovsBridgeCreate: connect.NewClient[v1.OvsBridgeCreateRequest, v1.OvsBridgeCreateResponse](
			httpClient,
			baseURL+NetworkServiceOvsBridgeCreateProcedure,
			connect.WithSchema(networkServiceOvsBridgeCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsBridgeDelete: connect.NewClient[v1.OvsBridgeDeleteRequest, v1.OvsBridgeDeleteResponse](
			httpClient,
			baseURL+NetworkServiceOvsBridgeDeleteProcedure,
			connect.WithSchema(networkServiceOvsBridgeDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsBridgeList: connect.NewClient[v1.OvsBridgeListRequest, v1.OvsBridgeListResponse](
			httpClient,
			baseURL+NetworkServiceOvsBridgeListProcedure,
			connect.WithSchema(networkServiceOvsBridgeListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsBridgeGet: connect.NewClient[v1.OvsBridgeGetRequest, v1.OvsBridgeGetResponse](
			httpClient,
			baseURL+NetworkServiceOvsBridgeGetProcedure,
			connect.WithSchema(networkServiceOvsBridgeGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsPortCreate: connect.NewClient[v1.OvsPortCreateRequest, v1.OvsPortCreateResponse](
			httpClient,
			baseURL+NetworkServiceOvsPortCreateProcedure,
			connect.WithSchema(networkServiceOvsPortCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsPortDelete: connect.NewClient[v1.OvsPortDeleteRequest, v1.OvsPortDeleteResponse](
			httpClient,
			baseURL+NetworkServiceOvsPortDeleteProcedure,
			connect.WithSchema(networkServiceOvsPortDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsPortList: connect.NewClient[v1.OvsPortListRequest, v1.OvsPortListResponse](
			httpClient,
			baseURL+NetworkServiceOvsPortListProcedure,
			connect.WithSchema(networkServiceOvsPortListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsPortGet: connect.NewClient[v1.OvsPortGetRequest, v1.OvsPortGetResponse](
			httpClient,
			baseURL+NetworkServiceOvsPortGetProcedure,
			connect.WithSchema(networkServiceOvsPortGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsPortUpdate: connect.NewClient[v1.OvsPortUpdateRequest, v1.OvsPortUpdateResponse](
			httpClient,
			baseURL+NetworkServiceOvsPortUpdateProcedure,
			connect.WithSchema(networkServiceOvsPortUpdateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsQosRuleCreate: connect.NewClient[v1.OvsQosRuleCreateRequest, v1.OvsQosRuleCreateResponse](
			httpClient,
			baseURL+NetworkServiceOvsQosRuleCreateProcedure,
			connect.WithSchema(networkServiceOvsQosRuleCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsQosRuleDelete: connect.NewClient[v1.OvsQosRuleDeleteRequest, v1.OvsQosRuleDeleteResponse](
			httpClient,
			baseURL+NetworkServiceOvsQosRuleDeleteProcedure,
			connect.WithSchema(networkServiceOvsQosRuleDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsQosRuleList: connect.NewClient[v1.OvsQosRuleListRequest, v1.OvsQosRuleListResponse](
			httpClient,
			baseURL+NetworkServiceOvsQosRuleListProcedure,
			connect.WithSchema(networkServiceOvsQosRuleListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		ovsQosRuleGet: connect.NewClient[v1.OvsQosRuleGetRequest, v1.OvsQosRuleGetResponse](
			httpClient,
			baseURL+NetworkServiceOvsQosRuleGetProcedure,
			connect.WithSchema(networkServiceOvsQosRuleGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		networkCreate: connect.NewClient[v1.NetworkCreateRequest, v1.NetworkCreateResponse](
			httpClient,
			baseURL+NetworkServiceNetworkCreateProcedure,
			connect.WithSchema(networkServiceNetworkCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		networkDelete: connect.NewClient[v1.NetworkDeleteRequest, v1.NetworkDeleteResponse](
			httpClient,
			baseURL+NetworkServiceNetworkDeleteProcedure,
			connect.WithSchema(networkServiceNetworkDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		networkList: connect.NewClient[v1.NetworkListRequest, v1.NetworkListResponse](
			httpClient,
			baseURL+NetworkServiceNetworkListProcedure,
			connect.WithSchema(networkServiceNetworkListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		networkGet: connect.NewClient[v1.NetworkGetRequest, v1.NetworkGetResponse](
			httpClient,
			baseURL+NetworkServiceNetworkGetProcedure,
			connect.WithSchema(networkServiceNetworkGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// networkServiceClient implements NetworkServiceClient.
type networkServiceClient struct {
	ovsBridgeCreate  *connect.Client[v1.OvsBridgeCreateRequest, v1.OvsBridgeCreateResponse]
	ovsBridgeDelete  *connect.Client[v1.OvsBridgeDeleteRequest, v1.OvsBridgeDeleteResponse]
	ovsBridgeList    *connect.Client[v1.OvsBridgeListRequest, v1.OvsBridgeListResponse]
	ovsBridgeGet     *connect.Client[v1.OvsBridgeGetRequest, v1.OvsBridgeGetResponse]
	ovsPortCreate    *connect.Client[v1.OvsPortCreateRequest, v1.OvsPortCreateResponse]
	ovsPortDelete    *connect.Client[v1.OvsPortDeleteRequest, v1.OvsPortDeleteResponse]
	ovsPortList      *connect.Client[v1.OvsPortListRequest, v1.OvsPortListResponse]
	ovsPortGet       *connect.Client[v1.OvsPortGetRequest, v1.OvsPortGetResponse]
	ovsPortUpdate    *connect.Client[v1.OvsPortUpdateRequest, v1.OvsPortUpdateResponse]
	ovsQosRuleCreate *connect.Client[v1.OvsQosRuleCreateRequest, v1.OvsQosRuleCreateResponse]
	ovsQosRuleDelete *connect.Client[v1.OvsQosRuleDeleteRequest, v1.OvsQosRuleDeleteResponse]
	ovsQosRuleList   *connect.Client[v1.OvsQosRuleListRequest, v1.OvsQosRuleListResponse]
	ovsQosRuleGet    *connect.Client[v1.OvsQosRuleGetRequest, v1.OvsQosRuleGetResponse]
	networkCreate    *connect.Client[v1.NetworkCreateRequest, v1.NetworkCreateResponse]
	networkDelete    *connect.Client[v1.NetworkDeleteRequest, v1.NetworkDeleteResponse]
	networkList      *connect.Client[v1.NetworkListRequest, v1.NetworkListResponse]
	networkGet       *connect.Client[v1.NetworkGetRequest, v1.NetworkGetResponse]
}

// OvsBridgeCreate calls pilab.cloud.agent.v1.NetworkService.OvsBridgeCreate.
func (c *networkServiceClient) OvsBridgeCreate(ctx context.Context, req *connect.Request[v1.OvsBridgeCreateRequest]) (*connect.Response[v1.OvsBridgeCreateResponse], error) {
	return c.ovsBridgeCreate.CallUnary(ctx, req)
}

// OvsBridgeDelete calls pilab.cloud.agent.v1.NetworkService.OvsBridgeDelete.
func (c *networkServiceClient) OvsBridgeDelete(ctx context.Context, req *connect.Request[v1.OvsBridgeDeleteRequest]) (*connect.Response[v1.OvsBridgeDeleteResponse], error) {
	return c.ovsBridgeDelete.CallUnary(ctx, req)
}

// OvsBridgeList calls pilab.cloud.agent.v1.NetworkService.OvsBridgeList.
func (c *networkServiceClient) OvsBridgeList(ctx context.Context, req *connect.Request[v1.OvsBridgeListRequest]) (*connect.Response[v1.OvsBridgeListResponse], error) {
	return c.ovsBridgeList.CallUnary(ctx, req)
}

// OvsBridgeGet calls pilab.cloud.agent.v1.NetworkService.OvsBridgeGet.
func (c *networkServiceClient) OvsBridgeGet(ctx context.Context, req *connect.Request[v1.OvsBridgeGetRequest]) (*connect.Response[v1.OvsBridgeGetResponse], error) {
	return c.ovsBridgeGet.CallUnary(ctx, req)
}

// OvsPortCreate calls pilab.cloud.agent.v1.NetworkService.OvsPortCreate.
func (c *networkServiceClient) OvsPortCreate(ctx context.Context, req *connect.Request[v1.OvsPortCreateRequest]) (*connect.Response[v1.OvsPortCreateResponse], error) {
	return c.ovsPortCreate.CallUnary(ctx, req)
}

// OvsPortDelete calls pilab.cloud.agent.v1.NetworkService.OvsPortDelete.
func (c *networkServiceClient) OvsPortDelete(ctx context.Context, req *connect.Request[v1.OvsPortDeleteRequest]) (*connect.Response[v1.OvsPortDeleteResponse], error) {
	return c.ovsPortDelete.CallUnary(ctx, req)
}

// OvsPortList calls pilab.cloud.agent.v1.NetworkService.OvsPortList.
func (c *networkServiceClient) OvsPortList(ctx context.Context, req *connect.Request[v1.OvsPortListRequest]) (*connect.Response[v1.OvsPortListResponse], error) {
	return c.ovsPortList.CallUnary(ctx, req)
}

// OvsPortGet calls pilab.cloud.agent.v1.NetworkService.OvsPortGet.
func (c *networkServiceClient) OvsPortGet(ctx context.Context, req *connect.Request[v1.OvsPortGetRequest]) (*connect.Response[v1.OvsPortGetResponse], error) {
	return c.ovsPortGet.CallUnary(ctx, req)
}

// OvsPortUpdate calls pilab.cloud.agent.v1.NetworkService.OvsPortUpdate.
func (c *networkServiceClient) OvsPortUpdate(ctx context.Context, req *connect.Request[v1.OvsPortUpdateRequest]) (*connect.Response[v1.OvsPortUpdateResponse], error) {
	return c.ovsPortUpdate.CallUnary(ctx, req)
}

// OvsQosRuleCreate calls pilab.cloud.agent.v1.NetworkService.OvsQosRuleCreate.
func (c *networkServiceClient) OvsQosRuleCreate(ctx context.Context, req *connect.Request[v1.OvsQosRuleCreateRequest]) (*connect.Response[v1.OvsQosRuleCreateResponse], error) {
	return c.ovsQosRuleCreate.CallUnary(ctx, req)
}

// OvsQosRuleDelete calls pilab.cloud.agent.v1.NetworkService.OvsQosRuleDelete.
func (c *networkServiceClient) OvsQosRuleDelete(ctx context.Context, req *connect.Request[v1.OvsQosRuleDeleteRequest]) (*connect.Response[v1.OvsQosRuleDeleteResponse], error) {
	return c.ovsQosRuleDelete.CallUnary(ctx, req)
}

// OvsQosRuleList calls pilab.cloud.agent.v1.NetworkService.OvsQosRuleList.
func (c *networkServiceClient) OvsQosRuleList(ctx context.Context, req *connect.Request[v1.OvsQosRuleListRequest]) (*connect.Response[v1.OvsQosRuleListResponse], error) {
	return c.ovsQosRuleList.CallUnary(ctx, req)
}

// OvsQosRuleGet calls pilab.cloud.agent.v1.NetworkService.OvsQosRuleGet.
func (c *networkServiceClient) OvsQosRuleGet(ctx context.Context, req *connect.Request[v1.OvsQosRuleGetRequest]) (*connect.Response[v1.OvsQosRuleGetResponse], error) {
	return c.ovsQosRuleGet.CallUnary(ctx, req)
}

// NetworkCreate calls pilab.cloud.agent.v1.NetworkService.NetworkCreate.
func (c *networkServiceClient) NetworkCreate(ctx context.Context, req *connect.Request[v1.NetworkCreateRequest]) (*connect.Response[v1.NetworkCreateResponse], error) {
	return c.networkCreate.CallUnary(ctx, req)
}

// NetworkDelete calls pilab.cloud.agent.v1.NetworkService.NetworkDelete.
func (c *networkServiceClient) NetworkDelete(ctx context.Context, req *connect.Request[v1.NetworkDeleteRequest]) (*connect.Response[v1.NetworkDeleteResponse], error) {
	return c.networkDelete.CallUnary(ctx, req)
}

// NetworkList calls pilab.cloud.agent.v1.NetworkService.NetworkList.
func (c *networkServiceClient) NetworkList(ctx context.Context, req *connect.Request[v1.NetworkListRequest]) (*connect.Response[v1.NetworkListResponse], error) {
	return c.networkList.CallUnary(ctx, req)
}

// NetworkGet calls pilab.cloud.agent.v1.NetworkService.NetworkGet.
func (c *networkServiceClient) NetworkGet(ctx context.Context, req *connect.Request[v1.NetworkGetRequest]) (*connect.Response[v1.NetworkGetResponse], error) {
	return c.networkGet.CallUnary(ctx, req)
}

// NetworkServiceHandler is an implementation of the pilab.cloud.agent.v1.NetworkService service.
type NetworkServiceHandler interface {
	// OVS Bridge Management
	OvsBridgeCreate(context.Context, *connect.Request[v1.OvsBridgeCreateRequest]) (*connect.Response[v1.OvsBridgeCreateResponse], error)
	OvsBridgeDelete(context.Context, *connect.Request[v1.OvsBridgeDeleteRequest]) (*connect.Response[v1.OvsBridgeDeleteResponse], error)
	OvsBridgeList(context.Context, *connect.Request[v1.OvsBridgeListRequest]) (*connect.Response[v1.OvsBridgeListResponse], error)
	OvsBridgeGet(context.Context, *connect.Request[v1.OvsBridgeGetRequest]) (*connect.Response[v1.OvsBridgeGetResponse], error)
	// OVS Port Management
	OvsPortCreate(context.Context, *connect.Request[v1.OvsPortCreateRequest]) (*connect.Response[v1.OvsPortCreateResponse], error)
	OvsPortDelete(context.Context, *connect.Request[v1.OvsPortDeleteRequest]) (*connect.Response[v1.OvsPortDeleteResponse], error)
	OvsPortList(context.Context, *connect.Request[v1.OvsPortListRequest]) (*connect.Response[v1.OvsPortListResponse], error)
	OvsPortGet(context.Context, *connect.Request[v1.OvsPortGetRequest]) (*connect.Response[v1.OvsPortGetResponse], error)
	OvsPortUpdate(context.Context, *connect.Request[v1.OvsPortUpdateRequest]) (*connect.Response[v1.OvsPortUpdateResponse], error)
	// OVS QoS Management
	OvsQosRuleCreate(context.Context, *connect.Request[v1.OvsQosRuleCreateRequest]) (*connect.Response[v1.OvsQosRuleCreateResponse], error)
	OvsQosRuleDelete(context.Context, *connect.Request[v1.OvsQosRuleDeleteRequest]) (*connect.Response[v1.OvsQosRuleDeleteResponse], error)
	OvsQosRuleList(context.Context, *connect.Request[v1.OvsQosRuleListRequest]) (*connect.Response[v1.OvsQosRuleListResponse], error)
	OvsQosRuleGet(context.Context, *connect.Request[v1.OvsQosRuleGetRequest]) (*connect.Response[v1.OvsQosRuleGetResponse], error)
	// Libvirt Network Management
	NetworkCreate(context.Context, *connect.Request[v1.NetworkCreateRequest]) (*connect.Response[v1.NetworkCreateResponse], error)
	NetworkDelete(context.Context, *connect.Request[v1.NetworkDeleteRequest]) (*connect.Response[v1.NetworkDeleteResponse], error)
	NetworkList(context.Context, *connect.Request[v1.NetworkListRequest]) (*connect.Response[v1.NetworkListResponse], error)
	NetworkGet(context.Context, *connect.Request[v1.NetworkGetRequest]) (*connect.Response[v1.NetworkGetResponse], error)
}

// NewNetworkServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNetworkServiceHandler(svc NetworkServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	networkServiceOvsBridgeCreateHandler := connect.NewUnaryHandler(
		NetworkServiceOvsBridgeCreateProcedure,
		svc.OvsBridgeCreate,
		connect.WithSchema(networkServiceOvsBridgeCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsBridgeDeleteHandler := connect.NewUnaryHandler(
		NetworkServiceOvsBridgeDeleteProcedure,
		svc.OvsBridgeDelete,
		connect.WithSchema(networkServiceOvsBridgeDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsBridgeListHandler := connect.NewUnaryHandler(
		NetworkServiceOvsBridgeListProcedure,
		svc.OvsBridgeList,
		connect.WithSchema(networkServiceOvsBridgeListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsBridgeGetHandler := connect.NewUnaryHandler(
		NetworkServiceOvsBridgeGetProcedure,
		svc.OvsBridgeGet,
		connect.WithSchema(networkServiceOvsBridgeGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsPortCreateHandler := connect.NewUnaryHandler(
		NetworkServiceOvsPortCreateProcedure,
		svc.OvsPortCreate,
		connect.WithSchema(networkServiceOvsPortCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsPortDeleteHandler := connect.NewUnaryHandler(
		NetworkServiceOvsPortDeleteProcedure,
		svc.OvsPortDelete,
		connect.WithSchema(networkServiceOvsPortDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsPortListHandler := connect.NewUnaryHandler(
		NetworkServiceOvsPortListProcedure,
		svc.OvsPortList,
		connect.WithSchema(networkServiceOvsPortListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsPortGetHandler := connect.NewUnaryHandler(
		NetworkServiceOvsPortGetProcedure,
		svc.OvsPortGet,
		connect.WithSchema(networkServiceOvsPortGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsPortUpdateHandler := connect.NewUnaryHandler(
		NetworkServiceOvsPortUpdateProcedure,
		svc.OvsPortUpdate,
		connect.WithSchema(networkServiceOvsPortUpdateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsQosRuleCreateHandler := connect.NewUnaryHandler(
		NetworkServiceOvsQosRuleCreateProcedure,
		svc.OvsQosRuleCreate,
		connect.WithSchema(networkServiceOvsQosRuleCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsQosRuleDeleteHandler := connect.NewUnaryHandler(
		NetworkServiceOvsQosRuleDeleteProcedure,
		svc.OvsQosRuleDelete,
		connect.WithSchema(networkServiceOvsQosRuleDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsQosRuleListHandler := connect.NewUnaryHandler(
		NetworkServiceOvsQosRuleListProcedure,
		svc.OvsQosRuleList,
		connect.WithSchema(networkServiceOvsQosRuleListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceOvsQosRuleGetHandler := connect.NewUnaryHandler(
		NetworkServiceOvsQosRuleGetProcedure,
		svc.OvsQosRuleGet,
		connect.WithSchema(networkServiceOvsQosRuleGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceNetworkCreateHandler := connect.NewUnaryHandler(
		NetworkServiceNetworkCreateProcedure,
		svc.NetworkCreate,
		connect.WithSchema(networkServiceNetworkCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceNetworkDeleteHandler := connect.NewUnaryHandler(
		NetworkServiceNetworkDeleteProcedure,
		svc.NetworkDelete,
		connect.WithSchema(networkServiceNetworkDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceNetworkListHandler := connect.NewUnaryHandler(
		NetworkServiceNetworkListProcedure,
		svc.NetworkList,
		connect.WithSchema(networkServiceNetworkListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	networkServiceNetworkGetHandler := connect.NewUnaryHandler(
		NetworkServiceNetworkGetProcedure,
		svc.NetworkGet,
		connect.WithSchema(networkServiceNetworkGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/pilab.cloud.agent.v1.NetworkService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NetworkServiceOvsBridgeCreateProcedure:
			networkServiceOvsBridgeCreateHandler.ServeHTTP(w, r)
		case NetworkServiceOvsBridgeDeleteProcedure:
			networkServiceOvsBridgeDeleteHandler.ServeHTTP(w, r)
		case NetworkServiceOvsBridgeListProcedure:
			networkServiceOvsBridgeListHandler.ServeHTTP(w, r)
		case NetworkServiceOvsBridgeGetProcedure:
			networkServiceOvsBridgeGetHandler.ServeHTTP(w, r)
		case NetworkServiceOvsPortCreateProcedure:
			networkServiceOvsPortCreateHandler.ServeHTTP(w, r)
		case NetworkServiceOvsPortDeleteProcedure:
			networkServiceOvsPortDeleteHandler.ServeHTTP(w, r)
		case NetworkServiceOvsPortListProcedure:
			networkServiceOvsPortListHandler.ServeHTTP(w, r)
		case NetworkServiceOvsPortGetProcedure:
			networkServiceOvsPortGetHandler.ServeHTTP(w, r)
		case NetworkServiceOvsPortUpdateProcedure:
			networkServiceOvsPortUpdateHandler.ServeHTTP(w, r)
		case NetworkServiceOvsQosRuleCreateProcedure:
			networkServiceOvsQosRuleCreateHandler.ServeHTTP(w, r)
		case NetworkServiceOvsQosRuleDeleteProcedure:
			networkServiceOvsQosRuleDeleteHandler.ServeHTTP(w, r)
		case NetworkServiceOvsQosRuleListProcedure:
			networkServiceOvsQosRuleListHandler.ServeHTTP(w, r)
		case NetworkServiceOvsQosRuleGetProcedure:
			networkServiceOvsQosRuleGetHandler.ServeHTTP(w, r)
		case NetworkServiceNetworkCreateProcedure:
			networkServiceNetworkCreateHandler.ServeHTTP(w, r)
		case NetworkServiceNetworkDeleteProcedure:
			networkServiceNetworkDeleteHandler.ServeHTTP(w, r)
		case NetworkServiceNetworkListProcedure:
			networkServiceNetworkListHandler.ServeHTTP(w, r)
		case NetworkServiceNetworkGetProcedure:
			networkServiceNetworkGetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNetworkServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNetworkServiceHandler struct{}

func (UnimplementedNetworkServiceHandler) OvsBridgeCreate(context.Context, *connect.Request[v1.OvsBridgeCreateRequest]) (*connect.Response[v1.OvsBridgeCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsBridgeCreate is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsBridgeDelete(context.Context, *connect.Request[v1.OvsBridgeDeleteRequest]) (*connect.Response[v1.OvsBridgeDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsBridgeDelete is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsBridgeList(context.Context, *connect.Request[v1.OvsBridgeListRequest]) (*connect.Response[v1.OvsBridgeListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsBridgeList is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsBridgeGet(context.Context, *connect.Request[v1.OvsBridgeGetRequest]) (*connect.Response[v1.OvsBridgeGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsBridgeGet is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsPortCreate(context.Context, *connect.Request[v1.OvsPortCreateRequest]) (*connect.Response[v1.OvsPortCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsPortCreate is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsPortDelete(context.Context, *connect.Request[v1.OvsPortDeleteRequest]) (*connect.Response[v1.OvsPortDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsPortDelete is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsPortList(context.Context, *connect.Request[v1.OvsPortListRequest]) (*connect.Response[v1.OvsPortListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsPortList is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsPortGet(context.Context, *connect.Request[v1.OvsPortGetRequest]) (*connect.Response[v1.OvsPortGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsPortGet is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsPortUpdate(context.Context, *connect.Request[v1.OvsPortUpdateRequest]) (*connect.Response[v1.OvsPortUpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsPortUpdate is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsQosRuleCreate(context.Context, *connect.Request[v1.OvsQosRuleCreateRequest]) (*connect.Response[v1.OvsQosRuleCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsQosRuleCreate is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsQosRuleDelete(context.Context, *connect.Request[v1.OvsQosRuleDeleteRequest]) (*connect.Response[v1.OvsQosRuleDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsQosRuleDelete is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsQosRuleList(context.Context, *connect.Request[v1.OvsQosRuleListRequest]) (*connect.Response[v1.OvsQosRuleListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsQosRuleList is not implemented"))
}

func (UnimplementedNetworkServiceHandler) OvsQosRuleGet(context.Context, *connect.Request[v1.OvsQosRuleGetRequest]) (*connect.Response[v1.OvsQosRuleGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.OvsQosRuleGet is not implemented"))
}

func (UnimplementedNetworkServiceHandler) NetworkCreate(context.Context, *connect.Request[v1.NetworkCreateRequest]) (*connect.Response[v1.NetworkCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.NetworkCreate is not implemented"))
}

func (UnimplementedNetworkServiceHandler) NetworkDelete(context.Context, *connect.Request[v1.NetworkDeleteRequest]) (*connect.Response[v1.NetworkDeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.NetworkDelete is not implemented"))
}

func (UnimplementedNetworkServiceHandler) NetworkList(context.Context, *connect.Request[v1.NetworkListRequest]) (*connect.Response[v1.NetworkListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.NetworkList is not implemented"))
}

func (UnimplementedNetworkServiceHandler) NetworkGet(context.Context, *connect.Request[v1.NetworkGetRequest]) (*connect.Response[v1.NetworkGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pilab.cloud.agent.v1.NetworkService.NetworkGet is not implemented"))
}
