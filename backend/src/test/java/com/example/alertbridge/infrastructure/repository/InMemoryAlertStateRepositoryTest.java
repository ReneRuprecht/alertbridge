package com.example.alertbridge.infrastructure.repository;

import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertStateRepository;
import com.example.alertbridge.domain.value.AlertStatus;
import fixtures.TestFixtures;
import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;

public class InMemoryAlertStateRepositoryTest {

    @Test
    void shouldSaveAndFindAlertByFingerprint() {
        AlertStateRepository repo = new InMemoryAlertStateRepository();
        AlertState alert = TestFixtures.firingAlert("fp-1");

        repo.save(alert);

        Optional<AlertState> found = repo.findByFingerprint(alert.fingerprint());

        assertThat(found).isPresent();
        assertThat(found.get().status()).isEqualTo(AlertStatus.FIRING);
    }

    @Test
    void shouldOverwriteExistingAlertWithSameFingerprint() {
        AlertStateRepository repo = new InMemoryAlertStateRepository();

        AlertState firing = TestFixtures.firingAlert("fp-1");
        repo.save(firing);

        AlertState resolved = TestFixtures.resolvedAlert("fp-1");
        repo.save(resolved);

        AlertState found = repo.findByFingerprint(firing.fingerprint()).orElseThrow();

        assertThat(found.status()).isEqualTo(AlertStatus.RESOLVED);
    }

    @Test
    void shouldReturnOnlyActiveAlerts() {
        AlertStateRepository repo = new InMemoryAlertStateRepository();

        repo.save(TestFixtures.firingAlert("fp-1"));
        repo.save(TestFixtures.resolvedAlert("fp-2"));
        repo.save(TestFixtures.firingAlert("fp-3"));

        List<AlertState> active = repo.findAllActive();

        assertThat(active).hasSize(2).allMatch(AlertState::isActive);
    }


}
