package com.example.alertbridge.alerts.unit.application;

import com.example.alertbridge.alerts.application.AlertApplicationMapper;
import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import org.junit.jupiter.api.Test;

import java.time.Instant;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

public class AlertApplicationMapperTest {
    @Test
    void shouldMapCommandToDomainSuccessfully_whenCommandIsValid() {

        ReceiveAlertCommand command = new ReceiveAlertCommand(
                "fp-123",
                "firing",
                "CPUHigh",
                "prod",
                "i-1",
                "job-a",
                "critical",
                Instant.parse("2026-01-01T12:00:00Z")
        );

        Alert alert = AlertApplicationMapper.toDomain(command);

        assertThat(alert.fingerprint().value()).isEqualTo("fp-123");
        assertThat(alert.status()).isEqualTo(AlertStatus.FIRING);

        assertThat(alert.labels().alertName()).isEqualTo("CPUHigh");
        assertThat(alert.labels().severity()).isEqualTo(AlertSeverity.CRITICAL);
        assertThat(alert.labels().environment()).isEqualTo("prod");
        assertThat(alert.labels().instance()).isEqualTo("i-1");
        assertThat(alert.labels().job()).isEqualTo("job-a");

        assertThat(alert.startsAt()).isEqualTo(command.startsAt());
    }

}
