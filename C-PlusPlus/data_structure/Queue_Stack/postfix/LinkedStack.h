#include <iostream>
#include <assert.h>
using namespace std;
template <typename T>
struct node{
	T data;
	node<T> *next;
};
template <typename T>
class LinkedStack{
		node<T> *top;
	public:
		LinkedStack(){
			top=0;
		}
		~LinkedStack(){
			destroy();
		}
		bool isEmpty(){
			return (top==NULL);
		}
		T showTop(){
			assert(!top==NULL);
			return top->data;
			
		}
		T pop();
		void push(T &item);
		void destroy();
};
template <typename T>
void LinkedStack<T>::destroy(){
	node<T> *temp;
	while(top != NULL){
		temp=top;
		top=top->next;
		delete temp;	
	}
	
}
template <typename T>
T LinkedStack<T>::pop(){
	node<T> *p=new node<T>;
	T item;
	assert(!isEmpty());
	p=top;
	top=top->next;
	item=p->data;
	delete p;
	return item;
}
template <typename T>
void LinkedStack<T>::push(T &item){
	node<T> *newnode=new node<T>;
	newnode->data=item;
	newnode->next=top;
	top=newnode;
}








