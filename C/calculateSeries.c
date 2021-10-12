#include <stdio.h>
#include<math.h>
int calculateSeries(int a , int n, char op);
void getValues();
int main(){
	getValues();
	return 0;
}
int calculateSeries(int a , int n, char op){
	int sum;
	int productSum=1;
	int addSum=0;
	int i;
	if(op=='P'|| op=='p')
	{
		for(i=1;i<=n;i++)
		{
			productSum=productSum*(pow(a,i));
		}
		sum=productSum;
	}
	else if(op=='S'|| op=='s')
	{
		for(i=1;i<=n;i++)
		{
			addSum=addSum+(pow(a,i));
		}
		sum=addSum;
	}		
	return sum;
}
void getValues(){
	int firstnumber;
	int pow;
	char character;
	int result=0;
	int k;
	int count=1;
	printf("Enter the base number:");
	scanf("%d",&firstnumber);
	printf("Enter the pow limit:");
	scanf("%d",&pow);
	printf("Enter the operation symbol P or S:");
	scanf(" %c",&character);
	switch(character){
		case'P':
		case'p':
			result=calculateSeries(firstnumber,pow,character);
			printf("The series :");
			for(k=1;k<=pow;k++)
			{	
				printf("%d^%d",firstnumber,k);
				count++;
				if(count <= pow)
				printf(" * ");
			}
			printf("\n");
			printf("The result of the series :%d",result);
			break;
		case'S':
		case's':
			result=calculateSeries(firstnumber,pow,character);
			printf("The series :");
			for(k=1;k<=pow;k++)
			{
				printf("%d^%d",firstnumber,k);
				count++;
				if(count <= pow)
				printf(" + ");
			}
			printf("\n");
			printf("The result of the series :%d",result);
			break;	
		default:
			printf("Wrong operation symbol!!!");
			break;
	}
}
