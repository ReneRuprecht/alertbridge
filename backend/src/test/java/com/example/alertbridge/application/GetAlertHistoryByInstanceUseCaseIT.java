package com.example.alertbridge.application;

import com.example.alertbridge.application.alertstate.GetAlertHistoryByInstanceUseCase;
import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
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
class GetAlertHistoryByInstanceUseCaseIT {

    @Container
    static PostgreSQLContainer postgres = new PostgreSQLContainer("postgres:18-alpine");
    GetAlertHistoryByInstanceUseCase useCase;
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
        useCase = new GetAlertHistoryByInstanceUseCase(repository);
    }


    @Test
    void shouldGetAlertsForInstance() {
        AlertEvent eventFP1 = TestFixtures.firingEvent("fp-1", "instance-1");
        AlertState stateFP1 = AlertState.fromEvent(eventFP1);

        AlertEvent eventFP2 = TestFixtures.firingEvent("fp-2", "instance-2");
        AlertState stateFP2 = AlertState.fromEvent(eventFP2);

        AlertEvent eventFP3 = TestFixtures.firingEvent("fp-3", "instance-1");
        AlertState stateFP3 = AlertState.fromEvent(eventFP3);

        repository.save(stateFP1);
        repository.save(stateFP2);
        repository.save(stateFP3);

        List<AlertEvent> events = this.useCase.execute(eventFP1.labels().instance());

        assertThat(events.size()).isEqualTo(2);
        assertThat(events.getFirst().fingerprint().value()).isEqualTo("fp-1");
        assertThat(events.get(1).fingerprint().value()).isEqualTo("fp-3");
    }

}