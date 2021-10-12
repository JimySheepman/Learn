#include <stdio.h>

struct employee{ //initilize struct
	int id;
	char surname[10];
	int noProj;
	char job[10];
};

void coderInfo(struct employee arr[], int size){ 
	int i, count=0;
	FILE * ptr; // files pointer
	ptr = fopen("hw.txt", "w"); // set file to write mode
	
	for(i=0; i<size; i++){ //check if number of project is greater than 10 and increase count to get the number of employees
		if(arr[i].noProj >= 10){
			fprintf(ptr, "%d %s %d %s \n", arr[i].id, arr[i].surname, arr[i].noProj, arr[i].job); //use fprintf to write into file
			count++;
		}	
	}
	printf("\n%d Coders info work more than 10 projects recorded into file! \n", count); //output number of employees working on 10 or more projects
	fclose (ptr);	
}

void sortEmployees(struct employee arr[], int size){
	
	int i, max=5, count=size, temp;
	char tempName[10];
	
	printf("\nSorted list of employee according to the number of projects: \n");
	printf("Surename           Number of Projects: \n");
	printf("-----------------------------\n");
	
	while(count > 0){ //use two loops to continously check the array
		
		for(i=0; i<5; i++){ // for loop to check if max is less than the current employee info
			if(arr[i].noProj > max){ // if max is less than, set max to the employee info and temp to i to keep track of the max laction in array
				max = arr[i].noProj;
				temp = i;
			}
		}
		
		printf("%s", arr[temp].surname);
		printf("		");
		printf("%d", max);
		printf("\n");
		
		arr[temp].noProj = 0; // set max array info to 0 so when checked again it wont be taken twice
		count--; // decrease count for while loop condition
		max = 0;	// set max to 0 for new itiration
	}
}

int main(){

struct employee emp[5];
int i;

	printf("Enter id, surname, number of projects and job of 5 employees:\n");
	
for(i=0; i<5;i++){ // take input into struct
	
	scanf("%d %s %d %s", &emp[i].id, &emp[i].surname, &emp[i].noProj, &emp[i].job);
}

coderInfo(emp, 5); //call functions
sortEmployees(emp,5);

}
