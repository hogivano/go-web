include .env

migrateup:
	migrate -path db/migration -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose down

migratedrop:
	migrate -path db/migration -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose drop

migrateadd:
	migrate create -ext sql -dir db/migration

testcase:
	go test go-web/test

.PHONY: migrateup migratedown migratedrop migrateadd testcase