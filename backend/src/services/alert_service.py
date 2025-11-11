import logging
import os
from typing import List, Optional

from datetime import datetime, timezone

import requests
from src.models.alert import Alert
from src.repository.psql_client import PSQLClient


class AlertService:
    def __init__(
        self,
        psql_client: PSQLClient,
        alertmanager_url: Optional[str] = None,
        logger: Optional[logging.Logger] = None,
    ):
        """
        Initialize the AlertService.
        """
        self.psql_client = psql_client
        self.alertmanager_url = alertmanager_url or os.getenv("ALERTMANAGER_URL", None)
        self.logger = logger or logging.getLogger("AlertService")

    def _init_db_data_from_alertmanager(self) -> int:
        """
        Fetch alerts from alertmanager (/api/v2/alerts) and store them in DB
        """
        if not self.alertmanager_url:
            self.logger.warning("ALERTMANAGER_URL not set")
            return 0

        try:
            resp = requests.get(self.alertmanager_url + "/api/v2/alerts")
            resp.raise_for_status()
            alerts_json = resp.json()
            if not alerts_json:
                self.logger.warning("Alerts from Alertmanager were empty")
                return 0

            if not isinstance(alerts_json, list):
                self.logger.warning("Invalid payload: expected a list of alerts")
                return 0

            processed_alerts = self.save_alerts_to_db(alerts_json)

            return processed_alerts
        except Exception as e:
            self.logger.error("Error fetching from Alertmanager: %s", e)
            return 0

    def fetch_alerts_from_db(self) -> list[Alert]:
        """
        Fetch alerts from db
        """
        try:
            rows = self.psql_client.read_alerts()
        except Exception as e:

            self.logger.error("Error reading from DB: %s", e)
            return []

        alerts: List[Alert] = []
        for row in rows:
            try:
                alert = Alert.from_psql(row)
            except Exception as e:
                self.logger.error("Could not map DB row to Alert: %s | %s", e, row)
                continue
            alerts.append(alert)

        self.logger.info("Fetched %d alerts", len(alerts))
        return alerts

    def process_webhook_payload(self, payload: dict) -> list[dict]:
        alerts_json = payload.get("alerts", [])

        if not isinstance(alerts_json, list) or not alerts_json:
            return []
        return alerts_json

    def save_alerts_to_db(self, alerts_json: list[dict]) -> int:
        alerts: list[Alert] = []

        for alert in alerts_json:
            alert_obj = Alert.from_json(alert)

            # Set ended_at only if resolved
            if alert_obj.status == "resolved":
                alert_obj.ended_at = datetime.now(timezone.utc)

            alert_obj.updated_at = datetime.now(timezone.utc)
            alerts.append(alert_obj)

        if not alerts:
            return 0

        return self.psql_client.save_alerts_batch(alerts)

