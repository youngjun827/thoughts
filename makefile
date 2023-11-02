SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

run:
	go run app/services/thoughts-api/main.go | go run app/tooling/logfmt/main.go