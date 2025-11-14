#include <iostream>
using namespace std;

class Number{
    int n;
    public:
        Number( int _n=0):n(_n){}
        Number& add(int);
        void print();
};

Number& Number::add(int x){
    n+=x;
    return *this;
}
void Number::print(){
    cout<<"n= "<<n<<endl;
}
int main()
{
    Number s(1);
    s.add(2).add(3).add(12).print();
    return 0;
}
