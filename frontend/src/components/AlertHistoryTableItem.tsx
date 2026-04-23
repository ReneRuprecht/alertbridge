import type { HistoryAlert } from '../types/Alert';
import { formatAlertTime, getSeverityStyle } from '../utils/Formatter';
import "./AlertHistoryTableItem.css"

interface AlertItemProps {
    alert: HistoryAlert;
    isActive: boolean
}

export default function AlertHistoryTableItem({ alert, isActive }: AlertItemProps) {
    return (
        <tr>
            <td>{alert.alertName}</td>
            <td>{alert.job}</td>
            <td className={getSeverityStyle(alert.severity)}>{alert.severity}</td>
            <td>{alert.description}</td>
            <td>{formatAlertTime(alert.receivedAt)}</td>
            <td>{formatAlertTime(alert.startsAt)}</td>
            <td className={isActive ? "firing" : "resolved"}>{alert.status}</td>
        </tr>
    );
}

