import json
import os
from typing import List

import psycopg2

from src.models import Alert

conn = psycopg2.connect(
    dbname=os.getenv("POSTGRES_DB", "alerts_db"),
    user=os.getenv("POSTGRES_USER", "root"),
    password=os.getenv("POSTGRES_PASSWORD", "root"),
    host=os.getenv("POSTGRES_HOST", "localhost"),
    port=int(os.getenv("POSTGRES_PORT", 5432)),
)
cursor = conn.cursor()


def save_alert(alert: Alert):
    cursor.execute(
        """
        INSERT INTO alerts (alertname, status, labels, fingerprint)
        VALUES (%s, %s, %s, %s)
        ON CONFLICT (fingerprint) DO UPDATE
        SET status = EXCLUDED.status,
            labels = EXCLUDED.labels
        """,
        (alert.alertname, alert.status, json.dumps(alert.labels), alert.fingerprint),
    )
    conn.commit()


def read_alerts() -> List[tuple]:
    cursor.execute("SELECT * FROM alerts;")
    alert_records = cursor.fetchall()

    return alert_records
