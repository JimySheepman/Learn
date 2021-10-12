#include <stdio.h>
#include <stdlib.h>
#include <math.h>
int main(){
	
	int number;	//Declare number to be an integer variable
	
	char ch;	//Declare ch to be an character variable
	
	double result=0;//Declare result to be an real number variable
	
	int final;//Declare final to be an integer variable for result
	
	printf("Enter an Integer:");//output string
	
	scanf("%d",&number);// input number
	
	printf("MENU\n");//output string
	
	printf("Enter P or p to compute 2nd Power\n");//output string menu option
	
	printf("Enter S or s to compute Square Root\n");//output string menu option
	
	printf("Enter A or a to compute Absolute value\n");//output string menu option
	
	printf("Enter a Character:");//output string
	
	scanf(" %c",&ch);// input character
	
	switch(ch){
		
		case 'P':
			
			result=pow(number,2);//pow function result
			
			printf("The second power of %d is %.2lf",number,result);//output abs result
			
			break;//exit segment
			
		case 'p':
			
			result=pow(number,2);//pow function result
			
			printf("The second power of %d is %.2lf",number,result);//output abs result
			
			break;//exit segment
			
		case 'S':
			
			if(number<0)//check number positive
			{
				printf("Error!");// error message
			}
			else
			{
				result=sqrt(number);//sqrt function result
				
				printf("The square root of %d is %.2lf",number,result);//output sqrt result
				
			}
			
			break;//exit segment
			
		case 's':
			
			if(number<0)//check number positive
			{
				printf("Error!");// error message
			}
			else
			{
				result=sqrt(number);//sqrt function result
				
				printf("The square root of %d is %.2lf",number,result);//output sqrt result
				
			}
			
			break;//exit segment
			
		case 'A':
			
			final=abs(number);//abs function result
			
			printf("The absolute value of %d is %d",number,final);//output abs result
			
			break;//exit segment
			
		case 'a':
			
			final=abs(number);//abs function result
			
			printf("The absolute value of %d is %d",number,final);//output abs result
			
			break;//exit segment
			
		default:
			
			printf("Error!");// error message
			
			break;//exit segment
			
	}	
	
	return 0;
}
