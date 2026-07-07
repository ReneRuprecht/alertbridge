package com.example.alertbridge.alerts.infrastructure.http;

import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusAlert;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusPayloadRequest;

import java.util.List;

public class AlertHttpMapper {

    public static ReceiveAlertsCommand toCommand(PrometheusPayloadRequest request) {

        List<ReceiveAlertCommand> alerts = request
                .alerts()
                .stream()
                .map(AlertHttpMapper::toCommand)
                .toList();

        return new ReceiveAlertsCommand(alerts);
    }

    private static ReceiveAlertCommand toCommand(PrometheusAlert alert) {
        return new ReceiveAlertCommand(
                alert.fingerprint(),
                alert.status(),
                alert.labels().get("alertname"),
                alert.labels().get("environment"),
                alert.labels().get("instance"),
                alert.labels().get("job"),
                alert.labels().get("severity"),
                alert.startsAt()
        );
    }
}
