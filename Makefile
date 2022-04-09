run:
	go run src/main/main.go

unit_test:
	go test -v -race -timeout 1000s -covermode=atomic -coverpkg=./... -coverprofile=unit_test.raw.out ./...