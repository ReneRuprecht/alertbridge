package com.example.alertbridge.alerts.unit.application;

import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateBatchWriterPort;
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

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoInteractions;

@ExtendWith(MockitoExtension.class)
public class ReceiveAlertsUseCaseTest {

    @Mock
    AlertBatchWriterPort alertBatchWriterPort;

    @Mock
    AlertCurrentStateBatchWriterPort alertCurrentStateBatchWriterPort;

    @InjectMocks
    ReceiveAlertsUseCase underTest;

    @Captor
    ArgumentCaptor<List<Alert>> historyCaptor;

    @Captor
    ArgumentCaptor<List<Alert>> currentStateCaptor;


    @Test
    void shouldMapAndCallPorts_whenAlertsArePresent() {

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


        verify(alertBatchWriterPort).saveAll(historyCaptor.capture());

        verify(alertCurrentStateBatchWriterPort).saveAll(currentStateCaptor.capture());


        List<Alert> historyAlerts = historyCaptor.getValue();
        List<Alert> currentStateAlerts = currentStateCaptor.getValue();


        assertThat(historyAlerts).hasSize(2);

        assertThat(currentStateAlerts).hasSize(2);


        Alert alert = historyAlerts.getFirst();

        assertThat(alert.fingerprint().value()).isEqualTo("fp1");

        assertThat(alert.status()).isEqualTo(AlertStatus.FIRING);

        assertThat(alert.labels().alertName()).isEqualTo("CPUHigh");
    }


    @Test
    void shouldNotCallPorts_whenAlertsAreEmpty() {

        underTest.receive(new ReceiveAlertsCommand(List.of()));


        verifyNoInteractions(alertBatchWriterPort, alertCurrentStateBatchWriterPort);
    }


    @Test
    void shouldNotCallPorts_whenAlertsAreNull() {

        underTest.receive(new ReceiveAlertsCommand(null));

        verifyNoInteractions(alertBatchWriterPort, alertCurrentStateBatchWriterPort);
    }
}
