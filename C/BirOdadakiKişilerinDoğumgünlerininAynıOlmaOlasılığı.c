#include<stdio.h>
int main(){
	int sayi,i;
	double a;
	double b;
	double olasilik=1;
	printf("Odadaki kisi sayisini giriniz:");
	scanf("%d",&sayi);
	for(i=1;i<=sayi;i++)
	{
		a=366-i+1;
		olasilik=olasilik*(a/366);	
	}
	olasilik=1-olasilik;
	olasilik=olasilik*100;
	printf("Olasilik yuzde %.1lf",olasilik);
	return 0;
}
