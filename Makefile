migrateup:
	migrate -path db/migration -database "mysql://root:qwertyuiop@tcp(localhost:3306)/goweb" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:qwertyuiop@tcp(localhost:3306)/goweb" -verbose down

migratedrop:
	migrate -path db/migration -database "mysql://root:qwertyuiop@tcp(localhost:3306)/goweb" -verbose drop

migrateadd:
	migrate create -ext sql -dir db/migration

.PHONY: migrateup migratedown migratedrop migrateadd