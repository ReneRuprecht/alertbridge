package com.example.alertbridge.application.alertstate;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertHistoryRepository;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class SaveAlertsToHistoryUseCase {

    private final AlertHistoryRepository historyRepository;

    public SaveAlertsToHistoryUseCase(AlertHistoryRepository historyRepository) {
        this.historyRepository = historyRepository;
    }

    public void execute(List<AlertEvent> events) {
        for (AlertEvent event : events) {
            save(event);
        }
    }

    private void save(AlertEvent event) {

        AlertState historyState = AlertState.fromEvent(event);

        if (!this.historyRepository.existsByAlertHash(historyState)) {
            this.historyRepository.save(historyState);
        }

    }
}
