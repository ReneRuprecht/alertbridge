import logging
from typing import Optional

from src.models.alert import Alert
from src.models.alert_frontend import FrontendAlert
from src.services.alerts.http_alert_service import HttpAlertService
from src.services.alerts.psql_alert_service import PSQLAlertService
from src.services.alerts.redis_http_service import RedisAlertService


class AlertService:
    def __init__(
        self,
        psql_service: PSQLAlertService,
        redis_service: RedisAlertService,
        http_alert_service: HttpAlertService,
        logger: Optional[logging.Logger] = None,
    ):
        """
        Initialize the AlertService.
        """
        self.psql_service = psql_service
        self.redis_service = redis_service
        self.http_service = http_alert_service

        self.logger = logger or logging.getLogger("AlertService")

    async def initial_setup_from_alertmanager(self) -> int:
        """
        Initialize the db with data from the alertmanager

        Return the number of inserted records
        """

        try:
            alerts_json = (
                await self.http_service.fetch_alerts_from_alertmanager()
            )
            if not alerts_json:
                return 0

            processed_alerts = await self.save_alerts(alerts_json)

            return processed_alerts
        except Exception as e:
            self.logger.error("Error fetching from Alertmanager: %s", e)
            return 0

    async def save_alerts_from_webhook(self, payload: dict) -> int:
        """
        Proccess the alertmanager webhook payload

        Return count of inserted alerts
        """
        alerts_json = self.http_service.extract_alerts(payload=payload)
        if not alerts_json:
            return 0

        return await self.save_alerts(alerts_json=alerts_json)

    async def save_alerts(self, alerts_json: list[dict]) -> int:
        """
        Save alerts to psql and redis

        Return count of inserted alerts
        """
        alerts = self.parse_alerts_json(alerts_json)

        if not alerts:
            return 0

        redis_result_count = await self.redis_service.save_alerts(alerts)
        psql_result_count = await self.psql_service.save_alerts(alerts)

        if redis_result_count != psql_result_count:
            self.logger.info("Redis / PSQL inserted alerts not the same!")

        return psql_result_count

    async def read_alerts_from_db(self):
        return await self.psql_service.read_alerts_from_db()

    def parse_alerts_json(self, alerts_json: list[dict]) -> list[Alert]:
        alerts: list[Alert] = []

        for alert in alerts_json:
            alerts.append(Alert.from_json(alert))

        return alerts

    async def read_alerts_from_redis(self) -> list[FrontendAlert]:
        return await self.redis_service.read_all_alerts_from_redis()
