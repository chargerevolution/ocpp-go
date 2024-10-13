package examples

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gregszalay/ocpp-messages-go/types/AuthorizeResponse"
	"github.com/gregszalay/ocpp-messages-go/wrappers"
	"github.com/sanity-io/litter"
)

func Run_callresult_unmarshal_example(json_string []byte) {
	// Create CALL JSON string
	fmt.Println("\n*******************************")
	fmt.Println("OCPP CALLRESULT message as JSON:")
	fmt.Println("*******************************")
	fmt.Printf("%s\n", json_string)

	// Unmarshal CALLRESULT JSON string
	fmt.Println("\n*******************************")
	var callresult wrappers.CALLRESULT
	call_result_unmarshal_err := callresult.UnmarshalJSON(json_string)
	if call_result_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal OCPP CALLRESULT message. Error: %s", call_result_unmarshal_err)
	} else {
		fmt.Println("OCPP CALLRESULT message as an OBJECT:")
		fmt.Println("*******************************")
		litter.Dump(callresult)
	}

	// Unmarshal payload
	fmt.Println("\n*******************************")
	var req AuthorizeResponse.AuthorizeResponseJson
	payload_unmarshal_err := req.UnmarshalJSON(callresult.GetPayloadAsJSON())
	if payload_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal CALLRESULT message payload. Error: %s", payload_unmarshal_err)
	} else {
		fmt.Println("Payload as an OBJECT:")
		fmt.Println("*******************************")
		litter.Dump(req)
	}
	fmt.Println("*******************************")
}

func Create_callresult_json_string() []byte {
	// Create example OCPP message object (e.g. BootNotificationRequest)
	auth_resp := AuthorizeResponse.AuthorizeResponseJson{
		IdTokenInfo: AuthorizeResponse.IdTokenInfoType{
			Status: AuthorizeResponse.AuthorizationStatusEnumType_1_Accepted,
			EvseId: []int{1},
		},
	}

	// All OCPP messages need to be wrapped into a CALL, CALLRESULT or CALLERROR type
	call_result_wrapper := wrappers.CALLRESULT{
		BaseMessage: wrappers.BaseMessage{
			MessageTypeId: wrappers.CALLRESULT_TYPE,
			MessageId:     uuid.NewString(),
		},
		Payload: auth_resp,
	}

	return call_result_wrapper.MarshalPretty()
}
