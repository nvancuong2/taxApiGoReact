import React, { useState } from 'react';
import { calculateTax } from '../api/taxApi';

function TaxCalculator() {
    const [income, setIncome] = useState('');
    const [tax, setTax] = useState(null);

    const handleCalculate = async () => {
        try {
            const result = await calculateTax(parseFloat(income));
            setTax(result.tax);
        } catch (error) {
            alert('Error calculating tax');
        }
    };

    return (
        <div>
            <h1>Tax Calculator</h1>
            <input
                type="number"
                placeholder="Enter income"
                value={income}
                onChange={(e) => setIncome(e.target.value)}
            />
            <button onClick={handleCalculate}>Calculate Tax</button>
            {tax !== null && <p>Calculated Tax: ${tax.toFixed(2)}</p>}
        </div>
    );
}

export default TaxCalculator;
