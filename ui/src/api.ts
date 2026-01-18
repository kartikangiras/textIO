export async function sendRequest(endpoint: string, payload: unknown) {
  try {
    const res = await fetch(endpoint, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    const data = await res.json();
    
    if (!res.ok) {
      throw new Error(data.error || 'Something went wrong');
    }
    
    return data.result;
  } catch (err: unknown) {
    console.error(err);
    if (err instanceof Error) {
      return `Error: ${err.message}`;
    }
    return 'Error: 
  }
}