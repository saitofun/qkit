
tidy:
	go mod tidy

build: tidy clean
	go build

test: build gen_config gen_supervisor gen_docker print_config create_user

gen_config:
	./poc gen

gen_supervisor:
	./poc supervisor

gen_docker:
	./poc docker

print_config:
	./poc config

create_user:
	./poc user

clean:
	rm -rf config/default.yml supervisor.conf poc

