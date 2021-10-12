/*
 * Queue.h
 *
 *  Created on: Nov 23, 2020
 *      Author: ziya
 */

#ifndef QUEUE_H_
#define QUEUE_H_

#include <iostream>
#include <cassert>

namespace std{

template<typename T>
struct Node{
	T 			data;
	Node<T>* 	next;
};

template <class T>
class Queue{
protected:
	Node<T> *queueFront;
	Node<T> *queueRear;
	int count;

public:
	Queue();
	virtual ~Queue();
	Queue(const Queue &other);
	Queue& operator=(const Queue &other);

	bool isEmpty() const;
	bool isFull() const;
	void clearQueue();
	T front() const;
	T rear() const;
	void add(const T&); //insert, push, enqueue, add
	T remove(); //Delete, serve, remove,dequeue, pop
	int size() const;
};

template <typename T>
Queue<T>::Queue() {
	queueFront = queueRear = NULL;
	count = 0;
}

template <typename T>
Queue<T>::~Queue() {
	clearQueue();
}

template <typename T>
Queue<T>::Queue(const Queue &other){
	queueFront = queueRear = NULL;
	count = 0;

	Node<T> *p = other.queueFront;

	while (p != NULL){
		add(p->data);
		p = p->next;
	}
}

template <typename T>
Queue<T>& Queue<T>::operator=(const Queue<T> &other){
	clearQueue();
	if (other.count == 0)
		return *this;

	Node<T> *p = other.queueFront;
	while(p != NULL){
		add(p->data);
		p = p->next;
	}
	return *this;
}


template <typename T>
bool Queue<T>::isEmpty() const{
	return queueFront == NULL;
}

template <typename T>
bool Queue<T>::isFull() const{
	return false;
}


template <typename T>
void Queue<T>::clearQueue(){
	Node<T> *p;

	while (queueFront != NULL){
		p = queueFront;
		queueFront = queueFront->next;
		delete p;
	}
	queueRear = NULL;
	count = 0;
}

template <typename T>
T Queue<T>::front() const{
	assert(!isEmpty());
	return queueFront->data;
}

template <typename T>
T Queue<T>::rear() const{
	assert(!isEmpty());
	return queueRear->data;
}

template <typename T>
void Queue<T>::add(const T &val){
	Node<T> *newNode = new Node<T>;
	newNode->data = val;
	newNode->next = NULL;

	if (queueFront == NULL){ // queue is empty
		queueFront = queueRear = newNode;
	}
	else{ // queue is not empty, add to rear
		queueRear->next = newNode;
		queueRear = queueRear->next;
	}
	count++;
}

template <typename T>
T Queue<T>::remove(){
	assert(!isEmpty());
	T temp = queueFront->data;
	Node<T> *p = queueFront;
	queueFront = queueFront->next;
	delete p;
	count--;
	return temp;
}

template <typename T>
int Queue<T>::size() const{
	return count;
}

}
#endif /* QUEUE_H_ */
