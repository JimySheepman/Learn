#include <iostream>
#include <cstring>
#include <cstdlib>
#include <string>
#include <sstream>
#include <algorithm>
#include <vector>
#include <fstream>
#include "Stack.hpp"
#include "Queue.hpp"

using namespace std;

vector<string>* split_str(string str) 
{
    vector<string> *string_list = new vector<string>;
    // Used to split string around spaces. 
    istringstream ss(str); 
  
    // Traverse through all words 
    do { 
        // Read a word 
        string word; 
        ss >> word;

        string_list->push_back(word);
  
        // While there is more to read 
    } while (ss); 

    return string_list;
} 

int perform_operation(char opt, int op1, int op2) {
    switch (opt) {
    case '+':
        return op1 + op2;
    case '-':
        return op1 - op2;
    case '*':
        return op1 * op2;
    case '/':
        return op1 / op2;
    default:
        exit(-1);
    }
}

// If precedence of opt1 higher than opt2, return true
bool cmp_precedence(char opt1, char opt2) {
    if (opt1 == '*' && opt2 == '/') {
        return true;
    }

    if (opt1 == '+' && opt2 == '-') {
        return true;
    }

    if (opt1 == opt2) {
        return true;
    }

    if ((opt1 == '*' || opt1 == '/') && (opt2 == '+' || opt2 == '-')) {
        return true;
    }

    return false;
}

void join(const vector<string>& v, char c, string& s) {

   s.clear();

   for (vector<string>::const_iterator p = v.begin();
        p != v.end(); ++p) {
      s += *p;
      if (p != v.end() - 1)
        s += c;
   }
}

char* infix2postfix(const char* infix_str) {
    Stack<char> operator_stack;
    char *postfix_str = new char[strlen(infix_str)] { 0 };

    bool skip = false;
    char tmp[10];
    for (int i = 0, number = 0; i < strlen(infix_str); i++) {
        number = 0;

        if (!skip && isdigit(infix_str[i])) {
            sscanf(&infix_str[i], "%d", &number);
            sprintf(tmp, "%d ", number);;
            strcat(postfix_str, tmp);
            skip = true;
        } else if (infix_str[i] == ' ') {
            skip = false;
        } else if (ispunct(infix_str[i])) {
            while (!operator_stack.IsEmptyStack() && cmp_precedence(operator_stack.Top(), infix_str[i])) {
                sprintf(tmp, "%c ", operator_stack.Pop());
                strcat(postfix_str, tmp);
            }
            operator_stack.Push(infix_str[i]);
        }
    }

    while (!operator_stack.IsEmptyStack()) {
        sprintf(tmp, "%c ", operator_stack.Pop());
        strcat(postfix_str, tmp);
    }

    return postfix_str;
}

char* infix2prefix(const char* infix_str) {
    vector<string> *splitted_str = split_str(string(infix_str));
    reverse(splitted_str->begin(), splitted_str->end());

    string joined_str;
    join(*splitted_str, ' ', joined_str);

    char *reversed_postfix = infix2postfix(joined_str.c_str());
    string reversed_postfix_str(reversed_postfix);

    vector<string> *splitted_postfix_str = split_str(reversed_postfix_str);
    reverse(splitted_postfix_str->begin(), splitted_postfix_str->end());
    string *joined_postfix_str = new string;
    join(*splitted_postfix_str, ' ', *joined_postfix_str);
    
    return (char*) joined_postfix_str->c_str();
}

int evaluate_postfix(const char* postfix_op) {
    Stack<int> number_stack;

    bool skip = false;
    for (int i = 0; i < strlen(postfix_op); i++) {
        int number = 0;

        if (!skip && isdigit(postfix_op[i])) {
            sscanf(&postfix_op[i], "%d", &number);
            number_stack.Push(number);
            skip = true;
        } else if (postfix_op[i] == ' ') {
            skip = false;
        } else if (ispunct(postfix_op[i])) {
            int op2 = number_stack.Pop();
            int op1 = number_stack.Pop();

            number_stack.Push(perform_operation(postfix_op[i], op1, op2));
        }
    }

    return number_stack.Pop();
}

int evaluate_prefix(const char* prefix_op) {
    Stack<int> number_stack;
    Stack<char> operator_stack;

    bool skip = false;
    int consec_num = 0;
    for (int i = 0; i < strlen(prefix_op); i++) {
        int number = 0;

        if (!skip && isdigit(prefix_op[i])) {
            sscanf(&prefix_op[i], "%d", &number);
            number_stack.Push(number);
            skip = true;
            consec_num++;

            if (consec_num == 2) {
                consec_num = 0;

                while (number_stack.Size() >= 2) {
                    int op1 = number_stack.Pop();
                    int op2 = number_stack.Pop();
                    number_stack.Push(perform_operation(operator_stack.Pop(), op1, op2));
                }
            }
        } else if (prefix_op[i] == ' ') {
            skip = false;
        } else if (ispunct(prefix_op[i])) {
            consec_num = 0;
            operator_stack.Push(prefix_op[i]);
        }
    }

    while (!operator_stack.IsEmptyStack()) {
        int op2 = number_stack.Pop();
        int op1 = number_stack.Pop();
        number_stack.Push(perform_operation(operator_stack.Pop(), op1, op2));
    }

    return number_stack.Pop();
}

int main(int argc, char **argv) {
    char *file_name = argv[1];
    
    std::ifstream question_file(file_name);
    
    std::string line;
    while (std::getline(question_file, line)) {
        char *infix_op = (char*) line.c_str();
        char *postfix_op = infix2postfix(infix_op);
        char *prefix_op = infix2prefix(infix_op);
        int result = evaluate_postfix(postfix_op);
        
        std::cout << "Infix form: " << infix_op << std::endl;
        std::cout << "Postfix form: " << postfix_op << std::endl;
        std::cout << "Prefix form: " << prefix_op << std::endl;
        std::cout << "Result: " << result << std::endl;
        std::cout << "===============================================" << std::endl;
    }

    return 0;
}
