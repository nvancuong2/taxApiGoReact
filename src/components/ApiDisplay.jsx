import React, { useState, useEffect } from 'react';

const ApiDisplay = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        // Replace 'http://localhost:5000/api' with your backend API URL
        fetch('http://localhost:5000/api')
            .then((response) => response.json())
            .then((data) => {
                setData(data);
                setLoading(false);
            })
            .catch((error) => {
                console.error('Error fetching data:', error);
                setLoading(false);
            });
    }, []);

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h1>API Data</h1>
            <ul>
                {data.map((item, index) => (
                    <li key={index}>{JSON.stringify(item)}</li>
                ))}
            </ul>
        </div>
    );
};

export default ApiDisplay;
