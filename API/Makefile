run : ## Run go program
	go run ./cmd/api

wire: ## Generate wire_gen.go on ubuntu
	cd pkg/di && wire
wirewin: ## Generate wire_gen.go on windows
	cd pkg/di ; wire

swag :
	swag init -g cmd/api/main.go -o ./cmd/api/docs

