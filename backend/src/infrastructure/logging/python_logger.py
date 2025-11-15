import logging

from src.application.ports.logger_port import LoggerPort


class PythonLogger(LoggerPort):
    def __init__(self, logger: logging.Logger):
        self._logger = logger

    def info(self, msg: str):
        self._logger.info(msg)

    def warining(self, msg: str):
        self._logger.warning(msg)

    def error(self, msg: str):
        self._logger.error(msg)
