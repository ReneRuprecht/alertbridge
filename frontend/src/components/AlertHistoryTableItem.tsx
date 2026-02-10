import type { AlertEvent } from '../types/Alert';
import { formatAlertTime } from '../utils/Formatter';

interface AlertItemProps {
  event: AlertEvent;
}

export default function AlertHistoryTableItem({ event }: AlertItemProps) {
  return (
    <tr>
      <td>{event.alertName}</td>
      <td>{event.job}</td>
      <td>{event.severity}</td>
      <td>{formatAlertTime(event.startsAt)}</td>
      <td>{formatAlertTime(event.receivedAt)}</td>
      <td>{event.status}</td>
    </tr>
  );
}
