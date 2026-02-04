package com.example.alertbridge.application.alertstate;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertHistoryRepository;
import com.example.alertbridge.domain.port.AlertStateRepository;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class SyncAlertsUseCase {
    private final AlertStateRepository repository;
    private final AlertHistoryRepository historyRepository;

    public SyncAlertsUseCase(AlertStateRepository repository,
                             AlertHistoryRepository historyRepository) {
        this.repository = repository;
        this.historyRepository = historyRepository;
    }

    public void execute(List<AlertEvent> events) {
        for (AlertEvent event : events) {
            syncSingle(event);
        }
    }

    private void syncSingle(AlertEvent event) {

        AlertState state = repository
                .findByFingerprint(event.fingerprint())
                .orElseGet(() -> AlertState.fromEvent(event));

        state.apply(event);
        this.repository.save(state);
        this.historyRepository.save(state);
    }

}
