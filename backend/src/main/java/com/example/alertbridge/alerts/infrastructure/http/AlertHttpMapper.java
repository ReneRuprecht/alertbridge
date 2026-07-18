package com.example.alertbridge.alerts.infrastructure.http;

import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.application.query.GetAlertHistoryByInstanceQuery;
import com.example.alertbridge.alerts.domain.model.AlertHistorySnapshot;
import com.example.alertbridge.alerts.domain.model.CurrentAlert;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusAlert;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusPayloadRequest;
import com.example.alertbridge.alerts.infrastructure.http.response.AlertHistoryItemResponse;
import com.example.alertbridge.alerts.infrastructure.http.response.CurrentAlertResponse;

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

    public static CurrentAlertResponse toAlertHistoryItemResponse(CurrentAlert alert) {
        return new CurrentAlertResponse(
                alert.fingerprint().value(),
                alert.status().name(),
                alert.alertName(),
                alert.severity().name(),
                alert.environment(),
                alert.instance(),
                alert.job(),
                alert.startsAt(),
                alert.lastUpdatedAt()
        );
    }

    public static GetAlertHistoryByInstanceQuery toGetAlertHistoryByInstanceQuery(String instance) {
        return new GetAlertHistoryByInstanceQuery(instance);
    }

    public static AlertHistoryItemResponse toAlertHistoryItemResponse(AlertHistorySnapshot snapshot) {
        return new AlertHistoryItemResponse(
                snapshot.fingerprint().value(),
                snapshot.status().name(),
                snapshot.alertName(),
                snapshot.severity().name(),
                snapshot.environment(),
                snapshot.instance(),
                snapshot.job(),
                snapshot.startsAt(),
                snapshot.receivedAt()
        );
    }
}
