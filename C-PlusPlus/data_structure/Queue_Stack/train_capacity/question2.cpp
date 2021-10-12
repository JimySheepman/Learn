#include <iostream>
#include <fstream>
#include "Queue.hpp"
#include "Stack.hpp"

using namespace std;

class Passenger
{
	public:
		string name;
		string surname;
		string destination;
		
		friend ostream &operator<<(ostream &os, Passenger &p)
		{
			os << p.name << " " << p.surname << " " << p.destination << endl;
			return os;
		}
};

int main(void)
{	
	Queue<Passenger> queue;
	int capacityOfAnkara, capacityOfIstanbul, capacityOfAdana;
	int countAnkara = 0, countIstanbul = 0, countAdana = 0;
	Passenger passenger;
	ifstream File;
	
	//Enter a capacity <Start>
	cout << "Enter capacity of Ankara train: ";
	cin >> capacityOfAnkara;
	
	cout << "Enter capacity of Istanbul train: ";
	cin >> capacityOfIstanbul;
	
	cout << "Enter capacity of Adana train: ";
	cin >> capacityOfAdana;
	//Enter a capacity <End>
	
	//Train Stack <Start>
	Stack<Passenger> istanbulStack(capacityOfIstanbul);
	Stack<Passenger> ankaraStack(capacityOfAnkara);
	Stack<Passenger> adanaStack(capacityOfAdana);
	//Train Stack <End>
	
	//Reading File <start>
	File.open("passengers.txt");
	
	for(int i = 0; i < 20; i++)
	{
		if(File.is_open())
		{
			File >> passenger.name;
			File >> passenger.surname;
			File >> passenger.destination;
			
			queue.insertQueue(passenger);
		}
	}
	
	File.close();
	//Reading File <End>
	
	//Save all passengers in stack <Start>
	while(!queue.isEmpty())
	{
		passenger = queue.deleteQueue();
		
		if(passenger.destination == "Ankara")
		{
			if(countAnkara < capacityOfAnkara)
			{
				ankaraStack.push(passenger);
				countAnkara++;
			}
			else
			{
				cout << "Ankara train is full!" << endl;
			}
		}
		
		else if(passenger.destination == "Istanbul")
		{
			if(countIstanbul < capacityOfIstanbul)
			{
				istanbulStack.push(passenger);
				countIstanbul++;
			}
			else
			{
				cout << "Istanbul train is full!" << endl;
			}
		}
		
		else if(passenger.destination == "Adana")
		{
			if(countAdana < capacityOfAdana)
			{
				adanaStack.push(passenger);
				countAdana++;
			}
			else
			{
				cout << "Adana train is full!" << endl;
			}
		}
	}
	//Save all passengers in stack <End>
	
	//Save all passengers in file <Start>
	ofstream ankaraFile("AnkaraTrain.txt");
	cout << "--Ankara Train--" << endl;
	while(!ankaraStack.isEmpty())
	{
		passenger = ankaraStack.pop();
		
		cout << passenger;
		ankaraFile << passenger;	
	}
	ankaraFile.close();
	cout << endl;
	
	ofstream istanbulFile("IstanbulTrain.txt");
	cout << "--Istanbul Train--" << endl;
	while(!istanbulStack.isEmpty())
	{
		passenger = istanbulStack.pop();
		
		cout << passenger;
		istanbulFile << passenger;
	}
	istanbulFile.close();
	cout << endl;
	
	ofstream adanaFile("AdanaTrain.txt");
	cout << "--Adana Train--" << endl;
	while(!adanaStack.isEmpty())
	{
		passenger = adanaStack.pop();
		
		cout << passenger;
		adanaFile << passenger;
	}
	adanaFile.close();
	//Save all passenger in file<End>

	return 0;
}
