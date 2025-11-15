from src.application.dto.alert_redis_dto import AlertRedisDto
from src.domain.entities.alert import Alert


def alert_to_redis(alert: Alert) -> str:
    dto = AlertRedisDto.from_domain(alert)
    return dto.to_json()
