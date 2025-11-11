import logging
from contextlib import asynccontextmanager

from fastapi import FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware
from src.repository.psql_client import PSQLClient
from src.services.alert_service import AlertService

logging.basicConfig(level=logging.INFO)

app = FastAPI()

origins = ["http://localhost:5173"]

app.add_middleware(CORSMiddleware, allow_origins=origins)

psql_client: PSQLClient = PSQLClient()
alert_service: AlertService = AlertService(psql_client=psql_client)


@asynccontextmanager
async def lifespan(app: FastAPI):
    logging.info("Starting postgres connection")
    await psql_client.create_pool(open=True)

    logging.info("Starting initial alert fetch from alertmanager")
    await alert_service._init_db_data_from_alertmanager()
    logging.info("Initial alert fetch completed")

    yield

    await psql_client.close_pool()


@app.get("/alerts")
async def alerts():
    alerts = await alert_service.fetch_alerts_from_db()

    return {"alerts": alerts}


@app.post("/api/v1/alerts")
async def alert_manager_webhook(request: Request):
    try:
        data = await request.json()

        alert_json = alert_service.process_webhook_payload(data)
        processed_alerts_count: int = await alert_service.save_alerts_to_db(alert_json)

        logging.info(
            "Processed %s alerts from Alertmanager webhook", processed_alerts_count
        )

    except Exception as e:
        logging.error("Failed to process webhook: %s", e)


app.router.lifespan_context = lifespan
