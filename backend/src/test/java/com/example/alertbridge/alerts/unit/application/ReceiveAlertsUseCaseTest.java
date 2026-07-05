package com.example.alertbridge.alerts.unit.application;

import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.ArgumentCaptor;
import org.mockito.Captor;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.time.Instant;
import java.util.List;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoInteractions;

@ExtendWith(MockitoExtension.class)
public class ReceiveAlertsUseCaseTest {

    @Mock
    AlertBatchWriterPort alertBatchWriterPort;

    @InjectMocks
    ReceiveAlertsUseCase underTest;

    @Captor
    private ArgumentCaptor<List<Alert>> captor;

    @Test
    void shouldMapAndCallPort_whenAlertsArePresent() {

        ReceiveAlertsCommand command = new ReceiveAlertsCommand(List.of(
                new ReceiveAlertCommand(
                        "fp1",
                        "firing",
                        "CPUHigh",
                        "prod",
                        "i-1",
                        "job-a",
                        "critical",
                        Instant.parse("2026-01-01T00:00:00Z")
                ), new ReceiveAlertCommand(
                        "fp2",
                        "resolved",
                        "CPUHigh",
                        "prod",
                        "i-2",
                        "job-b",
                        "critical",
                        Instant.parse("2026-01-01T00:00:00Z")
                )
        ));

        underTest.receive(command);

        verify(alertBatchWriterPort).saveAll(captor.capture());

        List<Alert> saved = captor.getValue();

        assertThat(saved.size()).isEqualTo(2);

        Alert alert = saved.getFirst();

        assertThat(alert.fingerprint().value()).isEqualTo("fp1");
        assertThat(alert.status()).isEqualTo(AlertStatus.FIRING);
        assertThat(alert.labels().alertName()).isEqualTo("CPUHigh");
    }

    @Test
    void shouldNotCallPort_whenAlertsAreEmpty() {

        underTest.receive(new ReceiveAlertsCommand(List.of()));

        verifyNoInteractions(alertBatchWriterPort);
    }

    @Test
    void shouldNotCallPort_whenAlertsAreNull() {

        underTest.receive(new ReceiveAlertsCommand(null));

        verifyNoInteractions(alertBatchWriterPort);
    }

}
