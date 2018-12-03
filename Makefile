NODE_BIN = $(shell npm bin)
IMPORT_PATH = $(shell echo `pwd` | sed "s|^$(GOPATH)/src/||g")
TARGET = $(shell echo $(IMPORT_PATH) | sed 's:.*/::')

GO_FILES = $(shell find . -type f -name "*.go")

PID = $(TARGET).pid
PROCESS_ID = $(shell cat $(PID))

BACKEND_PATH = "backend"
FRONTEND_PATH = "frontend"
DIST_PATH = "dist"

BACKEND_PORT = 9000
FRONTEND_PORT = 8000

build: clean $(TARGET)

clean: kill
	@printf "Cleaning up...\n"
	@if test -s $(BACKEND_PATH)/$(TARGET); \
	then \
		rm $(BACKEND_PATH)/$(TARGET); \
	fi;
	@if test -s $(DIST)/$(TARGET); \
	then \
		rm $(DIST)/$(TARGET); \
	fi;

$(TARGET): $(GO_FILES)
	@printf "Building '$(TARGET)'...\n"
	@cd $(BACKEND_PATH); \
		go build -race -o $(TARGET)

dev: clean start_backend
	cd $(FRONTEND_PATH); FRONTEND_PORT=$(FRONTEND_PORT) BACKEND_PORT=$(BACKEND_PORT) yarn serve &
	@printf "\nWaiting for changes...\n\n"
	@fswatch --event=Updated $(GO_FILES) | xargs -n1 -I{} make start_backend || make kill

start_backend: kill $(TARGET)
	@printf "Starting backend server...\n\n"
	@BACKEND_PORT=$(BACKEND_PORT) ./$(BACKEND_PATH)/$(TARGET) -debug & echo $$! > $(PID)

kill:
		kill $(PROCESS_ID) || true

