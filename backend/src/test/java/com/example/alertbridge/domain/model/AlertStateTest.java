package com.example.alertbridge.domain.model;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.value.AlertFingerprint;
import com.example.alertbridge.domain.value.AlertStartsAt;
import com.example.alertbridge.domain.value.AlertStatus;
import fixtures.TestFixtures;
import org.junit.jupiter.api.Test;

import java.time.Instant;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class AlertStateTest {

    @Test
    void shouldChangeStateFromFiringToResolved() {

        // given
        AlertEvent firing = TestFixtures.firingEvent("fp-1");
        AlertState state = AlertState.fromEvent(firing);

        AlertEvent resolved = TestFixtures.resolvedEvent("fp-1");

        // when
        state.apply(resolved);

        // then
        assertThat(state.status()).isEqualTo(AlertStatus.RESOLVED);
        assertThat(state.isActive()).isFalse();
    }

    @Test
    void shouldKeepFingerprintImmutable() {
        AlertEvent event = TestFixtures.firingEvent("fp-1");
        AlertState state = AlertState.fromEvent(event);

        assertThat(state.fingerprint()).isEqualTo(event.fingerprint());
    }

    @Test
    void shouldCreateAlertStateFromEvent() {
        AlertEvent event = new AlertEvent(
                new AlertFingerprint("fp-1"),
                TestFixtures.defaultLabels(),
                AlertStatus.FIRING,
                new AlertStartsAt(Instant.parse("2026-02-02T10:00:00Z"))
        );

        AlertState state = AlertState.fromEvent(event);

        assertThat(state.fingerprint().value()).isEqualTo("fp-1");
        assertThat(state.status()).isEqualTo(AlertStatus.FIRING);
        assertThat(state.labels()).isEqualTo(TestFixtures.defaultLabels());
        assertThat(state.startsAt().value()).isEqualTo(Instant.parse("2026-02-02T10:00:00Z"));
        assertThat(state.isActive()).isTrue();
    }

    @Test
    void shouldThrowIfEventFingerprintDoesNotMatch() {
        AlertState state = AlertState.fromEvent(
                TestFixtures.firingEvent("fp-1")
        );

        AlertEvent wrongEvent = new AlertEvent(
                new AlertFingerprint("fp-2"),
                TestFixtures.defaultLabels(),
                AlertStatus.RESOLVED,
                new AlertStartsAt(Instant.parse("2026-02-02T10:05:00Z"))
        );

        assertThrows(IllegalArgumentException.class,
                () -> state.apply(wrongEvent));
    }


    @Test
    void isActiveShouldReturnTrueOnlyForFiring() {
        AlertState firing = AlertState.fromEvent(TestFixtures.firingEvent("fp-1"));
        AlertState resolved = AlertState.fromEvent(TestFixtures.resolvedEvent("fp-1"));

        assertThat(firing.isActive()).isTrue();
        assertThat(resolved.isActive()).isFalse();
    }


}
