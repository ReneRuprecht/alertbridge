from src.domain.entities.alert import Alert
from src.domain.ports.alert_historical_repository_port import (
    AlertHistoricalRepositoryPort,
)


class SaveHistoricalAlertsUseCase:
    def __init__(self, db_repo: AlertHistoricalRepositoryPort):
        self.db_repo = db_repo

    async def execute(self, alerts: list[Alert]) -> int:
        return await self.db_repo.save_historical_alerts(alerts=alerts)
