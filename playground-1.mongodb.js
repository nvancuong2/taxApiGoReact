// Connect to the MongoDB instance using the connection URL
const connection = connect('mongodb+srv://nvancuong:PvxBhYFBgD55r0km@cluster0.qz2fljz.mongodb.net/taxdb');

// Verify the connection
if (connection) {
    console.log("Connected to MongoDB successfully!");

    // Query the taxreturns collection to verify the connection
    const results = connection.getCollection('taxreturns').find().toArray();
    console.log("Query Results:", results);
} else {
    console.error("Failed to connect to MongoDB!");
}