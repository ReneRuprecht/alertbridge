package com.example.alertbridge.api.alertstate;

import com.example.alertbridge.api.alertstate.dto.AlertViewDto;
import com.example.alertbridge.api.alertstate.mapper.AlertViewMapper;
import com.example.alertbridge.application.alertstate.GetCurrentAlertStateByFingerprintUseCase;
import com.example.alertbridge.application.alertstate.GetCurrentAlertStatesUseCase;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.value.AlertFingerprint;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/api/v1/alert-states")
public class CurrentAlertController {
    private final GetCurrentAlertStatesUseCase getCurrentAlertStatesUseCase;
    private final GetCurrentAlertStateByFingerprintUseCase getCurrentAlertStateByFingerprintUseCase;
    private final AlertViewMapper mapper;

    public CurrentAlertController(GetCurrentAlertStatesUseCase getCurrentAlertStatesUseCase,
                                  GetCurrentAlertStateByFingerprintUseCase getCurrentAlertStateByFingerprintUseCase,
                                  AlertViewMapper mapper) {
        this.getCurrentAlertStatesUseCase = getCurrentAlertStatesUseCase;
        this.getCurrentAlertStateByFingerprintUseCase = getCurrentAlertStateByFingerprintUseCase;
        this.mapper = mapper;
    }

    @GetMapping
    @CrossOrigin
    public List<AlertViewDto> getCurrentAlerts() {
        return this.getCurrentAlertStatesUseCase.execute().stream().map(mapper::toDto).toList();
    }

    @GetMapping("/{fingerprint}")
    @CrossOrigin
    public ResponseEntity<AlertViewDto> getAlert(@PathVariable String fingerprint) {
        AlertFingerprint fp = new AlertFingerprint(fingerprint);

        Optional<AlertState> state = this.getCurrentAlertStateByFingerprintUseCase.execute(fp);

        if (state.isEmpty()) {
            return ResponseEntity.notFound().build();
        }

        AlertViewDto dto = this.mapper.toDto(state.get());
        return ResponseEntity.ok(dto);

    }

}
