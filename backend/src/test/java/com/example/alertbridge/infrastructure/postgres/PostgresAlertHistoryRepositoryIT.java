package com.example.alertbridge.infrastructure.postgres;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.value.AlertStatus;
import fixtures.TestFixtures;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.data.jpa.test.autoconfigure.DataJpaTest;
import org.springframework.context.annotation.Import;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.DynamicPropertyRegistry;
import org.springframework.test.context.DynamicPropertySource;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;
import org.testcontainers.postgresql.PostgreSQLContainer;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

@DataJpaTest
@Testcontainers
@Import(PostgresAlertHistoryRepository.class)
@ActiveProfiles("postgres-it")
public class PostgresAlertHistoryRepositoryIT {

    @Container
    static PostgreSQLContainer postgres = new PostgreSQLContainer(
            "postgres:18-alpine").withExposedPorts(5432);
    @Autowired
    private PostgresAlertHistoryRepository historyRepository;

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
    void shouldPersistAndLoadAlertHistory() {
        AlertState state = TestFixtures.firingAlert("fp-1");
        AlertState stateResolved = TestFixtures.resolvedAlert("fp-1");

        this.historyRepository.save(state);
        this.historyRepository.save(stateResolved);

        List<AlertEvent> historyFromPostgres =
                this.historyRepository.findByFingerprint(state.fingerprint());

        assertThat(historyFromPostgres.size()).isEqualTo(2);
        assertThat(historyFromPostgres.getFirst().fingerprint()).isEqualTo(state.fingerprint);
        assertThat(historyFromPostgres.getFirst().status()).isEqualTo(AlertStatus.FIRING);

        assertThat(historyFromPostgres.get(1).status()).isEqualTo(AlertStatus.RESOLVED);
    }

}
