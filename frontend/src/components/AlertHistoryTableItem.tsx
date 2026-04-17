import type { HistoryAlert } from '../types/Alert';
import { formatAlertTime } from '../utils/Formatter';
import "./AlertHistoryTableItem.css"

interface AlertItemProps {
    alert: HistoryAlert;
    isActive: boolean
}

export default function AlertHistoryTableItem({ alert, isActive }: AlertItemProps) {
    return (
        <tr className={isActive ? "active-row" : ""}>
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

