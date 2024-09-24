build:
	go build -o ./bin/app ./cmd/main.go

lint:
	golangci-lint run -v --config=.golangci.yaml

lcm:
	go run cmd/main.go --mode lcm

geom-progression:
	go run cmd/main.go --mode geom-progression
