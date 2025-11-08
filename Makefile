HOST = 0.0.0.0
PORT = 8000

dev:
	fastapi dev src/main.py --host ${HOST} --port ${PORT}
