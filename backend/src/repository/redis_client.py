import json
import logging
from typing import Optional

import redis.asyncio as redis

from src.models.alert import Alert
from src.models.alert_frontend import FrontendAlert


class RedisClient:
    def __init__(
        self,
        logger: Optional[logging.Logger] = None,
        host: str = "localhost",
        port: int = 6379,
        username: Optional[str] = None,
        password: Optional[str] = None,
        db: Optional[int] = 0,
    ):
        self.logger = logger or logging.getLogger("RedisClient")
        self.host = host
        self.port = port
        self.username = username
        self.password = password
        self.db = db
        if username and password:
            self.conninfo = f"redis://{username}:{password}@{host}:{port}/{db}"
        else:
            self.conninfo = f"redis://{host}:{port}/{db}"
        self.client: Optional[redis.Redis] = None

    def connect(self):
        self.logger.info("Connect to redis server")
        if self.client:
            return

        try:
            pool = redis.ConnectionPool.from_url(url=self.conninfo)
            self.client = redis.Redis(
                connection_pool=pool,
                decode_responses=True,
                encoding="utf-8",
            )
            self.logger.info("Established connection to redis server")
        except Exception as e:
            self.logger.error("Connection to redis server failed: %s", e)

    async def _check_conn(
        self, strict: bool = False, show_log: bool = True
    ) -> bool:
        if not isinstance(self.client, redis.Redis):
            msg = "Redis not initialized, call connect before usage"
            if show_log:
                self.logger.error(msg)
            if strict:
                raise RuntimeError(msg)
            return False
        return True

    async def insert_alerts(self, alerts: list[Alert]):
        await self._check_conn()

        assert self.client is not None
        try:
            async with self.client.pipeline(transaction=True) as pipe:
                for alert in alerts:
                    ex = 120 if alert.status == "resolved" else 300
                    await pipe.set(
                        f"alert:{alert.fingerprint}",
                        alert.to_redis(),
                        ex=ex,
                    )
                await pipe.execute()

            self.logger.info(f"Inserted {len(alerts)}")

        except Exception as e:
            self.logger.info(f"Failed to insert alerts: {e}")

    async def read_alert(self, alert: Alert) -> Optional[FrontendAlert]:
        await self._check_conn()
        assert self.client is not None

        try:
            data = await self.client.get(alert.fingerprint)
            if isinstance(data, bytes):
                data = data.decode()
            if isinstance(data, str):
                data = json.loads(data)

            frontend_alert = FrontendAlert.from_dict(data)
            self.logger.info(f"READ READIS {frontend_alert.fingerprint}")
            return frontend_alert

        except Exception as e:
            self.logger.error(f"Error while reading alert: {e}")
            return None

    async def read_alerts_frontend(self) -> list[FrontendAlert]:
        await self._check_conn()
        assert self.client is not None

        try:
            keys = await self.client.keys("alert:*")

            async with self.client.pipeline() as pipe:
                for key in keys:
                    await pipe.get(key)
                value = await pipe.execute()

            alerts: list[FrontendAlert] = []
            for val in value:
                alerts.append(FrontendAlert.from_dict(json.loads(val)))

            return alerts

        except Exception as e:
            self.logger.error(f"Error while reading alert: {e}")
            return []
