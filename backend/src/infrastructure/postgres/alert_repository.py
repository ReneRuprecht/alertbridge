from psycopg_pool import AsyncConnectionPool

from src.application.ports.logger_factory_port import LoggerFactoryPort
from src.domain.entities.alert import Alert
from src.domain.ports.alert_historical_repository_port import (
    AlertHistoricalRepositoryPort,
)

from .mappers.alert_dto_mapper import row_to_alert
from .mappers.alerts_to_postgres_params_mapper import (
    alerts_to_postgres_params,
)


class AlertPostgresHistoricalRepository(AlertHistoricalRepositoryPort):
    def __init__(
        self, pool: AsyncConnectionPool, logger_factory: LoggerFactoryPort
    ):
        self.pool: AsyncConnectionPool = pool
        self.logger = logger_factory.get(name="AlertPostgresHistoryRepository")

    async def read_historical_alerts(self) -> list[Alert]:
        async with self.pool.connection() as conn:
            try:
                async with conn.cursor() as cur:
                    query = "SELECT * FROM alerts_history;"
                    await cur.execute(query=query)
                    rows = await cur.fetchall()

                    alerts: list[Alert] = []

                    for row in rows:
                        alerts.append(row_to_alert(row))

                    return alerts

            except Exception as e:
                self.logger.error(f"Read historical alerts error: {e}")
                return []

    async def save_historical_alerts(self, alerts: list[Alert]) -> int:
        async with self.pool.connection() as conn:
            try:
                async with conn.cursor() as cur:
                    query = """
                    INSERT INTO alerts_history (
                        alertname, status, labels, fingerprint,
                        starts_at, ended_at, updated_at
                    )
                    VALUES (%s,%s,%s,%s,%s,%s,%s)
                    ON CONFLICT (fingerprint, status, starts_at) DO NOTHING;
                    """

                    await cur.executemany(
                        query, params_seq=alerts_to_postgres_params(alerts)
                    )
                await conn.commit()

                self.logger.info(f"Inserted {len(alerts)} alerts")
                return len(alerts)

            except Exception as e:
                self.logger.error(f"save historical alerts error: {e}")
                await conn.rollback()

        return 0
