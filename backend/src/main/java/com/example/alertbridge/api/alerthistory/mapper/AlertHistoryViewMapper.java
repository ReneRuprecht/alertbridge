package com.example.alertbridge.api.alerthistory.mapper;

import com.example.alertbridge.api.alerthistory.dto.AlertHistoryViewDto;
import com.example.alertbridge.api.alerthistory.dto.value.AlertHistoryInfoDto;
import com.example.alertbridge.domain.event.AlertEvent;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class AlertHistoryViewMapper {

    private static AlertHistoryInfoDto toEventDto(AlertEvent event) {
        return new AlertHistoryInfoDto(
                event.labels().alertName().value(),
                event.labels().job().value(),
                event.labels().severity().name(),
                event.status().name(),
                event.startsAt().value(),
                event.receivedAt().value()
        );
    }


    public AlertHistoryViewDto toDto(List<AlertEvent> events) {

        AlertEvent event = events.getFirst();

        List<AlertHistoryInfoDto> dto = events
                .stream()
                .map(AlertHistoryViewMapper::toEventDto)
                .toList();

        return new AlertHistoryViewDto(
                event.fingerprint().value(),
                event.labels().instance().value(),
                dto
        );
    }
}
