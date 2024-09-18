import '../scss/styles.scss'; // Import SCSS file

// Import a function from another module
import { someFunction } from './module';

// Use the imported function
someFunction();

// ES6 example
const name = 'Webpack';
const greet = (name) => `Hello, ${name}!`;

console.log(greet(name));
