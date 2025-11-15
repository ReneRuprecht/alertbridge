from contextlib import asynccontextmanager

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from psycopg_pool import AsyncConnectionPool

from src.api.v1.alerts_router import router as alerts_router
from src.api.v1.dependencies.cache.cache import (
    get_insert_current_alerts_usecase,
)
from src.api.v1.dependencies.db.db import get_alert_repository
from src.api.v1.dependencies.http.http import get_fetch_alerts_usecase
from src.application.services.logger_factory import get_logger_factory
from src.domain.entities.alert import Alert
from src.infrastructure.http.client import shutdown_http_client
from src.infrastructure.postgres.alert_repository import (
    AlertPostgresHistoricalRepository,
)

app = FastAPI()

origins = ["http://localhost:5173"]

app.add_middleware(CORSMiddleware, allow_origins=origins)


async def fetch_and_save_current_alerts():
    fetch_alerts_usecase = await get_fetch_alerts_usecase()
    alerts: list[Alert] = await fetch_alerts_usecase.execute()
    insert_current_alerts_usecase = await get_insert_current_alerts_usecase()
    await insert_current_alerts_usecase.execute(alerts=alerts)


@asynccontextmanager
async def lifespan(app: FastAPI):
    logger_factory = await get_logger_factory()
    logger = logger_factory.get("App")
    logger.info("Starting postgres connection")
    await fetch_and_save_current_alerts()

    yield

    logger.info("Close pool")
    psql: AlertPostgresHistoricalRepository = await get_alert_repository()
    pool: AsyncConnectionPool = psql.pool
    await pool.close()
    await shutdown_http_client()


app.router.lifespan_context = lifespan

app.include_router(alerts_router)
