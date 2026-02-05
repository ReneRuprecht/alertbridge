# AlertBridge

## Beschreibung
Das Projekt stellt das Backend für eine zentrale Alert-Management-Plattform bereit.


Ziel ist es, Alerts (aktuell Prometheus Alertmanager) nicht nur entgegenzunehmen, 
sondern sie strukturiert aufzubereiten, dauerhaft zu speichern und komfortabel verwaltbar zu machen.

Im Gegensatz zur reinen Alertmanager-Oberfläche liegt der Fokus auf:
- übersichtliche grafischen Darstellung aktueller und historischer Alerts
- nachträglichen Analyse von Alert-Verläufen
- aktive Alert-Verwaltung

Das Backend dient dabei als zentrale Quelle der Wahrheit für:
- aktuelle Alert Zustände
- Alert Historie
- zukünftige Verwaltungsaktionen (z. B. Silencing, Quittieren, Eskalieren)

## Aktuelle Features

- Entgegennahme von Alerts (Alertmanager Webhook)
- Verwaltung des aktuellen Alert-Zustands (Redis)
- Persistierung der Alert-Historie (PostgreSQL)

#### Redis – Current Alert State

- Speichert den aktuellen Zustand pro Fingerprint
- Schneller Zugriff
- Kein historischer Kontext

#### PostgreSQL – Alert-Historie

- Speichert historische Alert-Events
- Grundlage für History-Ansichten

## Techstack

### Backend
- Java 21 & Spring Boot
- Redis
- Postgres

## Pipeline
Aktuell existiert ein Workflow, der das Backend testet. 
Mit dem Fortschreiten des Projekts wird dieser weiter ausgebaut.

Sobald Änderungen auf `main` oder `dev` erfolgen, werden die Test-Reports verschlüsselt in der Summary bereitgestellt.
Auf allen anderen Branches wird das Hochladen der Test-Reports übersprungen.

## Ausblick & Erweiterungen

Das Projekt ist bewusst modular aufgebaut und für folgende Erweiterungen vorbereitet:

### Grafische Alert-Übersicht

- Übersicht aktueller Alerts (z. B. gruppiert nach Umgebung, Service, Instanz)
- Detailansicht eines Alerts inkl. vollständiger Historie

### Alert-Historie & Analyse

- Historische Auswertung von Alerts
- Erkennung wiederkehrender Alerts
- Analyse von Dauer, Häufigkeit und Zeitfenstern

### Alert-Verwaltung

- Silencing von Alerts
- Quittieren von Alerts

### Aktives Versenden von Alerts

- Erzeugen und Versenden eigener Alerts
- Nutzung des Systems als zentrales Alert-Gateway
- Weiterleitung an externe Systeme (z.B. Discord etc.)

