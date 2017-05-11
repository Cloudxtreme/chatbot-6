build: docker-compose build

up:
	docker-compose up -d --force-recreate --remove-orphans

down:
	docker-compose down

api:
	docker-compose up -d --force-recreate api1 api2 api3
