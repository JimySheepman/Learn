// bad comments
// this code will do this thing (...) and that thing (...)
// ...and who knows what else...
somecode;

/**
 *  Good comment
 * @param {number} n is total number
 * @return {number}  Boolean type 
 */


function showPrimes(n) {
  nextPrime: for (let i = 2; i < n; i++) {
    for (let j = 2; j < i; j++) {
      if (i % j == 0) continue nextPrime;
    }
    console.log(i);
  }
}

function showPrimes(n) {
  for (let i = 2; i < n; i++) {
    if (!isPrime(i)) continue;
    console.log(i);
  }
}
function isPrime(n) {
  for (let i = 2; i < n; i++) {
    if (n % i == 0) return false;
  }
  return true;
}
