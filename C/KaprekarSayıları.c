#include<stdio.h>
int main(){
	int asilsayi;
	int sayi;
	int karesayi;
	int basamak=1;
	int sag=0;
	int sol=0;
	int sonuc;
	printf("sayi girin:");
	scanf("%d",&asilsayi);
	karesayi=asilsayi*asilsayi;
	printf("Sayinin karesi:%d\n",karesayi);
	sayi=karesayi;
    while(sayi>10)
	{
  		basamak++; 
  		sayi/=10;         
	}
        if(basamak%2==0){
        	if(basamak==2)
			{
				sol=karesayi/10;
				sag=karesayi%10;
				sonuc=sag+sol;
				if(asilsayi==sonuc)
				{
					printf("Girdiginiz %d sayisi %d + %d = %d bir Kaprekar sayidir..",asilsayi,sol,sag,sonuc);
				}
				else
				{
					printf("Girdiginiz %d sayisi bir Kaprekar sayidir degildir..",asilsayi);
				}
			}
			else if(basamak==4)
			{
				sol=karesayi/100;
				sag=karesayi%100;
				sonuc=sag+sol;
				if(asilsayi==sonuc)
				{
					printf("Girdiginiz %d sayisi %d + %d = %d bir Kaprekar sayidir..",asilsayi,sol,sag,sonuc);
				}
				else
				{
					printf("Girdiginiz %d sayisi bir Kaprekar sayidir degildir..",asilsayi);
				}
			}
			else if(basamak==6){
				sol=karesayi/1000;
				sag=karesayi%1000;
				sonuc=sag+sol;
				if(asilsayi==sonuc)
				{
					printf("Girdiginiz %d sayisi %d + %d = %d bir Kaprekar sayidir..",asilsayi,sol,sag,sonuc);
				}
				else
				{
					printf("Girdiginiz %d sayisi bir Kaprekar sayidir degildir..",asilsayi);
				}
			}
			else if(basamak==8){
				sol=karesayi/10000;
				sag=karesayi%10000;
				sonuc=sag+sol;
				if(asilsayi==sonuc)
				{
					printf("Girdiginiz %d sayisi %d + %d = %d bir Kaprekar sayidir..",asilsayi,sol,sag,sonuc);
				}
				else
				{
					printf("Girdiginiz %d sayisi bir Kaprekar sayidir degildir..",asilsayi);
				}
			}	
		}
		else{
			if(basamak==1){
				if(asilsayi==karesayi){
					printf("Girdiginiz %d sayisi %d bir Kaprekar sayidir..",asilsayi,karesayi);
				}	
			}
			else if(basamak==3){
				sol=karesayi/100;
				sag=karesayi%100;
				sonuc=sag+sol;
				if(asilsayi==sonuc)
				{
					printf("Girdiginiz %d sayisi %d + %d = %d bir Kaprekar sayidir..",asilsayi,sol,sag,sonuc);
				}
				else
				{
					printf("Girdiginiz %d sayisi bir Kaprekar sayidir degildir..",asilsayi);
				}
			}
			else if(basamak==5){
				sol=karesayi/1000;
				sag=karesayi%1000;
				sonuc=sag+sol;
				if(asilsayi==sonuc)
				{
					printf("Girdiginiz %d sayisi %d + %d = %d bir Kaprekar sayidir..",asilsayi,sol,sag,sonuc);
				}
				else
				{
					printf("Girdiginiz %d sayisi bir Kaprekar sayidir degildir..",asilsayi);
				}
			}
			else if(basamak==7){
				sol=karesayi/10000;
				sag=karesayi%10000;
				sonuc=sag+sol;
				if(asilsayi==sonuc)
				{
					printf("Girdiginiz %d sayisi %d + %d = %d bir Kaprekar sayidir..",asilsayi,sol,sag,sonuc);
				}
				else
				{
					printf("Girdiginiz %d sayisi bir Kaprekar sayidir degildir..",asilsayi);
				}
			}
		}		
		return 0;
}
