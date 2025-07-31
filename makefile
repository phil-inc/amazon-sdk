.PHONY: generate-go-client
generate-go-client:
	openapi-generator-cli generate \
		-i api.yaml \
		-g go \
		-o ./generated-clients/api