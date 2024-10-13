# OCPP 2.0.1 data types for GO

Data types for 
- CALL, CALLRESULT and CALLERROR message wrappers (custom built)
- all OCPP 2.0.1 message payloads (generated from JSON schemas with: https://github.com/atombender/go-jsonschema)

## Getting started

Run **example.go** to see the types working

## Re-generate types from schemas (if needed):

(See: https://github.com/omissis/go-jsonschema)
    
    go get github.com/atombender/go-jsonschema/...

    go install github.com/atombender/go-jsonschema@latest


    ./generate.sh
## More info

Based on: OCPP-2.0.1_part3_JSON_schemas (Download:
[OCPP 2.0.1](https://www.openchargealliance.org/downloads/))





