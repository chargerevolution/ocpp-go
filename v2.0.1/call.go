package messages

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CALL structure extending BaseMessage
type CALL struct {
	BaseMessage
	Action string // The name of the remote procedure or action
	// JSON Payload of the action, see: JSON Payload for more information.
	Payload interface{}
}

// UnmarshalJSON custom unmarshals the CALL from a JSON array
func (c *CALL) UnmarshalJSON(b []byte) error {
	var raw []interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("could not unmarshal CALL json: %w", err)
	}
	if err := c.UnmarshalBaseMessage(raw, 4); err != nil {
		return err
	}

	action, ok := raw[2].(string)
	if !ok || action == "" {
		return errors.New("action is not a valid non-empty string")
	}
	payload, ok := raw[3].(map[string]interface{})
	if !ok || payload == nil {
		return errors.New("payload is not a valid map[string]interface{}")
	}

	c.Action = action
	c.Payload = payload
	return nil
}

// Marshal generates a JSON array from the CALL
func (c *CALL) Marshal() []byte {
	messageArray := []interface{}{c.MessageTypeId, c.MessageId, c.Action, c.Payload}
	return c.BaseMessage.Marshal(messageArray)
}

// MarshalPretty generates a pretty JSON from the CALL
func (c *CALL) MarshalPretty() []byte {
	messageArray := []interface{}{c.MessageTypeId, c.MessageId, c.Action, c.Payload}
	return c.BaseMessage.MarshalPretty(messageArray)
}

// GetPayloadAsJSON returns the payload as JSON
func (c *CALL) GetPayloadAsJSON() []byte {
	reMarshalledPayload, err := json.MarshalIndent(c.Payload, "", " ")
	if err != nil {
		fmt.Println("Failed to re-marshal CALL payload to JSON")
		return []byte("")
	}
	return reMarshalledPayload
}
