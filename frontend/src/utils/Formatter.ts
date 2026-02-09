export function formatAlertInstance(instance: string | undefined): string {
  if (instance === undefined) return "";
  if (instance.trim() === "") return "";

  return instance.split(":")[0];
}

export function formatAlertTime(startsAt: string | undefined): string {
  if (startsAt === undefined) return "";
  if (startsAt.trim() === "") return "";

  const date = new Date(startsAt).toLocaleDateString('de-DE');
  const time = new Date(startsAt).toLocaleTimeString('de-DE', {
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: false
  });

  return date + " " + time;
}