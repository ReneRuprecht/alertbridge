from src.application.dtos.alert_redis_dto import AlertRedisDto


def redis_to_alert_dto(data: str) -> AlertRedisDto:
    dto = AlertRedisDto.from_json(data=data)
    return dto
