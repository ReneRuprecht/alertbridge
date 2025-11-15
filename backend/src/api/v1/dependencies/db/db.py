from src.application.services.logger_factory import get_logger_factory
from src.application.usecases.postgres.read_historical_alerts_usecase import (
    ReadHistoricalAlertsUseCase,
)
from src.application.usecases.postgres.save_alerts_usecase import (
    SaveHistoricalAlertsUseCase,
)
from src.config.services.db.postgres import get_postgres_db_pool
from src.infrastructure.postgres.alert_repository import (
    AlertPostgresHistoricalRepository,
)


async def get_alert_repository() -> AlertPostgresHistoricalRepository:
    logger_factory = await get_logger_factory()
    pool = await get_postgres_db_pool()
    await pool.open()
    await pool.wait()
    return AlertPostgresHistoricalRepository(
        pool=pool, logger_factory=logger_factory
    )


async def get_read_historical_alerts_usecase() -> ReadHistoricalAlertsUseCase:
    db_repo = await get_alert_repository()
    return ReadHistoricalAlertsUseCase(db_repo=db_repo)


async def get_save_historical_alerts_usecase() -> SaveHistoricalAlertsUseCase:
    db_repo = await get_alert_repository()
    return SaveHistoricalAlertsUseCase(db_repo=db_repo)
