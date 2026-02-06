import type { AlertState } from "../types/Alert";

interface AlertItemProps {
  alert: AlertState;
}

export default function AlertTableItem({ alert }: AlertItemProps) {
  return (
    <tr>
      <td>{alert.alertName}</td>
      <td>{alert.instance.split(':')[0]}</td>
      <td>{alert.job}</td>
      <td>{alert.severity}</td>
      <td>{new Date(alert.startsAt).toLocaleDateString()}</td>
      <td>{alert.status}</td>
    </tr>
  );
}
