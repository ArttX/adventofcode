const fs = require('fs');
const file = fs.readFileSync('./day_4.txt');

const numbers = file.toString().split('\n')[0].trim().split(',');
const boards = file.toString().trim().split('\r');

console.log(boards);
