import { useEffect, useState } from 'react';
import { getCurrentAlerts } from '../api/GetCurrentAlerts';
import AlertTableItem from './AlertTableItem';
import type { AlertState } from '../types/Alert';

export default function AlertTable() {
  const [alerts, setAlerts] = useState<AlertState[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    let loadingTimer: ReturnType<typeof setTimeout>;
    const minVisible = 500;
    const startTime = Date.now();

    async function loadAlerts() {
      try {
        const data = await getCurrentAlerts();
        setAlerts(data);
      } catch (e: any) {
        setError(e.message);
      } finally {
        const elapsed = Date.now() - startTime;
        const remaining = minVisible - elapsed;
        if (remaining > 0) {
          loadingTimer = setTimeout(() => setLoading(false), remaining);
        } else {
          setLoading(false);
        }
      }
    }

    loadAlerts();

    return () => clearTimeout(loadingTimer);
  }, []);

  if (loading) return <h1>Lädt aktuelle Alerts</h1>;
  if (error) return <h1>Fehler: {error}</h1>;

  return (
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Instance</th>
          <th>Job</th>
          <th>Severity</th>
          <th>Angefangen</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {alerts.map((alert) => (
          <AlertTableItem key={alert.fingerprint} alert={alert} />
        ))}
      </tbody>
    </table>
  );
}
