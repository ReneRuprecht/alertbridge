export interface Alert {
  alertname: string
  status: string
  labels: Record<string, string>
  fingerprint: string
}

export interface AlertResponse {
  alerts: Alert[]
}

const API_BASE = 'http://localhost:8000'

export async function fetchAlerts(): Promise<Alert[]> {
  const res = await fetch(`${API_BASE}/alerts`)

  if (!res.ok) {
    throw new Error(`Failed to fetch alerts: ${res.status}`)
  }

  const data: AlertResponse = await res.json()
  return data.alerts || []
}
