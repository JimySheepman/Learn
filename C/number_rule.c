#include <stdio.h>
#include <time.h>
#include <stdlib.h>
int main()
{
	int cStare[100];// array of current stare
	int nState[100];// array of next state
	int flag = 1;// time counter
	int printN;// print value
	int fNumber;// firt number
	int sNumber;// second number
	int tNumber;// third number
	int one = 0;//  counter
	int two = 1;// counter
	int three = 3;// counter
	int rNumber;// rule number
	int i;
	printf("Please enter the rule number (Betwenn 0 and 255): ");
	scanf("%d",&rNumber);
	printf("Current State    Next State\n");
	for(i=0;i<26;i++)
	{
		printf("-");
	}
	printf("0			0\n");
	printf("1			1\n");
	printf("10			1\n");
	printf("11			1\n");
	printf("100			1\n");
	printf("101			0\n");
	printf("110			0\n");
	printf("111			0\n");
	printf("\n");
	for(int i = 0; i < 100; i++)//reset process
	{
		cStare[i] = 0;
		nState[i] = 0;
	}
	cStare[49] = 1;
	for(int i =0; i < 100; i++)
	{
		printN = cStare[i];
		
		if(printN == 0)
		{
			printf("_ ");
		}
		
		else if(printN == 1)
		{
			printf("# ");
		}
	}
	printf("\n");
	do
	{
		while(one < 97)//algorithm
		{
			fNumber = cStare[one];
			sNumber = cStare[two];
			tNumber = cStare[three];
			if(fNumber == 0 && sNumber == 0 && tNumber == 1)
			{
				nState[two] = 1;
			}
			else if(fNumber == 0 && sNumber == 1 && tNumber == 0)
			{
				nState[two] = 1;
			}
			else if(fNumber == 0 && sNumber == 1 && tNumber == 1)
			{
				nState[two] = 1;
			}
			else if(fNumber == 1 && sNumber == 0 && tNumber == 0)
			{
				nState[two] = 1;
			}
			else
			{
				nState[two] = 0;
			}
			one++;
			two++;
			three++;
		}
		one = 0;
		two = 1;
		three = 3;
		for(int i = 0; i < 100; i++)//equel process
		{
			cStare[i] = nState[i];
		}
		for(int i =0; i < 100; i++)//print loop
		{
			printN = cStare[i];
			
			if(printN == 0)
			{
				printf("_ ");
			}
			
			else if(printN == 1)
			{
				printf("# ");
			}
		}
		printf("\n");
		flag++;
	}while(flag < 50);
	return 0;	
}
