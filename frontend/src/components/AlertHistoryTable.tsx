import { useMemo } from "react";
import type { AlertHistory, HistoryAlert } from "../types/Alert";
import AlertHistoryTableItem from "./AlertHistoryTableItem";

interface AlertHistoryProps {
    alertHistory: AlertHistory;
}

function getKeyFromAlert(alert: HistoryAlert): string {
    return `${alert.fingerprint}-${alert.startsAt}-${alert.job}`;
}

export default function AlertHistoryTable({ alertHistory }: AlertHistoryProps) {
    const activeAlerts = useMemo(() => {
        const map = new Map<string, HistoryAlert>();

        alertHistory.alerts.forEach(alert => {
            const key = getKeyFromAlert(alert)

            const existing = map.get(key);

            if (!existing) {
                map.set(key, alert);
            } else {
                map.delete(key)
            }

        });

        return map;
    }, [alertHistory.alerts]);

    return (
        <table>
            <thead>
                <tr>
                    <th>Alertname</th>
                    <th>Job</th>
                    <th>Schweregrad</th>
                    <th>Beschreibung</th>
                    <th>Erhalten</th>
                    <th>Angefangen</th>
                    <th>Status</th>
                </tr>
            </thead>
            <tbody>
                {alertHistory.alerts.map((alert) => {
                    const key = getKeyFromAlert(alert)
                    const isActive = activeAlerts.get(key)?.status.toLowerCase() === "firing";

                    return (
                        <AlertHistoryTableItem
                            key={`${alert.fingerprint}-${alert.receivedAt}`}
                            alert={alert}
                            isActive={isActive}
                        />
                    );
                })}
            </tbody>
        </table>
    );
}
