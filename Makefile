.PHONY: setup_db teardown_db crawl

setup_db:
	docker-compose up -d
	sleep 20 # wait mysql initialization
	mysql -h ${DB_HOST} --port ${DB_PORT} -u${DB_USER} -p${DB_PASSWORD} -e "create database ${DB_SCHEMA}"
teardown_db:
	docker-compose down
crawl:
	go run crawl/main.go
