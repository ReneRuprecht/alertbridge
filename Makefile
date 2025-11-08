HOST = 0.0.0.0
PORT = 8000

.PHONY: backend frontend dev clean up down fresh_dev

dev:
	make -j2 backend frontend

backend:
	fastapi dev backend/src/main.py --host ${HOST} --port ${PORT}

frontend:
	@cd frontend && npm run dev

clean:
	docker-compose down -v --remove-orphans

up:
	docker-compose up -d

down:
	docker-compose down

fresh_db:
	$(MAKE) clean
	$(MAKE) up
