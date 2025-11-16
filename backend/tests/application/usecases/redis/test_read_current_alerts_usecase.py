from unittest.mock import AsyncMock

import pytest

from src.application.dtos.alert_redis_dto import AlertRedisDto
from src.application.usecases.redis.read_current_alerts_usecase import (
    ReadCurrentAlertsUseCase,
)
from src.domain.ports.alert_cache_repository_port import (
    AlertCacheRepositoryPort,
)


@pytest.mark.asyncio
async def test_execute_returns_alerts_dto_from_cache():
    mock_repo = AsyncMock(spec=AlertCacheRepositoryPort)
    mock_repo.read_current_alerts.return_value = [
        AlertRedisDto(
            alertname="InstanceDown",
            status="firing",
            fingerprint="abc123",
            instance="server1",
        )
    ]

    usecase = ReadCurrentAlertsUseCase(cache_repo=mock_repo)
    alerts = await usecase.execute()

    assert isinstance(alerts, list)
    assert len(alerts) == 1
    assert isinstance(alerts[0], AlertRedisDto)
    assert alerts[0].alertname == "InstanceDown"
    assert alerts[0].fingerprint == "abc123"
    assert alerts[0].instance == "server1"
    assert alerts[0].status == "firing"

    mock_repo.read_current_alerts.assert_awaited_once()


@pytest.mark.asyncio
async def test_execute_returns_empty_list_when_cache_empty():
    mock_repo = AsyncMock(spec=AlertCacheRepositoryPort)
    mock_repo.read_current_alerts.return_value = []

    usecase = ReadCurrentAlertsUseCase(cache_repo=mock_repo)
    alerts = await usecase.execute()

    assert alerts == []
    mock_repo.read_current_alerts.assert_awaited_once()
