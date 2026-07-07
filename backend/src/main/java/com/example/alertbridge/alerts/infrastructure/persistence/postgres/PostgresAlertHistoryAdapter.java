package com.example.alertbridge.alerts.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryWriterPort;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.entity.AlertHistoryEntity;
import jakarta.transaction.Transactional;

import java.util.List;

public class PostgresAlertHistoryAdapter implements AlertHistoryWriterPort {

    private final AlertHistoryJpaRepository alertHistoryJpaRepository;

    public PostgresAlertHistoryAdapter(AlertHistoryJpaRepository alertHistoryJpaRepository) {
        this.alertHistoryJpaRepository = alertHistoryJpaRepository;
    }

    @Override
    @Transactional
    public void saveAll(List<Alert> alerts) {

        List<AlertHistoryEntity> entities = alerts
                .stream()
                .map(AlertHistoryEntityMapper::toEntity)
                .toList();

        for (AlertHistoryEntity entity : entities) {
            alertHistoryJpaRepository.saveWithoutDuplicateEventKey(
                    entity.getId(),
                    entity.getFingerprint(),
                    entity.getStatus(),
                    entity.getAlertName(),
                    entity.getSeverity(),
                    entity.getEnvironment(),
                    entity.getInstance(),
                    entity.getJob(),
                    entity.getStartsAt(),
                    entity.getReceivedAt(),
                    entity.getEventKey()
            );
        }

    }
}
