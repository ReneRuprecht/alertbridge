from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from src.services.alert_service import fetch_alerts, fetch_and_store_alerts

app = FastAPI()

origins = [
        "http://localhost:5173"
        ]

app.add_middleware(
        CORSMiddleware,
        allow_origins=origins
        )

@app.get("/")
async def root():
    alerts = fetch_and_store_alerts()

    return {"alerts": alerts}


@app.get("/alerts")
async def alerts():
    alerts = fetch_alerts()

    return {"alerts": alerts}
