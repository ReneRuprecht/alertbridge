from src.domain.entities.alert import Alert
from src.domain.ports.alert_cache_repository_port import (
    AlertCacheRepositoryPort,
)


class InsertCurrentAlertsUseCase:
    def __init__(self, cache_repo: AlertCacheRepositoryPort):
        self.cache_repo = cache_repo

    async def execute(self, alerts: list[Alert]) -> int:
        return await self.cache_repo.insert_current_alerts(alerts=alerts)
