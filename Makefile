# run server
server:
	cd backend && air

# run /bin/sh in backend
backend-sh:
	docker-compose run --rm backend /bin/sh
