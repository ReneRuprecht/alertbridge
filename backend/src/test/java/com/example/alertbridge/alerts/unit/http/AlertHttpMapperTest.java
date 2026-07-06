package com.example.alertbridge.alerts.unit.http;

import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.infrastructure.http.AlertHttpMapper;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusAlert;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusPayloadRequest;
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



}


