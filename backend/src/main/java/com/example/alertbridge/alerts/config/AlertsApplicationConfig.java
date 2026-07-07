package com.example.alertbridge.alerts.config;

import com.example.alertbridge.alerts.application.GetCurrentAlertsUseCase;
import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateReaderPort;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateWriterPort;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryWriterPort;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AlertsApplicationConfig {

    @Bean
    public ReceiveAlertsUseCase receiveAlertsUseCase(AlertHistoryWriterPort alertHistoryWriterPort,
                                                     AlertCurrentStateWriterPort alertCurrentStateWriterPort) {
        return new ReceiveAlertsUseCase(alertHistoryWriterPort, alertCurrentStateWriterPort);
    }

    @Bean
    public GetCurrentAlertsUseCase currentAlertsUseCase(AlertCurrentStateReaderPort alertCurrentStateReaderPort) {
        return new GetCurrentAlertsUseCase(alertCurrentStateReaderPort);
    }
}
