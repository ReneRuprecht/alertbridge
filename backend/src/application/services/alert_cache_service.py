from src.application.dto.alert_redis_dto import AlertRedisDto
from src.application.usecases.redis.read_current_alerts_usecase import (
    ReadCurrentAlertsUseCase,
)
from src.application.usecases.redis.save_current_alerts_usecase import (
    InsertCurrentAlertsUseCase,
)
from src.domain.entities.alert import Alert


class AlertCacheService:
    def __init__(
        self,
        read_current_alerts_usecase: ReadCurrentAlertsUseCase,
        insert_current_alerts_usecase: InsertCurrentAlertsUseCase,
    ):
        self.read_current_alerts_usecase = read_current_alerts_usecase
        self.insert_current_alerts_usecase = insert_current_alerts_usecase

    async def insert_current_alerts(self, alerts: list[Alert]) -> int:
        return await self.insert_current_alerts_usecase.execute(alerts=alerts)

    async def read_current_alerts(self) -> list[AlertRedisDto]:
        return await self.read_current_alerts_usecase.execute()
