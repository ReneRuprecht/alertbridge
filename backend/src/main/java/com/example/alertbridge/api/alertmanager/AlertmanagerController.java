package com.example.alertbridge.api.alertmanager;

import com.example.alertbridge.api.alertmanager.dto.AlertmanagerPayloadDto;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("api/v1/alerts")
public class AlertmanagerController {

    @PostMapping
    public void receive(@RequestBody AlertmanagerPayloadDto alertmanagerPayloadDto) {
        alertmanagerPayloadDto.alerts().forEach(alertDto -> {
            System.out.println(alertDto.toString());
        });
    }
}
