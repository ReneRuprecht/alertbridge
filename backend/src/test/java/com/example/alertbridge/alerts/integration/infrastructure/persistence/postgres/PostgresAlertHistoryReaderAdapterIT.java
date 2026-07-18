package com.example.alertbridge.alerts.integration.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.domain.model.AlertHistorySnapshot;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.AlertHistoryJpaRepository;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.PostgresAlertHistoryReaderAdapter;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.entity.AlertHistoryEntity;
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
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;


@DataJpaTest
@Testcontainers
@Import(PostgresAlertHistoryReaderAdapter.class)
public class PostgresAlertHistoryReaderAdapterIT {

    static PostgreSQLContainer postgres = new PostgreSQLContainer("postgres:18-alpine");

    @Autowired
    private PostgresAlertHistoryReaderAdapter postgresAlertHistoryReaderAdapter;

    @Autowired
    private AlertHistoryJpaRepository alertHistoryJpaRepository;

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
    void shouldReadAlertHistoryByInstance() {
        AlertHistoryEntity entity = new AlertHistoryEntity(
                UUID.randomUUID(),
                "fp1",
                "FIRING",
                "CPUHigh",
                "CRITICAL",
                "prod",
                "instance-01",
                "exporter",
                Instant.parse("2026-01-01T00:00:00Z"),
                Instant.parse("2026-01-01T00:00:00Z"),
                "key-1"
        );


        alertHistoryJpaRepository.save(entity);

        List<AlertHistorySnapshot> result = postgresAlertHistoryReaderAdapter.findByInstance(
                "instance-01");

        assertThat(result).hasSize(1);

        AlertHistorySnapshot snapshot = result.getFirst();

        assertThat(snapshot.instance()).isEqualTo("instance-01");
        assertThat(snapshot.alertName()).isEqualTo("CPUHigh");
    }

}
