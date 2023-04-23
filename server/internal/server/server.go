package server

import (
	"context"
	pb "github.com/xcheng85/go-ldm/server/api/v1"
	"google.golang.org/grpc"
)

// dependencies injection for LDMServer
type LDMManager interface {
	Write(*pb.Tile) (uint64, error)
	Read(offset uint64) (*pb.Tile, error)
}

// server is used to implement the interface in ldm_grpc.pb.go
// // LdmServer is the server API for Ldm service.
// // All implementations must embed UnimplementedLdmServer
// // for forward compatibility
// type LdmServer interface {
// 	WriteTile(context.Context, *WriteTileRequest) (*WriteTileResponse, error)
// 	ReadTile(context.Context, *ReadTileRequest) (*ReadTileResponse, error)
// 	ReadTileStream(*ReadTileRequest, Ldm_ReadTileStreamServer) error
// 	WriteTileStream(Ldm_WriteTileStreamServer) error
// 	mustEmbedUnimplementedLdmServer()
// }
type LDMServer struct {
	// embedded field type cannot be a pointer to an interface
	// *LDMManager
	LDMManager
	// https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
	pb.UnimplementedLdmServer
}

func (s *LDMServer) WriteTile(ctx context.Context, request *pb.WriteTileRequest) (*pb.WriteTileResponse, error) {
	// to do
	return &pb.WriteTileResponse{Offset: 0}, nil
}

func (s *LDMServer) ReadTile(ctx context.Context, request *pb.ReadTileRequest) (*pb.ReadTileResponse, error) {
	// to do
	return &pb.ReadTileResponse{Tile: nil}, nil
}

func (s *LDMServer) WriteTileStream(stream pb.Ldm_WriteTileStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		res, err := s.WriteTile(stream.Context(), req)
		if err != nil {
			return err
		}
		if err = stream.Send(res); err != nil {
			return err
		}
	}
}

func (s *LDMServer) ReadTileStream(request *pb.ReadTileRequest, stream pb.Ldm_ReadTileStreamServer) error {
	for {
		// A select blocks until one of its cases can run
		// Mostly default branch, default is not provided it will block
		// Done() return a channel to listen <-chan struct{}
		// user cancels, first branch will be hit
		select {
		case <-stream.Context().Done():
			return nil
		default:
			res, err := s.ReadTile(stream.Context(), request)
			// type assertion
			switch err.(type) {
			case nil:
			//  case api.ErrOffsetOutOfRange:
			// 	continue
			default:
				return err
			}
			if err = stream.Send(res); err != nil {
				return err
			}
			request.Offset++
		}
	}
}

func newLDMServer(ldm LDMManager) (ldmServer *LDMServer, err error) {
	ldmServer = &LDMServer{
		LDMManager: ldm,
	}
	return ldmServer, nil
}

func NewGRPCServer(ldm LDMManager) (*grpc.Server, error) {
	s := grpc.NewServer()
	ldmServer, err := newLDMServer(ldm)
	if err != nil {
		return nil, err
	}
	pb.RegisterLdmServer(s, ldmServer)
	return s, nil 
}
