package messages

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	CALL_TYPE       = 2
	CALLRESULT_TYPE = 3
	CALLERROR_TYPE  = 4
)

// BaseMessage defines common fields and functions for CALL and CALLERROR
type BaseMessage struct {
	MessageTypeId uint32      // Identifies the message type.
	MessageId     string      // Unique identifier for the message, matches request and response.
	Payload       interface{} // Payload of the message, could be JSON or other data types.
}

// UnmarshalBaseMessage handles the common unmarshalling logic for BaseMessage
func (b *BaseMessage) UnmarshalBaseMessage(raw []interface{}, expectedLen int) error {
	if len(raw) != expectedLen {
		return fmt.Errorf("invalid message length, expected %d fields", expectedLen)
	}

	messageTypeId, ok := raw[0].(float64) // JSON numbers are unmarshalled as float64
	if !ok {
		return errors.New("messageTypeId is not a valid number")
	}
	messageId, ok := raw[1].(string)
	if !ok || messageId == "" {
		return errors.New("messageId is not a valid non-empty string")
	}

	b.MessageTypeId = uint32(messageTypeId)
	b.MessageId = messageId

	return nil
}

// Marshal converts the BaseMessage into a JSON array
func (b *BaseMessage) Marshal(messageArray []interface{}) []byte {
	messageArrayJSON, err := json.Marshal(messageArray)
	if err != nil {
		fmt.Printf("Error: Could not marshal message: %s\n", err)
		return []byte("")
	}
	return messageArrayJSON
}

// MarshalPretty outputs the JSON array in a pretty indented format
func (b *BaseMessage) MarshalPretty(messageArray []interface{}) []byte {
	uglyJSON := b.Marshal(messageArray)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, uglyJSON, "", "    "); err != nil {
		fmt.Printf("Error: Could not prettify JSON message: %s\n", err)
		return uglyJSON
	}
	return prettyJSON.Bytes()
}
