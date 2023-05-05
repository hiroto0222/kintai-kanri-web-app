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