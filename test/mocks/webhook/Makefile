
VERSION?=dev
OUTPUT_DIR=_output
GATEWAY_DIR=.
GATEWAY_SOURCES=$(call get_sources,$(GATEWAY_DIR))

$(OUTPUT_DIR)/converter-linux-amd64: $(GATEWAY_SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(GATEWAY_DIR)/webhook.go


.PHONY: converter
converter: $(OUTPUT_DIR)/converter-linux-amd64

$(OUTPUT_DIR)/Dockerfile.converter: $(GATEWAY_DIR)/Dockerfile
	cp $< $@

converter-docker: $(OUTPUT_DIR)/converter-linux-amd64 $(OUTPUT_DIR)/Dockerfile.converter
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.converter \
		-t solo-io/conversion-webhook:$(VERSION)
