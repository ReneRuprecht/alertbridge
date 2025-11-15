from src.application.usecases.http.fetch_alerts_usecase import (
    FetchAlertsUseCase,
)
from src.config.services.http.http import ALERTMANAGER_URL
from src.infrastructure.alertmanager.alertmanager_api_client import (
    AlertmanagerApiClient,
)
from src.infrastructure.http.client import get_http_client


async def get_alertmanager_api_client() -> AlertmanagerApiClient:
    http_client = await get_http_client()
    return AlertmanagerApiClient(
        alertmanager_url=ALERTMANAGER_URL, http_client=http_client
    )


async def get_fetch_alerts_usecase() -> FetchAlertsUseCase:
    source = await get_alertmanager_api_client()
    return FetchAlertsUseCase(external_source=source)
