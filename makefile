# Makefile

server-start:
	nodemon --exec go run main.go --signal SIGTERM