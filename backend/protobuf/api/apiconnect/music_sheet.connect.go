// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/music_sheet.proto

package apiconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	api "github.com/furu2revival/musicbox/protobuf/api"
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
	// MusicSheetServiceName is the fully-qualified name of the MusicSheetService service.
	MusicSheetServiceName = "api.MusicSheetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MusicSheetServiceGetV1Procedure is the fully-qualified name of the MusicSheetService's GetV1 RPC.
	MusicSheetServiceGetV1Procedure = "/api.MusicSheetService/GetV1"
	// MusicSheetServiceCreateV1Procedure is the fully-qualified name of the MusicSheetService's
	// CreateV1 RPC.
	MusicSheetServiceCreateV1Procedure = "/api.MusicSheetService/CreateV1"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	musicSheetServiceServiceDescriptor        = api.File_api_music_sheet_proto.Services().ByName("MusicSheetService")
	musicSheetServiceGetV1MethodDescriptor    = musicSheetServiceServiceDescriptor.Methods().ByName("GetV1")
	musicSheetServiceCreateV1MethodDescriptor = musicSheetServiceServiceDescriptor.Methods().ByName("CreateV1")
)

// MusicSheetServiceClient is a client for the api.MusicSheetService service.
type MusicSheetServiceClient interface {
	GetV1(context.Context, *connect.Request[api.MusicSheetServiceGetV1Request]) (*connect.Response[api.MusicSheetServiceGetV1Response], error)
	CreateV1(context.Context, *connect.Request[api.MusicSheetServiceCreateV1Request]) (*connect.Response[api.MusicSheetServiceCreateV1Response], error)
}

// NewMusicSheetServiceClient constructs a client for the api.MusicSheetService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMusicSheetServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MusicSheetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &musicSheetServiceClient{
		getV1: connect.NewClient[api.MusicSheetServiceGetV1Request, api.MusicSheetServiceGetV1Response](
			httpClient,
			baseURL+MusicSheetServiceGetV1Procedure,
			connect.WithSchema(musicSheetServiceGetV1MethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createV1: connect.NewClient[api.MusicSheetServiceCreateV1Request, api.MusicSheetServiceCreateV1Response](
			httpClient,
			baseURL+MusicSheetServiceCreateV1Procedure,
			connect.WithSchema(musicSheetServiceCreateV1MethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// musicSheetServiceClient implements MusicSheetServiceClient.
type musicSheetServiceClient struct {
	getV1    *connect.Client[api.MusicSheetServiceGetV1Request, api.MusicSheetServiceGetV1Response]
	createV1 *connect.Client[api.MusicSheetServiceCreateV1Request, api.MusicSheetServiceCreateV1Response]
}

// GetV1 calls api.MusicSheetService.GetV1.
func (c *musicSheetServiceClient) GetV1(ctx context.Context, req *connect.Request[api.MusicSheetServiceGetV1Request]) (*connect.Response[api.MusicSheetServiceGetV1Response], error) {
	return c.getV1.CallUnary(ctx, req)
}

// CreateV1 calls api.MusicSheetService.CreateV1.
func (c *musicSheetServiceClient) CreateV1(ctx context.Context, req *connect.Request[api.MusicSheetServiceCreateV1Request]) (*connect.Response[api.MusicSheetServiceCreateV1Response], error) {
	return c.createV1.CallUnary(ctx, req)
}

// MusicSheetServiceHandler is an implementation of the api.MusicSheetService service.
type MusicSheetServiceHandler interface {
	GetV1(context.Context, *connect.Request[api.MusicSheetServiceGetV1Request]) (*connect.Response[api.MusicSheetServiceGetV1Response], error)
	CreateV1(context.Context, *connect.Request[api.MusicSheetServiceCreateV1Request]) (*connect.Response[api.MusicSheetServiceCreateV1Response], error)
}

// NewMusicSheetServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMusicSheetServiceHandler(svc MusicSheetServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	musicSheetServiceGetV1Handler := connect.NewUnaryHandler(
		MusicSheetServiceGetV1Procedure,
		svc.GetV1,
		connect.WithSchema(musicSheetServiceGetV1MethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	musicSheetServiceCreateV1Handler := connect.NewUnaryHandler(
		MusicSheetServiceCreateV1Procedure,
		svc.CreateV1,
		connect.WithSchema(musicSheetServiceCreateV1MethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.MusicSheetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MusicSheetServiceGetV1Procedure:
			musicSheetServiceGetV1Handler.ServeHTTP(w, r)
		case MusicSheetServiceCreateV1Procedure:
			musicSheetServiceCreateV1Handler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMusicSheetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMusicSheetServiceHandler struct{}

func (UnimplementedMusicSheetServiceHandler) GetV1(context.Context, *connect.Request[api.MusicSheetServiceGetV1Request]) (*connect.Response[api.MusicSheetServiceGetV1Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.MusicSheetService.GetV1 is not implemented"))
}

func (UnimplementedMusicSheetServiceHandler) CreateV1(context.Context, *connect.Request[api.MusicSheetServiceCreateV1Request]) (*connect.Response[api.MusicSheetServiceCreateV1Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.MusicSheetService.CreateV1 is not implemented"))
}