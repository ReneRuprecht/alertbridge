import { useNavigate } from 'react-router';
import type { AlertState } from '../types/Alert';
import { formatAlertInstance } from '../utils/Formatter';

interface AlertItemProps {
  alert: AlertState;
}

export default function AlertTableItem({ alert }: AlertItemProps) {
  const navigate = useNavigate();

  return (
    <tr
      onClick={() =>
        navigate(`/history/${formatAlertInstance(alert.instance)}`)
      }
    >
      <td>{alert.alertName}</td>
      <td>{formatAlertInstance(alert.instance)}</td>
      <td>{alert.job}</td>
      <td>{alert.severity}</td>
      <td>{formatAlertInstance(alert.startsAt)}</td>
      <td>{alert.status}</td>
      <td></td>
    </tr>
  );
}
