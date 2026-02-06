import type { AlertState } from "../types/Alert";

export async function getCurrentAlerts(): Promise<AlertState[]> {
  const response = await fetch("http://localhost:8080/api/v1/alert-states");
  if (!response.ok) throw new Error("Failed to fetch alert states");
  return response.json();
}
