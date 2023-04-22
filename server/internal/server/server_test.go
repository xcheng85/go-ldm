package server

import (
	"context"
	_ "io/ioutil"
	"net"
	"testing"
	"github.com/stretchr/testify/require"
	pb "github.com/xcheng85/go-ldm/server/api/v1"
	"google.golang.org/grpc"
)

func TestServer(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		client pb.LdmClient,
	){
		"write Tile succeeeds":
			testWriteTile,
	} {
		t.Run(scenario, func(t *testing.T) {
			// setupTest initialize and return the client
			client, teardown := setupTest(t, nil)
			defer teardown()
			fn(t, client)
		})
	}
}

// func as a first element, return multiple
func setupTest(t *testing.T, fn func()) (
	client pb.LdmClient,
	teardown func(),
) {
	t.Helper()
	// grpc server
	l, err := net.Listen("tcp", ":0")

	require.NoError(t, err)
	// grpc client connecting to server
	clientOptions := []grpc.DialOption{grpc.WithInsecure()}
	// grpc client connection
	cc, err := grpc.Dial(l.Addr().String(), clientOptions...)
	require.NoError(t, err)

	// dir, err := ioutil.TempDir("", "server-test")
	// require.NoError(t, err)

	// clog, err := log.NewLog(dir, log.Config{})
	// require.NoError(t, err)

	// config = &Config{
	// 	CommitLog: clog,
	// }
	if fn != nil {
		fn()
	}
	server, err := NewGRPCServer()
	require.NoError(t, err)
	// grpc server serve goroutine
	go func() {
		server.Serve(l)
	}()

	client = pb.NewLdmClient(cc)

	return client, func() {
		// stop the grpc server
		server.Stop()
		// stop grpc client connection
		cc.Close()
		// stop tcp
		l.Close()
		// clog.Remove()
	}
}

func testWriteTile(t *testing.T, client pb.LdmClient) {
	ctx := context.Background()

	tile := &pb.Tile{
		Value: []byte{},
		Offset: 10,
	}

	response, err := client.WriteTile(
		ctx,
		&pb.WriteTileRequest{
			Tile: tile,
		},
	)
	require.NoError(t, err)
	require.Equal(t, response.Offset, uint64(0))

	// consume, err := client.Consume(ctx, &api.ConsumeRequest{
	// 	Offset: produce.Offset,
	// })
	// require.NoError(t, err)
	// require.Equal(t, want.Value, consume.Record.Value)
	// require.Equal(t, want.Offset, consume.Record.Offset)
}

