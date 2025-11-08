HOST = 0.0.0.0
PORT = 8000

dev:
	fastapi dev src/main.py --host ${HOST} --port ${PORT}

clean:
	docker-compose down -v --remove-orphans

up:
	docker-compose up -d

down:
	docker-compose down

fresh_db:
	$(MAKE) clean
	$(MAKE) up
