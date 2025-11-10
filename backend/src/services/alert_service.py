import logging
import os
from typing import List, Optional

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

    def fetch_and_store_alerts_from_alertmanager(self) -> List[Alert]:
        """
        Fetch alerts from alertmanager (/api/v2/alerts) and store them in DB
        """
        if not self.alertmanager_url:
            self.logger.warning("ALERTMANAGER_URL not set")
            return []

        try:
            resp = requests.get(self.alertmanager_url + "/api/v2/alerts")
            resp.raise_for_status()
            data = resp.json()
        except Exception as e:
            self.logger.error("Error fetching from Alertmanager: %s", e)
            return []

        alerts: List[Alert] = []
        for alert_json in data:
            try:
                alert_obj = Alert.from_json(alert_json)
            except Exception as e:
                self.logger.warning(
                    "Could not parse alert JSON: %s | %s", e, alert_json
                )
                continue

            try:
                self.psql_client.save_alert(alert_obj)
            except Exception as e:
                self.logger.error("Error saving alert to DB: %s", e)

            alerts.append(alert_obj)

        self.logger.info("Fetched and stored %d alerts", len(alerts))
        return alerts

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
