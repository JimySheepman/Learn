#ifndef QUEUE_H
#define QUEUE_H

#include <iostream>
#include <cassert>
#include "commons.hpp"

template <class Type>
class Queue
{
public:
	const Queue<Type>& operator=(const Queue<Type>& otherStack);
	bool IsEmptyStack() const;
	void Push(const Type& newItem);
	Type Top() const;
	Type Pop();
	void InitializeStack();
	Queue();
	Queue(const Queue<Type>& otherStack);
	~Queue();

	friend std::ostream& operator<<(std::ostream& os, const Queue& stack) {
		node_t<Type> *temp = stack.top;

		while (temp != NULL) {
			os << temp->info << std::endl;
			temp = temp->link;
		}

		return os;
	}
private:
	node_t<Type> *top;
	void CopyStack(const Queue<Type>& otherStack);
};

template <class Type>
Queue<Type>::Queue()
{
	top = NULL;
}

template <class Type>
Queue<Type>::Queue(const Queue<Type>& otherStack)
{
	top = NULL;
	CopyStack(otherStack);
}

template <class Type>
Queue<Type>::~Queue()
{
	InitializeStack();
}

template <class Type>
const Queue<Type>& Queue<Type>::operator=(const Queue<Type>& otherStack)
{
	if (this != &otherStack)
		CopyStack(otherStack);
	return *this;
}

template <class Type>
void Queue<Type>::InitializeStack()
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
bool Queue<Type>::IsEmptyStack() const
{
	return (top == NULL);
}

template <class Type>
void Queue<Type>::Push(const Type& newItem)
{
	node_t<Type> *newNode = new node_t<Type>;

	newNode->info = newItem;
	newNode->link = top;
	top = newNode;
}

template <class Type>
Type Queue<Type>::Top() const
{
	assert(top != NULL);
	return top->info;
}

template <class Type>
Type Queue<Type>::Pop()
{
	node_t<Type> *temp, *temp2;

	if (top != NULL)
	{
		temp = top;
		while (temp->link != NULL) {
			temp2 = temp;
			temp = temp->link;
		}

		temp2->link = NULL;
		Type item = temp->info;
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
void Queue<Type>::CopyStack(const Queue<Type>& otherStack)
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
