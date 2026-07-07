package com.example.alertbridge.alerts.integration.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.infrastructure.persistence.postgres.AlertHistoryJpaRepository;
import com.example.alertbridge.alerts.infrastructure.persistence.postgres.entity.AlertHistoryEntity;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.data.jpa.test.autoconfigure.DataJpaTest;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.test.context.DynamicPropertyRegistry;
import org.springframework.test.context.DynamicPropertySource;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;
import org.testcontainers.postgresql.PostgreSQLContainer;

import java.time.Instant;
import java.util.UUID;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@DataJpaTest
@Testcontainers
public class AlertHistoryJpaRepositoryIT {
    @Container
    static PostgreSQLContainer postgres = new PostgreSQLContainer("postgres:18-alpine");
    @Autowired
    private AlertHistoryJpaRepository repository;
    @Autowired
    private JdbcTemplate jdbcTemplate;

    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", postgres::getJdbcUrl);
        registry.add("spring.datasource.username", postgres::getUsername);
        registry.add("spring.datasource.password", postgres::getPassword);
    }

    @Test
    void postgresContainerShouldBeRunning() {
        assertThat(postgres.isRunning()).isTrue();
    }

    @Test
    void shouldPersistAlertHistory() {

        AlertHistoryEntity entity = createEntity("event-1");

        repository.save(entity);

        AlertHistoryEntity saved = repository.findById(entity.getId()).orElseThrow();

        assertThat(saved.getFingerprint()).isEqualTo("fp-test-1");

        assertThat(saved.getStatus()).isEqualTo("FIRING");

        assertThat(saved.getEventKey()).isEqualTo("event-1");
    }


    @Test
    void shouldIgnoreDuplicateEventKey() {

        AlertHistoryEntity first = createEntity("same-event");

        AlertHistoryEntity second = createEntity("same-event");


        repository.saveWithoutDuplicateEventKey(
                first.getId(),
                first.getFingerprint(),
                first.getStatus(),
                first.getAlertName(),
                first.getSeverity(),
                first.getEnvironment(),
                first.getInstance(),
                first.getJob(),
                first.getStartsAt(),
                first.getReceivedAt(),
                first.getEventKey()
        );
        repository.saveWithoutDuplicateEventKey(
                second.getId(),
                second.getFingerprint(),
                second.getStatus(),
                second.getAlertName(),
                second.getSeverity(),
                second.getEnvironment(),
                second.getInstance(),
                second.getJob(),
                second.getStartsAt(),
                second.getReceivedAt(),
                second.getEventKey()
        );


        Integer count = jdbcTemplate.queryForObject(
                """
                        SELECT COUNT(*)
                        FROM alert_history
                        WHERE event_key = 'same-event'
                        """, Integer.class
        );


        assertThat(count).isEqualTo(1);
    }


    private AlertHistoryEntity createEntity(String eventKey) {

        return new AlertHistoryEntity(
                UUID.randomUUID(),
                "fp-test-1",
                "FIRING",
                "CPUHigh",
                "CRITICAL",
                "prod",
                "instance-1",
                "node-exporter",
                Instant.parse("2026-01-01T00:00:00Z"),
                Instant.parse("2026-01-01T00:01:00Z"),
                eventKey
        );
    }
}
