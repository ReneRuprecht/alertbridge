package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertThrows;

public class AlertFingerprintTest {


    @Test
    void shouldThrowExceptionIfNull() {
        assertThrows(IllegalArgumentException.class, () -> new AlertFingerprint(null));
    }

    @Test
    void shouldThrowExceptionIfBlank() {
        assertThrows(IllegalArgumentException.class, () -> new AlertFingerprint("   "));
    }

    @Test
    void shouldCreateSuccessfullyIfValid() {
        AlertFingerprint fp = new AlertFingerprint("fp-123");
        assert fp.value().equals("fp-123");
    }


}
