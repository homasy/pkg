// pkg/kafka/config.go

package kafka

import (
	"crypto/tls"
	"time"

	"github.com/IBM/sarama"
)

// KafkaConfig holds configuration for Kafka
type KafkaConfig struct {
	Brokers    []string
	Topic      string
	GroupID    string
	Username   string
	Password   string
	UseTLS     bool
	ClientType string // Added to distinguish between different services
}

// NewKafkaConfig creates a new Kafka configuration
func NewKafkaConfig(brokers []string, topic, groupID, username, password string, useTLS bool, clientType string) *KafkaConfig {
	return &KafkaConfig{
		Brokers:    brokers,
		Topic:      topic,
		GroupID:    groupID,
		Username:   username,
		Password:   password,
		UseTLS:     useTLS,
		ClientType: clientType,
	}
}

// NewSaramaConfig creates a new Sarama configuration
func (c *KafkaConfig) NewSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	
	// Set producer config
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	
	// Set consumer config
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	
	// Set client ID based on client type
	switch c.ClientType {
	case "patient-service":
		config.ClientID = "patient-service"
	case "appointment-service":
		config.ClientID = "appointment-service"
	case "hr-service":
		config.ClientID = "hr-service"
	case "issue-report-service":
		config.ClientID = "issue-report-service"
	case "laboratory-service":
		config.ClientID = "laboratory-service"
	case "user-service":
		config.ClientID = "user-service"
	case "medical-records-service":
		config.ClientID = "medical-records-service"
	case "notification-service":
		config.ClientID = "notification-service"
	case "billing-service":
		config.ClientID = "billing-service"
	case "pharmacy-service":
		config.ClientID = "pharmacy-service"
	case "supply-chain-service":
		config.ClientID = "supply-chain-service"
	case "analytics-service":
		config.ClientID = "analytics-service"
	case "ward-service":
		config.ClientID = "ward-service"
	case "route-permissions-service":
		config.ClientID = "route-permissions-service"
	default:
		config.ClientID = "unknown-service"
	}
	
	// Set version
	config.Version = sarama.V2_8_0_0
	
	// Set timeout
	config.Net.DialTimeout = 10 * time.Second
	config.Net.ReadTimeout = 10 * time.Second
	config.Net.WriteTimeout = 10 * time.Second
	
	// Set authentication if provided
	if c.Username != "" && c.Password != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = c.Username
		config.Net.SASL.Password = c.Password
		config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}
	
	// Set TLS if enabled
	if c.UseTLS {
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	
	return config
}