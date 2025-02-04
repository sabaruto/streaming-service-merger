services := $(subst ./configMaps/services/,,$(wildcard ./configMaps/services/*))
all: update-images

# -- Docker --

update-images: update-gateway-image update-service-images

update-gateway-image: gateway-image-build gateway-image-push gateway-image-pull

update-service-images: $(foreach wrd,$(services),$(wrd)-image-build $(wrd)-image-push $(wrd)-image-pull)

%-image-build:
	docker build -t sabaruto/${subst -image-build,,$@} --target ${subst -image-build,,$@} ./backend/

%-image-push:
	docker image push sabaruto/${subst -image-push,,$@}

%-image-pull:
	docker pull sabaruto/${subst -image-pull,,$@}


# -- Golang -- 

binaries:
	mkdir -p ./bin
	go build -o ./bin ./backend/cmd/*

# -- postgres --

service-initdb: $(foreach wrd,$(services),$(wrd)-initdb)

%-setup-db: ${subst -setup-db,,$@}-initdb ${subst -setup-db,,$@}-goose-up

%-setup-test-db: ${subst -setup-test-db,,$@}-initdb-test ${subst -setup-test-db,,$@}-goose-up-test

%-teardown-db: ${subst -setup-db,,$@}-initdb ${subst -setup-db,,$@}-goose-down

%-teardown-test-db: ${subst -setup-test-db,,$@}-initdb-test ${subst -setup-test-db,,$@}-goose-down-test

%-initdb:
	psql -U postgres -h localhost -f ./backend/lib/${subst -initdb,,$@}/postgres/initdb.sql

%-initdb-test:
	psql -U postgres -h localhost -f ./backend/lib/${subst -initdb-test,,$@}/postgres/test_initdb.sql

%-goose-up:
	goose up --env ./backend/lib/${subst -goose-up,,$@}/postgres/local.env --dir ./backend/lib/${subst -goose-up,,$@}/postgres/migrations/

%-goose-down:
	goose down --env ./backend/lib/${subst -goose-down,,$@}/postgres/local.env --dir ./backend/lib/${subst -goose-down,,$@}/postgres/migrations/

%-goose-up-test:
	goose up --env ./backend/lib/${subst -goose-up-test,,$@}/postgres/test.env --dir ./backend/lib/${subst -goose-up-test,,$@}/postgres/migrations/

%-goose-down-test:
	goose down --env ./backend/lib/${subst -goose-down-test,,$@}/postgres/test.env --dir ./backend/lib/${subst -goose-down-test,,$@}/postgres/migrations/

clean:
	rm -rf ./bin
