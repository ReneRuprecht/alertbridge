package com.example.alertbridge.alerts.config;

import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;
import com.example.alertbridge.alerts.infrastructure.persistence.inmemory.InMemoryAlertBatchWriterAdapter;
import org.springframework.boot.test.context.TestConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Profile;

@TestConfiguration
public class TestPersistenceConfig {

    @Bean
    @Profile("test")
    AlertBatchWriterPort alertBatchWriterPort() {
        return new InMemoryAlertBatchWriterAdapter();
    }
}
