import type { Alert } from "../domain/alert";

interface AlertItemProps {
  alert: Alert;
}

export function AlertItem({ alert }: AlertItemProps) {
  return (
    <div>
      <h3>{alert.alertName}</h3>
      <p>{alert.instance}</p>
      <p>{alert.job}</p>
      <p>{alert.severity}</p>
    </div>
  );
}
