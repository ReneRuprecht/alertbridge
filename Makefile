
.PHONY: up down clean backend

up:
	docker-compose up -d postgres redis

backend:
	docker-compose up -d backend

	@echo "Waiting for backend to be ready"
	@timeout=60; \
	elapsed=0; \
	while ! curl -o /dev/null --silent http://localhost:8080/actuator/health; do \
		if [ $$elapsed -ge $$timeout ]; then \
			echo "Timeout reached, backend not healthy"; \
			exit 1; \
		fi; \
		echo "Backend not ready yet..."; \
		sleep 5; \
		elapsed=$$((elapsed+2)); \
	done
	@echo "Backend is ready"

down:
	docker-compose down

clean:
	docker-compose down -v --remove-orphans
