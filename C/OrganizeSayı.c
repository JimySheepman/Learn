#include <stdio.h>
int main(){
	int sayi;
	int i,j;
	printf("Bir sayi giriniz : ");  
	scanf("%d",&sayi); 
	printf("Girdiginiz sayiya kadar olan organize sayilar:\n ");
	for ( i=1; i <sayi ; i++) {
		int sum = 0;
		for( j=1; j<i; j++){
			if(i%j==0){
				sum = sum + j;
			}
		}
		if(sum == i){
			printf("%d\n",i );
		}
	}	
	return 0;
}
