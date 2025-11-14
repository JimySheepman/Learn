#ifndef COMMONS_H
#define COMMONS_H

template <class Type>
struct node_t
{
	Type info;
	node_t<Type> *link;
};

#endif;