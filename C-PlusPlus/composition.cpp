#include <iostream>
#include <string>
using namespace std;

class Date
{
    int day, month, year;

public:
    Date(int d, int m, int y) : day(d), month(m), year(y) {}
    void print();
};

void Date::print()
{
    cout << day << "/" << month << "/" << year << endl;
}

class Person
{
    string name;
    Date date;

public:
    Person(string, int, int, int);
    void print();
};
void Person::print()
{
    cout << name << " ";
    date.print();
}
Person::Person(string s, int d, int m, int y) : date(d, m, y)
{
    name = s;
}

int main()
{
    Person human("merlins",03,03,1970);
    human.print();
    return 0;
}
