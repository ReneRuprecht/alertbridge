package com.example.alertbridge.alerts.application;

import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateBatchWriterPort;

import java.util.List;

public class ReceiveAlertsUseCase {

    private final AlertBatchWriterPort alertBatchWriterPort;
    private final AlertCurrentStateBatchWriterPort alertCurrentStateBatchWriterPort;

    public ReceiveAlertsUseCase(AlertBatchWriterPort alertBatchWriterPort,
                                AlertCurrentStateBatchWriterPort alertCurrentStateBatchWriterPort) {
        this.alertBatchWriterPort = alertBatchWriterPort;
        this.alertCurrentStateBatchWriterPort = alertCurrentStateBatchWriterPort;
    }

    public void receive(ReceiveAlertsCommand command) {

        if (command.alerts().isEmpty()) return;

        List<Alert> alerts = command
                .alerts()
                .stream()
                .map(AlertApplicationMapper::toDomain)
                .toList();

        alertBatchWriterPort.saveAll(alerts);

        alertCurrentStateBatchWriterPort.saveAll(alerts);
    }
}
