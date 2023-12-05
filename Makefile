.PHONY: all compose create-db drop-db migrate-up migrate-down start test

compose:
	docker-compose up

create-db:
	docker exec -it catalog_db_1 createdb --username=root --owner=root catalog_development

drop-db:
	docker exec -it catalog_db_1 dropdb catalog_development

migrate-up:
	migrate -path db/migration -database "postgresql://root:root@localhost:5433/catalog_development?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:root@localhost:5433/catalog_development?sslmode=disable" -verbose down

start:
	go run app/api/main.go

test:
	go test -cover -coverprofile=coverage.out -json $$(go list ./... | grep -Ev "app") > ./UT-Catalog-report_tms.json

test-cover:
	make test
	go tool cover -html=coverage.out

mock-init:
	mockery --all --dir ./ --output ./test/mock --case underscore
