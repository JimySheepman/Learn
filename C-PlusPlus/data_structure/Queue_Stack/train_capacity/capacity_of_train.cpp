#include<fstream>
#include<string>
#include<iostream>
#include "Stack.h"
#include"Queue.h"
using namespace std;
class Passenger
{
	public:

			string name;
			string surname;
			string destination;
		friend ostream &operator << (ostream &os,Passenger &obj)
    {
        os<<obj.name<<" "<<obj.surname<<" "<<obj.destination<<endl;
        return os;
    }
		
};
int main()
{
    Stack<Passenger> ankara;
    Stack<Passenger> istanbul;
    Stack<Passenger> adana;
	Queue<Passenger> all_passenger;
	Passenger c;
	int capa_ank,capa_ist,capa_ada;
	int countank=0;
	int countist=0;
	int countada=0;
	int i;
	cout<<"Enter capacity for Ankara train";
	cin>>capa_ank;
	cout<<"Enter capacity for Ýstanbul train";
	cin>>capa_ist;
    cout<<"Enter capacity for Adana train";
    cin>>capa_ada;
    fstream dosya;
 	dosya.open("Passenger.txt" ,ios::in);

       for(i=0 ; i<30 ; i++){
    if(dosya.is_open()){

		dosya<<c.name;
		dosya<<c.surname;
		dosya<<c.destination;
		all_passenger.insertQueue(c);
    }
		}
	dosya.close();

	 while(!all_passenger.isEmpty())
	 {
	 	c=all_passenger.deleteQueue();
	 	if(c.destination=="Ankara")
            {
	 		if(countank <= capa_ank){
	 			ankara.push(c);
	 			countank++;

			 }
			 else {
			 	cout<<"ANKARA train is filled with passengers"<<endl;
			 }

		 }
	 	else if(c.destination=="Ýstanbul")
	 	{
	 		if(countist <= capa_ist){
	 			istanbul.push(c);
	 			countist++;
			 }
		 }
			 else{
		 	cout<<"ISTANBUL train is filled with passengers"<<endl;
		 }
	 	 if(c.destination=="Adana")
	 	{
	 		if(countada <= capa_ada){
	 			adana.push(c);
	 			countada++;
			 }
		 }
		 else {
		 	cout<<"ADANA train is filled with passengers"<<endl;

		 }
	 }
		 ofstream istanbuldosya("istanbul.txt");
	 	 while(!istanbul.isEmpty())
	 	 { 
	 	 	c= istanbul.pop();
	 	 	istanbuldosya<<c;
	 	 	cout<<c;
		  }
			istanbuldosya.close();
			cout<< endl;
	  	ofstream adanadosya("adana.txt");
	  	while(!adana.isEmpty()){
	  		
	  		c= adana.pop();
	  		cout<<c;
	  		adanadosya<<c;
		  }
	 	adanadosya.close();
	 	cout<< endl;
	 	 ofstream ankaradosya("ankara.txt");
	 	 while(!ankara.isEmpty()){
	 	 	c = ankara.pop();
	 	 	cout<<c;
		  }
		ankaradosya.close();
		cout<<endl;
	return(0);	
}
