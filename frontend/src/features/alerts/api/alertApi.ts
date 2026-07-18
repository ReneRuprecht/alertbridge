import type { CurrentAlertsResponse } from "../types/currentAlertsResponse";

export async function getCurrentAlerts(): Promise<CurrentAlertsResponse> {
  const response = await fetch('/api/v1/alerts/current');

  if (!response.ok) {
    throw new Error('Failed to load current alerts');
  }

  return response.json();
}
