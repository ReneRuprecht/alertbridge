import { describe, expect, it } from "vitest";
import { formatAlertInstance, formatAlertTime } from "../../src/utils/Formatter";

describe("formatAlertTime", () => {
  it("returns empty string for invalid input", () => {
    expect(formatAlertTime(undefined)).toBe("");
    expect(formatAlertTime("")).toBe("");
  });

  it("formats date in 24h format", () => {
    expect(formatAlertTime("2026-02-08T10:00:00Z")).toContain("10:00");
  });

  it("formats date in dd.MM.yyy HH:mm:ss", () => {
    expect(formatAlertTime("2026-02-08T10:00:00Z")).toBe("8.2.2026 10:00:00");
  });
});


describe('formatAlertInstance', () => {
  it("returns empty string for invalid input", () => {
    expect(formatAlertInstance(undefined)).toBe("");
    expect(formatAlertInstance("")).toBe("");
  });

  it("returns string without port", () => {
    expect(formatAlertInstance("web_exporter:9000")).toBe("web_exporter");
  });

})