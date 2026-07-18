package com.example.alertbridge.alerts.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.model.AlertHistorySnapshot;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryReaderPort;

import java.util.List;

public class PostgresAlertHistoryReaderAdapter implements AlertHistoryReaderPort {

    private final AlertHistoryJpaRepository alertHistoryJpaRepository;

    public PostgresAlertHistoryReaderAdapter(AlertHistoryJpaRepository alertHistoryJpaRepository) {
        this.alertHistoryJpaRepository = alertHistoryJpaRepository;
    }


    @Override
    public List<AlertHistorySnapshot> findByInstance(String instance) {
        return alertHistoryJpaRepository
                .findByInstance(instance)
                .stream()
                .map(AlertHistoryEntityMapper::toSnapshot)
                .toList();
    }
}
