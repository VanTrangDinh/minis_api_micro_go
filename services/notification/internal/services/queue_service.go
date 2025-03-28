package services

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type QueueService struct {
	conn *amqp.Connection
}

func NewQueueService(conn *amqp.Connection) *QueueService {
	return &QueueService{
		conn: conn,
	}
}

// PublishNotification publishes a notification to the queue
func (s *QueueService) PublishNotification(notification interface{}) error {
	// Create a channel
	ch, err := s.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		"notifications", // queue name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Convert notification to JSON
	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	// Publish the message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// ConsumeNotifications starts consuming notifications from the queue
func (s *QueueService) ConsumeNotifications(handler func([]byte) error) error {
	// Create a channel
	ch, err := s.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		"notifications", // queue name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Start consuming
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	// Process messages
	go func() {
		for d := range msgs {
			if err := handler(d.Body); err != nil {
				// Log error but continue processing
				fmt.Printf("Error processing message: %v\n", err)
			}
			// Acknowledge message
			d.Ack(false)
		}
	}()

	return nil
}
