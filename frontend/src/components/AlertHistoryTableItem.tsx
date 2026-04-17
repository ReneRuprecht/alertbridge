import type { HistoryAlert } from '../types/Alert';
import { formatAlertTime } from '../utils/Formatter';

interface AlertItemProps {
    alert: HistoryAlert;
}

export default function AlertHistoryTableItem({ alert }: AlertItemProps) {
    return (
        <tr>
            <td>{alert.alertName}</td>
            <td>{alert.job}</td>
            <td>{alert.severity}</td>
            <td>{alert.description}</td>
            <td>{formatAlertTime(alert.receivedAt)}</td>
            <td>{formatAlertTime(alert.startsAt)}</td>
            <td>{alert.status}</td>
        </tr>
    );
}

