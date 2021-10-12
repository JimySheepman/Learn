#include <iostream>
#include "LinkedList.hpp"
#include "Process.hpp"

using namespace std;

int main()
{
    LinkedList<Process> LL; // ll is mean linkend list
    Process process1("", 54);
    Process process2("", 24);
    Process process3("", 45);
    Process process4("", 63);
    Process process5("", 14);

    cout << "Inserting ->" << endl;
    cout << process1 << endl;
    LL.prioritized_insert(process1);
    cout << "Inserting ->" << endl;
    cout << process2 << endl;
    LL.prioritized_insert(process2);
    cout << "Inserting ->" << endl;
    cout << process3 << endl;
    LL.prioritized_insert(process3);
    cout << "Inserting ->" << endl;
    cout << process4 << endl;
    LL.prioritized_insert(process4);
    cout << "Inserting ->" << endl;
    cout << process5 << endl;
    LL.prioritized_insert(process5);

    cout << "Currently, there are " << LL.length() << " processes in the linked list."
         << endl << endl
         << "Processes in the linked list that are sorted by their priorities are as follows:"
         << endl;
    cout << LL << endl;

    return 0;
}

