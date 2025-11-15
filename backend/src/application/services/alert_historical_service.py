from src.application.ports.logger_factory_port import LoggerFactoryPort
from src.application.usecases.http.fetch_alerts_usecase import (
    FetchAlertsUseCase,
)
from src.application.usecases.postgres.read_historical_alerts_usecase import (
    ReadHistoricalAlertsUseCase,
)
from src.application.usecases.postgres.save_historical_alerts_usecase import (
    SaveHistoricalAlertsUseCase,
)
from src.domain.entities.alert import Alert


class AlertHistoricalService:
    def __init__(
        self,
        read_historical_alerts_usecase: ReadHistoricalAlertsUseCase,
        save_historical_alerts_usecase: SaveHistoricalAlertsUseCase,
        fetch_alerts_usecase: FetchAlertsUseCase,
        logger_factory: LoggerFactoryPort,
    ):
        self._read_historical_alerts_usecase = read_historical_alerts_usecase
        self._save_historical_alerts_usecase = save_historical_alerts_usecase
        self._fetch_alerts_usecase = fetch_alerts_usecase
        self.logger = logger_factory.get(name="AlertHistoricalService")

    async def read_historical_alerts(self):
        return await self._read_historical_alerts_usecase.execute()

    async def save_historical_alerts(self, alerts: list[Alert]) -> int:
        return await self._save_historical_alerts_usecase.execute(
            alerts=alerts
        )
