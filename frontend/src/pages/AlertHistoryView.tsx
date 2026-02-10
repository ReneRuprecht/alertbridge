import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router';
import type { AlertHistory } from '../types/Alert';
import { getHistoryAlert } from '../api/GetHistoryAlert';
import AlertHistoryTable from '../components/AlertHistoryTable';
import { formatAlertInstance } from '../utils/Formatter';
import './AlertHistoryView.css';

export default function AlertHistoryView() {
  const { alertInstance } = useParams();
  const [history, setHistory] = useState<AlertHistory>();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    let loadingTimer: ReturnType<typeof setTimeout>;
    const minVisible = 500;
    const startTime = Date.now();

    async function loadAlertHistory() {
      try {
        if (!alertInstance || alertInstance === 'undefined') {
          throw new Error('Missing alertInstance');
        }
        const data = await getHistoryAlert(alertInstance);
        setHistory(data);
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

    loadAlertHistory();

    return () => clearTimeout(loadingTimer);
  }, []);

  if (loading) return <h1>Lädt History von {alertInstance}</h1>;
  if (error) return <h1>Fehler: {error}</h1>;
  if (!history) return <h1>Keine History gefunden</h1>;
  if (history.events.length === 0)
    return <h1>Keine Event History für {alertInstance} gefunden</h1>;

  return (
    <>
      <div onClick={() => navigate('/')} style={{ display: 'flex' }}>
        <button style={{ marginRight: 'auto' }}>BACK</button>
      </div>
      <h1>{formatAlertInstance(history.instance)}</h1>
      <AlertHistoryTable alertHistory={history}></AlertHistoryTable>
    </>
  );
}
