#include<stdio.h>
int main(){
	int numara=100;
	int numara2=0;
	int toplam_sayi=0;
	int birler_basamagi=0;
	int onlar_basamagi=0;
	int yuzler_basamagi=0;
	int miktar=0;
	while(numara!=999){
		yuzler_basamagi=numara/100;
		numara2=numara%100;
		onlar_basamagi=numara2/10;
		birler_basamagi=numara2%10;
		if(yuzler_basamagi == onlar_basamagi || yuzler_basamagi==birler_basamagi || onlar_basamagi==birler_basamagi)
		{
			numara++;
		}
		else{
			printf("%d ",numara);
			miktar++;
			numara++;
		}
		
	}
	printf("\nToplam %d sayi.",miktar);
}
