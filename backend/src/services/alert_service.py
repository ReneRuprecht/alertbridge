import os
from typing import List

import requests

from src.db import read_alerts, save_alert
from src.models import Alert

ALERTMANAGER_URL = os.environ.get("ALERTMANAGER_URL", None)


def fetch_and_store_alerts() -> list[Alert]:
    if not ALERTMANAGER_URL:
        return []

    r = requests.get(ALERTMANAGER_URL)
    data = r.json()

    alerts = []
    for alert_json in data:
        alert_obj = Alert.from_json(alert_json)
        save_alert(alert_obj)
        alerts.append(alert_obj)

    return alerts


def fetch_alerts() -> list[Alert]:
    alert_records = read_alerts()
    alerts: List[Alert] = []

    for row in alert_records:
        alert = Alert.from_psql(row)
        alerts.append(alert)

    return alerts
