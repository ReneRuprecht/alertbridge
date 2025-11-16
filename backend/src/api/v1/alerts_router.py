from fastapi import APIRouter, Depends, Request

from src.api.v1.deps import (
    get_alert_cache_service,
    get_alert_historical_service,
)
from src.api.v1.mappers.alertmanager_api_mapper import (
    alertmanager_payload_to_alerts,
)
from src.application.dtos.alert_redis_dto import AlertRedisDto
from src.application.ports.logger_factory_port import LoggerFactoryPort
from src.application.ports.logger_port import LoggerPort
from src.application.services.alert_cache_service import AlertCacheService
from src.application.services.alert_historical_service import (
    AlertHistoricalService,
)
from src.application.services.logger_factory import get_logger_factory
from src.domain.entities.alert import Alert

router = APIRouter(prefix="/api/v1/alerts")


@router.get("/")
async def alerts(
    alert_historical_service: AlertHistoricalService = Depends(
        get_alert_historical_service
    ),
):
    alerts = await alert_historical_service.read_historical_alerts()

    return {"alerts": alerts}


@router.post("/")
async def alert_manager_webhook(
    request: Request,
    alert_historical_service: AlertHistoricalService = Depends(
        get_alert_historical_service
    ),
    alert_cache_service: AlertCacheService = Depends(get_alert_cache_service),
    logger_factory: LoggerFactoryPort = Depends(get_logger_factory),
):
    logger: LoggerPort = logger_factory.get("AlertsRouterPost")
    try:
        payload = await request.json()
        alerts: list[Alert] = await alertmanager_payload_to_alerts(
            payload=payload
        )
        await alert_historical_service.save_historical_alerts(alerts)

        await alert_cache_service.insert_current_alerts(alerts)

    except Exception as e:
        logger.error(f"Error from post endpoint: {e}")


@router.get("/current")
async def alerts_current(
    alert_cache_service: AlertCacheService = Depends(get_alert_cache_service),
):
    alert_redis_dtos: list[AlertRedisDto] = (
        await alert_cache_service.read_current_alerts()
    )

    return {"alerts": alert_redis_dtos}
