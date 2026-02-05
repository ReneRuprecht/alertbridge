package com.example.alertbridge.api.alerthistory.mapper;

import com.example.alertbridge.api.alerthistory.dto.AlertHistoryViewDto;
import com.example.alertbridge.api.alerthistory.dto.value.AlertHistoryEventDto;
import com.example.alertbridge.domain.event.AlertEvent;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class AlertHistoryViewMapper {

    private static AlertHistoryEventDto toEventDto(AlertEvent event) {
        return new AlertHistoryEventDto(
                event.fingerprint().value(),
                event.labels().alertName().value(),
                event.labels().environment().value(),
                event.labels().instance().value(),
                event.labels().job().value(),
                event.labels().severity().name(),
                event.status().name(),
                event.startsAt().value()
        );
    }


    public AlertHistoryViewDto toDto(List<AlertEvent> events) {

        List<AlertHistoryEventDto> dto = events
                .stream()
                .map(AlertHistoryViewMapper::toEventDto)
                .toList();

        return new AlertHistoryViewDto(dto);
    }
}
