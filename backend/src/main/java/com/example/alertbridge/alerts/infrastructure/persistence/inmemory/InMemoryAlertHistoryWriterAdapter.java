package com.example.alertbridge.alerts.infrastructure.persistence.inmemory;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryWriterPort;

import java.util.ArrayList;
import java.util.List;

public class InMemoryAlertHistoryWriterAdapter implements AlertHistoryWriterPort {

    private final List<Alert> store = new ArrayList<>();

    @Override
    public void saveAll(List<Alert> alerts) {
        this.store.addAll(alerts);
    }
}
