package com.example.alertbridge.api.alertmanager;

import com.example.alertbridge.api.alertmanager.dto.AlertmanagerPayloadDto;
import com.example.alertbridge.api.alertmanager.mapper.AlertApiMapper;
import com.example.alertbridge.application.alertstate.SaveAlertsToHistoryUseCase;
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
    private final SaveAlertsToHistoryUseCase saveAlertsToHistoryUseCase;

    public AlertmanagerController(AlertApiMapper mapper,
                                  SyncAlertsUseCase syncAlertsUseCase,
                                  SaveAlertsToHistoryUseCase saveAlertsToHistoryUseCase) {
        this.mapper = mapper;
        this.syncAlertsUseCase = syncAlertsUseCase;
        this.saveAlertsToHistoryUseCase = saveAlertsToHistoryUseCase;
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
        this.saveAlertsToHistoryUseCase.execute(events);

        return ResponseEntity.ok().build();
    }
}
