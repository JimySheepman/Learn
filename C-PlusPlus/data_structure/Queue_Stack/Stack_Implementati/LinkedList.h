/*
 * LinkedList.h
 *
 *  	Created on: 	Oct 20, 2020
 *      Author: 		Ziya Karakaya
 *      Version: 		0.3
 *      Last Update on: Nov 17, 2020
 *      Description: 	This version is implemented only for teaching purpose.
 *      				May not be complete!
 *
 *		Update notes (v0.3):
 *      * I have added deleteFirst operation. Similarly you can add deleteLast.
 *      * I made this version more optimized by adding searchAndFindPrev function in protected part
 *        It takes 3 parameters:
 *        	* first parameter: value to be searched
 *        	* second parameter: to store its node address
 *        	* third parameter: to store its previous node address
 *        This way we search, assign node address and prev node address in one pass.
 *
 *      License:		GPLv3
 */

#ifndef LINKEDLIST_H_
#define LINKEDLIST_H_

#include <iostream>
#include <cassert>

namespace std {

template<typename T>
struct Node{
	T 			data;
	Node<T>* 	next;
};

template<typename T>
class LinkedList {
protected:
	Node<T>*   	head;
	Node<T>*   	tail;
	int 		count;
public:
	LinkedList();
	virtual ~LinkedList();
	LinkedList(const LinkedList &other);
	LinkedList& operator=(const LinkedList &other);

	bool isEmpty() const;
	int length() const;

	Node<T>* search(const T& val);
	void insertFirst(const T &val);
	void insertLast(const T &val);

	void insertAt(const int idx, const T &val);
	void insertAfter(const T &sVal, const T &val);
	void insertBefore(const T &sVal, const T &val);

	void deleteNode(const T &val);
	void deleteFirst();
	void clearList();

	T front() const;
	T back() const;

	template <typename S>
	friend ostream& operator<<(ostream& out, const LinkedList<S> &list);

	void reversePrint();

protected:
	bool searchAndFindPrev(const T &val, Node<T>** sNode, Node<T>** pNode);
	void recursiveReversePrint(Node<T> *p);

};

template <typename T>
LinkedList<T>::LinkedList() {
	head = tail = NULL;
	count = 0;
}

template <typename T>
LinkedList<T>::~LinkedList() {
	clearList();
}

template <typename T>
LinkedList<T>::LinkedList(const LinkedList &other) {
	head = tail = NULL;
	Node<T> *p = other.head;

	while (p != NULL){
		insertLast(p->data);
		p = p->next;
	}
	count = other.count;
}

template <typename T>
LinkedList<T>& LinkedList<T>::operator=(const LinkedList &other) {
	clearList();
	if (other.count == 0)
		return *this;

	Node<T> *p = other.head;
	while(p != NULL){
		insertLast(p->data);
		p = p->next;
	}
	return *this;
}


template <typename T>
bool LinkedList<T>::isEmpty() const{
	return (count == 0);
}

template <typename T>
int LinkedList<T>::length() const{
	return count;
}

template <typename T>
Node<T>* LinkedList<T>::search(const T &val){
	bool found = false;
	Node<T> *p = head;
	while ((p != NULL) && !found){
		if (p->data == val)
			found = true;
		else
			p = p->next;
	}
	return p;
}

template <typename T>
void LinkedList<T>::insertFirst(const T &val){
	Node<T> *p = new Node<T>;
	p->data= val;
	p->next = head;
	head = p;
	if (tail == NULL)
		tail = p;

	count++;
}

template <typename T>
void LinkedList<T>::insertLast(const T &val){
	Node<T> *p = new Node<T>;
	p->data = val;
	p->next = NULL;

	if (head != NULL){
		tail->next = p;
		tail = p;
	}
	else{
		head = tail = p;
	}
	count++;
}

template <typename T>
void LinkedList<T>::insertAt(const int idx, const T &val){
	if (idx > count){
//		cout << "Given index is greater than the number of element, so appending to list"<<endl;
		insertLast(val);
	}
	else if (idx <= 1){
//		cout << "Given index value is less than or equal to 1, so inserting as the first element"<<endl;
		insertFirst(val);
	}
	else{ // insert at given position
		Node<T> *prevNode = head;
		int cnt=2;
		while (cnt < idx){
			prevNode = prevNode->next;
		} // we found (idx-1)'th node

		Node<T> *newNode = new Node<T>;
		newNode->data = val;
		newNode->next = prevNode->next;
		prevNode->next = newNode;
	}
	count++;
}

template <typename T>
void LinkedList<T>::insertAfter(const T &sVal, const T &val){
	Node<T> *searchItem = search(sVal); // search for sVal in the list
	if (searchItem == NULL){
		cout << "The value "<<sVal<<" is not found in the list, so nothing will be inserted."<<endl;
	}
	else{ // we found searched item
		Node<T> *newNode = new Node<T>;
		newNode->data = val;
		newNode->next = searchItem->next;
		searchItem->next = newNode;
		count++;
	}
}

template <typename T>
void LinkedList<T>::insertBefore(const T &searchValue, const T &val){
	Node<T> *searchedNode=NULL;
	Node<T> *prevNode=NULL;
	// search for searchValue in the list, assign searchedNode and prevNode if found
	bool found = searchAndFindPrev(searchValue, &searchedNode, &prevNode);

	if (!found){
		cout << "The value "<<searchValue<<" is not found in the list, so nothing will be inserted."<<endl;
	}
	else{ // we found searched item
		Node<T> *newNode = new Node<T>;
		newNode->data = val;
		if (searchedNode == head) { // insert as first element
			newNode->next = head;
			head = newNode;
		} else {
			//Node<T> *prevNode = getPrevNodeOf(searchItem);
			newNode->next = prevNode->next;
			prevNode->next = newNode;
		}
		count++;
	}
}

template <typename T>
void LinkedList<T>::deleteNode(const T &val){
	Node<T> *delNode=NULL;
	Node<T> *prevNode=NULL;
	bool found = searchAndFindPrev(val, &delNode, &prevNode); // search for val in the list

	if (!found){ // not found in the list
		cout << "Item "<< val << " is not in the list, nothing to delete!"<<endl;
	}
	else{ // found in the list and delItem points to that node
		assert(delNode != NULL);

		// depending on the position, delete appropriately
		if (delNode == head){ // delete first node
			head = delNode->next; // or head = head->next; both are same
		}
		else {
			assert (prevNode != NULL);
			if (delNode == tail) { // if last item is to be deleted
				prevNode->next = NULL;
				tail = prevNode;
			} else {
				prevNode->next = delNode->next;
			}
		}
		delete delNode;
		count--;
	}
}

template <typename T>
void LinkedList<T>::deleteFirst(){
	if (!isEmpty()){
		Node<T> *p = head;
		head = head->next;
		delete p;
		count--;
	}
}

template <typename T>
void LinkedList<T>::clearList() {
	Node<T> *p;

	while (head != NULL){
		p = head;
		head = head->next;
		delete p;
	}
	tail = NULL;
	count = 0;
}

template <typename T>
T LinkedList<T>::front() const{
	assert(head != NULL);
	return head->data;
}

template <typename T>
T LinkedList<T>::back() const{
	assert(tail != NULL);
	return tail->data;
}

template <typename T>
ostream& operator<<(ostream& out, const LinkedList<T> &list){
	Node<T> *p = list.head;
	while(p != NULL){
		cout << p->data << " ";
		p = p->next;
	}
	return out;
}

template <typename T>
bool LinkedList<T>::searchAndFindPrev(const T &val, Node<T>** sNode, Node<T>** pNode){
	// search for val in the list, if found sNode points to searched node, and
	// pNode points to previous of that node
	// return true if found, else return false
	bool found = false;
	*sNode = head; // searched Node
	*pNode = NULL; // previous node

	while ((*sNode != NULL) && !found){
		if ((*sNode)->data == val){
			found = true;
		}
		else{
			*pNode = *sNode;
			*sNode = (*sNode)->next;
		}
	}

	return found;
}

template<class T>
void LinkedList<T>::reversePrint() {
	recursiveReversePrint(head);
}


template<class T>
void LinkedList<T>::recursiveReversePrint(Node<T> *p) {
	if (p != NULL) {
		recursiveReversePrint(p->next); //recursive call
		cout << p->data << " ";
	}
}

} /* namespace std */

#endif /* LINKEDLIST_H_ */
