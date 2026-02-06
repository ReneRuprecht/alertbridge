import { useEffect, useState } from "react";
import { getCurrentAlerts } from "../api/GetCurrentAlerts";
import AlertTableItem from "./AlertTableItem";
import type { AlertState } from "../types/Alert";

export default function AlertTable() {
  const [alerts, setAlerts] = useState<AlertState[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    getCurrentAlerts()
      .then((data) => { setAlerts(data); setLoading(false); })
      .catch((err) => { setError(err.message); setLoading(false); });
  }, []);

  if (loading) return <p>Lädt aktuelle Alerts</p>;
  if (error) return <p>Fehler: {error}</p>;

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
        {alerts.map(alert => (
          <AlertTableItem
            key={alert.fingerprint}
            alert={alert}
          />
        ))}
      </tbody>
    </table>
  );
}
