/*
 * Notes:
 * For output operator overload to work, the user must implement the << operator for the template type.
 * Similarly, to search function to work the user must implement the == operator for the template type.
 */

#include <iostream>
#include "LinkedList.hpp"
#include "Process.hpp"

using namespace std;
int main() {
    Process processes[5] = {
        Process(NULL, 54),
        Process(NULL, 24),
        Process(NULL, 45),
        Process(NULL, 63),
        Process(NULL, 14)
    };

    LinkedList<Process> ready_list;

    for (int i = 0; i < 5; i++) {
    	cout<<"Inserting ->"<<endl<<processes[i]<<endl;
        ready_list.insert_prioritized(processes[i]);
    }
    cout<<"Currently there are "<<ready_list.length()<<" processes in the linked list."<<endl<<endl;
	cout<<"Processes in the linked list that are sorted by their priorities are as follows:"<<endl;
    cout<<ready_list;

    return 0;
}
