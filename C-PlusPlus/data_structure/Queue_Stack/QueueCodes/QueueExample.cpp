/*
Write a program which determines which customer are to be served in a bank.
Sample Input:
Input (A)dd, (R)emove or (Q)uit: A
Enter customer no: 323
Input (A)dd, (R)emove or (Q)uit: A
Enter customer no: 468
Input (A)dd, (R)emove or (Q)uit: R
Customer to be served is 323
Input (A)dd, (R)emove or (Q)uit: Q
Have a nice day.
*/

#include <iostream>

#include "LQueue.h"

using namespace std;

int main()
{
	Queue<int> q;
	int cno;
	char ch;
	do
	{
		cout<<"Input (A)dd, (R)emove or (Q)uit:";
		cin>>ch;

		switch(ch)
		{
			case 'A':
			case 'a': cout <<"Enter customer no:";
					  cin>>cno;
					  q.add(cno);
					  break;
			case 'R':
			case 'r':
					  if(!q.isEmpty())
					  {
						cno=q.remove();
					  	cout<<"Customer to be served is "<<cno<<endl;
					  }
					  else
					  	cerr<<"No customer to serve yet"<<endl;
					  break;
			case 'Q':
			case 'q': cout<<"Have a nice day."<<endl;
					  break;
			default : cerr<<"Invalid Input!!"<<endl;
		}

	}while( (ch!='Q') && (ch!='q') );

}

