#include <iostream>
#include <string>
#include <cctype>
#include "LStack.h" // linked list based stack

using namespace std;

// returns the precedence of operator.. higher value for higher precedence
// if you need, you can extend with additional operators
int prec(char ch) {
	if (ch == '*' || ch == '/')
		return 2;
	else if (ch == '+' || ch == '-')
		return 1;
	else
		return -1;
}

// this version does not accepts parentheses within expression
void infixToPostfix(string expr)
{
	Stack<char> myStack;
	char ch;
	int sz = expr.length();

	for(int i=0;i < sz;i++){
		ch = expr[i];

		if(isalnum(ch)){
			cout<<ch;
		}
		else {
			while(!myStack.isEmpty() && prec(ch) <= prec(myStack.peek())){
				cout << myStack.pop();
			}
			myStack.push(ch);
		}
	}
	while (!myStack.isEmpty()){
		cout << myStack.pop();
	}
	cout << endl;
}

// this version extended to handle parentheses within expression
void infixToPostfixExtended(string expr)
{
	Stack<char> myStack;
	char ch;
	int sz = expr.length();

	for(int i=0;i < sz;i++)
	{
		ch = expr[i];

		if (isalnum(ch)) {
			cout << ch;
		}
		else if (ch == '(') {
			myStack.push(ch);
		}
		else if (ch == ')') {
			while (!myStack.isEmpty() && myStack.peek() != '(') {
				cout << myStack.pop();
			}
			if (myStack.peek() == '(') {
				myStack.pop();
			}
		}
		else {
			while(!myStack.isEmpty() && prec(ch) <= prec(myStack.peek())){
				cout << myStack.pop();
			}
			myStack.push(ch);
		}
	}
	while (!myStack.isEmpty()){
		cout << myStack.pop();
	}
	cout << endl;
}

//int main(){
//	string exp = "A+B*C+D";
//	cout << "Infix="<<exp<<", Postfix=";
//	infixToPostfix(exp);
//	cout << "Infix="<<exp<<", Postfix=";
//	infixToPostfixExtended(exp);
//
//	exp = "(A+B)*(C+D)";
//	cout << "Infix="<<exp<<", Postfix=";
//	infixToPostfix(exp);
//	cout << "Infix="<<exp<<", Postfix=";
//	infixToPostfixExtended(exp);
//
//	exp = "A*B+C*D";
//	cout << "Infix="<<exp<<", Postfix=";
//	infixToPostfix(exp);
//	cout << "Infix="<<exp<<", Postfix=";
//	infixToPostfixExtended(exp);
//
//	exp = "A+B+C+D";
//	cout << "Infix="<<exp<<", Postfix=";
//	infixToPostfix(exp);
//	cout << "Infix="<<exp<<", Postfix=";
//	infixToPostfixExtended(exp);
//
//	return 0;
//}





