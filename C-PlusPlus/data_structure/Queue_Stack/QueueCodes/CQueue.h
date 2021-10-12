/*
 * IQueue.h
 *
 *  Created on: Nov 25, 2020
 *  Author: Ziya Karakaya
 *  Version: 0.1
 *  Description:
 *  	Queue class implemented by composition technique. This class is for educational purpose only.
 *  	Remember that there are two ways of code reuse: inheritance and composition.
 *  	When you use composition, there wouldn't be parent-child relationship between two classes,
 * 		but if you use inheritance, there would be.
 */

#ifndef QUEUE_H_
#define QUEUE_H_

#include <iostream>
#include <cassert>
#include "LinkedList.h"

namespace std{

template <typename T>
class Queue{
protected:
	LinkedList<T> myList;
public:
	bool isEmpty() const{
		return myList.isEmpty();
	}
	bool isFull() const{
		return false;
	}

	void clearQueue(){
		myList.clearList();
	}

	T front() const{
		return myList.front();
	}

	T rear() const{
		return myList.back();
	}

	void add(const T &val){ //insert, push, enqueue, add
		myList.insertLast(val);
	}
	T remove(){ //Delete, serve, remove, dequeue, pop
		T item = myList.front();
		myList.deleteFirst();
		return item;
	}
	int size() const{
		return myList.length();
	}
};

}
#endif /* QUEUE_H_ */
