#ifndef LINKEDLIST_H
#define LINKEDLIST_H

#include <iostream>
#include <cassert>

template <class T>
struct Node
{
	T data;
	Node<T> *next;
};

template <class T>
class LinkedList
{
public:
	void initialize();
	bool is_empty();
	int length();
	void destroy();
	T front();
	T back();
	const Node<T>* search(T& search_item);


	/* TODO: Homework 1 Addition */
	void insert_descending(T& new_item);


	void insert_head(T& new_item);
    void insert_last(T& new_item);
	void delete_node(T& delete_item);
	LinkedList();
	~LinkedList();
	
	template <class S>
	friend std::ostream& operator<<(std::ostream &os,LinkedList<S> &list)
	{
		Node<S> *p = new Node<S>;
		p=list.head;
		while (p != NULL)
		{
			os<<p->data<<" ";
			p=p->next;
		}
		return os;
	}
protected:
	int count;
	Node<T> *head;
	Node<T> *last;
};

template <class T>
const Node<T>* LinkedList<T>::search(T& search_item)
{
	Node<T> *temp = head;

	for (int i = 0; i < count; i++)
	{
		if (temp->data == search_item)
			return temp;
		temp = temp->next;
	}
	return NULL;
}

template <class T>
void LinkedList<T>::insert_head(T& new_item)
{
	Node<T> *newNode = new Node<T>;
	newNode->data = new_item;
	newNode->next = head;
	head = newNode;

	if (last == NULL)
		last = head;

	count++;
}

template <class T>
void LinkedList<T>::insert_last(T& new_item)
{
	Node<T> *newNode = new Node<T>;

	newNode->data = new_item;
	newNode->next = NULL;

	if (head == NULL)
	{
		head = last = newNode;
	}
	else
	{
		last->next = newNode;
		last = newNode;
	}

	count++;
}

/*TODO:  Homework 1 */
template <class T>
void LinkedList<T>::insert_descending(T& new_item) {
	Node<T> *WorkerPointer=new Node<T>;
	WorkerPointer->data=new_item;
	WorkerPointer->next=NULL;
	 if(this->is_empty())
	 {
	 	head=WorkerPointer;
	 	count++;
	 }
	 else
	 {
	 	if(head->data<WorkerPointer->data)
	 	{
	 		WorkerPointer->next=head;
	 		head=WorkerPointer;
	 		count++;
		}
		else
		{
			Node<T> *FirstLook;
			Node<T> *SecondLook=head;
			do{
				FirstLook=SecondLook;
				SecondLook=FirstLook->next;
			}while(SecondLook && WorkerPointer->data<SecondLook->data);
			if(!SecondLook)
			{
				SecondLook=WorkerPointer;
				FirstLook->next=SecondLook;
				SecondLook->next=NULL;
				count++;
			}
			else
			{
				FirstLook->next=WorkerPointer;
				WorkerPointer->next=SecondLook;
				count++;
			}
		}
	 }

	
}

template <class T>
void LinkedList<T>::delete_node(T& delete_item)
{
	Node<T> *current, *trailCurrent;
	bool found;

	if (head == NULL)
		std::cout << "Cannot delete from empty list." << std::endl;
	else
	{
		if (head->data == delete_item)
		{
			current = head;
			head = head->next;
			count--;
			if (head == NULL)
				last = NULL;
			delete current;
		}
		else
		{
			found = false;
			trailCurrent = head;
			current = head->next;

			while (current != NULL && !found)
			{
				if (current->data != delete_item)
				{
					trailCurrent = current;
					current = current->next;
				}
				else
				{
					found = true;
				}
			}
			if (found)
			{
				trailCurrent->next = current->next;
				count--;
				if (last == current)
					last = trailCurrent;
				delete current;
			}
			else
			{
				std::cout << "Item to be deleted is not in the list." << std::endl;
			}
		}
	}
}

template <class T>
bool LinkedList<T>::is_empty()
{
	return (head == NULL);
}

template <class T>
LinkedList<T>::LinkedList()
{
	head = NULL;
	last = NULL;
	count = 0;
}

template <class T>
void LinkedList<T>::destroy()
{
	Node<T> *temp;

	while (head != NULL)
	{
		temp = head;
		head = head->next;
		delete temp;
	}

	last = NULL;
	count = 0;
}

template <class T>
void LinkedList<T>::initialize()
{
	destroy();
}

template <class T>
int LinkedList<T>::length()
{
	return count;
}

template <class T>
T LinkedList<T>::front()
{
	assert(head != NULL);
	return head->data;
}

template <class T>
T LinkedList<T>::back()
{
	assert(last != NULL);
	return last->data;
}

template <class T>
LinkedList<T>::~LinkedList()
{
	destroy();
}


#endif
