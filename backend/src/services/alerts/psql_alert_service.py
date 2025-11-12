import json
import logging
from typing import Optional

from src.models.alert import Alert
from src.repository.psql_client import PSQLClient


class PSQLAlertService:
    def __init__(
        self,
        psql_client: PSQLClient,
        logger: Optional[logging.Logger] = None,
    ):
        self.psql_client = psql_client
        self.logger = logger or logging.getLogger("PSQLAlertService")

    async def read_alerts_from_db(self) -> list[Alert]:
        """
        Fetch alerts from db

        Return list of alerts
        """
        try:
            rows = await self.psql_client.read_alerts()
        except Exception as e:

            self.logger.error("Error reading from DB: %s", e)
            return []

        alerts: list[Alert] = []
        for row in rows:
            try:
                alert = Alert.from_psql(row)
            except Exception as e:
                self.logger.error(
                    "Could not map DB row to Alert: %s | %s", e, row
                )
                return []
            alerts.append(alert)

        self.logger.info("Fetched %d alerts", len(alerts))
        return alerts

    async def save_alerts(self, alerts: list[Alert]) -> int:
        """
        Save alerts to db

        Return count of inserted alerts
        """
        values = [
            (
                alert.alertname,
                alert.status,
                json.dumps(alert.labels),
                alert.fingerprint,
                alert.starts_at,
                alert.ended_at,
                alert.updated_at,
            )
            for alert in alerts
        ]
        if not values:
            return 0

        return await self.psql_client.save_alerts(values=values)
