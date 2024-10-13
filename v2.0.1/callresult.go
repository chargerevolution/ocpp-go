package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CALLRESULT represents a call result message.
type CALLRESULT struct {
	BaseMessage
	Payload interface{} `json:"payload"`
}

// UnmarshalJSON for CALLRESULT to parse JSON data into the struct.
func (j *CALLRESULT) UnmarshalJSON(b []byte) error {
	var raw []interface{}
	err := json.Unmarshal(b, &raw)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return err
	}
	if len(raw) != 3 {
		return fmt.Errorf("invalid CALLRESULT message length")
	}

	// MessageTypeId: must be a number.
	messageTypeId, ok := raw[0].(float64)
	if !ok {
		return fmt.Errorf("CALLRESULT data[0] is not a number")
	}

	// MessageId: must be a non-empty string.
	messageId, ok := raw[1].(string)
	if !ok || messageId == "" {
		return fmt.Errorf("CALLRESULT data[1] is not a valid string")
	}

	// Payload: must be a map.
	payload, ok := raw[2].(map[string]interface{})
	if !ok || payload == nil {
		return fmt.Errorf("CALLRESULT data[2] is not a map[string]interface{}")
	}

	// Assign parsed values to the struct fields.
	*j = CALLRESULT{
		BaseMessage: BaseMessage{
			MessageTypeId: uint32(messageTypeId),
			MessageId:     messageId,
		},
		Payload: payload,
	}
	return nil
}

// Marshal marshals the CALLRESULT into JSON format.
func (c *CALLRESULT) Marshal() []byte {
	messageArray := [...]interface{}{CALLRESULT_TYPE, c.MessageId, c.Payload}
	messageArrayJSON, err := json.Marshal(messageArray)
	if err != nil {
		fmt.Printf("Error: Could not marshal CALLRESULT message: %s\n", err)
		return []byte("")
	}
	return messageArrayJSON
}

// MarshalPretty marshals the CALLRESULT into pretty JSON format.
func (c *CALLRESULT) MarshalPretty() []byte {
	uglyJSON := c.Marshal()
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(uglyJSON), "", "    "); err != nil {
		return []byte("")
	}
	return prettyJSON.Bytes()
}

// GetPayloadAsJSON returns the payload marshaled as JSON.
func (c *CALLRESULT) GetPayloadAsJSON() []byte {
	if c == nil {
		fmt.Println("CALLRESULT object is empty")
		return []byte("")
	}
	// Re-marshal payload only
	reMarshalledPayload, reMarshalledErr := json.MarshalIndent(c.Payload, "", " ")
	if reMarshalledErr != nil {
		fmt.Println("Failed to re-marshal CALLRESULT payload to JSON")
		return []byte("")
	}
	return reMarshalledPayload
}
