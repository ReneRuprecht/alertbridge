from unittest.mock import AsyncMock

import pytest

from src.application.ports.alert_external_source_port import (
    AlertExternalSourcePort,
)
from src.application.usecases.http.fetch_alerts_usecase import (
    FetchAlertsUseCase,
)
from src.domain.entities.alert import Alert


@pytest.mark.asyncio
async def test_execute_returns_alerts():
    mock_external_source = AsyncMock(spec=AlertExternalSourcePort)
    mock_external_source.fetch_alerts.return_value = [
        Alert(
            alertname="HighCPU",
            status="firing",
            labels={},
            fingerprint="abc123",
            starts_at=None,
            ended_at=None,
            updated_at=None,
        ),
        Alert(
            alertname="InstanceDown",
            status="firing",
            labels={"instance": "server1"},
            fingerprint="abc123",
            starts_at=None,
            ended_at=None,
            updated_at=None,
        ),
    ]

    usecase = FetchAlertsUseCase(external_source=mock_external_source)
    alerts = await usecase.execute()

    assert isinstance(alerts, list)
    assert len(alerts) == 2
    assert isinstance(alerts[0], Alert)
    assert alerts[0].alertname == "HighCPU"

    assert isinstance(alerts[1], Alert)
    assert alerts[1].alertname == "InstanceDown"
    assert alerts[1].labels["instance"] == "server1"

    mock_external_source.fetch_alerts.assert_awaited_once()


@pytest.mark.asyncio
async def test_execute_returns_empty_list_if_no_alerts():
    mock_external_source = AsyncMock(spec=AlertExternalSourcePort)
    mock_external_source.fetch_alerts.return_value = []

    usecase = FetchAlertsUseCase(external_source=mock_external_source)
    alerts = await usecase.execute()

    assert alerts == []
    assert isinstance(alerts, list)

    mock_external_source.fetch_alerts.assert_awaited_once()
