import { Routes, Route } from 'react-router';
import { AlertDashboardPage } from './features/alerts/pages/AlertDashboardPage';
import { AlertDetailsPage } from './features/alerts/pages/AlertDetailsPage';

export function AppRouter() {
  return (
    <Routes>
      <Route path="/" element={<AlertDashboardPage />} />

      <Route path="/:instance" element={<AlertDetailsPage />} />
    </Routes>
  );
}
