package com.example.alertbridge.alerts.application;

import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryWriterPort;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateWriterPort;

import java.util.List;

public class ReceiveAlertsUseCase {

    private final AlertHistoryWriterPort alertHistoryWriterPort;
    private final AlertCurrentStateWriterPort alertCurrentStateWriterPort;

    public ReceiveAlertsUseCase(AlertHistoryWriterPort alertHistoryWriterPort,
                                AlertCurrentStateWriterPort alertCurrentStateWriterPort) {
        this.alertHistoryWriterPort = alertHistoryWriterPort;
        this.alertCurrentStateWriterPort = alertCurrentStateWriterPort;
    }

    public void receive(ReceiveAlertsCommand command) {

        if (command.alerts().isEmpty()) return;

        List<Alert> alerts = command
                .alerts()
                .stream()
                .map(AlertApplicationMapper::toDomain)
                .toList();

        alertHistoryWriterPort.saveAll(alerts);

        alertCurrentStateWriterPort.saveAll(alerts);
    }
}
