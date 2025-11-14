#include <stdio.h>
int main()
{
	int numara;
	int numara2;
	int numara3=0;
	printf("Bir sayi giriniz: ");
	scanf("%d",&numara);	
	while(numara>0){
		numara2=numara%10;
		numara=numara/10;
		if(numara2>numara3)
		{
			numara3=numara2;
		}
	}
	printf("En buyuk hanenin degeri: %d",numara3);
	return 0;	
}
