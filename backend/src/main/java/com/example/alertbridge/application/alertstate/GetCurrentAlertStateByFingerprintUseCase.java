package com.example.alertbridge.application.alertstate;

import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertStateRepository;
import com.example.alertbridge.domain.value.AlertFingerprint;
import org.springframework.stereotype.Component;

import java.util.Optional;

@Component
public class GetCurrentAlertStateByFingerprintUseCase {

    private final AlertStateRepository repository;

    public GetCurrentAlertStateByFingerprintUseCase(AlertStateRepository repository) {
        this.repository = repository;
    }


    public Optional<AlertState> execute(AlertFingerprint fingerprint) {
        return repository.findByFingerprint(fingerprint).filter(AlertState::isActive);
    }
}
