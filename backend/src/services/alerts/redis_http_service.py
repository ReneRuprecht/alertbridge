import logging
from typing import Optional

from src.models.alert import Alert
from src.models.alert_frontend import FrontendAlert
from src.repository.redis_client import RedisClient


class RedisAlertService:
    def __init__(
        self,
        redis_client: RedisClient,
        logger: Optional[logging.Logger] = None,
    ):
        self.redis_client = redis_client
        self.logger = logger or logging.getLogger("RedisAlertService")

    async def read_all_alerts_from_redis(self) -> list[FrontendAlert]:
        """
        Read the alerts from redis that are needed for the frontend

        Return list of frontendalerts
        """

        alerts_dict = await self.redis_client.read_alerts_frontend()

        alerts: list[FrontendAlert] = []

        for alert in alerts_dict:
            alerts.append(FrontendAlert.from_dict(alert))

        return alerts

    async def save_alerts(self, alerts: list[Alert]) -> int:
        """
        Save the alerts to redis

        Return count of inserted alerts
        """

        alerts_dict: list[dict] = []

        for alert in alerts:
            alerts_dict.append(alert.to_redis_minimal_dict())

        processed_items = await self.redis_client.save_alerts(alerts_dict)

        return processed_items
