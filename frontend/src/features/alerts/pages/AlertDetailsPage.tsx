import { useParams } from 'react-router';
import { useAlertHistory } from '../hooks/useAlertHistory';
import { AlertHistoryList } from '../components/AlertHistoryList';

export function AlertDetailsPage() {
  const { instance } = useParams<{ instance: string }>();

  const { history, loading, error } = useAlertHistory(instance ?? '');

  return (
    <>
      <h1>{instance}</h1>

      {loading && <div>Loading...</div>}
      {error && <div>Error loading history</div>}

      <AlertHistoryList alerts={history} />
    </>
  );
}
