package com.example.alertbridge.alerts.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.ports.AlertHistoryWriterPort;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class PostgresConfig {

    @Bean
    AlertHistoryWriterPort alertBatchWriterPort(AlertHistoryJpaRepository alertHistoryJpaRepository) {
        return new PostgresAlertHistoryAdapter(alertHistoryJpaRepository);
    }
}
