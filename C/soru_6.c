#include<stdio.h>
int main(){
	int i;
	double dizi[10];
	double a;
	double sonuc;
	dizi[0]=1;
	for(i=0;i<10;i++)
	{
		a=dizi[i];
		dizi[i+1]=1+(1/a);

	}
	sonuc=dizi[10];
	printf("sonuc %lf",sonuc);
	return 0;
}
