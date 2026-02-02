package com.example.alertbridge.api.alertmanager;

import com.example.alertbridge.api.alertmanager.dto.AlertmanagerPayloadDto;
import com.example.alertbridge.api.alertmanager.mapper.AlertApiMapper;
import com.example.alertbridge.domain.event.AlertEvent;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("api/v1/alerts")
public class AlertmanagerController {

    private final AlertApiMapper mapper;

    public AlertmanagerController(AlertApiMapper mapper) {
        this.mapper = mapper;
    }

    @PostMapping
    public void receive(@RequestBody AlertmanagerPayloadDto alertmanagerPayloadDto) {
        alertmanagerPayloadDto.alerts().forEach(alertDto -> {
            AlertEvent event = mapper.toEvent(alertDto);
            System.out.println(event.toString());
        });
    }
}
