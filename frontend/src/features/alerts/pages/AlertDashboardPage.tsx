import { useCurrentAlerts } from '../hooks/useCurrentAlerts';
import { AlertList } from '../components/AlertList';

export function AlertDashboardPage() {
  const { alerts, loading, error } = useCurrentAlerts();

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error.message}</div>;
  }

  return <AlertList alerts={alerts} />;
}
