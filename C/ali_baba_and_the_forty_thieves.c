#include<stdio.h>
void calculateGold(int,int);
int main(){
	int x;
	int y;
	int result=0;
	printf("How many rooms? ");
	scanf("%d",&x);
	printf("Which thief is the last one? ");
	scanf("%d",&y);
	calculateGold(x,y);
	return 0;
}
void calculateGold(int a,int b){
	int totalGold=0;
	int aliBabaGold=0;
	int lastThief=0;
	int sum=0;
	int i;
	int j;
	int k;
	int l;
	int m;
	for(i=1;i<a;i++)
	{
		if(a>1)
		{
			for(j=1;j<=40;j++)
			{
			totalGold+=j;
			}	
		}
		else if(a==1)
		{
			for(k=1;k<=b;k++)
			{
			totalGold+=k;
			}
		}	
	}
	for(l=0;l<a-1;l++)
	{
		aliBabaGold+=41;
	}
	if(a>=1)
	{
		for(m=1;m<=b;m++)
		{
			lastThief+=m;
		}
	}

	sum=totalGold+aliBabaGold+lastThief;
	printf("Total gold coins %d\n",sum);
	printf("Ali baba got %d coins",aliBabaGold);
}
