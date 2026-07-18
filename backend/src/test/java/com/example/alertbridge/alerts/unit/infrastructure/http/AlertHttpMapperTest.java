package com.example.alertbridge.alerts.unit.infrastructure.http;

import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.application.query.GetAlertHistoryByInstanceQuery;
import com.example.alertbridge.alerts.domain.model.AlertHistorySnapshot;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import com.example.alertbridge.alerts.infrastructure.http.AlertHttpMapper;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusAlert;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusPayloadRequest;
import com.example.alertbridge.alerts.infrastructure.http.response.AlertHistoryItemResponse;
import org.junit.jupiter.api.Test;

import java.time.Instant;
import java.util.List;
import java.util.Map;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class AlertHttpMapperTest {

    private static PrometheusPayloadRequest createPayload() {
        return new PrometheusPayloadRequest(List.of(createAlert("fp1")));
    }

    private static PrometheusAlert createAlert(String fingerprint) {
        return new PrometheusAlert(
                fingerprint, "firing", Map.of(
                "alertname",
                "CPUHigh",
                "environment",
                "prod",
                "instance",
                "i-1",
                "job",
                "job-a",
                "severity",
                "critical"
        ), Instant.parse("2026-01-01T00:00:00Z")
        );
    }

    @Test
    void shouldMapSingleAlertCorrectly() {

        PrometheusPayloadRequest request = createPayload();

        ReceiveAlertsCommand command = AlertHttpMapper.toCommand(request);

        assertThat(command.alerts().size()).isEqualTo(1);

        ReceiveAlertCommand mapped = command.alerts().getFirst();

        assertThat(mapped.fingerprint()).isEqualTo("fp1");
        assertThat(mapped.status()).isEqualTo("firing");
        assertThat(mapped.alertName()).isEqualTo("CPUHigh");
        assertThat(mapped.environment()).isEqualTo("prod");
        assertThat(mapped.instance()).isEqualTo("i-1");
        assertThat(mapped.job()).isEqualTo("job-a");
        assertThat(mapped.severity()).isEqualTo("critical");
        assertThat(mapped.startsAt()).isEqualTo(Instant.parse("2026-01-01T00:00:00Z"));
    }

    @Test
    void shouldMapMultipleAlerts() {

        PrometheusAlert a1 = createAlert("fp1");
        PrometheusAlert a2 = createAlert("fp2");

        PrometheusPayloadRequest request = new PrometheusPayloadRequest(List.of(a1, a2));

        ReceiveAlertsCommand command = AlertHttpMapper.toCommand(request);

        assertThat(command.alerts().size()).isEqualTo(2);
        assertThat(command.alerts())
                .extracting(ReceiveAlertCommand::fingerprint)
                .containsExactly("fp1", "fp2");
    }

    @Test
    void shouldReturnEmptyCommand_whenAlertsAreEmpty() {

        PrometheusPayloadRequest request = new PrometheusPayloadRequest(List.of());

        ReceiveAlertsCommand command = AlertHttpMapper.toCommand(request);

        assertTrue(command.alerts().isEmpty());
    }

    @Test
    void shouldReturnEmptyCommand_whenAlertsAreNull() {

        PrometheusPayloadRequest request = new PrometheusPayloadRequest(null);

        ReceiveAlertsCommand command = AlertHttpMapper.toCommand(request);

        assertTrue(command.alerts().isEmpty());
    }

    @Test
    void shouldCreateQueryFromInstance() {
        String instance = "instance-01";

        GetAlertHistoryByInstanceQuery query =
                AlertHttpMapper.toGetAlertHistoryByInstanceQuery(instance);

        assertThat(query.instance())
                .isEqualTo("instance-01");
    }


    @Test
    void shouldMapSnapshotToResponse() {
        Instant startsAt = Instant.parse("2026-01-01T10:00:00Z");
        Instant receivedAt = Instant.parse("2026-01-01T10:01:00Z");

        AlertHistorySnapshot snapshot = new AlertHistorySnapshot(
                new AlertFingerprint("fp1"),
                AlertStatus.FIRING,
                "CPUHigh",
                AlertSeverity.CRITICAL,
                "production",
                "instance-01",
                "node-exporter",
                startsAt,
                receivedAt
        );

        AlertHistoryItemResponse response =
                AlertHttpMapper.toAlertHistoryItemResponse(snapshot);

        assertThat(response.fingerprint())
                .isEqualTo("fp1");

        assertThat(response.status())
                .isEqualTo("FIRING");

        assertThat(response.alertName())
                .isEqualTo("CPUHigh");

        assertThat(response.severity())
                .isEqualTo("CRITICAL");

        assertThat(response.environment())
                .isEqualTo("production");

        assertThat(response.instance())
                .isEqualTo("instance-01");

        assertThat(response.job())
                .isEqualTo("node-exporter");

        assertThat(response.startsAt())
                .isEqualTo(startsAt);

        assertThat(response.receivedAt())
                .isEqualTo(receivedAt);
    }

}


