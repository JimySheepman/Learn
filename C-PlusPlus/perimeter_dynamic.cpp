#include <iostream>
using namespace std;

class Square{
    int side;
    public:
        void print();
        int perimeter();
};

void Square::print(){
    cout<<"Enter the side: ";
    cin>>side;
}

int Square::perimeter(){
    return 4*side;
}


int main()
{
    Square* p=new Square;
    p -> print();
    cout<<"Perimeter is "<<p->perimeter() <<" cm."<<endl;
    delete p;
    return 0;
}
