DB_URL=postgresql://admin:password123@postgres:5432/kintai-kanri-db?sslmode=disable

# run server
server:
	cd backend && air

# run /bin/sh in backend
backend-sh:
	docker-compose run --rm backend /bin/sh

# run db migrate-up
migrate-up:
	docker-compose exec backend migrate -path db/migrations -database "$(DB_URL)" -verbose up

# run db migrate-down
migrate-down:
	docker-compose exec backend migrate -path db/migrations -database "$(DB_URL)" -verbose down

# run db migrate-up 1
migrate-up-1:
	docker-compose exec backend migrate -path db/migrations -database "$(DB_URL)" -verbose up 1

# run db migrate-down 1
migrate-down-1:
	docker-compose exec backend migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

# run sqlc generate
sqlc-generate:
	docker-compose exec backend sqlc generate

# run tests
test:
	docker-compose exec backend go test -v -cover ./... -coverprofile=cover.out

# check test coverage
see-coverage:
	cd backend && go tool cover -html=cover.out -o cover.html && open cover.html

# generate mock store
mock-store:
	cd backend && mockgen --package mockdb --destination db/mock/store.go github.com/hiroto0222/kintai-kanri-web-app/db/sqlc Store
