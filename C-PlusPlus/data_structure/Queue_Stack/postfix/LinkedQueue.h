#include <iostream>
#include <assert.h>
using namespace std;
template <typename T>
struct Node{
	T data;
	Node<T> *next;
};
template <typename T>
class LinkedQueue{
		Node<T> *front,*rear;
	public:
		LinkedQueue(){
			front=rear=NULL;
		}
		bool isEmpty(){
			return (front==NULL);
		}
		void destroyQueue();
		~LinkedQueue(){
			destroyQueue();
		}
		T qfront(){
			assert(!isEmpty());
			return front->data;
		}
		T qRear(){
			assert(!isEmpty());
			return rear->data;
		}
		void push(T& item);
		T pop();
};
template <typename T>
void LinkedQueue<T>::push(T& item){
	Node<T> *p=new Node<T>;
	p->data=item;
	p->next=NULL;
	if(front==NULL){
		front=rear=p;
	}else{
		rear->next=p;
		rear=p;
	}
	
}
template<typename T>
T LinkedQueue<T>::pop(){
	Node<T> *p;
	T temp;
	assert(!isEmpty());
	p=front;
	temp=p->data;
	front=front->next;
	delete p;
	if(front==NULL)
		rear=NULL;
	return temp;
}
template <typename T>
void LinkedQueue<T>::destroyQueue(){
	Node<T> *p;
	while(front!=NULL){
		p=front;
		front=front->next;
		delete p;
	}
	rear=NULL;
}










