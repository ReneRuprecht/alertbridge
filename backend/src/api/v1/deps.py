from src.api.v1.dependencies.cache.cache import (
    get_insert_current_alerts_usecase,
    get_read_current_alerts_usecase,
)
from src.api.v1.dependencies.db.db import (
    get_read_historical_alerts_usecase,
    get_save_historical_alerts_usecase,
)
from src.api.v1.dependencies.http.http import get_fetch_alerts_usecase
from src.application.services.alert_cache_service import AlertCacheService
from src.application.services.alert_historical_service import (
    AlertHistoricalService,
)
from src.application.services.logger_factory import get_logger_factory


async def get_alert_cache_service() -> AlertCacheService:
    read_current_alerts_usecase = await get_read_current_alerts_usecase()
    insert_current_alerts_usecase = await get_insert_current_alerts_usecase()
    return AlertCacheService(
        read_current_alerts_usecase=read_current_alerts_usecase,
        insert_current_alerts_usecase=insert_current_alerts_usecase,
    )


async def get_alert_historical_service() -> AlertHistoricalService:
    read_historical_alerts_usecase = await get_read_historical_alerts_usecase()
    save_historical_alerts_usecase = await get_save_historical_alerts_usecase()
    fetch_alert_usecase = await get_fetch_alerts_usecase()
    logger_factory = await get_logger_factory()
    return AlertHistoricalService(
        read_historical_alerts_usecase=read_historical_alerts_usecase,
        save_historical_alerts_usecase=save_historical_alerts_usecase,
        fetch_alerts_usecase=fetch_alert_usecase,
        logger_factory=logger_factory,
    )
