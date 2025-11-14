#include<stdio.h>
int main(){
	int birler;
	int onlar;
	int carpim;
	int i;
	int direnc=0;
	for(i=10;i<100;i++)
	{
		birler=i%10;
		onlar=i/10;
		carpim=onlar*birler;
		direnc++;
		if(carpim>10)
		{
			while(carpim>10)
			{
				birler=carpim%10;
				onlar=carpim/10;
				carpim=onlar*birler;
				direnc++;
			}
		}
		printf("%d sayisinin carpim direnci = %d\n",i,direnc);
		direnc=0;
	}
}
