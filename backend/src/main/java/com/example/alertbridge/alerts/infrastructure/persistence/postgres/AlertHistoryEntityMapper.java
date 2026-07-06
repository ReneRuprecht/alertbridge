package com.example.alertbridge.alerts.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.entity.AlertHistoryEntity;

import java.util.UUID;

public class AlertHistoryEntityMapper {
    public static AlertHistoryEntity toEntity(Alert alert) {
        return new AlertHistoryEntity(
                UUID.randomUUID(),
                alert.fingerprint().value(),
                alert.status().name(),
                alert.labels().alertName(),
                alert.labels().severity().value(),
                alert.labels().environment(),
                alert.labels().instance(),
                alert.labels().job(),
                alert.startsAt(),
                alert.receivedAt(),
                createEventKey(alert)
        );

    }

    private static String createEventKey(Alert alert) {
        return String.format(
                "%s|%s|%s",
                alert.fingerprint().value(),
                alert.status().name(),
                alert.startsAt()
        );
    }

}
