package com.example.alertbridge.alerts.domain.ports;

import com.example.alertbridge.alerts.domain.model.AlertHistorySnapshot;

import java.util.List;

public interface AlertHistoryReaderPort {

    List<AlertHistorySnapshot> findByInstance(String instance);
}
