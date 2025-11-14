#include <iostream>
#include <limits>
#include <string>

using namespace std;

int main()
{
    // comments
    /* comments */

    const float PI = 3.14;

    cout << "Hello World!" << endl;
    cout << "Here is some text.\n";
    cout << "Here \tis some\\ text.\n";

    int myNum = 15;
    cout << myNum << endl;

    int myNum2 = 15;
    myNum2 = 25;
    cout << myNum2 << endl;

    int x = 5, y = 6, z = 50;
    cout << x + y + z << endl;

    int a, b, c;
    a = b = c = 50;
    cout << a + b + c << endl;

    cout << "Type a number: ";
    cin >> x;
    cout << "Your number is: " << x << endl;

    float myFloatNum = 5.99;
    double myDoubleNum = 9.98;
    char myLetter = 'D';
    bool myBoolean = true;
    string myText = "Hello";

    cout << "Addition: " << x + x << endl;
    cout << "Subtraction: " << x - x << endl;
    cout << "Multiplication: " << x * x << endl;
    cout << "Division: " << x / x << endl;
    cout << "Modulus: " << x % x << endl;
    cout << "Increment: " << ++x << endl;

    x += 3;
    cout << "x+=3: " << x << endl;
    x -= 3;
    cout << "x-=3: " << x << endl;
    x *= 3;
    cout << "x*=3: " << x << endl;
    x /= 3;
    cout << "x/=3: " << x << endl;
    x %= 3;
    cout << "x%=3: " << x << endl;
    x &= 3;
    cout << "x&=3: " << x << endl;
    x |= 3;
    cout << "x|=3: " << x << endl;
    x ^= 3;
    cout << "x^=3: " << x << endl;
    x >>= 3;
    cout << "x>>=3: " << x << endl;
    x <<= 3;
    cout << "x<<=3: " << x << endl;

    return 0;
}