from abc import ABC, abstractmethod


class LoggerPort(ABC):

    @abstractmethod
    def info(self, msg: str):
        pass

    @abstractmethod
    def warining(self, msg: str):
        pass

    @abstractmethod
    def error(self, msg: str):
        pass
