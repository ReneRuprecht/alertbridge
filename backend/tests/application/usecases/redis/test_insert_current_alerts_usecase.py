from unittest.mock import AsyncMock

import pytest

from src.application.usecases.redis.insert_current_alerts_usecase import (
    InsertCurrentAlertsUseCase,
)
from src.domain.entities.alert import Alert
from src.domain.ports.alert_cache_repository_port import (
    AlertCacheRepositoryPort,
)


@pytest.mark.asyncio
async def test_insert_current_alerts_returns_repo_count_and_passes_input():
    mock_repo = AsyncMock(spec=AlertCacheRepositoryPort)
    mock_repo.insert_current_alerts.return_value = 1

    usecase = InsertCurrentAlertsUseCase(cache_repo=mock_repo)

    alerts_to_save = [
        Alert(
            alertname="InstanceDown",
            status="firing",
            labels={"instance": "server1"},
            fingerprint="abc123",
            starts_at=None,
            ended_at=None,
            updated_at=None,
        )
    ]

    processed_alerts = await usecase.execute(alerts=alerts_to_save)

    assert isinstance(processed_alerts, int)
    assert processed_alerts == 1

    mock_repo.insert_current_alerts.assert_awaited_once_with(
        alerts=alerts_to_save
    )


@pytest.mark.asyncio
async def test_insert_current_alerts_with_empty_list():
    mock_repo = AsyncMock(spec=AlertCacheRepositoryPort)
    mock_repo.insert_current_alerts.return_value = 0

    usecase = InsertCurrentAlertsUseCase(cache_repo=mock_repo)

    empty_alerts = []

    processed_alerts = await usecase.execute(alerts=empty_alerts)

    assert isinstance(processed_alerts, int)
    assert processed_alerts == 0

    mock_repo.insert_current_alerts.assert_awaited_once_with(
        alerts=empty_alerts
    )
