#include<stdio.h>
int main(){
	int sayi1;
	int sayi2;
	int kalan=1,bolum,bolunen;
	printf("Birinci saiyyi giriniz: ");
	scanf("%d",&sayi1);
	printf("Ikinci saiyyi giriniz: ");
	scanf("%d",&sayi2);
	if(sayi1>sayi2)
	{
		kalan=sayi1%sayi2;
		while(kalan!=0)
		{
			bolum=kalan;
			bolunen=sayi2;
			kalan=bolunen%bolum;
		}
		printf("OBEB leri :%d",bolum);	
	}
	else if(sayi2>sayi1)
	{
		kalan=sayi2%sayi1;
		while(kalan!=0)
		{
			bolum=kalan;
			bolunen=sayi1;
			kalan=bolunen%bolum;
		}
		printf("OBEB leri :%d",bolum);
	}
	else
	{
		printf("hata....!");
	}
	return 0;
}
