export async function calculateTax(income) {
    try {
        const response = await fetch('/api/calculate-tax', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ income }),
        });
        if (!response.ok) {
            throw new Error('Failed to calculate tax');
        }
        return await response.json();
    } catch (error) {
        console.error(error);
        throw error;
    }
}
