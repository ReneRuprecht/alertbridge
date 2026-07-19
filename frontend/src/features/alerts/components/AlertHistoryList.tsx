import type { AlertHistory } from '../types/alertHistory';
import { AlertHistoryItem } from './AlertHistoryItem';

interface AlertHistoryListProps {
  alerts: AlertHistory[];
}

export function AlertHistoryList({ alerts }: AlertHistoryListProps) {
  return (
    <div>
      {alerts.map((alert) => (
        <AlertHistoryItem key={alert.fingerprint} alert={alert} />
      ))}
    </div>
  );
}
