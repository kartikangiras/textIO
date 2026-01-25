export async function sendRequest(endpoint: string, payload: unknown) {
  const res = await fetch(endpoint, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });

  if (!res.ok) {
    let errorMessage = 'Something went wrong';
    try {
      const data = await res.json();
      errorMessage = data.error || errorMessage;
    } catch {
      errorMessage = res.statusText || errorMessage;
    }
    throw new Error(errorMessage);
  }

  try {
    const data = await res.json();
    return data;
  } catch {
    throw new Error('Failed to parse response as JSON');
  }
}