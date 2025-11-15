from abc import ABC, abstractmethod

from src.domain.entities.alert import Alert


class AlertExternalSourcePort(ABC):

    @abstractmethod
    async def fetch_alerts(self) -> list[Alert]:
        pass
