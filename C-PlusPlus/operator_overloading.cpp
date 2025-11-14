#include <iostream>
using namespace std;

class Point
{
private:
    int x, y;

public:
    Point(int x = 0, int y = 0) { assingPoint(x, y); }
    void assingPoint(int x, int y);
    friend Point operator+(const Point &, const Point &);
    friend Point operator+(const int i, const Point &);
    void show();
};

void Point::assingPoint(int i, int j)
{
    x = i;
    y = j;
}

Point operator+(const Point &n1, const Point &n2)
{
    Point point;
    point.x = n1.x + n2.x;
    point.y = n1.y + n2.y;
    return point;
}

Point operator+(const int i, const Point & n){
    int x= n.x+i;
    int y=n.y+i;
    return Point(x,y);
}

void Point::show(){
    cout<<"("<<x<<","<<y<<")"<<endl;
}

int main(){
    Point n1(1,1), n2(4,4),n3;
    n1.show();
    n2.show();
    n3=n1+n2;
    n3.show();
    n2 = 13+n2;
    n2.show();
    n3.show();
    return 0;
}