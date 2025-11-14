#include <iostream>
using namespace std;

class One{
    public:
        static int number;
        One(){
            number++;
        }

};
int One::number=0;
int main()
{
    One n1,n2,n3;
    cout<<"Count of object: "<<n1.number<<endl;
    return 0;
}
