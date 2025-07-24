// pkg/client/billing_client.go

package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	billingpb "homasy/billing-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// BillingClient is a client for the Billing service
type BillingClient struct {
	client     billingpb.BillingServiceClient
	conn       *grpc.ClientConn
	serverAddr string
	mu         sync.Mutex
	connected  bool
}

// NewBillingClient creates a new billing client
func NewBillingClient(serverAddr string) *BillingClient {
	return &BillingClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the billing service
func (c *BillingClient) Connect() error {
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
		
		log.Printf("Failed to connect to Billing service (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2 // Exponential backoff
	}
	
	if err != nil {
		return fmt.Errorf("failed to connect to Billing service after %d attempts: %v", maxRetries, err)
	}
	
	c.conn = conn
	c.client = billingpb.NewBillingServiceClient(conn)
	c.connected = true
	
	log.Printf("Connected to billing service at %s", c.serverAddr)
	return nil
}

// Close closes the connection to the Billing service
func (c *BillingClient) CloseBillingConnection() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection to billing service: %v", err)
	}

	c.connected = false
	return nil
}


func (c *BillingClient) CreateServiceRecord(ctx context.Context, req *billingpb.CreateServiceRecordRequest) (*billingpb.CreateServiceRecordResponse, error) {
    if err := c.Connect(); err != nil {
        return nil, err
    }
    return c.client.CreateServiceRecord(ctx, req)
}

func (c *BillingClient) UpdateServiceRecord(ctx context.Context, req *billingpb.UpdateServiceRecordStatusRequest) (*billingpb.UpdateServiceRecordStatusResponse, error) {
    if err := c.Connect(); err != nil {
        return nil, err
    }
    return c.client.UpdateServiceRecordStatus(ctx, req)
}