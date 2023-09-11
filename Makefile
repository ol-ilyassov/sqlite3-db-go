
EXECUTABLE_NAME := "sqlite3-app"

.SILENT:

# build: builds app.
.PHONY: build
build:
	cd cmd/app && \
		go build -o ../../bin/${EXECUTABLE_NAME}
	@echo "- make build finished"

# build: builds app and runs it.
.PHONY: run
run: build
	cd bin && \
		./${EXECUTABLE_NAME}
	@echo "- make run finished"

# clear: clears bin folder.
.PHONY: clear
clear: 
	rm -rf bin