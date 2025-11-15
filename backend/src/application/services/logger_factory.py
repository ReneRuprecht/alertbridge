import logging

from src.infrastructure.logging.logger_factory import LoggerFactory

_logger_factory: LoggerFactory | None = None


async def get_logger_factory() -> LoggerFactory:
    global _logger_factory
    if _logger_factory is None:
        logging.basicConfig(
            level=logging.INFO,
            format="%(asctime)s | %(name)s | %(levelname)s | %(message)s",
        )
        _logger_factory = LoggerFactory()
    return _logger_factory
