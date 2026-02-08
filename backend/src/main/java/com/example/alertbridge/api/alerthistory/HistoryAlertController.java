package com.example.alertbridge.api.alerthistory;

import com.example.alertbridge.api.alerthistory.dto.AlertHistoryViewDto;
import com.example.alertbridge.api.alerthistory.mapper.AlertHistoryViewMapper;
import com.example.alertbridge.application.alertstate.GetAlertHistoryByInstanceUseCase;
import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.value.AlertInstance;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/v1/alert-history")
public class HistoryAlertController {
    private final GetAlertHistoryByInstanceUseCase getAlertHistoryByInstanceUseCase;
    private final AlertHistoryViewMapper mapper;

    public HistoryAlertController(GetAlertHistoryByInstanceUseCase getAlertHistoryByInstanceUseCase,
                                  AlertHistoryViewMapper mapper) {
        this.getAlertHistoryByInstanceUseCase = getAlertHistoryByInstanceUseCase;
        this.mapper = mapper;
    }

    @GetMapping("/{instance}")
    public ResponseEntity<AlertHistoryViewDto> getAlertHistory(@PathVariable String instance) {
        AlertInstance in = new AlertInstance(instance);

        List<AlertEvent> events = this.getAlertHistoryByInstanceUseCase.execute(in);

        if (events.isEmpty()) {
            return ResponseEntity.notFound().build();
        }

        AlertHistoryViewDto dto = this.mapper.toDto(events);
        return ResponseEntity.ok(dto);
    }

}
