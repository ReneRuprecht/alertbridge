package com.example.alertbridge.alerts.infrastructure.cache.redis;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.model.CurrentAlert;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;

import java.time.Instant;

public class AlertCurrentStateMapper {
    public static AlertCurrentStateRedis toState(Alert alert) {

        return new AlertCurrentStateRedis(
                alert.fingerprint().value(),
                alert.status().name(),
                alert.labels().alertName(),
                alert.labels().severity().name(),
                alert.labels().environment(),
                alert.labels().instance(),
                alert.labels().job(),
                alert.startsAt(),
                Instant.now()
        );
    }

    public static CurrentAlert toDomain(AlertCurrentStateRedis source) {
        return new CurrentAlert(
                new AlertFingerprint(source.fingerprint()),
                AlertStatus.valueOf(source.status().toUpperCase()),
                source.alertName(),
                AlertSeverity.of(source.severity().toUpperCase()),
                source.environment(),
                source.instance(),
                source.job(),
                source.startsAt(),
                source.lastUpdatedAt()
        );
    }
}
