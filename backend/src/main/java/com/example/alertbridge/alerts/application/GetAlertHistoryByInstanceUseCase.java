package com.example.alertbridge.alerts.application;

import com.example.alertbridge.alerts.application.query.GetAlertHistoryByInstanceQuery;
import com.example.alertbridge.alerts.domain.model.AlertHistorySnapshot;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryReaderPort;

import java.util.List;

public class GetAlertHistoryByInstanceUseCase {

    private final AlertHistoryReaderPort alertHistoryReaderPort;

    public GetAlertHistoryByInstanceUseCase(AlertHistoryReaderPort alertHistoryReaderPort) {
        this.alertHistoryReaderPort = alertHistoryReaderPort;
    }

    public List<AlertHistorySnapshot> getHistoryByInstance(GetAlertHistoryByInstanceQuery query) {
        return alertHistoryReaderPort.findByInstance(query.instance());
    }
}
