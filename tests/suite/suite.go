package suite

import (
	"context"
	"net"
	"sso/internal/config"
	"strconv"
	"testing"

	ssov1 "github.com/r0mbeg/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc client

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient // grpc-client for auth service (not sso)
}

const (
	grpcHost = "localhost"
)

// test migration:
// go run ./cmd/migrator/main.go --storage-path=./storage/sso.db --migrations-path=./migrations --migrations-table=migrations_test
func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath("../config/local.yaml")

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(context.Background(), // client connection
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials())) // using insecure connection for tests
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: ssov1.NewAuthClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
