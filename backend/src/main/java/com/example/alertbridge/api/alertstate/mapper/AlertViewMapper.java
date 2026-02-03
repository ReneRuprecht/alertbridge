package com.example.alertbridge.api.alertstate.mapper;

import com.example.alertbridge.api.alertstate.dto.AlertViewDto;
import com.example.alertbridge.domain.model.AlertState;
import org.springframework.stereotype.Component;

@Component
public class AlertViewMapper {
    public AlertViewDto toDto(AlertState state) {
        return new AlertViewDto(
                state.fingerprint().value(),
                state.labels().alertName().value(),
                state.labels().environment().value(),
                state.labels().instance().value(),
                state.labels().job().value(),
                state.labels().severity().name(),
                state.status().name(),
                state.startsAt().value()
        );

    }
}
