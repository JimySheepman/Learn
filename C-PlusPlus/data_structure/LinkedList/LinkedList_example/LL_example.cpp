#include<iostream>
#include "LinkedList.cpp"
using namespace std;

int main(){
	
	LinkedList<int> list;
	int num,i;
	for(i=0;i<10;i++)
	{
		cout<<"Enter an integer:";
		cin>>num;
		list.insert_descending(num);
	}
	cout<<list;
	
	}
	

