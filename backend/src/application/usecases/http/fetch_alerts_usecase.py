from src.application.ports.alert_external_source_port import (
    AlertExternalSourcePort,
)
from src.domain.entities.alert import Alert


class FetchAlertsUseCase:

    def __init__(self, external_source: AlertExternalSourcePort):
        self.external_source = external_source

    async def execute(self) -> list[Alert]:

        return await self.external_source.fetch_alerts()
