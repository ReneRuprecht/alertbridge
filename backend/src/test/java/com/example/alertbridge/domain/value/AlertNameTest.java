package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class AlertNameTest {

    @Test
    void shouldThrowIfNullOrBlank() {
        assertThrows(IllegalArgumentException.class, () -> new AlertName(null));
        assertThrows(IllegalArgumentException.class, () -> new AlertName(" "));
    }

    @Test
    void shouldKeepValueIfValid() {
        AlertName name = new AlertName("CPUHigh");
        assertThat(name.value()).isEqualTo("CPUHigh");
    }

}
