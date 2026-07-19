import { describe, expect, it } from "vitest";
import type { CurrentAlertItemResponse } from "../types/currentAlertsResponse";
import { AlertMapper } from "./alertMapper";

describe("alertMapper", () => {
  it("should map response to domain alert", () => {
    const response: CurrentAlertItemResponse = {
      fingerprint: "fp1",
      status: "firing",
      alert_name: "InstanceDown",
      severity: "critical",
      environment: "prod",
      instance: "node-01:9100",
      job: "node_exporter",
      starts_at: "2026-01-01T00:00:00Z",
      last_updated_at: "2026-01-01T00:05:00Z",
    };

    const alert = AlertMapper.toAlert(response);

    expect(alert).toEqual({
      fingerprint: "fp1",
      status: "firing",
      alertName: "InstanceDown",
      severity: "critical",
      environment: "prod",
      instance: "node-01:9100",
      job: "node_exporter",
      startsAt: "2026-01-01T00:00:00Z",
      lastUpdatedAt: "2026-01-01T00:05:00Z",
    });
  });
});
