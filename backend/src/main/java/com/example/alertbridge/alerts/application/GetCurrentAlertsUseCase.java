package com.example.alertbridge.alerts.application;

import com.example.alertbridge.alerts.domain.model.CurrentAlert;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateReaderPort;

import java.util.List;

public class GetCurrentAlertsUseCase {

    private final AlertCurrentStateReaderPort alertCurrentStateReaderPort;

    public GetCurrentAlertsUseCase(AlertCurrentStateReaderPort alertCurrentStateReaderPort) {
        this.alertCurrentStateReaderPort = alertCurrentStateReaderPort;
    }

    public List<CurrentAlert> getCurrentAlerts() {
        return this.alertCurrentStateReaderPort.findCurrentAlerts();
    }

}
