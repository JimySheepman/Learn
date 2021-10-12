#include<stdio.h>
int main(){
	int sayi;
	int i;
	int bolum;
	int sonuc=0;
	printf("Sayi girin:");
	scanf("%d",&sayi);
	for(i=1;i<sayi;i++){
		bolum=sayi % i;
		if(bolum==0){
			sonuc+=i;
		}
		
	}
	if(sonuc<sayi){
		printf("Bu sayi noksan sayidir..");
	}
	else
	{
		printf("Bu sayi noksan sayi degildir..");
	}	
	return 0;
}
