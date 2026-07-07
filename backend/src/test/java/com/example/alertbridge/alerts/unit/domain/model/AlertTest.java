package com.example.alertbridge.alerts.unit.domain.model;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertStartsAtException;
import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertLabels;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import org.junit.jupiter.api.Test;

import java.time.Instant;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.assertj.core.api.AssertionsForClassTypes.assertThatThrownBy;

public class AlertTest {

    @Test
    void shouldCreateAlert_whenAllValuesAreValid() {
        AlertFingerprint fingerprint = new AlertFingerprint("fp123");
        AlertLabels labels = new AlertLabels(
                "CPUHigh",
                AlertSeverity.CRITICAL,
                "prod",
                "server-1",
                "node-exporter"
        );
        Instant startsAt = Instant.parse("2026-01-01T12:00:00Z");

        Alert alert = Alert.create(fingerprint, AlertStatus.FIRING, labels, startsAt);

        assertThat(alert.fingerprint()).isEqualTo(fingerprint);
        assertThat(alert.status()).isEqualTo(AlertStatus.FIRING);
        assertThat(alert.labels()).isEqualTo(labels);
        assertThat(alert.startsAt()).isEqualTo(startsAt);
        assertThat(alert.receivedAt()).isNotNull();


    }

    @Test
    void shouldSetReceivedAtToCurrentTime_whenCreatingAlert() {
        Instant before = Instant.now();

        Alert alert = Alert.create(
                new AlertFingerprint("fp123"),
                AlertStatus.FIRING,
                new AlertLabels("CPUHigh", AlertSeverity.CRITICAL, null, null, null),
                Instant.now()
        );

        Instant after = Instant.now();

        assertThat(alert.receivedAt()).isBetween(before, after);
    }

    @Test
    void shouldThrowException_whenStartsAtIsNull() {
        assertThatThrownBy(() -> Alert.create(
                new AlertFingerprint("24ad9e973e22bdce"),
                AlertStatus.FIRING,
                new AlertLabels("CPUHigh", AlertSeverity.CRITICAL, null, null, null),
                null
        ))
                .isInstanceOf(InvalidAlertStartsAtException.class)
                .hasMessage("StartsAt must not be null");
    }


}
