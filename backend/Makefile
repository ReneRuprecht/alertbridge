
up:
	docker compose up -d

down:
	docker compose down

clean:
	docker compose down -v

test-unit:
	@output=$$(go test ./... 2>&1); \
	if [ $$? -eq 0 ]; then \
		echo "\033[0;32mOK\033[0m"; \
	else \
		echo "$$output"; \
		echo "\033[0;31mFAIL\033[0m"; \
	fi

test:
	@output=$$(go test -tags=integration ./... 2>&1); \
	if [ $$? -eq 0 ]; then \
		echo "\033[0;32mOK\033[0m"; \
	else \
		echo "$$output"; \
		echo "\033[0;31mFAIL\033[0m"; \
	fi
