export function formatAlertInstance(instance: string | undefined): string {
  if (instance === undefined) return "";
  if (instance.trim() === "") return "";

  return instance.split(":")[0];
}

export function formatAlertStartsAt(startsAt: string | undefined): string {
  if (startsAt === undefined) return "";
  if (startsAt.trim() === "") return "";

  const date = new Date(startsAt).toLocaleDateString();
  const time = new Date(startsAt).toLocaleTimeString();

  return date + " " + time;
}