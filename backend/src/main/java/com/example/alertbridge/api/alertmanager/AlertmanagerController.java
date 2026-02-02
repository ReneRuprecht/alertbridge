package com.example.alertbridge.api.alertmanager;

import com.example.alertbridge.api.alertmanager.dto.AlertDto;
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
        for (AlertDto alert : alertmanagerPayloadDto.alerts) {
            System.out.printf("Fingerprint %s\n", alert.fingerprint);
            System.out.printf("Status %s\n", alert.status);
            System.out.println("Labels");
            System.out.printf("     Alertname %s\n", alert.labels.alertName);
            System.out.printf("     Environment %s\n", alert.labels.environment);
            System.out.printf("     Instance %s\n", alert.labels.instance);
            System.out.printf("     Job %s\n", alert.labels.job);
            System.out.printf("     Severity %s\n", alert.labels.severity);
            System.out.printf("StartsAt %s\n", alert.startsAt);
        }


    }
}
