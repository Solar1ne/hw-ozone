test-cover:
	go test -cover ./internal/domain/services/cart_service_test.go


.PHONY: run-all

run-all:
	go run cmd/app/main.go
