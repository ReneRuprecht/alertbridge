package com.example.alertbridge.aplication;

import com.example.alertbridge.application.alertstate.GetCurrentAlertStateByFingerprintUseCase;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertStateRepository;
import com.example.alertbridge.domain.value.AlertFingerprint;
import com.example.alertbridge.infrastructure.repository.InMemoryAlertStateRepository;
import fixtures.TestFixtures;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;

public class GetCurrentAlertStateByFingerprintUseCaseTest {


    private AlertStateRepository repository;
    private GetCurrentAlertStateByFingerprintUseCase useCase;

    @BeforeEach
    void setUp() {
        repository = new InMemoryAlertStateRepository();
        useCase = new GetCurrentAlertStateByFingerprintUseCase(repository);
    }

    @Test
    void shouldReturnActiveAlertIfPresent() {
        AlertState activeAlert = TestFixtures.firingAlert("fp-123");
        repository.save(activeAlert);

        Optional<AlertState> result = useCase.execute(new AlertFingerprint("fp-123"));

        assertThat(result).isPresent();
        assertThat(result.get().isActive()).isTrue();
        assertThat(result.get().fingerprint().value()).isEqualTo("fp-123");
    }

    @Test
    void shouldReturnEmptyIfAlertIsResolved() {
        AlertState resolvedAlert = TestFixtures.resolvedAlert("fp-456");
        repository.save(resolvedAlert);

        Optional<AlertState> result = useCase.execute(new AlertFingerprint("fp-456"));

        assertThat(result).isEmpty();
    }

    @Test
    void shouldReturnEmptyIfAlertDoesNotExist() {
        Optional<AlertState> result = useCase.execute(new AlertFingerprint("fp-missing"));

        assertThat(result).isEmpty();
    }


}
