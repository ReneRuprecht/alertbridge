package com.example.alertbridge.infrastructure.postgres.mapper;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.value.*;
import com.example.alertbridge.infrastructure.postgres.AlertHistoryEntity;

public class AlertHistoryMapper {

    public static AlertHistoryEntity toEntity(AlertState state) {
        AlertHistoryEntity entity = new AlertHistoryEntity();
        entity.setFingerprint(state.fingerprint().value());
        entity.setStatus(state.status());
        entity.setStartsAt(state.startsAt().value());
        entity.setAlertName(state.labels().alertName().value());
        entity.setInstance(state.labels().instance().value());
        entity.setSeverity(state.labels().severity().name());
        entity.setJob(state.labels().job().value());
        entity.setEnvironment(state.labels().environment().value());
        return entity;
    }

    public static AlertEvent toDomain(AlertHistoryEntity entity) {
        return new AlertEvent(
                new AlertFingerprint(entity.getFingerprint()),
                new AlertLabels(
                        new AlertName(entity.getAlertName()),
                        new AlertEnvironment(entity.getEnvironment()),
                        new AlertInstance(entity.getInstance()),
                        new AlertJob(entity.getJob()),
                        AlertSeverity.fromString(entity.getSeverity())
                ),
                entity.getStatus(),
                new AlertStartsAt(entity.getStartsAt())
        );
    }
}
