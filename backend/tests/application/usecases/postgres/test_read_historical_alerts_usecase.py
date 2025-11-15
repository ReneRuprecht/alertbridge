from unittest.mock import AsyncMock

import pytest

from src.application.usecases.postgres.read_historical_alerts_usecase import (
    ReadHistoricalAlertsUseCase,
)
from src.domain.entities.alert import Alert
from src.domain.ports.alert_historical_repository_port import (
    AlertHistoricalRepositoryPort,
)


@pytest.mark.asyncio
async def test_execute_returns_alerts_from_repo():
    mock_repo = AsyncMock(spec=AlertHistoricalRepositoryPort)
    mock_repo.read_historical_alerts.return_value = [
        Alert(
            alertname="InstanceDown",
            status="firing",
            labels={"instance": "server1"},
            fingerprint="abc123",
        )
    ]

    usecase = ReadHistoricalAlertsUseCase(db_repo=mock_repo)
    alerts = await usecase.execute()

    assert isinstance(alerts, list)
    assert len(alerts) == 1
    assert isinstance(alerts[0], Alert)
    assert alerts[0].alertname == "InstanceDown"
    assert alerts[0].labels["instance"] == "server1"
    assert alerts[0].fingerprint == "abc123"
    assert alerts[0].status == "firing"

    mock_repo.read_historical_alerts.assert_awaited_once()


@pytest.mark.asyncio
async def test_execute_returns_empty_list_when_repo_empty():
    mock_repo = AsyncMock(spec=AlertHistoricalRepositoryPort)
    mock_repo.read_historical_alerts.return_value = []

    usecase = ReadHistoricalAlertsUseCase(db_repo=mock_repo)
    alerts = await usecase.execute()

    assert alerts == []
    mock_repo.read_historical_alerts.assert_awaited_once()
