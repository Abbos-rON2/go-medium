swag-init:
	swag init -g rest/server.go -o api/docs

migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/social?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/social?sslmode=disable' down

migration:
	migrate create -ext sql -dir ./migrations -seq -digits 3 ${name}