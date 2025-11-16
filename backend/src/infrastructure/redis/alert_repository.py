import redis.asyncio as redis

from src.application.dtos.alert_redis_dto import AlertRedisDto
from src.application.ports.logger_factory_port import LoggerFactoryPort
from src.domain.entities.alert import Alert
from src.domain.ports.alert_cache_repository_port import (
    AlertCacheRepositoryPort,
)
from src.infrastructure.redis.mappers.alert_to_redis import alert_to_redis
from src.infrastructure.redis.mappers.redis_to_alert_dto import (
    redis_to_alert_dto,
)


class AlertRedisCacheRepository(AlertCacheRepositoryPort):
    def __init__(
        self, cache_client: redis.Redis, logger_factory: LoggerFactoryPort
    ):
        self.cache_client: redis.Redis = cache_client
        self.logger = logger_factory.get("AlertRedisCacheRepository")

    async def insert_current_alerts(self, alerts: list[Alert]) -> int:

        try:
            async with self.cache_client.pipeline(transaction=True) as pipe:
                for alert in alerts:
                    ex = 120 if alert.status == "resolved" else 300
                    await pipe.set(
                        f"alert:{alert.fingerprint}",
                        alert_to_redis(alert),
                        ex=ex,
                    )
                await pipe.execute()

            self.logger.info(f"Inserted {len(alerts)}")
            return len(alerts)

        except Exception as e:
            self.logger.info(f"Failed to insert alerts: {e}")
            return 0

    async def read_current_alerts(self) -> list[AlertRedisDto]:
        dtos: list[AlertRedisDto] = []

        try:
            keys = await self.cache_client.keys("alert:*")

            async with self.cache_client.pipeline() as pipe:
                for key in keys:
                    await pipe.get(key)
                value = await pipe.execute()

            for val in value:
                dtos.append(redis_to_alert_dto(data=val))

            return dtos

        except Exception as e:
            self.logger.error(f"Error while reading alert: {e}")
            return []
