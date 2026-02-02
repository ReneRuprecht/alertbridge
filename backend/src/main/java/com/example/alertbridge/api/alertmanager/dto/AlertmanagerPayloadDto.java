package com.example.alertbridge.api.alertmanager.dto;


import java.util.List;

public record AlertmanagerPayloadDto(List<AlertDto> alerts) {
}
