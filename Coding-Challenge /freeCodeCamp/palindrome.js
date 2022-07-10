function palindrome(str) {
  let a = str.toLowerCase().replace(/[^0-9a-z]/gi, "");
  console.log(a);
  let flag = false;
  let splitString = a.split("");
  let convertReverseStrinngArray = splitString.reverse();
  let reverseString = convertReverseStrinngArray.join("");
  console.log(reverseString);
  for (let i = 0; i < a.length; i++) {
    if (reverseString[i] == a[i]) {
      flag = true;
    } else {
      flag = false;
      return false;
    }
  }
  return flag;
}

// Task and Test case

/* 
  palindrome("eye") should return a boolean.
  palindrome("eye") should return true.
  palindrome("_eye") should return true.
  palindrome("race car") should return true.
  palindrome("not a palindrome") should return false.
  palindrome("A man, a plan, a canal. Panama") should return true.
  palindrome("never odd or even") should return true.
  palindrome("nope") should return false.
  palindrome("almostomla") should return false.
  palindrome("My age is 0, 0 si ega ym.") should return true.
  palindrome("1 eye for of 1 eye.") should return false.
  palindrome("0_0 (: /-\ :) 0-0") should return true.
  palindrome("five|\_/|four") should return false. */
