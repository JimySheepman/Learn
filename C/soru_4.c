#include<stdio.h>
int main(){
	int aractipi;
	int saat;
	int aracsayisi1=0;
	int aracsayisi2=0;
	int aracsayisi3=0;
	int aracsayisi4=0;
	double toplampara=0;
	double ucret=0;			
	printf("Aractipi giriniz:");
	scanf("%d",&aractipi);	
	do{
		if(aractipi>0 && aractipi<5){
			printf("Aractipi giriniz:");
			scanf("%d",&aractipi);	
			if(aractipi==1)
			{
				printf("Kalinan sureyi girin:");
				scanf("%d",&saat);
				if(saat>=0 && saat<3)
				{
					ucret=15*0.5;
					toplampara+=ucret;	
					aracsayisi1++;
				}
				else if(saat>2 && saat<7)
				{
					ucret=25.5*0.5;
					toplampara+=ucret;	
					aracsayisi1++;
				}
				else if(saat>6 && saat<11)
				{
					ucret=30*0.5;
					toplampara+=ucret;	
					aracsayisi1++;
				}if(saat>10)
				{
					ucret=37.5*0.5;
					toplampara+=ucret;	
					aracsayisi1++;
				}
			}
			else if(aractipi==2)
			{
				printf("Kalinan sureyi girin:");
				scanf("%d",&saat);
				if(saat>=0 && saat<=2)
				{
					ucret=15*1;
					toplampara+=ucret;	
					aracsayisi2++;
				}
				else if(saat>=3 && saat<=6)
				{
					ucret=25.5*1;
					toplampara+=ucret;	
					aracsayisi2++;
				}
				else if(saat>=7 && saat<=10)
				{
					ucret=30*1;
					toplampara+=ucret;	
					aracsayisi2++;
				}if(saat>=11)
				{
					ucret=37.5*1;
					toplampara+=ucret;	
					aracsayisi2++;
				}
				else
				{
					printf("Yanlis saat girdiniz..");
				}
			}
			else if(aractipi==3)
			{
				printf("Kalinan sureyi girin:");
				scanf("%d",&saat);
				if(saat>=0 && saat<=2)
				{
					ucret=15*1.5;
					toplampara+=ucret;	
					aracsayisi3++;
				}
				else if(saat>=3 && saat<=6)
				{
					ucret=25.5*1.5;
					toplampara+=ucret;	
					aracsayisi3++;
				}
				else if(saat>=7 && saat<=10)
				{
					ucret=30*1.5;
					toplampara+=ucret;	
					aracsayisi3++;
				}if(saat>=11)
				{
					ucret=37.5*1.5;
					toplampara+=ucret;	
					aracsayisi3++;
				}
				else
				{
					printf("Yanlis saat girdiniz..");
				}
			}
			else if(aractipi==4)
			{
				printf("Kalinan sureyi girin:");
				scanf("%d",&saat);
				if(saat>=0 && saat<=2)
				{
					ucret=15*2;
					toplampara+=ucret;	
					aracsayisi4++;
				}
				else if(saat>=3 && saat<=6)
				{
					ucret=25.5*2;
					toplampara+=ucret;	
					aracsayisi4++;
				}
				else if(saat>=7 && saat<=10)
				{
					ucret=30*2;
					toplampara+=ucret;	
					aracsayisi4++;
				}if(saat>=11)
				{
					ucret=37.5*2;
					toplampara+=ucret;	
					aracsayisi4++;
				}
				else
				{
					printf("Yanlis saat girdiniz..");
				}
			}
		}
		else
		printf("Yanlis aractipi...");
		
	}while(aractipi!=0);
	printf("Motosiklet sayisi:%d\n",aracsayisi1);
	printf("Otomobil sayisi:%d\n",aracsayisi2);
	printf("Minibus sayisi:%d\n",aracsayisi3);
	printf("Otobus sayisi:%d\n",aracsayisi4);
	printf("Gunluk toplam kazanc:%.3lf",toplampara);
	return 0;
}
