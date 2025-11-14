#include <iostream>
#include <string>
#include <fstream>
#include "LinkedStack.h"
#include "LinkedQueue.h"
using namespace std;

int check(int x, int y, char ex){
	switch(ex)
	{
		case'+': return x+y;
		case'-': return x-y;
		case'/': return x/y;
		case'*': return x*y;	
	}
}
	int proporty(cahr k){
	switch(ex)
	{
		case'+': return 1;
		case'-': return 1;
		case'/': return 2;
		case'*': return 2;
		case'\n':return 3;	
		default : return 0;
	}
		return;
	}
    stirng infix_to_prefix(string a){
    	LinkedStack<char> prefix;
    	LinkedStack<char> brackets;

    	char c;
    	for(int i=a.lenght()-1;i>=0;i--){
    		c=a[i];
    		if(proporty(c)>=1){
    		if else(proporty(c==1)	
    			if else(proporty(c==2)	
    				if else(proporty(c==3)	
			}
			else{
				prefix.pop(c);
			}
		}
    	
	}
	int calculater(){
	}
}
 int main(){
 	string readline;
 	string line;
 	ifstream questions("questions.txt");
 	if(questions.is_open()){
	 
 	do{
 		getline(questions,readline);
 		cout<<"Infix From:"<<readline<<endl;
 		line=fonkisyona(readline);
 		cout<<"Prefix Form:"<<infix_to_prefix(line)<<endl;
 		cout<<"Result:"<<calculate(line)<<endl;
 		cout<<"----------------------------------------"<<endl;
 		cout<<"----------------------------------------"<<endl;
	 }while(!questions.eof());
}
else cout<<"Not founded text file :D lol";
	 questions.close();
	 return 0;
 }
