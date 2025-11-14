/*
 * AStack.h
 *
 *  Created on: Nov 11, 2020
 *      Author: Ziya Karakaya
 *      Stack class implemented using Array notation
 */

#ifndef STACK_H_
#define STACK_H_
using namespace std;
#include <cassert>
#include <iostream>

template<class T>
class Stack {
protected:
	T *arr; //array hold elements
	int top; //index of first available place in array
	int sz; //size of array
public:
	Stack(int size = 100);
	Stack(const Stack<T>&);
	virtual ~Stack();

	bool isEmpty();
	bool isFull();
	int size();
	void clear();
	void push(const T&);
	T pop();	// retrieve top element and remove it from stack
	T peek(); //retrieve top element. You can name it as getTop
};

template<class T>
Stack<T>::Stack(int stackSize) {
	sz = stackSize;
	arr = new T[sz];
	top = 0;
}
template<class T>
Stack<T>::~Stack() {
	delete[] arr;
}

template<class T>
bool Stack<T>::isEmpty() {
	return (top == 0);
}

template<class T>
bool Stack<T>::isFull() {
	return (top == sz);
}

template<class T>
int Stack<T>::size() {
	return top;
}

template<class T>
void Stack<T>::clear() {
	top = 0;
}
template<class T>
void Stack<T>::push(const T &data) {
	if (!isFull()) {
		arr[top] = data;
		top++;
	} else {
		cout << "Stack Full" << std::endl;
	}
}
template<class T>
T Stack<T>::pop() {
	assert(!isEmpty());
// If function evaluates to false, abort (FOR DEBUGGING)
	top--;
	return arr[top];
}

template<class T>
T Stack<T>::peek() {
	assert(!isEmpty());
// If function evaluates to false, abort (FOR DEBUGGING)
	return arr[top - 1];
}

template<class T>
Stack<T>::Stack(const Stack<T> &st) {
	delete[] arr;
	sz = st.sz;
	top = st.top;
	arr = new T[sz];
	for (int n = 0; n < top; n++) {
		arr[n] = st.arr[n];
	}
}




#endif /* STACK_H_ */
