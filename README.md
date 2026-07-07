# Alertbridge
Ein Service zum Empfangen, Speichern und Verwalten von Prometheus Alertmanager Alerts.

Der Service nimmt eingehende Webhooks vom Prometheus Alertmanager entgegen,
speichert die vollständige Alert-Historie in PostgreSQL und verwaltet den aktuellen Alert-Zustand in Redis.

## Problem
Der Alertmanager liefert viele Daten, aber es ist schwer, daraus eine schnelle und übersichtliche Darstellung aller Server und ihrer aktuellen Alerts zu bekommen.

## Idee

Ziel ist eine zentrale Plattform zur Darstellung und Auswertung von Alerts.
Der Alertmanager sendet Alerts an ein Backend, welches diese entgegennimmt, verarbeitet und speichert.
Die gespeicherten Daten sollen anschließend strukturiert abgefragt und visualisiert werden können.

## Ziel

### Backend
- Empfang von Alerts über HTTP
- Speicherung und Verarbeitung der Alerts
- Bereitstellung von APIs zur Abfrage aktiver und historischer Alerts

### Frontend
- Übersicht aller Server mit Anzahl aktiver Alerts
- Detailansicht pro Server (aktive + historische Alerts)

### Deployment
- Bereitstellung über Docker
- Konfiguration über Environment Variables


## Techstack
- Java 21
- Spring Boot
- PostgreSQL
- Redis
- Docker Compose
- Testcontainer
- Github Actions

## Aktueller Funktionsumfang

### Implementiert
- Empfang von Prometheus Alertmanager Webhooks
- Speicherung der Alert-Historie in PostgreSQL
- Speicherung der aktuellen Alerts in Redis
- Unit- und Integrationstests
- Testcontainer für PostgreSQL und Redis
- CI-Pipeline für automatisierte Tests

### In Arbeit
- Read API für aktuelle Alerts
- Read API für Alert-Historie
