package com.example.alertbridge.alerts.application;

import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertLabels;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;

public class AlertApplicationMapper {

    public static Alert toDomain(ReceiveAlertCommand command) {
        return Alert.create(
                new AlertFingerprint(command.fingerprint()),
                AlertStatus.of(command.status()),
                new AlertLabels(
                        command.alertName(),
                        AlertSeverity.of(command.severity()),
                        command.environment(),
                        command.instance(),
                        command.job()
                ),
                command.startsAt()
        );
    }
}
