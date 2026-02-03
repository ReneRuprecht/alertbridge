package com.example.alertbridge.api.mapper;

import com.example.alertbridge.api.alertmanager.dto.AlertDto;
import com.example.alertbridge.api.alertmanager.dto.AlertLabelsDto;
import com.example.alertbridge.api.alertmanager.dto.AlertSeverityDto;
import com.example.alertbridge.api.alertmanager.dto.AlertStatusDto;
import com.example.alertbridge.api.alertmanager.mapper.AlertApiMapper;
import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.value.AlertStatus;
import fixtures.TestFixtures;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.time.Instant;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class AlertApiMapperTest {

    private AlertApiMapper mapper;

    @BeforeEach
    void setUp() {
        mapper = new AlertApiMapper();
    }

    @Test
    void shouldMapDtoToEventSuccessfully() {
        AlertDto dto = TestFixtures.alertDto("fp-123", AlertStatusDto.FIRING, "prod");

        AlertEvent event = mapper.toEvent(dto);

        assertThat(event.fingerprint().value()).isEqualTo("fp-123");
        assertThat(event.status()).isEqualTo(AlertStatus.FIRING);
        assertThat(event.labels().environment().value()).isEqualTo("prod");
        assertThat(event.startsAt().value()).isEqualTo(Instant.parse(dto.startsAt()));
    }

    @Test
    void shouldMapResolvedStatus() {
        AlertDto dto = TestFixtures.alertDto("fp-456", AlertStatusDto.RESOLVED, "prod");

        AlertEvent event = mapper.toEvent(dto);

        assertThat(event.status()).isEqualTo(AlertStatus.RESOLVED);
    }

    @Test
    void shouldMapUnknownStatusForInvalidValue() {
        AlertDto dto = TestFixtures.alertDto("fp-789", AlertStatusDto.fromString("test"), "prod");

        AlertEvent event = mapper.toEvent(dto);

        assertThat(event.status()).isEqualTo(AlertStatus.UNKNOWN);
    }

    @Test
    void shouldThrowIfFingerprintInvalid() {
        AlertDto dto = TestFixtures.alertDto(null, AlertStatusDto.FIRING, "prod");

        assertThrows(IllegalArgumentException.class, () -> mapper.toEvent(dto));
    }

    @Test
    void shouldThrowIfLabelFieldsInvalid() {
        AlertLabelsDto labelsDto = new AlertLabelsDto(
                "",
                "",
                "",
                "",
                AlertSeverityDto.fromString("INFO")
        );
        AlertDto dto = new AlertDto(
                "fp-001",
                AlertStatusDto.fromString("firing"),
                labelsDto,
                Instant.now().toString()
        );

        assertThrows(IllegalArgumentException.class, () -> mapper.toEvent(dto));
    }


}
