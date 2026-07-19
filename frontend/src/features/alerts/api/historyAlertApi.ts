import type { AlertHistoryResponse } from '../types/historyAlertsResponse';

export async function getAlertHistory(instance: string): Promise<AlertHistoryResponse> {
  const response = await fetch(`/api/v1/alerts/history?instance=${instance}`);

  if (!response.ok) {
    throw new Error('Failed to load alert history');
  }

  return response.json();
}
