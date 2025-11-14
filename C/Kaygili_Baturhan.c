#include <stdio.h>
#include <math.h>
int main(){
	int number;	//Declare number to be an integer variable
	char ch;	//Declare ch to be an character variable
	double result=0;//Declare result to be an real number variable
	int res;//Declare res to be an integer variable for result
	printf("Enter an Integer:");//display text
	scanf("%d",&number);// input number
	printf("MENU\n");// display menu
	printf("Enter P or p to compute 2nd Power\n");//display menu 1.tab
	printf("Enter S or s to compute Square Root\n");//display menu 2.tab
	printf("Enter A or a to compute Absolute value\n");//display menu 3.tab
	
	printf("Enter a Character:");//display text
	scanf(" %c",&ch);	// input character
	
	if(ch=='P'|| ch=='p')//character check
	{
		result=pow(number,2);//pow function result
		printf("The second power of %d is %.2lf",number,result);//display result
	}
	else if(ch=='S' || ch=='s') //character check
	{
		if(number<0)//check number positive
		{
		printf("Error!");// error message
		}
		else
		{
		result=sqrt(number);//sqrt function result
		printf("The square root of %d is %.2lf",number,result);//display result
		}
	}
	else if(ch=='A'|| ch=='a')//character check
	{
		res=abs(number);//abs function result
		printf("The absolute value of %d is %d",number,res);//display result
	}
	else
	{
		printf("Error!");// error message
	}

	return 0;
}
