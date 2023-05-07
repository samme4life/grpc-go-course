## Adds Git hooks to correct location and configures Git to use SSH over HTTPS
config:
	git config --global --replace-all url."git@github.com:".insteadOf "https://github.com/"

.PHONY: mod
mod: config ## Initialise the go module
	rm -rf go.sum go.mod vendor
	go mod init github.com/samme4life/grpc-go-course
	go mod tidy

install: ## installs dependencies
	go mod download
	go mod vendor