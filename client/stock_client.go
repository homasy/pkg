// pkg/client/user_client.go

package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	supplypb "github.com/homasy/pkg/shared/supply-chain-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// StockClient is a client for the supply chain service
type StockClient struct {
	client     supplypb.SupplyChainServiceClient
	conn       *grpc.ClientConn
	serverAddr string
	mu         sync.Mutex
	connected  bool
}

// NewStockClient creates a new stock client
func NewStockClient(serverAddr string) *StockClient {
	return &StockClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the Supply Chain service
func (c *StockClient) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connected {
		return nil
	}

	// Set up connection with retry
	var err error
	var conn *grpc.ClientConn

	// Retry options
	maxRetries := 5
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		// Connect with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err = grpc.DialContext(
			ctx,
			c.serverAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)

		if err == nil {
			break
		}

		log.Printf("Failed to connect to supply chain service (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2 // Exponential backoff
	}

	if err != nil {
		return fmt.Errorf("failed to connect to supply chain service after %d attempts: %v", maxRetries, err)
	}

	c.conn = conn
	c.client = supplypb.NewSupplyChainServiceClient(conn)
	c.connected = true

	log.Printf("Connected to supply chain service at %s", c.serverAddr)
	return nil
}

// Close closes the connection to the supply chain service
func (c *StockClient) CloseSupplyChainConnection() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection to supply chain service: %v", err)
	}

	c.connected = false
	return nil
}

// Reduce Stock Item
func (c *StockClient) ReduceStockItem(ctx context.Context, req *supplypb.ReduceStockItemRequest) (*supplypb.ReduceStockItemResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.ReduceStockItem(ctx, req)
}

// ListStockItems lists stock items
func (c *StockClient) ListStockItems(ctx context.Context, req *supplypb.ListStockItemsRequest) (*supplypb.ListStockItemsResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	return c.client.ListStockItems(ctx, req)
}

// GetMedicalScheme retrieves a medical scheme by ID
func (c *StockClient) GetMedicalScheme(ctx context.Context, id int32) (*supplypb.MedicalScheme, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}

	req := &supplypb.GetMedicalSchemeRequest{Id: id}
	resp, err := c.client.GetMedicalScheme(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get medical scheme: %v", err)
	}

	return resp.GetMedicalScheme(), nil
}
