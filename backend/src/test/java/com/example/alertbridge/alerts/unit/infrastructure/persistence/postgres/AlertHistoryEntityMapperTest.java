package com.example.alertbridge.alerts.unit.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.model.AlertHistorySnapshot;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertLabels;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.AlertHistoryEntityMapper;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.entity.AlertHistoryEntity;
import org.junit.jupiter.api.Test;

import java.time.Instant;
import java.util.UUID;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

public class AlertHistoryEntityMapperTest {

    @Test
    void shouldMapAlertToHistoryEntity() {

        Alert alert = Alert.create(
                new AlertFingerprint("fp1"),
                AlertStatus.FIRING,
                new AlertLabels("CPUHigh", AlertSeverity.CRITICAL, "prod", "instance-1", "job-a"),
                Instant.parse("2026-01-01T00:00:00Z")
        );

        AlertHistoryEntity entity = AlertHistoryEntityMapper.toEntity(alert);

        assertThat(entity.getFingerprint()).isEqualTo("fp1");

        assertThat(entity.getStatus()).isEqualTo("FIRING");

        assertThat(entity.getAlertName()).isEqualTo("CPUHigh");

        assertThat(entity.getEventKey()).isNotBlank();
    }
    @Test
    void shouldMapEntityToSnapshot() {

        AlertHistoryEntity entity = new AlertHistoryEntity(
                UUID.randomUUID(),
                "fp1",
                "FIRING",
                "CPUHigh",
                "CRITICAL",
                "prod",
                "instance-01",
                "node-exporter",
                Instant.parse("2026-01-01T00:00:00Z"),
                Instant.parse("2026-01-01T00:01:00Z"),
                "key-01"
        );

        AlertHistorySnapshot snapshot =
                AlertHistoryEntityMapper.toSnapshot(entity);

        assertThat(snapshot.fingerprint().value()).isEqualTo("fp1");
        assertThat(snapshot.status()).isEqualTo(AlertStatus.FIRING);
        assertThat(snapshot.alertName()).isEqualTo("CPUHigh");
        assertThat(snapshot.severity()).isEqualTo(AlertSeverity.CRITICAL);
        assertThat(snapshot.environment()).isEqualTo("prod");
        assertThat(snapshot.instance()).isEqualTo("instance-01");
        assertThat(snapshot.job()).isEqualTo("node-exporter");
        assertThat(snapshot.startsAt()).isEqualTo("2026-01-01T00:00:00Z");
        assertThat(snapshot.receivedAt()).isEqualTo("2026-01-01T00:01:00Z");
    }
}
