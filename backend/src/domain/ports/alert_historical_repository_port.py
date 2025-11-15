from abc import ABC, abstractmethod

from src.domain.entities.alert import Alert


class AlertHistoricalRepositoryPort(ABC):
    @abstractmethod
    async def save_historical_alerts(self, alerts: list[Alert]) -> int:
        pass

    @abstractmethod
    async def read_historical_alerts(self) -> list[Alert]:
        pass
