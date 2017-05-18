all:
	docker-compose up -d --force-recreate --no-deps api && docker-compose logs -f api
