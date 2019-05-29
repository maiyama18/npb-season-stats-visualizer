setup_db:
	docker-compose up -d
	sleep 15 # wait mysql initialization
	mysql -h ${DB_HOST} --port ${DB_PORT} -u${DB_USER} -p${DB_PASSWORD} -e "create database ${DB_SCHEMA}"
teardown_db:
	docker-compose down
