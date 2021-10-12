#include <iostream>
using namespace std;
// comment out whichever version you want to work with
//#include "AStack.h" // Stack implementation with array
//#include "LStack.h" // Stack implementation with Linked List
#include "IStack.h" // Stack implementation by inheriting LinkedList class


class Student {
public:
	string name;
	int grade;
	friend istream& operator>>(istream &is, Student &s) {
		is >> s.name >> s.grade;
		return is;
	}
	friend ostream& operator<<(ostream &os, Student &s) {
		os << s.name << endl;
		return os;
	}
};

int main() {
	int max=-1;
	Student s;
	Stack<Student> st;

	for (int i=1 ; i <= 5 ; i++) {
		cout<<"Enter student name and grade:";
		cin >> s;
		if (s.grade > max) {
			max=s.grade;
			st.clear();
			st.push(s);
		} else if (s.grade == max) {
			st.push(s);
		}
	}
	cout << "Top Grade="<<max<<endl;
	cout << "Students who received top grade:"<<endl;
	while(!st.isEmpty()) {
		s = st.pop();
		cout << s;
	}
	return 0;
}

// other two examples are given below. Only one main function must be uncommented.


//int main(){
//	// Test given expression for parentheses correctness
//	Stack<char> parser;
//	bool valid = true;
//	string expr;
//	char next;
//	cout << "Input an expression:";
//	//cin >> expr;
//	getline(cin, expr);
//	int len = expr.length();
//
//	for (int i=0; valid && i < len; i++){
//		next = expr[i];
//		if (next == '(' ){
//			parser.push(next);
//		}
//		else if (next == ')' ){
//			if (parser.isEmpty()){
//				valid = false;
//			}
//			else{
//				parser.pop();
//			}
//		}
//	}
//	if (valid && parser.isEmpty()){
//		cout << "Correct paranthesis"<<endl;
//	}
//	else{
//		cout << "Incorrect paranthesis"<<endl;
//	}
//	return 0;
//}


//int main() {
//	Stack<int> myStack;
//	int n, x;
//	cout << "How many numbers (less than 10): ";
//	cin >> n;
//	for (int i=0; i < n; i++){
//		cout << "Enter an integer:";
//		cin >> x;
//		myStack.push(x);
//	}
//
//	while (!myStack.isEmpty()){
//		cout << myStack.pop();
//	}
//
//	return 0;
//}
