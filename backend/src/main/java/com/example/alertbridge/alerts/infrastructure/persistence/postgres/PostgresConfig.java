package com.example.alertbridge.alerts.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.ports.AlertBatchWriterPort;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class PostgresConfig {

    @Bean
    AlertBatchWriterPort alertBatchWriterPort(AlertHistoryJpaRepository alertHistoryJpaRepository) {
        return new PostgresAlertHistoryRepository(alertHistoryJpaRepository);
    }
}
