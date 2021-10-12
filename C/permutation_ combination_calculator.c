#include<stdio.h>
int Fact(int a)//calculate factorial fuctions
{
	int fact=1;
	int i;
	for(i=1;i<=a;i++)
	{
		fact=fact*i;
	}
	return fact;
}
double Calculate_C(int a,int b,int c)//calculate combination functions
{
	return a/(b*c);
}
double Calculate_P(int a,int b)//calculate permutation functions
{
	return a/b;
}
int main(){
	int Nnumber;
	int Rnumber;
	int Nfact=0;
	int Rfact=0;
	int N_difference_R=0;
	int N_Rfact=0;
	int error=0;
	while(error!=1)//loop
	{
	printf("Enter N value:");
	scanf("%d",&Nnumber);
	printf("Enter R value:");
	scanf("%d",&Rnumber);
	if(Nnumber>Rnumber && Nnumber>0 && Rnumber>0)//check value
	{
		Nfact=Fact(Nnumber);//call functions
		Rfact=Fact(Rnumber);//call functions
		N_difference_R=Nnumber-Rnumber;
		N_Rfact=Fact(N_difference_R);//call functions
		printf("C value : C(%d,%d) = %.2lf\n",Nnumber,Rnumber,Calculate_C(Nfact,Rfact,N_Rfact));
		printf("R value : R(%d,%d) = %.2lf\n",Nnumber,Rnumber,Calculate_P(Nfact,N_Rfact));	
		printf("\n");
	}
	else if(Rnumber>Nnumber && Nnumber>0 && Rnumber>0)//check value
	{
		printf("N value should be greater than R value ! Try again.\n");
		printf("\n");
	}
	else if(Rnumber<0 || Nnumber<0)//check value
	{
		printf("Error : N or R value should be equal greater than zero. Bye!");// error message
		error=1;// exit conditions
		printf("\n");
	}
	}
	return 0;
}
