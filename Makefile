ifneq ($(MODE),ci)
SOURCES=$(shell go list -f '{{$$dir := .Dir}}{{range .GoFiles}}{{$$dir}}/{{.}} {{end}}' ./...)
endif

.PHONY: check
check:
	go vet ./...
	go test -race ./...

DOCS_TARGET_DIR=docs

.PHONY: docs
docs: $(SOURCES)
ifeq (,$(shell which gomarkdoc))
	$(error gomarkdoc is not installed.  Run go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest)
endif
	gomarkdoc --output "$(DOCS_TARGET_DIR)/Home.md" ./pkg/loggo
