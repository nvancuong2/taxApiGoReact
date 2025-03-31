import axios from 'axios';

const API_URL = 'http://localhost:8080/api'; // Adjust the URL as needed

export const submitTaxForm = async (companyData) => {
    try {
        const response = await axios.post(`${API_URL}/calculate-tax`, companyData);
        return response.data;
    } catch (error) {
        console.error('Error submitting tax form:', error);
        throw error;
    }
};