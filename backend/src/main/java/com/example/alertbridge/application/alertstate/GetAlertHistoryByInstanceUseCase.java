package com.example.alertbridge.application.alertstate;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.port.AlertHistoryRepository;
import com.example.alertbridge.domain.value.AlertInstance;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class GetAlertHistoryByInstanceUseCase {
    private final AlertHistoryRepository repository;

    public GetAlertHistoryByInstanceUseCase(AlertHistoryRepository repository) {
        this.repository = repository;
    }


    public List<AlertEvent> execute(AlertInstance instance) {
        return repository.findByAlertInstance(instance);
    }
}

