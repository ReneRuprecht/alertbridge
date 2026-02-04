package com.example.alertbridge.application.alertstate;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.port.AlertHistoryRepository;
import com.example.alertbridge.domain.value.AlertFingerprint;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class GetAlertHistoryByFingerprintUseCase {
    private final AlertHistoryRepository repository;

    public GetAlertHistoryByFingerprintUseCase(AlertHistoryRepository repository) {
        this.repository = repository;
    }


    public List<AlertEvent> execute(AlertFingerprint fingerprint) {
        return repository.findByFingerprint(fingerprint);
    }
}

