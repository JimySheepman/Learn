#include <iostream>
using namespace std;

class Ulke
{
    string baskent,plaka,dil;
    long nufus;
    int telefon;
    public:
        Ulke(){}
        Ulke(long n):nufus(n){}
        Ulke(int t):telefon(t){}
        Ulke(string b):baskent(b){}
        Ulke(string b, string p):baskent(b),plaka(p){}
        Ulke(string b, string p, string d):baskent(b),plaka(p),dil(d){}

        void print(){
            cout<<baskent<<" "<<plaka<<" "<<dil<<" "<<nufus<<" "<<telefon<<" "<<endl;
        }
};

int main()
{
    Ulke turkiye("Ankara","TR");
    Ulke almanya;
    Ulke ingiltere("Londra");
    Ulke kosova(2000000);
    Ulke ispanya("madrid","E","Ispanyolca");
    Ulke estonya(372);

    turkiye.print();
    almanya.print();
    ingiltere.print();
    kosova.print();
    ispanya.print();
    estonya.print();
    return 0;
}