import type { Alert } from '../types/Alert';
import { formatAlertInstance, formatAlertStatus, formatAlertTime } from '../utils/Formatter';

interface ActiveAlertItemProps {
    alert: Alert;
}

export default function ActiveAlertTableItem({ alert }: ActiveAlertItemProps) {

    return (
        <tr>
            <td>{alert.alertName}</td>
            <td>{formatAlertInstance(alert.instance)}</td>
            <td>{formatAlertTime(alert.startsAt)}</td>
            <td>{formatAlertStatus(alert.status)}</td>
            <td></td>
        </tr>
    );
}

