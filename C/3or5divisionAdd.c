#include<stdio.h>
int main(){
	int number;//number input
	int i;//loop value 
	int sum=0;//to find the sum value
	printf("Enter an integer value: ");
	scanf("%d",&number);
	if(number>=0)//check number for error
	{
		for(i=0;i<=number;i++)//calculate loop
		{
			if(i%3==0||i%5==0)//calculate conditions
			{
				sum+=i;//calculate operation
			}
		}
		/*this loop works like this
		example user input is 16;
		i=0   for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=1   for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=2   for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=3   for check to i mod 3=0 or i mod 5=0 yes,so sum =0+3
		i=4   for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=5   for check to i mod 3=0 or i mod 5=0 yes,so sum =0+3+5
		i=6   for check to i mod 3=0 or i mod 5=0 yes,so sum =0+3+5+6
		i=7   for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=8   for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=9   for check to i mod 3=0 or i mod 5=0 yes,so sum =0+3+5+6+9
		i=10  for check to i mod 3=0 or i mod 5=0 yes,so sum =0+3+5+6+9+10
		i=11  for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=12  for check to i mod 3=0 or i mod 5=0 yes,so sum =0+3+5+6+9+10+12
		i=13  for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=14  for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		i=15  for check to i mod 3=0 or i mod 5=0 yes,so sum =0+3+5+6+9+10+12+15
		i=16  for check to i mod 3=0 or i mod 5=0 nope,so sum =0
		so, summation is 60
		*/
		printf("The summation is %d\n",sum);
	}
	else//error message
	{
		printf("invalid input...!\n");
	}
return 0;
}
