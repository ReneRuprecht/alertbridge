package com.example.alertbridge.application;

import com.example.alertbridge.application.alertstate.SaveAlertsToHistoryUseCase;
import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.value.AlertStatus;
import com.example.alertbridge.infrastructure.postgres.JpaAlertHistoryRepository;
import com.example.alertbridge.infrastructure.postgres.PostgresAlertHistoryRepository;
import fixtures.TestFixtures;
import org.junit.jupiter.api.BeforeEach;
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
class SaveAlertsToHistoryUseCaseIT {
    @Container
    static PostgreSQLContainer postgres = new PostgreSQLContainer("postgres:18-alpine");
    SaveAlertsToHistoryUseCase useCase;
    PostgresAlertHistoryRepository repository;
    @Autowired
    JpaAlertHistoryRepository jpaAlertHistoryRepository;

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

    @BeforeEach
    void setup() {
        repository = new PostgresAlertHistoryRepository(jpaAlertHistoryRepository);
        useCase = new SaveAlertsToHistoryUseCase(repository);
    }


    @Test
    void shouldSaveAlertOnlyOnce() {
        AlertEvent event = TestFixtures.firingEvent("fp-1");

        this.useCase.execute(List.of(event));

        List<AlertEvent> historyFromPostgres = this.repository.findByAlertInstance(event
                .labels()
                .instance());

        assertThat(historyFromPostgres.size()).isEqualTo(1);
        assertThat(historyFromPostgres.getFirst().status()).isEqualTo(AlertStatus.FIRING);
    }

}


