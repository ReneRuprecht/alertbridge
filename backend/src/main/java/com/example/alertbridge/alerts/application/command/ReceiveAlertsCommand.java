package com.example.alertbridge.alerts.application.command;

import java.util.List;

public record ReceiveAlertsCommand(List<ReceiveAlertCommand> alerts) {

    public ReceiveAlertsCommand {
        if (alerts == null) {
            alerts = List.of();
        }
    }
}
