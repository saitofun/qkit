
export PG_TEST_DB_NAME=test
export PG_TEST_DB_USER=test_user
export PG_TEST_DB_PASSWD=test_passwd
export PG_TEST_HOSTNAME='postgres://$(PG_TEST_DB_USER):$(PG_TEST_DB_PASSWD)@127.0.0.1:5432'
export PG_TEST_MASTER_EP='$(PG_TEST_HOSTNAME)/$(PG_TEST_DB_NAME)'
export PG_TEST_SLAVE_EP=$(PG_TEST_HOSTNAME)

print_envs:
	@echo $(PG_TEST_DB_NAME)
	@echo $(PG_TEST_DB_USER)
	@echo $(PG_TEST_DB_PASSWD)
	@echo $(PG_TEST_HOSTNAME)
	@echo $(PG_TEST_MASTER_EP)
	@echo $(PG_TEST_SLAVE_EP)

test: print_envs
	go test -v .

init_db: print_envs
	@echo "init database"
	psql $(PG_TEST_HOSTNAME) -c 'create database $(PG_TEST_DB_NAME)'
	psql $(PG_TEST_HOSTNAME) -c 'create schema $(PG_TEST_DB_NAME)'
	@echo "finished"

start_postgres:
	docker-compose -f docker-compose-pg.yaml up

clean:
	docker-compose -f docker-compose-pg.yaml down
	rm -rf data log

