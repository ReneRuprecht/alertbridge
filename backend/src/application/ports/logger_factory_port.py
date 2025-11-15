from abc import ABC, abstractmethod

from src.application.ports.logger_port import LoggerPort


class LoggerFactoryPort(ABC):

    @abstractmethod
    def get(self, name: str) -> LoggerPort:
        pass
