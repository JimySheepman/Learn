function romanDigit(numberDigits) {
  let unitDigit = ["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"];
  let tenth = ["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"];
  let hundredsDigit = [
    "",
    "C",
    "CC",
    "CCC",
    "CD",
    "D",
    "DC",
    "DCC",
    "DCCC",
    "CM",
  ];
  let thousandsDigit = [
    "",
    "M",
    "MM",
    "MMM",
    "MMMM",
    "MMMMM",
    "MMMMMM",
    "MMMMMMM",
    "MMMMMMMM",
    "MMMMMMMMM",
  ];
  return String(
    thousandsDigit[numberDigits[4]] +
      hundredsDigit[numberDigits[3]] +
      tenth[numberDigits[2]] +
      unitDigit[numberDigits[1]]
  );
}
function convertToRoman(num) {
  let numberDigits = {
    1: 0,
    2: 0,
    3: 0,
    4: 0,
  };
  if (num > 0 && num <= 9999) {
    convertNumber = num;
    let romasNumber;
    for (let i = 1; i < 5; i++) {
      if (convertNumber != 0) {
        numberDigits[i] = convertNumber % 10;
        convertNumber = Math.floor(convertNumber / 10);
      }
    }
    return romanDigit(numberDigits);
  } else {
    return 0;
  }
}

/* console.log(convertToRoman(36))
console.log(convertToRoman(2));
console.log(convertToRoman(3));
console.log(convertToRoman(4));
console.log(convertToRoman(5));
console.log(convertToRoman(9));
console.log(convertToRoman(12));
console.log(convertToRoman(16));
console.log(convertToRoman(29));
console.log(convertToRoman(99));
console.log(convertToRoman(400));
console.log(convertToRoman(798));
console.log(convertToRoman(1000));
console.log(convertToRoman(2014));
 */