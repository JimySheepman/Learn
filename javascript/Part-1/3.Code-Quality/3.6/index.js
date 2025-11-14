// Transpilers
height = height ?? 100;

height = height !== undefined && height !== null ? height : 100;

// Polyfills

if (!Math.trunc) {
  Math.trunc = function (number) {
    return number < 0 ? Math.ceil(number) : Math.floor(number);
  };
}
