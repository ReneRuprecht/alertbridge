package com.example.alertbridge.alerts.config;

import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;
import com.example.alertbridge.alerts.infrastructure.persistence.inmemory.InMemoryAlertBatchWriterAdapter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AlertsApplicationConfig {

    @Bean
    public InMemoryAlertBatchWriterAdapter inMemoryAlertBatchWriterAdapter() {
        return new InMemoryAlertBatchWriterAdapter();
    }

    @Bean
    public ReceiveAlertsUseCase receiveAlertsUseCase(AlertBatchWriterPort alertBatchWriterPort) {
        return new ReceiveAlertsUseCase(alertBatchWriterPort);
    }
}
