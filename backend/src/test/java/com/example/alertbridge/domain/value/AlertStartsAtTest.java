package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import java.time.Instant;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class AlertStartsAtTest {

    @Test
    void shouldThrowIfNull() {
        assertThrows(IllegalArgumentException.class, () -> new AlertStartsAt(null));
    }

    @Test
    void shouldKeepValueIfValid() {
        Instant now = Instant.now();
        AlertStartsAt startsAt = new AlertStartsAt(now);
        assertThat(startsAt.value()).isEqualTo(now);
    }

}
