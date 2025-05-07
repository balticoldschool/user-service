run:
	go run .

tidy:
	go mod tidy

generate-api:
	oapi-codegen -config api-generator-config.yaml resource/swagger/openapi.yaml
	$(MAKE) tidy
