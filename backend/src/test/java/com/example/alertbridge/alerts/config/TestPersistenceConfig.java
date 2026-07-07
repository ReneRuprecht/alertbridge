package com.example.alertbridge.alerts.config;

import com.example.alertbridge.alerts.domain.ports.AlertHistoryWriterPort;
import com.example.alertbridge.alerts.infrastructure.persistence.inmemory.InMemoryAlertHistoryWriterAdapter;
import org.springframework.boot.test.context.TestConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Profile;

@TestConfiguration
public class TestPersistenceConfig {

    @Bean
    @Profile("test")
    AlertHistoryWriterPort alertBatchWriterPort() {
        return new InMemoryAlertHistoryWriterAdapter();
    }
}
