package com.example.alertbridge.alerts.config;

import com.example.alertbridge.alerts.application.GetAlertHistoryByInstanceUseCase;
import com.example.alertbridge.alerts.application.GetCurrentAlertsUseCase;
import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateReaderPort;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateWriterPort;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryReaderPort;
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
    public GetCurrentAlertsUseCase getCurrentAlertsUseCase(AlertCurrentStateReaderPort alertCurrentStateReaderPort) {
        return new GetCurrentAlertsUseCase(alertCurrentStateReaderPort);
    }

    @Bean
    public GetAlertHistoryByInstanceUseCase getAlertHistoryByInstanceUseCase(AlertHistoryReaderPort alertCurrentStateReaderPort) {
        return new GetAlertHistoryByInstanceUseCase(alertCurrentStateReaderPort);
    }
}
