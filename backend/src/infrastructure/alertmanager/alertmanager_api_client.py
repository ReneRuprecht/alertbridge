import aiohttp

from src.domain.entities.alert import Alert
from src.domain.ports.alert_external_source_port import AlertExternalSourcePort
from src.infrastructure.alertmanager.mappers.alertmanager_mapper import (
    external_alert_json_to_domain,
)


class AlertmanagerApiClient(AlertExternalSourcePort):
    def __init__(
        self, alertmanager_url: str, http_client: aiohttp.ClientSession
    ):
        self.url = alertmanager_url
        self.http = http_client

    async def fetch_alerts(self) -> list[Alert]:

        async with self.http.get(self.url + "/api/v2/alerts") as resp:
            response = await resp.json()

            alerts: list[Alert] = []
            for data in response:
                alerts.append(external_alert_json_to_domain(data))

            return alerts
