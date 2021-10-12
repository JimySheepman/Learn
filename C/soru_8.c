#include<stdio.h>
int main(){
	int sayi1;
	int sayi2;
	int birler1;
	int onlar1;
	int birler2;
	int onlar2;
	double sonuc1;
	double sonuc2;
	int i,j;
	for(i=10;i<100;i++)
	{
		sayi1=i;
		birler1=sayi1%10;
		onlar1=sayi1/10;
		if(onlar1 != birler1)
		{
			for(j=10;j<100;j++)
			{
				sayi2=j;
				birler2=sayi2%10;
				onlar2=sayi2/10;
				if(sayi1!=sayi2 && onlar2!=birler2)
				{
					sonuc1=sayi1/sayi2;
					if(onlar1 == onlar2)
					{
						sonuc2=birler1/birler2;
					}
					if(onlar1 == birler2)
					{
						sonuc2=birler1/onlar2;
					}
					if(birler1 == onlar2)
					{
						sonuc2=onlar1/birler2;
					}
					if(birler1 == birler2)
					{
						sonuc2=onlar1/onlar2;
					}
					if(sonuc1==sonuc2)
					{
						printf("%d ve %d uykucu yontemine uyar..\n",sayi1,sayi2);
					}
				}
			
			}
		}
	}
	
	
	
	
	
}
