#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SIZE				12
#define NUMBER_OF_THREADS	7

void *sorter(void *params);
void *merger(void *params);	

char *list[SIZE] ={"KARAMAN", "IZMIR","ANKARA", "KONYA", "KARABUK", "ISTANBUL", "GAZIANTEP", "MERSIN", "ZONGULDAK","VAN", "BOLU", "ORDU"};

char *result[SIZE];

typedef struct 
{
	int from_index;
	int to_index;
	int finish;
} parameters;

int main (int argc, const char * argv[]) 
{
	int i,k,a;
	char c='.';
	
	printf("Before Sorting: ");
    for (i = 0; i < SIZE; i++)
		printf("%s  ",list[i]);
	printf("\n");
	
	for(i=0;i<54;i++)
		printf("%c",c);
	printf("\n");
	
	pthread_t cities[NUMBER_OF_THREADS];
	
	parameters *data = (parameters *) malloc (sizeof(parameters));
	data->from_index = 0;
	data->to_index = (SIZE/4) - 1;
	pthread_create(&cities[0], 0, sorter, data);
	
	
	data = (parameters *) malloc (sizeof(parameters));
	data->from_index = (SIZE/4);
	data->to_index = SIZE/2 - 1;
	pthread_create(&cities[1], 0, sorter, data);
	
	data = (parameters *) malloc (sizeof(parameters));
	data->from_index = (SIZE/2);
	data->to_index = SIZE*3/4 - 1;
	pthread_create(&cities[2], 0, sorter, data);
	
	data = (parameters *) malloc (sizeof(parameters));
	data->from_index = (SIZE*3/4);
	data->to_index = SIZE- 1;
	pthread_create(&cities[3], 0, sorter, data);
	
	for (i = 0; i < NUMBER_OF_THREADS - 3; i++)
		pthread_join(cities[i], NULL);
		
	for (a=0; a<4; a++) 
	{
		printf("After %d. Sorting Thread Work: ",a+1);
		for (k = a*SIZE/4; k < (a+1)*SIZE/4; k++) 
		{
			printf("%s  ",list[k]);
		}
		printf("\n");
	}
	
	for(i=0;i<54;i++)
		printf("%c",c);
	printf("\n");
	
	data = (parameters *) malloc(sizeof(parameters));
	data->from_index = 0;
	data->to_index = (SIZE/4);
	data->finish = (SIZE/2);
	pthread_create(&cities[4], 0, merger, data);
	pthread_join(cities[4], NULL);
	
	printf("After first Merging Thread Work: ");
	for (i = 0; i < SIZE/2; i++)
		printf("%s  ",result[i]);
	printf("\n");
	
	data = (parameters *) malloc(sizeof(parameters));
	data->from_index = (SIZE/2);
	data->to_index = (SIZE*3/4);
	data->finish = SIZE;
	pthread_create(&cities[5], 0, merger, data);
	pthread_join(cities[5], NULL);
	
	printf("After second Merging Thread Work: ");	
	for (i = SIZE/2; i < SIZE; i++)
		printf("%s  ",result[i]);
	printf("\n");
	
	for (i = 0; i < SIZE; i++)
		list[i]=result[i];
		
	data = (parameters *) malloc(sizeof(parameters));
	data->from_index = 0;
	data->to_index = (SIZE/2);
	data->finish = SIZE;
	pthread_create(&cities[6], 0, merger, data);
	pthread_join(cities[6], NULL);
	
	for(i=0;i<54;i++)
		printf("%c",c);
	printf("\n");
	
	printf("After third Merging Thread Work: ");
	for (i = 0; i < SIZE; i++)
		printf("%s  ",result[i]);

    return 0;
}
void *sorter(void *params)
{
	parameters* p = (parameters *)params;
	int begin = p->from_index;
	int end = p->to_index;
	int i, key, j; 
    for (i = begin+1; i <=end; i++) 
	{ 
        key = list[i]; 
        j = i - 1; 
        while (j >= begin && strlen(list[j])> strlen(key)) 
		{ 
            list[j + 1] = list[j]; 
            j = j - 1; 
        } 
        list[j + 1] = key; 
    } 
	pthread_exit(0);
}    
void *merger(void *params)
{
	parameters* p = (parameters *)params;
	int i,j;
	i = p->from_index;
	j = p->to_index;
	int position = i;
	while (i < p->to_index && j < p->finish)
	{
		if (strlen(list[i]) <= strlen(list[j]))
		{
			result[position++] = list[i];
			i++;
		}
		else 
		{
			result[position++] = list[j];
			j++;
		}
	}
	if (i < p->to_index) 
	{
		while (i < p->to_index) 
		{
			result[position] = list[i];
			position++;
			i++;
		}
	}
	else 
	{
		while (j < p->finish) 
		{
			result[position] = list[j];
			position++;
			j++;
		}
	}
	pthread_exit(0);
}
