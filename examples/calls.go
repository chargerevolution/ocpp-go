package examples

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gregszalay/ocpp-messages-go/types/BootNotificationRequest"
	"github.com/gregszalay/ocpp-messages-go/wrappers"
	"github.com/sanity-io/litter"
)

func Run_call_unmarshal_example(json_string []byte) {
	// Create CALL JSON string
	fmt.Println("\n*******************************")
	fmt.Println("OCPP CALL message as JSON:")
	fmt.Println("*******************************")
	fmt.Printf("%s\n", json_string)

	// Unmarshal CALL JSON string
	fmt.Println("\n*******************************")
	var call wrappers.CALL
	call_unmarshal_err := call.UnmarshalJSON(json_string)
	if call_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal OCPP CALL message. Error: %s", call_unmarshal_err)
	} else {
		fmt.Println("OCPP CALL message as an OBJECT:")
		fmt.Println("*******************************")
		litter.Dump(call)
	}

	// Unmarshal payload
	fmt.Println("\n*******************************")
	var req BootNotificationRequest.BootNotificationRequestJson
	payload_unmarshal_err := req.UnmarshalJSON(call.GetPayloadAsJSON())
	if payload_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal CALL message payload. Error: %s", payload_unmarshal_err)
	} else {
		fmt.Println("Payload as an OBJECT:")
		fmt.Println("*******************************")
		litter.Dump(req)
	}
	fmt.Println("*******************************")
}

func Create_call_json_string() []byte {
	// Create example OCPP message object (e.g. BootNotificationRequest)
	boot_req := BootNotificationRequest.BootNotificationRequestJson{
		Reason: "PowerUp",
		ChargingStation: BootNotificationRequest.ChargingStationType{
			Model:      "super-charger-6000",
			VendorName: "WattsUp",
		},
	}

	// All OCPP messages need to be wrapped into a CALL, CALLRESULT or CALLERROR type
	call_wrapper := wrappers.CALL{
		BaseMessage: wrappers.BaseMessage{
			MessageTypeId: wrappers.CALL_TYPE,
			MessageId:     uuid.NewString(),
		},
		Action:  "BootNotification",
		Payload: boot_req,
	}

	return call_wrapper.MarshalPretty()
}
