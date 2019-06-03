.PHONY: setup_db teardown_db empty_db login_db build_crawler build_server crawl server

boot_db:
	docker-compose up -d
	sleep 20 # wait mysql initialization
create_schema:
	mysql -h ${DB_HOST} --port ${DB_PORT} -u${DB_USER} -p${DB_PASSWORD} -e "create database ${DB_SCHEMA}"
teardown_db:
	docker-compose down
empty_db:
	mysql -h ${DB_HOST} --port ${DB_PORT} -u${DB_USER} -p${DB_PASSWORD} -e "drop database ${DB_SCHEMA}; create database ${DB_SCHEMA}"
login_db:
	mysql -h ${DB_HOST} --port ${DB_PORT} -u${DB_USER} -p${DB_PASSWORD} ${DB_SCHEMA}
build_crawler:
	GOOS=linux GOARCH=amd64 go build -o crawler crawl/main.go
build_server:
	GOOS=linux GOARCH=amd64 go build -o server api/main.go
crawl:
	go run crawl/main.go
server:
	go run api/main.go
