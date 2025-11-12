import logging
import os
from typing import List, Optional, Tuple

from psycopg_pool import AsyncConnectionPool


class PSQLClient:
    def __init__(
        self,
        dbname: str = os.getenv("POSTGRES_DB", "postgres"),
        user: str = os.getenv("POSTGRES_USER", "postgres"),
        password: str = os.getenv("POSTGRES_PASSWORD", "postgres"),
        host: str = os.getenv("POSTGRES_HOST", "localhost"),
        port: int = int(os.getenv("POSTGRES_PORT", 5432)),
        logger: Optional[logging.Logger] = None,
    ):
        self.logger = logger or logging.getLogger("PSQLClient")
        self.dbname = dbname
        self.user = user
        self.password = password
        self.host = host
        self.port = port

        self.conninfo = f"postgres://{user}:{password}@{host}:{port}/{dbname}"
        self.pool: Optional[AsyncConnectionPool] = None

    async def create_pool(
        self, min_size: int = 1, max_size: int = 4, open: bool = False
    ):
        """
        Create AsyncConnectionPool
        """
        if self._check_pool(show_log=False):
            self.logger.info("Pool is already created")
            return

        self.logger.info("Pool Initializing")
        self.pool = AsyncConnectionPool(
            conninfo=self.conninfo,
            min_size=min_size,
            max_size=max_size,
            open=open,
        )

        if open:
            self.logger.info("Wait for pool to be ready")
            await self.pool.wait()

    async def close_pool(self):
        """
        Close AsyncConnectionPool
        """
        if not self._check_pool(show_log=False):
            return
        assert self.pool is not None

        await self.pool.close()
        self.logger.info("Pool closed")

    def _check_pool(self, strict: bool = False, show_log: bool = True) -> bool:
        if not isinstance(self.pool, AsyncConnectionPool):
            msg = "Pool not initialized, call create_pool before usage"
            if show_log:
                self.logger.error(msg)
            if strict:
                raise RuntimeError(msg)
            return False
        return True

    async def read_alerts(self) -> List[Tuple]:
        """
        Read alerts from database
        """
        self._check_pool(strict=True)
        assert self.pool is not None

        async with self.pool.connection() as conn:
            try:
                async with conn.cursor() as cur:
                    query = "SELECT * FROM alerts_history;"
                    await cur.execute(query=query)
                    rows = await cur.fetchall()
                    return rows
            except Exception as e:
                self.logger.error("Error while fetching alerts: %s", e)
                return []

    async def save_alerts(self, values: list[tuple]) -> int:
        """
        Save multiple alerts to db
        """
        self._check_pool(strict=True)
        assert self.pool is not None

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

                    await cur.executemany(query, values)

                await conn.commit()
                self.logger.info("Inserted %d alerts", len(values))
                return len(values)

            except Exception as e:
                self.logger.error("Insertion of alerts failed: %s", e)
                await conn.rollback()

        return 0
