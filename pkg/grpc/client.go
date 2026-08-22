package grpc

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/adm87/onyx-server/pkg/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type Client struct {
	cfg  *config.GrpcConfig
	log  *zap.Logger
	conn *grpc.ClientConn
}

func NewClient(cfg *config.GrpcConfig, log *zap.Logger) *Client {
	return &Client{
		cfg: cfg,
		log: log,
	}
}

func (c *Client) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *Client) Connect(opts ...grpc.DialOption) error {
	addr := net.JoinHostPort(c.cfg.Host, strconv.Itoa(c.cfg.Port))
	c.log.Info("Connecting to gRPC server", zap.String("address", addr))

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		c.log.Error("Failed to connect to gRPC server", zap.String("address", addr), zap.Error(err))
		return err
	}

	conn.Connect()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.ConnectTimeoutSeconds)*time.Second)
	defer cancel()

	for {
		state := conn.GetState()
		if state == connectivity.Ready {
			break
		}
		if !conn.WaitForStateChange(ctx, state) {
			conn.Close()

			c.log.Error("Timeout waiting for gRPC connection to become ready", zap.String("address", addr))
			return context.DeadlineExceeded
		}
	}

	c.conn = conn
	c.log.Info("Successfully connected to gRPC server", zap.String("address", addr))
	return nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
