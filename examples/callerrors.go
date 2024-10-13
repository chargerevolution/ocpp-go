package examples

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gregszalay/ocpp-messages-go/wrappers"
	"github.com/sanity-io/litter"
)

func Run_callerror_unmarshal_example(json_string []byte) {
	// Create CALL JSON string
	fmt.Println("\n*******************************")
	fmt.Println("OCPP CALLERROR message as JSON:")
	fmt.Println("*******************************")
	fmt.Printf("%s\n", json_string)

	// Unmarshal CALLERROR JSON string
	fmt.Println("\n*******************************")
	var callerror wrappers.CALLERROR
	callerror_result_unmarshal_err := callerror.UnmarshalJSON(json_string)
	if callerror_result_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal OCPP CALLERROR message. Error: %s", callerror_result_unmarshal_err)
	} else {
		fmt.Println("OCPP CALLERROR message as an OBJECT:")
		fmt.Println("*******************************")
		litter.Dump(callerror)
	}
}

func Create_callerror_json_string() []byte {
	// Create a CALLERROR instance with the necessary fields
	callError := wrappers.CALLERROR{
		BaseMessage: wrappers.BaseMessage{
			MessageTypeId: wrappers.CALLERROR_TYPE,
			MessageId:     uuid.NewString(),
		},
		ErrorCode:        string(wrappers.FormatViolation),
		ErrorDescription: "Some really descriptive words about this error",
		ErrorDetails:     "Don't even bother debugging, just go on stackoverflow :D",
	}

	// Marshal the CALLERROR into a pretty JSON string
	return callError.MarshalPretty()
}
