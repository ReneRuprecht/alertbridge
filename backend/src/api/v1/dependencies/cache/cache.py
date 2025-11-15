from src.application.ports.logger_factory_port import LoggerFactoryPort
from src.application.services.logger_factory import get_logger_factory
from src.application.usecases.redis.insert_current_alerts_usecase import (
    InsertCurrentAlertsUseCase,
)
from src.application.usecases.redis.read_current_alerts_usecase import (
    ReadCurrentAlertsUseCase,
)
from src.config.services.cache.redis import get_redis_cache_client
from src.domain.ports.alert_cache_repository_port import (
    AlertCacheRepositoryPort,
)
from src.infrastructure.redis.alert_repository import AlertRedisCacheRepository


async def get_alert_cache_repository() -> AlertCacheRepositoryPort:
    cache_client = await get_redis_cache_client()
    logger_factory: LoggerFactoryPort = await get_logger_factory()

    return AlertRedisCacheRepository(
        cache_client=cache_client, logger_factory=logger_factory
    )


async def get_read_current_alerts_usecase() -> ReadCurrentAlertsUseCase:
    cache_repo = await get_alert_cache_repository()
    return ReadCurrentAlertsUseCase(cache_repo=cache_repo)


async def get_insert_current_alerts_usecase() -> InsertCurrentAlertsUseCase:
    cache_repo = await get_alert_cache_repository()
    return InsertCurrentAlertsUseCase(cache_repo=cache_repo)
