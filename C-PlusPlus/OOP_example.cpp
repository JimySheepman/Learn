#include <iostream>
#include <string.h>
using namespace std;
//abstract class
//base class
class Animal {
protected:
	string name, diet;
	float dailyCaloriesToBeTaken, expectedLifeTime;
public:
	// constructor part
	Animal(){}
	// constructor part
	Animal(string name, string diet, float dailyCaloriesToBeTaken)
	{
		this->name = name;
		this->diet = diet;
		this->dailyCaloriesToBeTaken = dailyCaloriesToBeTaken;
		expectedLifeTime = 0;
	}
	//calculateExpectedLife function part
	virtual void  calculateExpectedLife(float expectedLifeTime)=0;
	//display function part
	virtual void display()=0;
};
//sup class
class Mammal :public Animal {
	float brainSize;
public:
	// constructor part
	Mammal(){}
	// constructor part
	Mammal(string name, string diet, float dailyCaloriesToBeTaken, float brainSize) :Animal(name, diet, dailyCaloriesToBeTaken), brainSize(brainSize)	{}
	//calculateExpectedLife function part
	void calculateExpectedLife(float expectedLifeTime)
	{
		this->expectedLifeTime=(expectedLifeTime * 20.0) / dailyCaloriesToBeTaken;
	}
	//display function part
	void display()
	{
		cout << "Name: " << this->name << endl;
		cout << "Diet: " << this->diet << endl;
		cout << "Brain Size: " << this->brainSize << " m3" << endl;
		cout << "Expected Lifetime: " << this->expectedLifeTime << " year" << endl;
		cout << endl;
	}
};
//sup class
class Bird :public Animal
{
	float wingSize;
public:
	// constructor part
	Bird(){}
	// constructor part
	Bird(string name, string diet, float dailyCaloriesToBeTaken, float wingSize) :Animal(name, diet, dailyCaloriesToBeTaken), wingSize(wingSize) {}
	//calculateExpectedLife function part
	void calculateExpectedLife(float expectedLifeTime)
	{
		this->expectedLifeTime = (expectedLifeTime * 10.0) / dailyCaloriesToBeTaken;
	}
	//display function part
	void display()
	{
		cout << "Name: " << this->name <<endl;
		cout << "Diet: " << this->diet << endl;
		cout << "Wing Size: " << this->wingSize << "m" << endl;
		cout << "Expected Lifetime: " << this->expectedLifeTime << " year" << endl;
		cout << endl;
	}
};
//main part
int main()
{
	Animal *arr[10];//object array
	int i = 0;
	int flag;
	int bool_flag;
	char ch='a';
	string sname, sdiet;
	float brain_size, calorieshoultake, caloriestake, wing_size;
	while (ch != 'e')//menu loop
	{
		cout << "What do you want to do?(1 for add an animal, 2 for remove last added animal,3 display all): ";
		cin >> flag;
		if (flag == 1)//menu tab 1
		{
			cout << "Which species do you want to add ? (1 for mammals, 2 for birds) ";
			cin >> bool_flag;
			if (bool_flag == 1)//mammal check
			{
				cout << "Enter its name: ";
				cin >> sname;
				cout << "Enter its diet: ";
				cin >> sdiet;
				cout << "Enter its brain size: ";
				cin >> brain_size;
				cout << "Enter the daily calories it should take: ";
				cin >> calorieshoultake;
				cout << "Enter the daily calories it takes: ";
				cin >> caloriestake;
				Mammal *mammal=new Mammal(sname, sdiet, calorieshoultake, brain_size);//Mammal object
				arr[i] = mammal;
				arr[i]->calculateExpectedLife(caloriestake);
				i++;
			}
			else if (bool_flag == 2)//bird check
			{
				cout << "Enter its name: ";
				cin >> sname;
				cout << "Enter its diet: ";
				cin >> sdiet;
				cout << "Enter its brain size: ";
				cin >> wing_size;
				cout << "Enter the daily calories it should take: ";
				cin >> calorieshoultake;
				cout << "Enter the daily calories it takes: ";
				cin >> caloriestake;
				Bird* bird = new Bird(sname, sdiet, calorieshoultake, wing_size);//Bird object
				arr[i] = bird;
				arr[i]->calculateExpectedLife(caloriestake);
				i++;
			}
		}
		else if (flag == 2) //menu tab 2 for remove
		{
			cout << "Removed Animal:" << endl;
			arr[i - 1]->display();
			i--;
		}
		else if (flag == 3) //menu tab 3 for all display
		{
			for (int k = 0; k < i; k++)//display loop
			{
				arr[k]->display();
			}
		}
		cout << "Enter e for exit:";//exit 
		cin >> ch;
	} 
	return 0;
}
