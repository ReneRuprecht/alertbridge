from src.domain.entities.alert import Alert
from src.domain.ports.alert_historical_repository_port import (
    AlertHistoricalRepositoryPort,
)


class ReadHistoricalAlertsUseCase:
    def __init__(self, db_repo: AlertHistoricalRepositoryPort):
        self.db_repo = db_repo

    async def execute(self) -> list[Alert]:
        return await self.db_repo.read_historical_alerts()
