import logging

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from src.repository.psql_client import PSQLClient
from src.services.alert_service import AlertService

logging.basicConfig(level=logging.INFO)

app = FastAPI()

origins = ["http://localhost:5173"]

app.add_middleware(CORSMiddleware, allow_origins=origins)

psql_client: PSQLClient = PSQLClient()
alert_service: AlertService = AlertService(psql_client=psql_client)


@app.get("/")
async def root():
    alerts = alert_service.fetch_and_store_alerts_from_alertmanager()

    return {"alerts": alerts}


@app.get("/alerts")
async def alerts():
    alerts = alert_service.fetch_alerts_from_db()

    return {"alerts": alerts}
