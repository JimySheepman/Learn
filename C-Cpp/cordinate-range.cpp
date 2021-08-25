#include <iostream>
#include <math.h>
using namespace std;
struct cordinate
{
    int x;
    int y;
};

void read(struct cordinate &a)
{
    cout << "Enter the cordinate(x, y): ";
    cin >> a.x >> a.y;
}
float calculate(struct cordinate &a, struct cordinate &b){
    return sqrt(pow((a.x-b.x),2)+pow((a.y-b.y),2));
}
int main()
{
    struct cordinate cord1, cord2 ;
    cout<<"Enter the cordinate information:"<<endl;
    read(cord1);
    read(cord2);
    cout<<"range is = "<<calculate(cord1,cord2)<<endl;

    return 0;
}
