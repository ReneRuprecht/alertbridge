import { useEffect, useState } from 'react';
import { getCurrentAlerts } from '../api/alertApi';
import type { Alert } from '../domain/alert';

export function useCurrentAlerts() {
  const [alerts, setAlerts] = useState<Alert[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    async function loadAlerts() {
      try {
        setError(null);
        const data = await getCurrentAlerts();
        setAlerts(data);
      } catch (e) {
        setError(e instanceof Error ? e : new Error('Failed to load current alerts'));
      } finally {
        setLoading(false);
      }
    }

    loadAlerts();
  }, []);

  return {
    alerts,
    loading,
    error,
  };
}
