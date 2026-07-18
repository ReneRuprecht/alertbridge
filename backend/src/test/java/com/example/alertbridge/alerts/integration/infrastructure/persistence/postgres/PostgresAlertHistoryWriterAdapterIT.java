package com.example.alertbridge.alerts.integration.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertLabels;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.PostgresAlertHistoryWriterAdapter;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.data.jpa.test.autoconfigure.DataJpaTest;
import org.springframework.context.annotation.Import;
import org.springframework.test.context.DynamicPropertyRegistry;
import org.springframework.test.context.DynamicPropertySource;
import org.testcontainers.junit.jupiter.Testcontainers;
import org.testcontainers.postgresql.PostgreSQLContainer;

import java.time.Instant;
import java.util.List;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@DataJpaTest
@Testcontainers
@Import(PostgresAlertHistoryWriterAdapter.class)
public class PostgresAlertHistoryWriterAdapterIT {

    static PostgreSQLContainer postgres = new PostgreSQLContainer("postgres:18-alpine");

    @Autowired
    private PostgresAlertHistoryWriterAdapter postgresAlertHistoryWriterAdapter;

    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", postgres::getJdbcUrl);
        registry.add("spring.datasource.username", postgres::getUsername);
        registry.add("spring.datasource.password", postgres::getPassword);
    }

    @BeforeAll
    static void beforeAll() {
        postgres.start();
    }

    @AfterAll
    static void afterAll() {
        postgres.stop();
    }

    @Test
    void postgresContainerShouldBeRunning() {
        assertThat(postgres.isRunning()).isTrue();
    }

    @Test
    void shouldPersistAndLoadAlertHistory() {

        Alert alert = Alert.create(
                new AlertFingerprint("fp1"),
                AlertStatus.FIRING,
                new AlertLabels("CPUHigh", AlertSeverity.CRITICAL, "prod", "instance-1", "job-a"),
                Instant.parse("2026-01-01T00:00:00Z")
        );

        postgresAlertHistoryWriterAdapter.saveAll(List.of(alert));
    }


}
