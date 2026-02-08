import type { AlertHistory } from "../types/Alert";

export async function getHistoryAlert(instance: string): Promise<AlertHistory> {
  const base_api_url = import.meta.env.VITE_API_BASE_URL;
  const response = await fetch(base_api_url + "/api/v1/alert-history/" + instance);
  if (!response.ok) throw new Error(`Failed to fetch alert history ${instance}`);
  return response.json();
}
