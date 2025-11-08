import os

import requests
from fastapi import FastAPI

from src.models import Alert

app = FastAPI()

ALERTMANAGER_URL = os.environ.get("ALERTMANAGER_URL", None)


@app.get("/")
async def root():
    if not ALERTMANAGER_URL:
        return {"message": "alertmanager url is missing"}

    r = requests.get(url=ALERTMANAGER_URL)
    response = r.json()

    alerts = []

    for alert in response:
        alert_obj = Alert.from_json(alert)
        alerts.append(alert_obj)

    return {"alerts": alerts}
