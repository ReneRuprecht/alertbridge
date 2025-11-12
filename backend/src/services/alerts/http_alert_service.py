import logging
import os
from typing import Optional

import aiohttp


class HttpAlertService:
    def __init__(
        self,
        alertmanager_url: str = "",
        logger: Optional[logging.Logger] = None,
    ):

        self.alertmanager_url = alertmanager_url or os.getenv(
            "ALERTMANAGER_URL", None
        )
        self.logger = logger or logging.getLogger("HttpAlertService")

    async def fetch_alerts_from_alertmanager(self) -> list[dict]:
        """
        Fetch alerts from alertmanager /api/v2/alerts

        Return the alerts as json or none
        """
        if not self.alertmanager_url:
            self.logger.warning("ALERTMANAGER_URL not set")
            return []

        try:
            alerts_json = []

            async with aiohttp.ClientSession() as session:
                async with session.get(
                    self.alertmanager_url + "/api/v2/alerts"
                ) as resp:
                    alerts_json = await resp.json()

            if not alerts_json:
                self.logger.warning("Alerts from Alertmanager were empty")
                return []

            if not isinstance(alerts_json, list):
                self.logger.warning(
                    "Invalid payload: expected a list of alerts"
                )
                return []

            return alerts_json

        except Exception as e:
            self.logger.error("Error fetching from Alertmanager: %s", e)
            return []

    def extract_alerts(self, payload: dict) -> list[dict]:
        """
        Extract alerts from the alertmanager webhook payload

        Return alerts from the payload
        """
        alerts_json = payload.get("alerts", [])

        if not isinstance(alerts_json, list) or not alerts_json:
            return []
        return alerts_json
