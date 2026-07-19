import type { Alert } from '../domain/alert';
import { AlertItem } from './AlertItem';

interface AlertListProps {
  alerts: Alert[];
  onSelect: (alert: Alert) => void;
}

export function AlertList({ alerts, onSelect }: AlertListProps) {
  return (
    <div>
      {alerts.map((alert) => (
        <AlertItem key={alert.fingerprint} alert={alert} onClick={() => onSelect(alert)} />
      ))}
    </div>
  );
}
