from fastapi import FastAPI

from src.services.alert_service import fetch_alerts, fetch_and_store_alerts

app = FastAPI()


@app.get("/")
async def root():
    alerts = fetch_and_store_alerts()

    return {"alerts": alerts}


@app.get("/alerts")
async def alerts():
    alerts = fetch_alerts()

    return {"alerts": alerts}
