#include <iostream>
#include <string>
#include "LStack.h"
#include <cstring>
#pragma warning(disable : 4996)
using	namespace std;

int main() {
	Stack<string>  backward_stack;
	Stack<string>  forward_stack;
	char char_conditions ='a';
	string web_page_name = "";
	bool flag = true;
	while (flag) {
		cout << "Choose(N)ew page, go (B)ack one page, go (F)orward one page (Q)uit: ";
		cin >> char_conditions;
		if (char_conditions == 'N' || char_conditions == 'n') {
			if (web_page_name == "") {
				cout << "Enter new page address: ";
				cin >> web_page_name;

				char* cstr = new char[web_page_name.length() + 1];
				strcpy(cstr, web_page_name.c_str());// C4996
				char* p = strtok(cstr, ".");// C4996
				p = strtok(NULL, ".");// C4996
				web_page_name = p;
				delete[] cstr;

				cout << "You are currently visiting " + web_page_name +".\n" ;
				if (!forward_stack.isEmpty()) {
					forward_stack.clear();
					cout << "Forward Stack cleared.\n";
				}
			}
			else {
				backward_stack.push(web_page_name);
				cout << "Enter new page address: ";
				cin >> web_page_name;

				char* cstr = new char[web_page_name.length() + 1];
				strcpy(cstr, web_page_name.c_str());// C4996
				char* p = strtok(cstr, ".");// C4996
				p = strtok(NULL, ".");// C4996
				web_page_name = p;
				delete[] cstr;

				cout << "You are currently visiting " + web_page_name + ".\n";
				if (!forward_stack.isEmpty()) {
					forward_stack.clear();
					cout << "Forward Stack cleared.\n";
				}
			}

		}
		else if (char_conditions == 'B' || char_conditions == 'b') {
			if (backward_stack.isEmpty()) {
				cout << "Error! There is no page before.\n";
				cout << "You are currently visiting " + web_page_name + ".\n";
			}
			else {
				forward_stack.push(web_page_name);
				cout << "You have gone back one page.\n";
				web_page_name = backward_stack.pop();
				cout << "You are currently visiting " + web_page_name + ".\n";
			}
		}
		else if (char_conditions == 'F' || char_conditions == 'f') {
			if (forward_stack.isEmpty()) {
				cout << "Error! There is no Forward page.\n";
					cout << "You are currently visiting " + web_page_name + ".\n";
			}
			else {
				cout << "You have gone to the Forward page.\n";
				web_page_name = forward_stack.pop();
				cout << "You are currently visiting " + web_page_name + ".\n";
			}
		}
		else if (char_conditions == 'Q' || char_conditions == 'q')
		{
			flag = false;
		}
		else
		{
			cout << "You Entered the Wrong Character, Try Again!!!";
		}
	}
	cout << "Thank you for using this simple browser. Have a nice day.\n";

}
