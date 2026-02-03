package com.example.alertbridge.application.alertstate;

import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertStateRepository;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class GetCurrentAlertStatesUseCase {

    private final AlertStateRepository repository;

    public GetCurrentAlertStatesUseCase(AlertStateRepository repository) {
        this.repository = repository;
    }

    public List<AlertState> execute() {
        return this.repository.findAllActive();
    }

}