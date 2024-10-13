package messages

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CALLERROR structure extending BaseMessage
type CALLERROR struct {
	BaseMessage
	ErrorCode        string // Error code from the RPC framework
	ErrorDescription string // Optional description
	ErrorDetails     string // JSON object with error details
}

// UnmarshalJSON custom unmarshals the CALLERROR from a JSON array
func (c *CALLERROR) UnmarshalJSON(b []byte) error {
	var raw []interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("could not unmarshal CALLERROR json: %w", err)
	}
	if err := c.UnmarshalBaseMessage(raw, 5); err != nil {
		return err
	}

	errorCode, ok := raw[2].(string)
	if !ok || errorCode == "" {
		return errors.New("errorCode is not a valid non-empty string")
	}
	errorDescription, ok := raw[3].(string)
	if !ok {
		return errors.New("errorDescription is not a valid string")
	}
	errorDetails, ok := raw[4].(string)
	if !ok {
		return errors.New("errorDetails is not a valid string")
	}

	c.ErrorCode = errorCode
	c.ErrorDescription = errorDescription
	c.ErrorDetails = errorDetails
	return nil
}

// Marshal generates a JSON array from the CALLERROR
func (c *CALLERROR) Marshal() []byte {
	messageArray := []interface{}{c.MessageTypeId, c.MessageId, c.ErrorCode, c.ErrorDescription, c.ErrorDetails}
	return c.BaseMessage.Marshal(messageArray)
}

// MarshalPretty generates a pretty JSON from the CALLERROR
func (c *CALLERROR) MarshalPretty() []byte {
	messageArray := []interface{}{c.MessageTypeId, c.MessageId, c.ErrorCode, c.ErrorDescription, c.ErrorDetails}
	return c.BaseMessage.MarshalPretty(messageArray)
}

// RPC framework error codes
type rpcFrameworkErrorCode string

const (
	FormatViolation               rpcFrameworkErrorCode = "FormatViolation"
	GenericError                  rpcFrameworkErrorCode = "GenericError"
	InternalError                 rpcFrameworkErrorCode = "InternalError"
	MessageTypeNotSupported       rpcFrameworkErrorCode = "MessageTypeNotSupported"
	NotImplemented                rpcFrameworkErrorCode = "NotImplemented"
	NotSupported                  rpcFrameworkErrorCode = "NotSupported"
	OccurrenceConstraintViolation rpcFrameworkErrorCode = "OccurrenceConstraintViolation"
	PropertyConstraintViolation   rpcFrameworkErrorCode = "PropertyConstraintViolation"
	ProtocolError                 rpcFrameworkErrorCode = "ProtocolError"
	RpcFrameworkError             rpcFrameworkErrorCode = "RpcFrameworkError"
	SecurityError                 rpcFrameworkErrorCode = "SecurityError"
	TypeConstraintViolation       rpcFrameworkErrorCode = "TypeConstraintViolation"
)
