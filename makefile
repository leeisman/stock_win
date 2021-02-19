p=$(shell pwd)

scrape:
	PROJ_DIR=$p GOFLAGS=-mod=vendor go run ./main.go scrape

migrate-up-dev:
	goose -dir $p/deployment/migrations mysql 'root:root@tcp(localhost:3306)/stock_win?parseTime=true' up

migrate-down-dev: