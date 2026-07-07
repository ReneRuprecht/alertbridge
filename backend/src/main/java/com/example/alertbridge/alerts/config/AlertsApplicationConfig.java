package com.example.alertbridge.alerts.config;

import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateBatchWriterPort;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AlertsApplicationConfig {

    @Bean
    public ReceiveAlertsUseCase receiveAlertsUseCase(AlertBatchWriterPort alertBatchWriterPort,
                                                     AlertCurrentStateBatchWriterPort alertCurrentStateBatchWriterPort) {
        return new ReceiveAlertsUseCase(alertBatchWriterPort, alertCurrentStateBatchWriterPort);
    }
}
