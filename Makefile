# Nombre del binario final
BINARY_NAME=mi-api

# Ruta al main
MAIN_PATH=./cmd/api

.PHONY: build run clean

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

clean:
	del $(BINARY_NAME).exe 2> NUL || true
