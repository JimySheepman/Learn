#include<stdio.h>
#include<math.h>	
int main()
{
	int sayi;
	int sonuc=0;
	int i=0;
	int toplam=0;
	int kalan;
	printf("Ikilik taban da bir sayi giriniz:");
	scanf("%d",&sayi);
		while(sayi>0)
		{
		        kalan=sayi%10;
			sonuc=kalan*pow(2,i);
			toplam=toplam+sonuc;
			sayi=sayi/10;
			i++;
		}
	
	sonuc=toplam;

	printf("Onluk tabandaki karsiligi:%d",sonuc);
	return 0;
}

