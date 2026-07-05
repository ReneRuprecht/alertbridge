package com.example.alertbridge.alerts.infrastructure.persistence.inmemory;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;

import java.util.ArrayList;
import java.util.List;

public class InMemoryAlertBatchWriterAdapter implements AlertBatchWriterPort {

    private final List<Alert> store = new ArrayList<>();

    @Override
    public void saveAll(List<Alert> alerts) {
        this.store.addAll(alerts);
    }
}
