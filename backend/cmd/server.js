const express = require('express');
const swaggerJsDoc = require('swagger-jsdoc');
const swaggerUi = require('swagger-ui-express');
const cors = require('cors'); // Import CORS middleware

const app = express();

// Enable CORS
app.use(cors());

// Swagger configuration for API documentation
const swaggerOptions = {
    swaggerDefinition: {
        openapi: '3.0.0',
        info: {
            title: 'API Documentation',
            version: '1.0.0',
            description: 'API documentation for the backend',
        },
        servers: [
            {
                url: 'http://localhost:8080',
            },
        ],
    },
    apis: ['./docs/**/*.js'], // Corrected path to match the actual location of documentation files
};

const swaggerDocs = swaggerJsDoc(swaggerOptions);
app.use('/swagger', swaggerUi.serve, swaggerUi.setup(swaggerDocs)); // Updated endpoint to /swagger

app.listen(8080, () => {
    console.log('Server is running on http://localhost:8080');
    console.log('Swagger docs available at http://localhost:8080/swagger/index.html'); // Updated log message
});
