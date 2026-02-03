package com.example.alertbridge.aplication;

import com.example.alertbridge.application.alertstate.SyncAlertsUseCase;
import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertStateRepository;
import com.example.alertbridge.domain.value.AlertFingerprint;
import com.example.alertbridge.domain.value.AlertStatus;
import com.example.alertbridge.infrastructure.repository.InMemoryAlertStateRepository;
import fixtures.TestFixtures;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

public class SyncAlertsUseCaseTest {

    private AlertStateRepository repository;
    private SyncAlertsUseCase useCase;

    @BeforeEach
    void setUp() {
        repository = new InMemoryAlertStateRepository();
        useCase = new SyncAlertsUseCase(repository);
    }

    @Test
    void shouldCreateAndUpdateAlertState() {

        AlertEvent firing = TestFixtures.firingEvent("fp-1");
        AlertEvent resolved = TestFixtures.resolvedEvent("fp-1");

        useCase.execute(List.of(firing));
        useCase.execute(List.of(resolved));

        AlertState state = this.repository.findByFingerprint(new AlertFingerprint("fp-1")).orElseThrow();

        assertThat(state.status()).isEqualTo(AlertStatus.RESOLVED);
    }


}
