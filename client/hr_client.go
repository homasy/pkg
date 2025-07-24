// pkg/client/hr_client.go

package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"strconv"

	hrpb "homasy-backend/services/human-resource-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// HRClient is a client for the HR service
type HRClient struct {
	client     hrpb.HRServiceClient
	conn       *grpc.ClientConn
	serverAddr string
	mu         sync.Mutex
	connected  bool
}

// NewPatientClient creates a new hr client
func NewHRClient(serverAddr string) *HRClient {
	return &HRClient{
		serverAddr: serverAddr,
	}
}

// Connect connects to the hr service
func (c *HRClient) Connect() error {
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
		
		log.Printf("Failed to connect to HR service (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
		retryDelay *= 2 // Exponential backoff
	}
	
	if err != nil {
		return fmt.Errorf("failed to connect to HR service after %d attempts: %v", maxRetries, err)
	}
	
	c.conn = conn
	c.client = hrpb.NewHRServiceClient(conn)
	c.connected = true
	
	log.Printf("Connected to staff service at %s", c.serverAddr)
	return nil
}

// Close closes the connection to the HR service
func (c *HRClient) CloseHRConnection() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection to hr service: %v", err)
	}

	c.connected = false
	return nil
}


// GetStaff gets a staff by ID
func (c *HRClient) GetStaff(ctx context.Context, staffID string) (*hrpb.GetStaffResponse, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	staffIDInt, err := strconv.Atoi(staffID)
	if err != nil {
		return nil, fmt.Errorf("invalid staffID: %v", err)
	}
	req := &hrpb.GetStaffRequest{
		StaffId: int32(staffIDInt),
	}
	

	return c.client.GetStaff(ctx, req)
}

// CheckStaffExists checks if a staff exists
func (c *HRClient) CheckStaffExists(ctx context.Context, staffID string) (bool, error) {
	staff, err := c.GetStaff(ctx, staffID)
	if err != nil {
		return false, err
	}

	return staff != nil, nil
}