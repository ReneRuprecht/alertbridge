package com.example.alertbridge.api.alerthistory;

import com.example.alertbridge.api.alerthistory.dto.AlertHistoryViewDto;
import com.example.alertbridge.api.alerthistory.mapper.AlertHistoryViewMapper;
import com.example.alertbridge.application.alertstate.GetAlertHistoryByFingerprintUseCase;
import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.value.AlertFingerprint;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/v1/alert-history")
public class HistoryAlertController {
    private final GetAlertHistoryByFingerprintUseCase getAlertHistoryByFingerprintUseCase;
    private final AlertHistoryViewMapper mapper;

    public HistoryAlertController(GetAlertHistoryByFingerprintUseCase getAlertHistoryByFingerprintUseCase,
                                  AlertHistoryViewMapper mapper) {
        this.getAlertHistoryByFingerprintUseCase = getAlertHistoryByFingerprintUseCase;
        this.mapper = mapper;
    }

    @GetMapping("/{fingerprint}")
    public ResponseEntity<AlertHistoryViewDto> getAlert(@PathVariable String fingerprint) {
        AlertFingerprint fp = new AlertFingerprint(fingerprint);

        List<AlertEvent> events = this.getAlertHistoryByFingerprintUseCase.execute(fp);

        if (events.isEmpty()) {
            return ResponseEntity.notFound().build();
        }

        AlertHistoryViewDto dto = this.mapper.toDto(events);
        return ResponseEntity.ok(dto);
    }

}
