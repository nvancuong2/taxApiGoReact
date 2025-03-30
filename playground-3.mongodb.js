/* global use, db */
// MongoDB Playground
// To disable this template go to Settings | MongoDB | Use Default Template For Playground.
// Make sure you are connected to enable completions and to be able to run a playground.
// Use Ctrl+Space inside a snippet or a string literal to trigger completions.
// The result of the last command run in a playground is shown on the results panel.
// By default the first 20 documents will be returned with a cursor.
// Use 'console.log()' to print to the debug output.
// For more documentation on playgrounds please refer to
// https://www.mongodb.com/docs/mongodb-vscode/playgrounds/
use('taxdb');
db.getCollection('taxreturns').insertOne({
    expenses: 5000,
    net_income: 45000,
    taxable_income: 40000,
    created_at: new Date(),
    business_number: "BN123456",
    revenue: 50000,
    tax_payable: 8000,
    refund: 2000,
    updated_at: new Date(),
    tax_year: 2023
});
