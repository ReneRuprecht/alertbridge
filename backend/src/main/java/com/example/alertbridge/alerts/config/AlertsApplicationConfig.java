package com.example.alertbridge.alerts.config;

import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryWriterPort;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateWriterPort;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AlertsApplicationConfig {

    @Bean
    public ReceiveAlertsUseCase receiveAlertsUseCase(AlertHistoryWriterPort alertHistoryWriterPort,
                                                     AlertCurrentStateWriterPort alertCurrentStateWriterPort) {
        return new ReceiveAlertsUseCase(alertHistoryWriterPort, alertCurrentStateWriterPort);
    }
}
