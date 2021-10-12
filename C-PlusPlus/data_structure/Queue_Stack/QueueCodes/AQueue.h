#ifndef QUEUE_H_
#define QUEUE_H_

#include <iostream>
#include <cassert>

namespace std{

template <class T>
class Queue
{
	int maxQueueSize; //Hold the maximum size of the queue
	int count; //Hold current number of items in the queue
	int queueFront, queueRear; //Hold the front and rear indices on the array
	T *data;  //Pointer to create the dynamic array that will hold all elements
public:
	
	//Constructor: Sets the initial values of the variables, size 100 by default
	Queue(int = 100);
	
	//Destructor: Free the entire array from the memory
	~Queue();
	
	//Checks whether the queue is empty or not
	bool isEmpty();
	
	//Checks whether the queue is full or not
	bool isFull();
	
	//Sets the queue to its initial state
	void clearQueue(); // Reset Queue
	
	//Returns the element that is in the very front of the queue
	T front();
	
	//Returns the element that is in the end/rear of the queue
	T rear();
	
	//Inserts an element to the end of the queue if it is not full
	void add(T&); //insert, push, enqueue, add
	
	//Deletes an element from the begining/front of the queue and returns its value
	T remove(); //Delete, serve, remove,dequeue, pop
	
	//Returns the number of element in the queue
	int size();
	
};

template <class T>
Queue<T>::Queue(int size)
{
	maxQueueSize = size; //set  max to be the size entered by the user, or 100 by default
	data = new T[maxQueueSize]; //Create an array of T values of size max
	count = 0;
	queueFront = 0;
	queueRear = maxQueueSize-1; //Not equal to zero, after the update operation it will be zero
}


template <class T>
Queue<T>::~Queue()
{
	delete [] data;
}

template <class T>
void Queue<T>::clearQueue()
{
	queueFront=count=0;
	queueRear = maxQueueSize-1;	
}

template <class T>
T Queue<T>::front()
{
	assert(!isEmpty());
	return data[queueFront];	
}      

template <class T>
T Queue<T>::rear()
{
	assert(!isEmpty());
	return data[queueRear];	
}

template <class T>
void Queue<T>::add(T& item)
{
	if(!isFull())
	{
		queueRear = ( queueRear + 1 ) % maxQueueSize;
		data[queueRear] = item;
		count++;
	}
	else
		cerr<<"No space left in queue!"<<endl;
}

template <class T>
T Queue<T>::remove()
{
	assert(!isEmpty());
	
	T temp;
	temp = data[queueFront];
	queueFront = ( queueFront + 1 ) % maxQueueSize;
	count--;
	return temp;
}
   
//Checks whether the queue is empty or not
template <class T>
bool Queue<T>::isEmpty()
{
	return count==0;
}

//Checks whether the queue is full or not
template <class T>
bool Queue<T>::isFull()
{
	return count==maxQueueSize;
}

//Checks whether the queue is full or not
template <class T>
int Queue<T>::size()
{
	return count;
}

}
#endif
