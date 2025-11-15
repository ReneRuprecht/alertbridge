import logging

from src.application.ports.logger_factory_port import LoggerFactoryPort
from src.application.ports.logger_port import LoggerPort
from src.infrastructure.logging.python_logger import PythonLogger


class LoggerFactory(LoggerFactoryPort):

    def get(self, name: str) -> LoggerPort:
        logger = logging.getLogger(name)
        return PythonLogger(logger)
