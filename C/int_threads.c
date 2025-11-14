#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#define SIZE				12
#define NUMBER_OF_THREADS	7
void *sorter(void *params);
void *merger(void *params);
int list[SIZE] = {7,12,19,3,18,4,2,6,15,8,13,1};
int result[SIZE];

typedef struct
{
	int from_index;
	int to_index;
	int finish;
} parameters;

int main (int argc, const char * argv[]) 
{
	int i,k;
	printf("Before Sorting: ");
    for (i = 0; i < SIZE; i++)
		printf("%d  ",list[i]);
	printf("\n");
	
	pthread_t workers[NUMBER_OF_THREADS];
	
	parameters *data = (parameters *) malloc (sizeof(parameters));
	data->from_index = 0;
	data->to_index = (SIZE/4) - 1;
	pthread_create(&workers[0], 0, sorter, data);
	
	data = (parameters *) malloc (sizeof(parameters));
	data->from_index = (SIZE/4);
	data->to_index = SIZE/2 - 1;
	pthread_create(&workers[1], 0, sorter, data);
	
	data = (parameters *) malloc (sizeof(parameters));
	data->from_index = (SIZE/2);
	data->to_index = SIZE*3/4 - 1;
	pthread_create(&workers[2], 0, sorter, data);
	
	data = (parameters *) malloc (sizeof(parameters));
	data->from_index = (SIZE*3/4);
	data->to_index = SIZE- 1;
	pthread_create(&workers[3], 0, sorter, data);
	
	for (i = 0; i < NUMBER_OF_THREADS - 3; i++)
		pthread_join(workers[i], NULL);
	
	int a;
	for (a=0; a<4; a++) {
	
		printf("After %d. Sorting Thread Work: ",a+1);
		for (k = a*SIZE/4; k < (a+1)*SIZE/4; k++) {
		
			printf("%d  ",list[k]);
		}
		printf("\n");
	}
	data = (parameters *) malloc(sizeof(parameters));
	data->from_index = 0;
	data->to_index = (SIZE/4);
	data->finish = (SIZE/2);
	pthread_create(&workers[4], 0, merger, data);
	pthread_join(workers[4], NULL);
	printf("After first Merging Thread Work: ");
	for (i = 0; i < SIZE/2; i++)
		printf("%d  ",result[i]);
	printf("\n");
		
	data = (parameters *) malloc(sizeof(parameters));
	data->from_index = (SIZE/2);
	data->to_index = (SIZE*3/4);
	data->finish = SIZE;
	pthread_create(&workers[5], 0, merger, data);
	pthread_join(workers[5], NULL);
	
	printf("After second Merging Thread Work: ");	
	for (i = SIZE/2; i < SIZE; i++)
		printf("%d  ",result[i]);
	printf("\n");
	
	for (i = 0; i < SIZE; i++)
		list[i]=result[i];

	data = (parameters *) malloc(sizeof(parameters));
	data->from_index = 0;
	data->to_index = (SIZE/2);
	data->finish = SIZE;
	pthread_create(&workers[6], 0, merger, data);

	pthread_join(workers[6], NULL);

		printf("After third Merging Thread Work: ");
	for (i = 0; i < SIZE; i++)
		printf("%d  ",result[i]);
	printf("\n");
	
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
        while (j >= begin && list[j] > key) 
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
	printf("%d , %d\n",i,j);
	int position = i;
	
	while (i < p->to_index && j < p->finish) 
	{
		if (list[i] <= list[j]) 
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

