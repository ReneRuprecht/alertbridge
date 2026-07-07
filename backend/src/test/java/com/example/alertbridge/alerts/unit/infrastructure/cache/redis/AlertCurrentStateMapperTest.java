package com.example.alertbridge.alerts.unit.infrastructure.cache.redis;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.model.CurrentAlert;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertLabels;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import com.example.alertbridge.alerts.infrastructure.cache.redis.AlertCurrentStateMapper;
import com.example.alertbridge.alerts.infrastructure.cache.redis.AlertCurrentStateRedis;
import org.junit.jupiter.api.Test;

import java.time.Instant;

import static org.assertj.core.api.Assertions.assertThat;

public class AlertCurrentStateMapperTest {
    @Test
    void shouldMapAlertToRedisState() {

        Alert alert = Alert.create(
                new AlertFingerprint("fp1"),
                AlertStatus.FIRING,
                new AlertLabels("CPUHigh", AlertSeverity.CRITICAL, "prod", "instance-1", "job-a"),
                Instant.parse("2026-01-01T00:00:00Z")
        );

        AlertCurrentStateRedis result = AlertCurrentStateMapper.toState(alert);

        assertThat(result.fingerprint()).isEqualTo("fp1");

        assertThat(result.status()).isEqualTo("FIRING");

        assertThat(result.alertName()).isEqualTo("CPUHigh");

        assertThat(result.environment()).isEqualTo("prod");

        assertThat(result.job()).isEqualTo("job-a");
    }


    @Test
    void shouldMapRedisStateToCurrentAlert() {

        AlertCurrentStateRedis state = new AlertCurrentStateRedis(
                "fp1",
                "FIRING",
                "CPUHigh",
                "CRITICAL",
                "prod",
                "instance-1",
                "job-a",
                Instant.parse("2026-01-01T00:00:00Z"),
                Instant.parse("2026-01-01T00:01:00Z")
        );

        CurrentAlert result = AlertCurrentStateMapper.toDomain(state);

        assertThat(result.fingerprint().value()).isEqualTo("fp1");

        assertThat(result.status()).isEqualTo(AlertStatus.FIRING);

        assertThat(result.alertName()).isEqualTo("CPUHigh");

        assertThat(result.severity()).isEqualTo(AlertSeverity.CRITICAL);

        assertThat(result.updatedAt()).isEqualTo(Instant.parse("2026-01-01T00:01:00Z"));
    }
}
