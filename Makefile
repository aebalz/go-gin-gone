.PHONY: dev

dev:
	nodemon --exec go run ./main.go --signal SIGTERM