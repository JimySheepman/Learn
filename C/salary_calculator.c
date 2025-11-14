#include <stdio.h> 
int main() {
	char name[30]; // crate variable for name  length max 30 character
	char surname[30]; // crate variable for surname length max 30 character
	int weekly_working_hours = 0; // crate variable for weekly working hours
	char payment_mode; // crate variable for payment method
	int km = 0; // crate variable for extra km 
	int children_numbers = 0; // crate variable for children numbers
	int weekly_salary = 0;// crate variable for weekly salary
	double salary = 0; // crate variable for monthly salary

	printf("Enter you name and surname: ");
	scanf(" %s%s",name,surname);
	printf("----------------------------\n");
	printf("Enter the number of hours you are working for a week: ");
	scanf(" %d",&weekly_working_hours);
	if (weekly_working_hours>0 && weekly_working_hours<110) {//control of weekly working hours
		if (weekly_working_hours < 20) // calculate of salary  
		{
			weekly_salary = weekly_working_hours * 50;
			salary = weekly_salary * 4;
		}
		else if (weekly_working_hours >= 20 && weekly_working_hours <= 45)
		{
			weekly_salary = weekly_working_hours * 75;
			salary = weekly_salary * 4;
		}
		else
		{
			weekly_salary = weekly_working_hours * 100;
			salary = weekly_salary * 4;
		}
		printf("Select extra payment method. Enter K/k for km, C/c for amount of children:");
		scanf(" %c", &payment_mode);
		switch (payment_mode) //  control of payment method and calculate of salary
		{
		case 'K':
		case 'k':
			printf("Enter the km that you came to work:");
			scanf(" %d",&km );
			if (km<25)
			{
				salary = salary * 1.04;
			}
			else if(km>=25 && km<=50)
			{
				salary = salary * 1.08;
			}
			else
			{
				salary = salary * 1.12;
			}
			break;
		case 'C':
		case 'c':
			printf("Enter the amount of children you have: ");
			scanf(" %d",&children_numbers);
			if (children_numbers>0 && children_numbers<2)
			{
				salary = salary * 1.05;
			}
			else if (children_numbers >= 2 && children_numbers <= 4)
			{
				salary = salary * 1.09;
			}
			else
			{
				salary = salary * 1.15;
			}
			break;
		default:
			printf("You entered incorrectly, please try again !!! ");
		}
		printf("\nDear %s %s, your salary for this month is: $%.2lf", name,surname,salary ); // result line
	}
	else
	{
		printf("You entered incorrectly, please try again !!! ");
	}

	return 0;
}
