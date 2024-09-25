VERSION = 2.0

generate-client:
	@echo "Generating client for version $(VERSION)"
	oapi-codegen -package oscp_v20 ./$(VERSION)/api/$(VERSION).spec.yaml > ./$(VERSION)/client.go
	cd $(VERSION) && go mod tidy