from src.application.dto.alert_redis_dto import AlertRedisDto
from src.domain.ports.alert_cache_repository_port import (
    AlertCacheRepositoryPort,
)


class ReadCurrentAlertsUseCase:
    def __init__(self, cache_repo: AlertCacheRepositoryPort):
        self.cache_repo = cache_repo

    async def execute(self) -> list[AlertRedisDto]:
        return await self.cache_repo.read_current_alerts()
