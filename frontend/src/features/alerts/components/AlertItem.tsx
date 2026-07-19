import type { Alert } from '../domain/alert';

interface AlertItemProps {
  alert: Alert;
  onClick: () => void;
}

export function AlertItem({ alert, onClick }: AlertItemProps) {
  return (
    <div onClick={onClick}>
      <h3>{alert.alertName}</h3>
      <p>{alert.instance}</p>
      <p>{alert.job}</p>
      <p>{alert.severity}</p>
    </div>
  );
}
