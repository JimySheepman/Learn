#include <stdio.h>
void polynomial (int Fdegree, int arr[]){
	int flag_point = 0;
	Fdegree=Fdegree+ 1;
	printf("f(x)=");
	do{
		Fdegree--;
		if(flag_point != 0){
			if(arr[Fdegree] >= 0){
				printf("+");
			}
		}
		printf("%d", arr[Fdegree]);
		if(Fdegree > 0){
			printf("x^%d", Fdegree);
		}
		flag_point++;
	}while(Fdegree != 0);
	printf("\n");
}
void find_derivative(int array[], int dvt, int size){
	int flag = 0;
	int newDegree;
	int newSize;
	int newDerivative = dvt;
	int i;
	printf("%d. derivative: df(x)= " ,dvt);
	while(size >= dvt)
	{
		newSize = size;
		for( i = 0; i < dvt; i++){
			array[size] = newSize *  array[size];
			newSize--;
		}
		if(flag != 0){
			if(array[size] >= 0){
				printf("+");
			}
		}
		printf("%d", array[size]);
		if(size > 1){
			newDegree = size - dvt;
			if(newDegree == 0){
			}
			else{
				printf("x^%d", newDegree);
			}
		}
		flag++;
		if(dvt == size){
			break;
		}
		size--;
	}
}
int main (){
	int polynomial_degree;
	int derivative_degree;
	int size_degree;
	int function_array[10];
	printf("Enter the degree of polynomial equation (1-10): ");
	scanf("%d", &polynomial_degree);
	size_degree = polynomial_degree;
	while(polynomial_degree !=-1){
		printf("Enter the coefficient of x to the power %d : ", polynomial_degree);
		scanf("%d", &function_array[polynomial_degree]);
		polynomial_degree--;
	}
	polynomial(size_degree, function_array);
	printf("Which degree of derivative you want to calculate? ");
	scanf("%d", &derivative_degree);
	find_derivative(function_array, derivative_degree, size_degree);
	return 0;
}
