import logging
from contextlib import asynccontextmanager

from fastapi import Depends, FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware

from src.config.service import (
    get_alert_service,
    get_psql_client,
    get_redis_client,
)
from src.services.alert_service import AlertService

logging.basicConfig(level=logging.INFO)

app = FastAPI()

origins = ["http://localhost:5173"]

app.add_middleware(CORSMiddleware, allow_origins=origins)


@asynccontextmanager
async def lifespan(app: FastAPI):
    logging.info("Starting postgres connection")
    await get_psql_client().create_pool(open=True)
    get_redis_client().connect()

    logging.info("Starting initial alert fetch from alertmanager")
    await get_alert_service().initial_setup_from_alertmanager()
    logging.info("Initial alert fetch completed")

    yield

    await get_psql_client().close_pool()


@app.get("/alerts")
async def alerts(alert_service: AlertService = Depends(get_alert_service)):
    alerts = await alert_service.read_alerts_from_db()

    return {"alerts": alerts}


@app.get("/alerts/minimal")
async def alerts_minimal(
    alert_service: AlertService = Depends(get_alert_service),
):
    alerts = await alert_service.read_alerts_from_redis()

    return {"alerts": alerts}


@app.post("/api/v1/alerts")
async def alert_manager_webhook(
    request: Request, alert_service: AlertService = Depends(get_alert_service)
):
    try:
        data = await request.json()

        processed_alerts_count: int = (
            await alert_service.save_alerts_from_webhook(data)
        )

        logging.info(
            "Processed %s alerts from Alertmanager webhook",
            processed_alerts_count,
        )

    except Exception as e:
        logging.error("Failed to process webhook: %s", e)


app.router.lifespan_context = lifespan
