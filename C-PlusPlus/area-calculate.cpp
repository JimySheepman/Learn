#include <iostream>

using namespace std;

int main()
{
    int a;
    int b;
    int r;
    double result;
    char ansewrs = '0';
    do
    {
        cout << "choose your ans:" << endl;
        cout << "1 : rectangle  area" << endl;
        cout << "2 : square  area" << endl;
        cout << "3 : circle  area" << endl;
        cout << "4 : exit " << endl;
        cin >> ansewrs;

        if (ansewrs == '1')
        {
            cout << "enter the a and b: ";
            cin >> a >> b;
            result = a * b;
            cout << "Result area: " << result << endl;
        }
        else if (ansewrs == '2')
        {
            cout << "enter the a: ";
            cin >> a;
            result = a * a;
            cout << "Result area: " << result << endl;
        }
        else if (ansewrs == '3')
        {
            cout << "enter the r: ";
            cin >> r;
            result = 2 * r * 3.14;
            cout << "Result area: " << result << endl;
        }

    } while (ansewrs != '4');

    return 0;
}