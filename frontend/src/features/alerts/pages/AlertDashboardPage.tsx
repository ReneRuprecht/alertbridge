import { useCurrentAlerts } from '../hooks/useCurrentAlerts';
import { AlertList } from '../components/AlertList';
import { useNavigate } from 'react-router';

export function AlertDashboardPage() {
  const { alerts, loading, error } = useCurrentAlerts();
  const navigate = useNavigate();

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error.message}</div>;
  }

  return <AlertList alerts={alerts} onSelect={(alert) => navigate(`/${alert.instance}`)} />;
}
