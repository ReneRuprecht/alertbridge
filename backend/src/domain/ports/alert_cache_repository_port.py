from abc import ABC, abstractmethod

from src.application.dtos.alert_redis_dto import AlertRedisDto
from src.domain.entities.alert import Alert


class AlertCacheRepositoryPort(ABC):
    @abstractmethod
    async def insert_current_alerts(self, alerts: list[Alert]) -> int:
        pass

    @abstractmethod
    async def read_current_alerts(self) -> list[AlertRedisDto]:
        pass
