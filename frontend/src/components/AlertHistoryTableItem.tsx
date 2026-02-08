import type { AlertEvent } from "../types/Alert";
import { formatAlertStartsAt } from "../utils/formatter";

interface AlertItemProps {
  event: AlertEvent;
}

export default function AlertHistoryTableItem({ event }: AlertItemProps) {
  return (
    <tr>
      <td>{event.alertName}</td>
      <td>{event.job}</td>
      <td>{event.severity}</td>
      <td>{formatAlertStartsAt(event.startsAt)}</td>
      <td>{event.status}</td>
    </tr>
  );
}
