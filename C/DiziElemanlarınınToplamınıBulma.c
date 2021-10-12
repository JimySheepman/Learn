#include <stdio.h>

int main(){
	int dizi[20];
	int toplam=0;
	int i=0;
	int j,k;
	while(i!=20)
	{
		printf("Dizinin %d elemanini giriniz: ",i+1);
		scanf("%d",&dizi[i]);
		i++;
	}
	for(j=0;j<20;j++)
	{
		for(k=0;k<20;k++)
		{
			if(dizi[j]+dizi[k]==20)
			{
				printf("%d ve %d elemanlarin toplami = 20\n",j+1,k+1);
			}	
		}
	}

return 0;
}
