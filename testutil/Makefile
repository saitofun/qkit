export PG_TEST_DB_NAME=test
export PG_TEST_DB_USER=test_user
export PG_TEST_DB_PASSWD=test_passwd
export PG_TEST_HOSTNAME='postgres://$(PG_TEST_DB_USER):$(PG_TEST_DB_PASSWD)@127.0.0.1:5432'
export PG_TEST_MASTER_EP='$(PG_TEST_HOSTNAME)/$(PG_TEST_DB_NAME)'
export PG_TEST_SLAVE_EP=$(PG_TEST_HOSTNAME)


pg_envs:
	@echo "=== print env variable ==="
	@echo 'PG_TEST_DB_NAME   = $(PG_TEST_DB_NAME)'
	@echo 'PG_TEST_DB_USER   = $(PG_TEST_DB_USER)'
	@echo 'PG_TEST_DB_PASSWD = $(PG_TEST_DB_PASSWD)'
	@echo 'PG_TEST_HOSTNAME  = $(PG_TEST_HOSTNAME)'
	@echo 'PG_TEST_MASTER_EP = $(PG_TEST_MASTER_EP)'
	@echo 'PG_TEST_SLAVE_EP  = $(PG_TEST_SLAVE_EP)'
	@echo "=== print env variable end  ===\n"

pg_start:
	@if [[ $$(pg_isready -h localhost) != "localhost:5432 - accepting connections" ]] ; \
	then \
		echo "=== start postgres server ==="; \
		docker-compose -f docker-compose-pg.yaml up -d ; \
		echo "=== init database ===" ; \
		for i in {1..5} ; \
		do \
			if [[ $$(pg_isready -h localhost) =~ "accepting connections" ]] ; \
			then \
				psql $(PG_TEST_HOSTNAME) -c 'create database $(PG_TEST_DB_NAME)' && \
				psql $(PG_TEST_HOSTNAME) -c 'create schema $(PG_TEST_DB_NAME)' ; \
				break ; \
			else \
				echo "server not ready, retry in 10 second" ; sleep 10 ; \
			fi \
		done ; \
		if [[ $$(pg_isready -h localhost) != "localhost:5432 - accepting connections" ]] ; \
		then \
			echo "=== database init failed ==="  ; \
			exit 1;  \
		fi \
	fi ; \


CONF_POSTGRES_ROOT ?= ../conf/postgres
SQLX_ROOT ?= ../kit/sqlx

test_conf_postgres: pg_envs pg_start
	@echo "root: $(CONF_POSTGRES_ROOT)"
	@cd $(CONF_POSTGRES_ROOT) && go test -v .

test_sqlx: pg_envs pg_start
	@echo "root: $(SQLX_ROOT)"
	@cd $(SQLX_ROOT) && go test -v ./...

test: test_conf_postgres test_sqlx
	@cd ../kit/enum                      && go test ./...
	@cd ../kit/enumgen                   && go test ./...
	@cd ../kit/httptransport/client      && go test ./...
	@cd ../kit/httptransport/handlers    && go test ./...
	@cd ../kit/httptransport/httpx       && go test ./...
	@cd ../kit/httptransport/mock        && go test ./...
	@cd ../kit/httptransport/transformer && go test ./...
	@cd ../kit/httptransport             && go test .
	@cd ../kit/kit                       && go test ./...
	@cd ../kit/metax                     && go test ./...
	@cd ../kit/modelgen                  && go test ./...
	@cd ../kit/statusx                   && go test ./...
	@cd ../kit/statusxgen                && go test ./...
	@cd ../kit/validator                 && go test ./...
	@cd ../x/contextx                    && go test ./...
	@cd ../x/mapx                        && go test ./...
	@cd ../x/pkgx                        && go test ./...
	@cd ../x/reflectx                    && go test ./...
	@cd ../x/typesx                      && go test ./...
	@echo "========TEST PASSED========"


test_and_clean: test
	@docker-compose -f docker-compose-pg.yaml down
	@rm -rf data

fmt:
	go fmt ../...
