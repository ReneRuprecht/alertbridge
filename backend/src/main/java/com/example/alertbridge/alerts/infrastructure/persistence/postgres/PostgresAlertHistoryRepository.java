package com.example.alertbridge.alerts.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.entity.AlertHistoryEntity;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class PostgresAlertHistoryRepository implements AlertBatchWriterPort {

    private final AlertHistoryJpaRepository alertHistoryJpaRepository;

    public PostgresAlertHistoryRepository(AlertHistoryJpaRepository alertHistoryJpaRepository) {
        this.alertHistoryJpaRepository = alertHistoryJpaRepository;
    }

    @Override
    public void saveAll(List<Alert> alerts) {

        List<AlertHistoryEntity> entities = alerts
                .stream()
                .map(AlertHistoryEntityMapper::toEntity)
                .toList();

        alertHistoryJpaRepository.saveAll(entities);

    }
}
