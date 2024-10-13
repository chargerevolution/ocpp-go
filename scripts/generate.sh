#!/bin/bash

fatal() {
	echo "FATAL: $*"
	exit 1
}

OCPP_VERSION="$1"
case "$OCPP_VERSION" in
"1.6" | "2.0.1") ;;
*)
	fatal "Please provide a valid OCPP protocol version"
	;;
esac

MESSAGE_TYPE="$2"
if [ -z "${MESSAGE_TYPE}" ]; then fatal "Please provide message type"; fi

SCHEMA_FILE="v$OCPP_VERSION/docs/schemas/$MESSAGE_TYPE.json"
if ! [ -e "${SCHEMA_FILE}" ]; then fatal "File $SCHEMA_FILE not found"; fi

DESTINATION_DIR="${3:-$PWD/v$OCPP_VERSION/generated}"
mkdir -p "$DESTINATION_DIR"

# Install jsonschema-go converter
if ! go-jsonschema -h >>/dev/null; then
	go get github.com/atombender/go-jsonschema/... || fatal "Unable to download go-jsonschema"
	go install github.com/atombender/go-jsonschema@latest || fatal "Unable to install go-jsonschema"
fi

GO_PACKAGE_NAME="messages"
echo "$DESTINATION_DIR/${MESSAGE_TYPE}.go"
go-jsonschema -p "$GO_PACKAGE_NAME" "$SCHEMA_FILE" >"$DESTINATION_DIR/${MESSAGE_TYPE}.go"
