import { useEffect, useState } from 'react';
import { getActiveAlerts } from '../api/GetActiveAlerts';
import ActiveAlertTableItem from './ActiveAlertTableItem';
import type { ActiveAlerts, Alert } from '../types/Alert';

export default function ActiveAlertTable() {
    const [alerts, setAlerts] = useState<Alert[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {

        const fetchActiveAlerts = async () => {
            try {
                const data: ActiveAlerts = await getActiveAlerts()
                setAlerts(data.alerts)
            }
            catch (err: any) {
                console.error(err)
                setError(err.message)
                setAlerts([])
            }
            finally {
                setLoading(false)
            }
        }

        fetchActiveAlerts()

    }, []);

    if (loading) return <h1>Lädt aktuelle Alerts</h1>;
    if (error) return <h1>Fehler: {error}</h1>;
    if (alerts === undefined || alerts.length === 0) return <h1>Keine Alerts</h1>;

    return (
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Job</th>
                    <th>Instance</th>
                    <th>Angefangen</th>
                    <th>Status</th>
                </tr>
            </thead>
            <tbody>
                {alerts.map((alert) => (
                    <ActiveAlertTableItem key={alert.fingerprint} alert={alert} />
                ))}
            </tbody>
        </table>
    );
}
