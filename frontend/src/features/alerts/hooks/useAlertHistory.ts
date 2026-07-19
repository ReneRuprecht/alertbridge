import { useEffect, useState } from 'react';
import { getAlertHistory } from '../api/historyAlertApi';
import type { AlertHistory } from '../types/alertHistory';

export function useAlertHistory(instance: string) {
  const [history, setHistory] = useState<AlertHistory[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    async function loadHistory() {
      try {
        const data = await getAlertHistory(instance);
        setHistory(data);
      } catch (e) {
        setError(e instanceof Error ? e : new Error('Unknown error'));
      } finally {
        setLoading(false);
      }
    }

    loadHistory();
  }, [instance]);

  return {
    history,
    loading,
    error,
  };
}
