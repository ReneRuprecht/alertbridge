package com.example.alertbridge.aplication;

import com.example.alertbridge.application.alertstate.GetCurrentAlertStatesUseCase;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertStateRepository;
import com.example.alertbridge.infrastructure.repository.InMemoryAlertStateRepository;
import fixtures.TestFixtures;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

public class GetCurrentAlertStatesUseCaseTest {

    private AlertStateRepository repository;
    private GetCurrentAlertStatesUseCase useCase;

    @BeforeEach
    void setUp() {
        repository = new InMemoryAlertStateRepository();
        useCase = new GetCurrentAlertStatesUseCase(repository);
    }

    @Test
    void shouldReturnOnlyActiveAlerts() {
        // Arrange: create firing and resolved alerts
        AlertState firing1 = TestFixtures.firingAlert("fp-1");
        AlertState firing2 = TestFixtures.firingAlert("fp-2");
        AlertState resolved = TestFixtures.resolvedAlert("fp-3");

        repository.save(firing1);
        repository.save(firing2);
        repository.save(resolved);

        List<AlertState> activeAlerts = useCase.execute();

        assertThat(activeAlerts)
                .hasSize(2)
                .extracting(AlertState::fingerprint)
                .containsExactlyInAnyOrder(
                        firing1.fingerprint(),
                        firing2.fingerprint()
                );

        assertThat(activeAlerts).allMatch(AlertState::isActive);
    }


}
