package com.example.alertbridge.alerts.unit.domain.value;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertFingerprintException;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;


class AlertFingerprintTest {

    @Test
    void shouldThrowException_whenValueIsNull() {
        assertThatThrownBy(() -> new AlertFingerprint(null))
                .isInstanceOf(InvalidAlertFingerprintException.class)
                .hasMessageContaining("Fingerprint must not be null or blank");
    }

    @Test
    void shouldThrowException_whenValueIsBlank() {
        assertThatThrownBy(() -> new AlertFingerprint("   "))
                .isInstanceOf(InvalidAlertFingerprintException.class)
                .hasMessageContaining("Fingerprint must not be null or blank");
    }

    @Test
    void shouldCreate_whenValueIsValid() {
        AlertFingerprint fp = new AlertFingerprint("24ad9e973e22bdce");

        assertThat(fp.value()).isEqualTo("24ad9e973e22bdce");
    }
}
