# Install required dependencies
[working-directory: 'frontend']
deps:
    npm install

# Build the project
build:
    npm run build --prefix frontend/
    go build -ldflags="-s -w" -trimpath -o ./bin/hajimari ./cmd/hajimari/main.go

# Run the program (depends on build)
run: build
    ./bin/hajimari

# Run frontend and backend in dev mode (parallel, but frontend waits for backend)
dev: dev-backend dev-frontend

dev-backend:
    air &

[working-directory: 'frontend']
dev-frontend: dev-backend
    npm run dev -- --open &

# Format the project with gofmt
fmt:
    gofmt -l -w -s .

# Lint code with golangci-lint
lint:
    golangci-lint run

# Run the tests
test:
    go test -cover ./...
