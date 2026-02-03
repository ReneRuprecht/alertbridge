package com.example.alertbridge.api.alertmanager;

import com.example.alertbridge.api.alertmanager.dto.AlertmanagerPayloadDto;
import com.example.alertbridge.api.alertmanager.mapper.AlertApiMapper;
import com.example.alertbridge.application.alertstate.SyncAlertsUseCase;
import com.example.alertbridge.domain.event.AlertEvent;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("api/v1/alerts")
public class AlertmanagerController {

    private final AlertApiMapper mapper;
    private final SyncAlertsUseCase syncAlertsUseCase;

    public AlertmanagerController(AlertApiMapper mapper, SyncAlertsUseCase syncAlertsUseCase) {
        this.mapper = mapper;
        this.syncAlertsUseCase = syncAlertsUseCase;
    }

    @PostMapping
    public ResponseEntity<Void> receive(@RequestBody AlertmanagerPayloadDto alertmanagerPayloadDto) {
        List<AlertEvent> events = alertmanagerPayloadDto
                .alerts()
                .stream()
                .map(this.mapper::toEvent)
                .toList();

        System.out.printf("Sync %d events\n", events.size());

        this.syncAlertsUseCase.execute(events);

        return ResponseEntity.ok().build();
    }
}
