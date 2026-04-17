import type { AlertHistory } from "../types/Alert";
import AlertHistoryTableItem from "./AlertHistoryTableItem";

interface AlertHistoryProps {
    alertHistory: AlertHistory;
}
export default function AlertHistoryTable({ alertHistory }: AlertHistoryProps) {
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
                {alertHistory.alerts.map((alert) => (
                    <AlertHistoryTableItem
                        key={`${alert.fingerprint}-${alert.receivedAt}`}
                        alert={alert}
                    ></AlertHistoryTableItem>
                ))}
            </tbody>
        </table>
    );
}
