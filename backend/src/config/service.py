from src.repository.psql_client import PSQLClient
from src.repository.redis_client import RedisClient
from src.services.alert_service import AlertService
from src.services.alerts.http_alert_service import HttpAlertService
from src.services.alerts.psql_alert_service import PSQLAlertService
from src.services.alerts.redis_http_service import RedisAlertService

psql_client: PSQLClient = PSQLClient()
redis_client: RedisClient = RedisClient()

psql_alert_service: PSQLAlertService = PSQLAlertService(
    psql_client=psql_client
)
redis_alert_service: RedisAlertService = RedisAlertService(
    redis_client=redis_client
)
http_alert_service: HttpAlertService = HttpAlertService()


alert_service: AlertService = AlertService(
    psql_service=psql_alert_service,
    redis_service=redis_alert_service,
    http_alert_service=http_alert_service,
)


def get_psql_client():
    return psql_client


def get_redis_client():
    return redis_client


def get_alert_service():
    return alert_service
