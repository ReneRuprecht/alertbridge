import type { AlertHistory } from "../types/Alert";
import AlertHistoryTableItem from "./AlertHistoryTableItem";

interface AlertHistoryProps {
  alertHistory: AlertHistory
}
export default function AlertHistoryTable({ alertHistory }: AlertHistoryProps) {

  return (
    <table>
      <thead>
        <tr>
          <th>Alertname</th>
          <th>Job</th>
          <th>Severity</th>
          <th>Angefangen</th>
          <th>Erhalten</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {alertHistory.events.map(event => (
          <AlertHistoryTableItem key={`${alertHistory.fingerprint}-${event.receivedAt}`} event={event}></AlertHistoryTableItem>
        ))}
      </tbody>
    </table>
  )
}