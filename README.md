# Problem
Der Alertmanager liefert viele Daten, aber es ist schwer, daraus eine schnelle und übersichtliche Darstellung aller Server und ihrer aktuellen Alerts zu bekommen.

# Idee

Ziel ist eine zentrale Plattform zur Darstellung und Auswertung von Alerts.
Der Alertmanager sendet Alerts an ein Backend, welches diese entgegennimmt, verarbeitet und speichert.
Die gespeicherten Daten sollen anschließend strukturiert abgefragt und visualisiert werden können.

# Ziel

## Backend
- Empfang von Alerts über HTTP
- Speicherung und Verarbeitung der Alerts
- Bereitstellung von APIs zur Abfrage aktiver und historischer Alerts

## Frontend
- Übersicht aller Server mit Anzahl aktiver Alerts
- Detailansicht pro Server (aktive + historische Alerts)

## Deployment
- Bereitstellung über Docker
- Konfiguration über Environment Variables
