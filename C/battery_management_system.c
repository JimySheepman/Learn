#include <stdio.h>

int main(){
	double battery_array[20]={3.5,2.8,3.4,3.3,2.0,1.8,3.4,3.6,3.7,3.2,3.5,2.9,3.1,3.5,3.4,2.6,3.0,3.4,2.9,3.3};
	int i,sturdy=0,charge=0,degis=0;
	double best_case,average_case,worst_case;
	for(i=0;i<20;i++){
		if(battery_array[i]>3.6 ){
			printf("BAT%d sturdy\n",i+1);
			sturdy++;
		}
		else if(battery_array[i]<3.6 && battery_array[i]>=3.2){
			printf("BAT%d must be charged\n",i+1);
			charge++;
		}
		else{
			printf("BAT%d must be alteration\n",i+1);
			degis++;
		}
	}
	best_case= (sturdy*100)/20;
	average_case= (charge*100)/20;
	worst_case=(degis*100)/20;
	printf("------------------SISTEM ANALIZI------------------\n");
	printf("%.0f percent of the system is in good condition\n",best_case);
	printf("%.0f percent of the system should be charged\n",average_case);
	printf("%.0f percent of the system should be alteration\n",worst_case);
	if(worst_case>49){
		printf("\n\nThe system is inoperable.\n");
	}
	else{
		printf("\n\nThe system is operational.\n");
	}
}
