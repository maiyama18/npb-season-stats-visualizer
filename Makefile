.PHONY: setup_db teardown_db crawl

setup_db:
	docker-compose up -d
	sleep 20 # wait mysql initialization
	mysql -h ${DB_HOST} --port ${DB_PORT} -u${DB_USER} -p${DB_PASSWORD} -e "create database ${DB_SCHEMA}"
teardown_db:
	docker-compose down
empty_db:
	mysql -h ${DB_HOST} --port ${DB_PORT} -u${DB_USER} -p${DB_PASSWORD} -e "drop database ${DB_SCHEMA}; create database ${DB_SCHEMA}"
login_db:
	mysql -h ${DB_HOST} --port ${DB_PORT} -u${DB_USER} -p${DB_PASSWORD} ${DB_SCHEMA}
crawl:
	go run crawl/main.go
server:
	go run api/main.go