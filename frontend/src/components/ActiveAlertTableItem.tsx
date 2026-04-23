import { useNavigate } from 'react-router';
import type { Alert } from '../types/Alert';
import { formatAlertInstance, formatAlertTime, getSeverityStyle } from '../utils/Formatter';
import "./ActiveAlertTableItem.css"

interface ActiveAlertItemProps {
    alert: Alert;
}

export default function ActiveAlertTableItem({ alert }: ActiveAlertItemProps) {
    const navigate = useNavigate();

    return (
        <tr
            onClick={() =>
                navigate(`/${formatAlertInstance(alert.instance)}`)
            }
        >
            <td>{alert.alertName}</td>
            <td>{alert.job}</td>
            <td>{formatAlertInstance(alert.instance)}</td>
            <td>{formatAlertTime(alert.startsAt)}</td>
            <td className={getSeverityStyle(alert.severity)}>{alert.severity}</td>
            <td></td>
        </tr>
    );
}

