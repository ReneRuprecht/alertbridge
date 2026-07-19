import type { AlertHistory } from "../domain/alertHistory";

interface AlertHistoryItemProps {
  alert: AlertHistory;
}

export function AlertHistoryItem({ alert }: AlertHistoryItemProps) {
  return (
    <div>
      <div>{alert.alertName}</div>
      <div>{alert.instance}</div>
      <div>{alert.status}</div>
      <div>{alert.startsAt.toLocaleString()}</div>
      <div>{alert.receivedAt.toLocaleString()}</div>
    </div>
  );
}
