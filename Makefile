# Variáveis
APP_NAME=go-stock
CMD_PATH=./cmd
BUILD_DIR=./bin

# Comando padrão
.PHONY: all
all: run

# Rodar a aplicação
.PHONY: run
run:
	go run $(CMD_PATH)/main.go

# Build
.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

# Testes
.PHONY: test
test:
	go test ./... -v

# Limpar build
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Lint (se tiver golangci-lint instalado)
.PHONY: lint
lint:
	golangci-lint run

# Baixar dependências
.PHONY: deps
deps:
	go mod tidy