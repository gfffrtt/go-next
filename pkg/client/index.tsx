export const getTargets = (id: string) => {
  const targets = document.querySelectorAll(`[data-id="${id}"]`);
  if (targets.length === 0) throw new Error(`Targets ${id} not found`);
  return targets;
};

export const getTarget = (id: string) => {
  const target = document.querySelector(`[data-id="${id}"]`);
  if (!target) throw new Error(`Target element ${id} not found`);
  return target;
};

export const getJson = <
  T extends Record<string, unknown> = Record<string, unknown>
>(
  id: string
): T => {
  const data = document.querySelector(`[data-client-component-id="${id}"]`);
  if (!data || !data.textContent)
    throw new Error(`Data element ${id} not found`);
  return JSON.parse(data.textContent);
};
