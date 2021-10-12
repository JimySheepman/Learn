#ifndef STACK_H
#define STACK_H

#include <iostream>
#include <cassert>
#include "commons.hpp"

template <class Type>
class Stack
{
public:
	const Stack<Type>& operator=(const Stack<Type>& otherStack);
	bool IsEmptyStack() const;
	void Push(const Type& newItem);
	Type Top() const;
	Type Pop();
	void InitializeStack();
	int Size();
	Stack();
	Stack(const Stack<Type>& otherStack);
	~Stack();

	friend std::ostream& operator<<(std::ostream& os, const Stack& stack) {
		node_t<Type> *temp = stack.top;

		while (temp != NULL) {
			os << temp->info << std::endl;
			temp = temp->link;
		}

		return os;
	}
private:
	node_t<Type> *top;
	int size;
	void CopyStack(const Stack<Type>& otherStack);
};

template <class Type>
Stack<Type>::Stack()
{
	top = NULL;
	size = 0;
}

template <class Type>
Stack<Type>::Stack(const Stack<Type>& otherStack)
{
	top = NULL;
	size = 0;
	CopyStack(otherStack);
}

template <class Type>
Stack<Type>::~Stack()
{
	InitializeStack();
}

template <class Type>
int Stack<Type>::Size() {
	return size;
}

template <class Type>
const Stack<Type>& Stack<Type>::operator=(const Stack<Type>& otherStack)
{
	if (this != &otherStack)
		CopyStack(otherStack);
	return *this;
}

template <class Type>
void Stack<Type>::InitializeStack()
{
	node_t<Type> *temp;

	while (top != NULL)
	{
		temp = top;
		top = top->link;
		delete temp;
	}
}

template <class Type>
bool Stack<Type>::IsEmptyStack() const
{
	return (top == NULL);
}

template <class Type>
void Stack<Type>::Push(const Type& newItem)
{
	node_t<Type> *newNode = new node_t<Type>;

	newNode->info = newItem;
	newNode->link = top;
	top = newNode;
	size++;
}

template <class Type>
Type Stack<Type>::Top() const
{
	assert(top != NULL);
	return top->info;
}

template <class Type>
Type Stack<Type>::Pop()
{
	node_t<Type> *temp;

	if (top != NULL)
	{
		temp = top;
		top = top->link;
		Type item = temp->info;
		size--;
		delete temp;
		return item;
	}
	else
	{
		std::cerr << "Cannot remove from an empty stack." << std::endl;
		assert(top == NULL);
	}
}

template <class Type>
void Stack<Type>::CopyStack(const Stack<Type>& otherStack)
{
	node_t<Type> *newNode, *current, *last;

	if (top != NULL)
		InitializeStack();

	if (otherStack.top == NULL)
		top = NULL;
	else
	{
		current = otherStack.top;
		top = new node_t<Type>;
		top->info = current->info;
		top->link = NULL;
		last = top;
		current = current->link;

		while (current != NULL)
		{
			newNode = new node_t<Type>;
			newNode->info = current->info;
			newNode->link = NULL;
			last->link = newNode;
			last = newNode;
			current = current->link;
		}
	}
}

#endif
