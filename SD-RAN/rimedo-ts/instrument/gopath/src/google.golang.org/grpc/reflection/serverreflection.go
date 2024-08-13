//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:19
/*
Package reflection implements server reflection service.

The service implemented is defined in:
https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:

	import "google.golang.org/grpc/reflection"

	s := grpc.NewServer()
	pb.RegisterYourOwnServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	s.Serve(lis)
*/
package reflection

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:37
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:37
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:37
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:37
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:37
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:37
)

import (
	"io"
	"sort"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	v1alphagrpc "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	v1alphapb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// GRPCServer is the interface provided by a gRPC server. It is implemented by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:55
// *grpc.Server, but could also be implemented by other concrete types. It acts
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:55
// as a registry, for accumulating the services exposed by the server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:58
type GRPCServer interface {
	grpc.ServiceRegistrar
	ServiceInfoProvider
}

var _ GRPCServer = (*grpc.Server)(nil)

// Register registers the server reflection service on the given gRPC server.
func Register(s GRPCServer) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:66
	_go_fuzz_dep_.CoverTab[194102]++
													svr := NewServer(ServerOptions{Services: s})
													v1alphagrpc.RegisterServerReflectionServer(s, svr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:68
	// _ = "end of CoverTab[194102]"
}

// ServiceInfoProvider is an interface used to retrieve metadata about the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
// services to expose.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
// The reflection service is only interested in the service names, but the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
// signature is this way so that *grpc.Server implements it. So it is okay
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
// for a custom implementation to return zero values for the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
// grpc.ServiceInfo values in the map.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:71
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:83
type ServiceInfoProvider interface {
	GetServiceInfo() map[string]grpc.ServiceInfo
}

// ExtensionResolver is the interface used to query details about extensions.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:87
// This interface is satisfied by protoregistry.GlobalTypes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:87
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:87
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:87
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:87
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:87
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:94
type ExtensionResolver interface {
	protoregistry.ExtensionTypeResolver
	RangeExtensionsByMessage(message protoreflect.FullName, f func(protoreflect.ExtensionType) bool)
}

// ServerOptions represents the options used to construct a reflection server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:99
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:99
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:99
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:99
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:99
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:105
type ServerOptions struct {
	// The source of advertised RPC services. If not specified, the reflection
	// server will report an empty list when asked to list services.
	//
	// This value will typically be a *grpc.Server. But the set of advertised
	// services can be customized by wrapping a *grpc.Server or using an
	// alternate implementation that returns a custom set of service names.
	Services	ServiceInfoProvider
	// Optional resolver used to load descriptors. If not specified,
	// protoregistry.GlobalFiles will be used.
	DescriptorResolver	protodesc.Resolver
	// Optional resolver used to query for known extensions. If not specified,
	// protoregistry.GlobalTypes will be used.
	ExtensionResolver	ExtensionResolver
}

// NewServer returns a reflection server implementation using the given options.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:121
// This can be used to customize behavior of the reflection service. Most usages
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:121
// should prefer to use Register instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:121
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:121
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:121
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:121
// Notice: This function is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:121
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:129
func NewServer(opts ServerOptions) v1alphagrpc.ServerReflectionServer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:129
	_go_fuzz_dep_.CoverTab[194103]++
													if opts.DescriptorResolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:130
		_go_fuzz_dep_.CoverTab[194106]++
														opts.DescriptorResolver = protoregistry.GlobalFiles
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:131
		// _ = "end of CoverTab[194106]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:132
		_go_fuzz_dep_.CoverTab[194107]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:132
		// _ = "end of CoverTab[194107]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:132
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:132
	// _ = "end of CoverTab[194103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:132
	_go_fuzz_dep_.CoverTab[194104]++
													if opts.ExtensionResolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:133
		_go_fuzz_dep_.CoverTab[194108]++
														opts.ExtensionResolver = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:134
		// _ = "end of CoverTab[194108]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:135
		_go_fuzz_dep_.CoverTab[194109]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:135
		// _ = "end of CoverTab[194109]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:135
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:135
	// _ = "end of CoverTab[194104]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:135
	_go_fuzz_dep_.CoverTab[194105]++
													return &serverReflectionServer{
		s:		opts.Services,
		descResolver:	opts.DescriptorResolver,
		extResolver:	opts.ExtensionResolver,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:140
	// _ = "end of CoverTab[194105]"
}

type serverReflectionServer struct {
	v1alphagrpc.UnimplementedServerReflectionServer
	s		ServiceInfoProvider
	descResolver	protodesc.Resolver
	extResolver	ExtensionResolver
}

// fileDescWithDependencies returns a slice of serialized fileDescriptors in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:150
// wire format ([]byte). The fileDescriptors will include fd and all the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:150
// transitive dependencies of fd with names not in sentFileDescriptors.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:153
func (s *serverReflectionServer) fileDescWithDependencies(fd protoreflect.FileDescriptor, sentFileDescriptors map[string]bool) ([][]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:153
	_go_fuzz_dep_.CoverTab[194110]++
													var r [][]byte
													queue := []protoreflect.FileDescriptor{fd}
													for len(queue) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:156
		_go_fuzz_dep_.CoverTab[194112]++
														currentfd := queue[0]
														queue = queue[1:]
														if sent := sentFileDescriptors[currentfd.Path()]; len(r) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:159
			_go_fuzz_dep_.CoverTab[194114]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:159
			return !sent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:159
			// _ = "end of CoverTab[194114]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:159
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:159
			_go_fuzz_dep_.CoverTab[194115]++
															sentFileDescriptors[currentfd.Path()] = true
															fdProto := protodesc.ToFileDescriptorProto(currentfd)
															currentfdEncoded, err := proto.Marshal(fdProto)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:163
				_go_fuzz_dep_.CoverTab[194117]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:164
				// _ = "end of CoverTab[194117]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:165
				_go_fuzz_dep_.CoverTab[194118]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:165
				// _ = "end of CoverTab[194118]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:165
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:165
			// _ = "end of CoverTab[194115]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:165
			_go_fuzz_dep_.CoverTab[194116]++
															r = append(r, currentfdEncoded)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:166
			// _ = "end of CoverTab[194116]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:167
			_go_fuzz_dep_.CoverTab[194119]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:167
			// _ = "end of CoverTab[194119]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:167
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:167
		// _ = "end of CoverTab[194112]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:167
		_go_fuzz_dep_.CoverTab[194113]++
														for i := 0; i < currentfd.Imports().Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:168
			_go_fuzz_dep_.CoverTab[194120]++
															queue = append(queue, currentfd.Imports().Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:169
			// _ = "end of CoverTab[194120]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:170
		// _ = "end of CoverTab[194113]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:171
	// _ = "end of CoverTab[194110]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:171
	_go_fuzz_dep_.CoverTab[194111]++
													return r, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:172
	// _ = "end of CoverTab[194111]"
}

// fileDescEncodingContainingSymbol finds the file descriptor containing the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:175
// given symbol, finds all of its previously unsent transitive dependencies,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:175
// does marshalling on them, and returns the marshalled result. The given symbol
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:175
// can be a type, a service or a method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:179
func (s *serverReflectionServer) fileDescEncodingContainingSymbol(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:179
	_go_fuzz_dep_.CoverTab[194121]++
													d, err := s.descResolver.FindDescriptorByName(protoreflect.FullName(name))
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:181
		_go_fuzz_dep_.CoverTab[194123]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:182
		// _ = "end of CoverTab[194123]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:183
		_go_fuzz_dep_.CoverTab[194124]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:183
		// _ = "end of CoverTab[194124]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:183
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:183
	// _ = "end of CoverTab[194121]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:183
	_go_fuzz_dep_.CoverTab[194122]++
													return s.fileDescWithDependencies(d.ParentFile(), sentFileDescriptors)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:184
	// _ = "end of CoverTab[194122]"
}

// fileDescEncodingContainingExtension finds the file descriptor containing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:187
// given extension, finds all of its previously unsent transitive dependencies,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:187
// does marshalling on them, and returns the marshalled result.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:190
func (s *serverReflectionServer) fileDescEncodingContainingExtension(typeName string, extNum int32, sentFileDescriptors map[string]bool) ([][]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:190
	_go_fuzz_dep_.CoverTab[194125]++
													xt, err := s.extResolver.FindExtensionByNumber(protoreflect.FullName(typeName), protoreflect.FieldNumber(extNum))
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:192
		_go_fuzz_dep_.CoverTab[194127]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:193
		// _ = "end of CoverTab[194127]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:194
		_go_fuzz_dep_.CoverTab[194128]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:194
		// _ = "end of CoverTab[194128]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:194
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:194
	// _ = "end of CoverTab[194125]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:194
	_go_fuzz_dep_.CoverTab[194126]++
													return s.fileDescWithDependencies(xt.TypeDescriptor().ParentFile(), sentFileDescriptors)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:195
	// _ = "end of CoverTab[194126]"
}

// allExtensionNumbersForTypeName returns all extension numbers for the given type.
func (s *serverReflectionServer) allExtensionNumbersForTypeName(name string) ([]int32, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:199
	_go_fuzz_dep_.CoverTab[194129]++
													var numbers []int32
													s.extResolver.RangeExtensionsByMessage(protoreflect.FullName(name), func(xt protoreflect.ExtensionType) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:201
		_go_fuzz_dep_.CoverTab[194133]++
														numbers = append(numbers, int32(xt.TypeDescriptor().Number()))
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:203
		// _ = "end of CoverTab[194133]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:204
	// _ = "end of CoverTab[194129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:204
	_go_fuzz_dep_.CoverTab[194130]++
													sort.Slice(numbers, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:205
		_go_fuzz_dep_.CoverTab[194134]++
														return numbers[i] < numbers[j]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:206
		// _ = "end of CoverTab[194134]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:207
	// _ = "end of CoverTab[194130]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:207
	_go_fuzz_dep_.CoverTab[194131]++
													if len(numbers) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:208
		_go_fuzz_dep_.CoverTab[194135]++

														if _, err := s.descResolver.FindDescriptorByName(protoreflect.FullName(name)); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:210
			_go_fuzz_dep_.CoverTab[194136]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:211
			// _ = "end of CoverTab[194136]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:212
			_go_fuzz_dep_.CoverTab[194137]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:212
			// _ = "end of CoverTab[194137]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:212
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:212
		// _ = "end of CoverTab[194135]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:213
		_go_fuzz_dep_.CoverTab[194138]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:213
		// _ = "end of CoverTab[194138]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:213
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:213
	// _ = "end of CoverTab[194131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:213
	_go_fuzz_dep_.CoverTab[194132]++
													return numbers, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:214
	// _ = "end of CoverTab[194132]"
}

// listServices returns the names of services this server exposes.
func (s *serverReflectionServer) listServices() []*v1alphapb.ServiceResponse {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:218
	_go_fuzz_dep_.CoverTab[194139]++
													serviceInfo := s.s.GetServiceInfo()
													resp := make([]*v1alphapb.ServiceResponse, 0, len(serviceInfo))
													for svc := range serviceInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:221
		_go_fuzz_dep_.CoverTab[194142]++
														resp = append(resp, &v1alphapb.ServiceResponse{Name: svc})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:222
		// _ = "end of CoverTab[194142]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:223
	// _ = "end of CoverTab[194139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:223
	_go_fuzz_dep_.CoverTab[194140]++
													sort.Slice(resp, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:224
		_go_fuzz_dep_.CoverTab[194143]++
														return resp[i].Name < resp[j].Name
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:225
		// _ = "end of CoverTab[194143]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:226
	// _ = "end of CoverTab[194140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:226
	_go_fuzz_dep_.CoverTab[194141]++
													return resp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:227
	// _ = "end of CoverTab[194141]"
}

// ServerReflectionInfo is the reflection service handler.
func (s *serverReflectionServer) ServerReflectionInfo(stream v1alphagrpc.ServerReflection_ServerReflectionInfoServer) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:231
	_go_fuzz_dep_.CoverTab[194144]++
													sentFileDescriptors := make(map[string]bool)
													for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:233
		_go_fuzz_dep_.CoverTab[194145]++
														in, err := stream.Recv()
														if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:235
			_go_fuzz_dep_.CoverTab[194149]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:236
			// _ = "end of CoverTab[194149]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:237
			_go_fuzz_dep_.CoverTab[194150]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:237
			// _ = "end of CoverTab[194150]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:237
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:237
		// _ = "end of CoverTab[194145]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:237
		_go_fuzz_dep_.CoverTab[194146]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:238
			_go_fuzz_dep_.CoverTab[194151]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:239
			// _ = "end of CoverTab[194151]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:240
			_go_fuzz_dep_.CoverTab[194152]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:240
			// _ = "end of CoverTab[194152]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:240
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:240
		// _ = "end of CoverTab[194146]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:240
		_go_fuzz_dep_.CoverTab[194147]++

														out := &v1alphapb.ServerReflectionResponse{
			ValidHost:		in.Host,
			OriginalRequest:	in,
		}
		switch req := in.MessageRequest.(type) {
		case *v1alphapb.ServerReflectionRequest_FileByFilename:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:247
			_go_fuzz_dep_.CoverTab[194153]++
															var b [][]byte
															fd, err := s.descResolver.FindFileByPath(req.FileByFilename)
															if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:250
				_go_fuzz_dep_.CoverTab[194160]++
																b, err = s.fileDescWithDependencies(fd, sentFileDescriptors)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:251
				// _ = "end of CoverTab[194160]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:252
				_go_fuzz_dep_.CoverTab[194161]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:252
				// _ = "end of CoverTab[194161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:252
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:252
			// _ = "end of CoverTab[194153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:252
			_go_fuzz_dep_.CoverTab[194154]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:253
				_go_fuzz_dep_.CoverTab[194162]++
																out.MessageResponse = &v1alphapb.ServerReflectionResponse_ErrorResponse{
					ErrorResponse: &v1alphapb.ErrorResponse{
						ErrorCode:	int32(codes.NotFound),
						ErrorMessage:	err.Error(),
					},
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:259
				// _ = "end of CoverTab[194162]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:260
				_go_fuzz_dep_.CoverTab[194163]++
																out.MessageResponse = &v1alphapb.ServerReflectionResponse_FileDescriptorResponse{
					FileDescriptorResponse: &v1alphapb.FileDescriptorResponse{FileDescriptorProto: b},
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:263
				// _ = "end of CoverTab[194163]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:264
			// _ = "end of CoverTab[194154]"
		case *v1alphapb.ServerReflectionRequest_FileContainingSymbol:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:265
			_go_fuzz_dep_.CoverTab[194155]++
															b, err := s.fileDescEncodingContainingSymbol(req.FileContainingSymbol, sentFileDescriptors)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:267
				_go_fuzz_dep_.CoverTab[194164]++
																out.MessageResponse = &v1alphapb.ServerReflectionResponse_ErrorResponse{
					ErrorResponse: &v1alphapb.ErrorResponse{
						ErrorCode:	int32(codes.NotFound),
						ErrorMessage:	err.Error(),
					},
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:273
				// _ = "end of CoverTab[194164]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:274
				_go_fuzz_dep_.CoverTab[194165]++
																out.MessageResponse = &v1alphapb.ServerReflectionResponse_FileDescriptorResponse{
					FileDescriptorResponse: &v1alphapb.FileDescriptorResponse{FileDescriptorProto: b},
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:277
				// _ = "end of CoverTab[194165]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:278
			// _ = "end of CoverTab[194155]"
		case *v1alphapb.ServerReflectionRequest_FileContainingExtension:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:279
			_go_fuzz_dep_.CoverTab[194156]++
															typeName := req.FileContainingExtension.ContainingType
															extNum := req.FileContainingExtension.ExtensionNumber
															b, err := s.fileDescEncodingContainingExtension(typeName, extNum, sentFileDescriptors)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:283
				_go_fuzz_dep_.CoverTab[194166]++
																out.MessageResponse = &v1alphapb.ServerReflectionResponse_ErrorResponse{
					ErrorResponse: &v1alphapb.ErrorResponse{
						ErrorCode:	int32(codes.NotFound),
						ErrorMessage:	err.Error(),
					},
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:289
				// _ = "end of CoverTab[194166]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:290
				_go_fuzz_dep_.CoverTab[194167]++
																out.MessageResponse = &v1alphapb.ServerReflectionResponse_FileDescriptorResponse{
					FileDescriptorResponse: &v1alphapb.FileDescriptorResponse{FileDescriptorProto: b},
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:293
				// _ = "end of CoverTab[194167]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:294
			// _ = "end of CoverTab[194156]"
		case *v1alphapb.ServerReflectionRequest_AllExtensionNumbersOfType:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:295
			_go_fuzz_dep_.CoverTab[194157]++
															extNums, err := s.allExtensionNumbersForTypeName(req.AllExtensionNumbersOfType)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:297
				_go_fuzz_dep_.CoverTab[194168]++
																out.MessageResponse = &v1alphapb.ServerReflectionResponse_ErrorResponse{
					ErrorResponse: &v1alphapb.ErrorResponse{
						ErrorCode:	int32(codes.NotFound),
						ErrorMessage:	err.Error(),
					},
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:303
				// _ = "end of CoverTab[194168]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:304
				_go_fuzz_dep_.CoverTab[194169]++
																out.MessageResponse = &v1alphapb.ServerReflectionResponse_AllExtensionNumbersResponse{
					AllExtensionNumbersResponse: &v1alphapb.ExtensionNumberResponse{
						BaseTypeName:		req.AllExtensionNumbersOfType,
						ExtensionNumber:	extNums,
					},
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:310
				// _ = "end of CoverTab[194169]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:311
			// _ = "end of CoverTab[194157]"
		case *v1alphapb.ServerReflectionRequest_ListServices:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:312
			_go_fuzz_dep_.CoverTab[194158]++
															out.MessageResponse = &v1alphapb.ServerReflectionResponse_ListServicesResponse{
				ListServicesResponse: &v1alphapb.ListServiceResponse{
					Service: s.listServices(),
				},
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:317
			// _ = "end of CoverTab[194158]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:318
			_go_fuzz_dep_.CoverTab[194159]++
															return status.Errorf(codes.InvalidArgument, "invalid MessageRequest: %v", in.MessageRequest)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:319
			// _ = "end of CoverTab[194159]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:320
		// _ = "end of CoverTab[194147]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:320
		_go_fuzz_dep_.CoverTab[194148]++

														if err := stream.Send(out); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:322
			_go_fuzz_dep_.CoverTab[194170]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:323
			// _ = "end of CoverTab[194170]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:324
			_go_fuzz_dep_.CoverTab[194171]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:324
			// _ = "end of CoverTab[194171]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:324
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:324
		// _ = "end of CoverTab[194148]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:325
	// _ = "end of CoverTab[194144]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:326
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/reflection/serverreflection.go:326
var _ = _go_fuzz_dep_.CoverTab
