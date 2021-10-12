#include<stdio.h>

int main(){
	int coursNumber;
	int totalAkts;
	int i;
	double gpa=0;
	int cAkts=0;
	double divison=0;
	double sum=0;
	double sum1=0;
	double result=0;
	printf("Enter the number of course: ");
	scanf("%d",&coursNumber);
	printf("Enter the total akts: ");
	scanf("%d",&totalAkts);
	
	for(i=1;i<=coursNumber;i++){
		printf("Enter the course akts and divison value:");
		scanf("%d%lf",&cAkts,&divison);
		sum=cAkts*divison;
		sum1+=sum;
	}
	result=(sum1*4)/(totalAkts*4);
	printf("your gpa are %.2lf .",result);
	
}
