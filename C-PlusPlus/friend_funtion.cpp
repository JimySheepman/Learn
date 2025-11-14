#include <iostream>
using namespace std;
class Try
{
    int no;

public:
    Try(int n = 0) : no(n) {}
    friend void print(Try &);
};

void print(Try &x)
{
    cout << "number = " << x.no << endl;
}

class B;
class A {
    int x;
    public:
        A(int _x):x(_x){}
        friend class B;
};

class B{
    int y;
    public:
        void fon(A& nsn);
        int cons() const;
};

void B::fon(A& nsn){
    y=nsn.x;
    cout<<"y= "<<y<<endl;
}
int B::cons() const{
    return y*y*y;
}
int main()
{
    Try case1(7), case2(12);
    A a(10);
    B b;
    print(case1);
    print(case2);
    b.fon(a);
    cout<< "cons = "<<b.cons()<<endl;
    return 0;
}


