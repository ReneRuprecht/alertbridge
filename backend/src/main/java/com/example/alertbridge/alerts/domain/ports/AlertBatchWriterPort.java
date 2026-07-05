package com.example.alertbridge.alerts.domain.ports;

import com.example.alertbridge.alerts.domain.model.Alert;

import java.util.List;

public interface AlertBatchWriterPort {
    void saveAll(List<Alert> alerts);
}
