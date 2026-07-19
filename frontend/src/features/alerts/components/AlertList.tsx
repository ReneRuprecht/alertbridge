import type { Alert } from '../types/alert';
import { AlertItem } from './AlertItem';

interface AlertListProps {
  alerts: Alert[];
}

export function AlertList({ alerts }: AlertListProps) {
  return (
    <div>
      {alerts.map((alert) => (
        <AlertItem key={alert.fingerprint} alert={alert} />
      ))}
    </div>
  );
}
