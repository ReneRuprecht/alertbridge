from src.application.dtos.alert_redis_dto import AlertRedisDto
from src.domain.entities.alert import Alert


def test_from_domain_to_dto():
    alert = Alert(
        alertname="InstanceDown",
        status="firing",
        labels={"instance": "server1"},
        fingerprint="abc123",
    )

    dto = AlertRedisDto.from_domain(alert)
    assert dto.alertname == "InstanceDown"
    assert dto.instance == "server1"
    assert dto.status == "firing"
    assert dto.fingerprint == "abc123"

    json_str = dto.to_json()
    assert isinstance(json_str, str)

    dto_from_json = AlertRedisDto.from_json(json_str)
    assert dto_from_json == dto
