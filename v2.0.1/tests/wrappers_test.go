package messages

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"home/greg/ampmates/ocpp-messages-go"
	"github.com/stretchr/testify/assert"
)

func Create_callerror_json_string() []byte {
	// Create a CALLERROR instance with the necessary fields
	callError := messages.CALLERROR{
		BaseMessage: messages.BaseMessage{
			MessageTypeId: messages.CALLERROR_TYPE,
			MessageId:     uuid.NewString(),
		},
		ErrorCode:        string(messages.FormatViolation),
		ErrorDescription: "Some really descriptive words about this error",
		ErrorDetails:     "Don't even bother debugging, just go on stackoverflow :D",
	}

	// Marshal the CALLERROR into a pretty JSON string
	return callError.MarshalPretty()
}

func Create_call_json_string() []byte {
	// Create example OCPP message object (e.g. BootNotificationRequest)
	boot_req := messages.BootNotificationRequest.BootNotificationRequestJson{
		Reason: "PowerUp",
		ChargingStation: messages.BootNotificationRequest.ChargingStationType{
			Model:      "super-charger-6000",
			VendorName: "WattsUp",
		},
	}

	// All OCPP messages need to be wrapped into a CALL, CALLRESULT or CALLERROR type
	call_wrapper := messages.CALL{
		BaseMessage: messages.BaseMessage{
			MessageTypeId: messages.CALL_TYPE,
			MessageId:     uuid.NewString(),
		},
		Action:  "BootNotification",
		Payload: boot_req,
	}

	return call_wrapper.MarshalPretty()
}

func Create_callresult_json_string() []byte {
	// Create example OCPP message object (e.g. BootNotificationRequest)
	auth_resp := messages.AuthorizeResponse.AuthorizeResponseJson{
		IdTokenInfo: messages.AuthorizeResponse.IdTokenInfoType{
			Status: messages.AuthorizeResponse.AuthorizationStatusEnumType_1_Accepted,
			EvseId: []int{1},
		},
	}

	// All OCPP messages need to be wrapped into a CALL, CALLRESULT or CALLERROR type
	call_result_wrapper := messages.CALLRESULT{
		BaseMessage: messages.BaseMessage{
			MessageTypeId: messages.CALLRESULT_TYPE,
			MessageId:     uuid.NewString(),
		},
		Payload: auth_resp,
	}

	return call_result_wrapper.MarshalPretty()
}

func TestCALL_MarshalUnmarshal(t *testing.T) {
	bootReq := messages.BootNotificationRequest.BootNotificationRequestJson{
		Reason: "PowerUp",
		ChargingStation: messages.BootNotificationRequest.ChargingStationType{
			Model:      "super-charger-6000",
			VendorName: "WattsUp",
		},
	}

	call := messages.CALL{
		BaseMessage: messages.BaseMessage{
			MessageTypeId: messages.CALL_TYPE,
			MessageId:     uuid.NewString(),
		},
		Action:  "BootNotification",
		Payload: bootReq,
	}

	// Test Marshaling
	marshaled, err := json.Marshal(call)
	assert.NoError(t, err, "Expected no error during marshaling")
	assert.NotEmpty(t, marshaled, "Marshalled JSON should not be empty")

	// Test Unmarshaling
	var unmarshaledCall messages.CALL
	err = unmarshaledCall.UnmarshalJSON(marshaled)
	assert.NoError(t, err, "Expected no error during unmarshaling")
	assert.Equal(t, call.MessageId, unmarshaledCall.MessageId, "Message IDs should match")
	assert.Equal(t, call.Action, unmarshaledCall.Action, "Actions should match")

	// Verify Payload
	payloadJSON := call.GetPayloadAsJSON()
	var payload messages.BootNotificationRequest.BootNotificationRequestJson
	err = json.Unmarshal(payloadJSON, &payload)
	assert.NoError(t, err, "Expected no error during payload unmarshaling")
	assert.Equal(t, bootReq.Reason, payload.Reason, "Payload reasons should match")
}

func TestCALLERROR_MarshalUnmarshal(t *testing.T) {
	callError := CALLERROR{
		BaseMessage: messages.BaseMessage{
			MessageTypeId: messages.CALLERROR_TYPE,
			MessageId:     uuid.NewString(),
		},
		ErrorCode:        string(messages.FormatViolation),
		ErrorDescription: "Some descriptive words about this error",
		ErrorDetails:     "Don't even bother debugging, just go on StackOverflow :D",
	}

	// Test Marshaling
	marshaledError, err := json.Marshal(callError)
	assert.NoError(t, err, "Expected no error during marshaling")
	assert.NotEmpty(t, marshaledError, "Marshalled JSON should not be empty")

	// Test Unmarshaling
	var unmarshaledError CALLERROR
	err = unmarshaledError.UnmarshalJSON(marshaledError)
	assert.NoError(t, err, "Expected no error during unmarshaling")
	assert.Equal(t, callError.MessageId, unmarshaledError.MessageId, "Message IDs should match")
	assert.Equal(t, callError.ErrorCode, unmarshaledError.ErrorCode, "Error codes should match")
}

func TestCALLRESULT_MarshalUnmarshal(t *testing.T) {
	payload := map[string]interface{}{"status": "Accepted"}

	callResult := CALLRESULT{
		BaseMessage: BaseMessage{
			MessageTypeId: CALLRESULT_TYPE,
			MessageId:     uuid.NewString(),
		},
		Payload: payload,
	}

	// Test Marshaling
	marshaledResult, err := json.Marshal(callResult)
	assert.NoError(t, err, "Expected no error during marshaling")
	assert.NotEmpty(t, marshaledResult, "Marshalled JSON should not be empty")

	// Test Unmarshaling
	var unmarshaledResult CALLRESULT
	err = unmarshaledResult.UnmarshalJSON(marshaledResult)
	assert.NoError(t, err, "Expected no error during unmarshaling")
	assert.Equal(t, callResult.MessageId, unmarshaledResult.MessageId, "Message IDs should match")
	assert.Equal(t, payload, unmarshaledResult.Payload, "Payloads should match")
}

func TestCreateCallJSON(t *testing.T) {
	expectedAction := "BootNotification"

	callJSON := Create_call_json_string()
	var call CALL
	err := call.UnmarshalJSON(callJSON)
	assert.NoError(t, err, "Expected no error during unmarshaling")
	assert.Equal(t, expectedAction, call.Action, "Actions should match")
}
