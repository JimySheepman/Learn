#include <ctype.h> 
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void keycopy(char check[],char input[]){
	char ch,chr;
    int i=0;
    int t=0;
    while(strlen(check)!=strlen(input)){
    	
		ch=check[i];
		strncat(check,&ch,1);
		if(i>strlen(check)-1){
			i=0;
		}
    	i++;
	}
	 while (check[t]) 
			{ 
        		chr = check[t]; 
        		printf("%c", toupper(chr)); 
        		check[t]=toupper(chr);
        		t++; 
   			} 
}
void encrpt(char text[],char keys[],char textc[]){
	int a=0;
	int b=0;
	int c;
	char ca;
	char cb;
	char ch;
	char chiphertext[100];
	char table[][26] = {
{'M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A',},
{'B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M',},
{'C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B',},
{'D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C',},
{'F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D',},
{'E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F',},
{'G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E',},
{'H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G',},
{'K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H',},
{'J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K',},
{'I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J',},
{'S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I',},
{'Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S',},
{'N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z',},
{'O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N',},
{'P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O',},
{'Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P',},
{'R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q',},
{'L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R',},
{'T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L',},
{'U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T',},
{'X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U',},
{'W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X',},
{'V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W',},
{'Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V',},
{'A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y',},
};
printf("Output (phase-1): ");
for(c=0;c<strlen(text);c++)
{
	ca=text[c];
	a=ca%65;
	cb=keys[c];
	b=cb%65;
	ch=table[a][b];
	printf("%c",ch);
	textc[c]=ch;
}

}
void ciphertextsplit(char arr[],char arr1[]){
 	char  leftHalf[100];
	char rightHalf[100];
	char ch='0';
	char ck,cj;
    int length, mid, i, k;
	length=strlen(arr);
	if(length%2==0){
		mid = length/2;
    for(i = 0; i < mid; i++) {
        leftHalf[i]= arr[i];
    }
    leftHalf[i] = '\0';
    for(i = mid, k = 0; i <= length; i++, k++) {
 rightHalf[k]= arr[i];
    }
    printf("Left half : %s\n",leftHalf);
    printf("Right half : %s\n",rightHalf);
		
	}
	else{
		strncat(arr, &ch, 1);
		length=strlen(arr);
		mid = length/2;
    for(i = 0; i < mid; i++) {
        leftHalf[i]= arr[i];
    }
    leftHalf[i] = '\0';
    for(i = mid, k = 0; i <= length; i++, k++) {
 rightHalf[k]= arr[i];
    }
    printf("Left half : %s\n",leftHalf);
    printf("Right half : %s\n",rightHalf);
	}
	printf("\n");
	printf("\n");
	 	for(i=0;i<length/2;i++){
 			ck=leftHalf[i];
 			strncat(arr1, &ck, 1);
 			cj=rightHalf[i];
 			strncat(arr1, &cj, 1);
	 }
	 printf("Ciphertext: %s\n",arr1);
	
}
void decryptiontext(char arr[],char arr1[]){
	char group1[50]="";
	char group2[50]="";
	int i,size;
	char ch,ck;
	char c='0';
	size=strlen(arr);
	for(i=0;i<size;i+=2)
	{
		ch=arr[i];
 		strncat(group1, &ch, 1);
 		ck=arr[i+1];
 		strncat(group2, &ck, 1);
	}
	printf("\n");
	printf("group-1: %s\n",group1);
	printf("group-2: %s\n",group2);
	strcat(arr1,(strcat(group1,group2)));
	for(i=0;i<size;i++){
		if(arr1[i]==c){
			arr1[i]='\0';
		}
	}
	printf("Output(phase-2): %s\n",arr1);
}
void founttext(char arr[],char arr1[]){
	int i;
	int k;
	int b=0;
	int a=0;
	int size=0;
	char ch,ck;
	char text[100]="";
	char table[][26] = {
{'M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A',},
{'B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M',},
{'C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B',},
{'D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C',},
{'F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D',},
{'E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F',},
{'G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E',},
{'H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G',},
{'K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H',},
{'J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K',},
{'I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J',},
{'S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I',},
{'Z','N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S',},
{'N','O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z',},
{'O','P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N',},
{'P','Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O',},
{'Q','R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P',},
{'R','L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q',},
{'L','T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R',},
{'T','U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L',},
{'U','X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T',},
{'X','W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U',},
{'W','V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X',},
{'V','Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W',},
{'Y','A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V',},
{'A','M','B','C','D','F','E','G','H','K','J','I','S','Z','N','O','P','Q','R','L','T','U','X','W','V','Y',},
};
	size=strlen(arr);
	printf("size:%d\n",size);
	for(k=0;k<size;k++)
	{
		ch=arr[k];
		ck=arr1[k];//key
		a=ck%65;
		for(i=0;i<26;i++)
		{
			if(table[a][i]==ch)
			{
				b=i+65;
				ck=b;
				strncat(text, &ck, 1);	
			}
		}
	}
	printf("Plaintext: %s\n",text);
}
int main(){
	char plaintext[100];
	char plaintext2[100]="";
	char ctext[100]="";
	char ciphertext[100]="";
	char key[100];
	char key2[100]="";
	int flag=0;
	int i,j,t=0;
	char *ptr;
	int count=0;
	int size = 0;
	int sizekey=0;
	char chr;
	do{
		printf("Simple Cipher:\n");
		printf("[1] Encrypt\n");
		printf("[2] Decrypt\n");
		printf("[3] Exit\n");
		printf("Selection  ");
		scanf("%d",&flag);
		if(flag==1){
			printf("\n");
			printf("Enter text: ");
			fgets (plaintext, 100, stdin);
			scanf ("%[^\n]%*c", plaintext); 
			size = strlen(plaintext);
			printf("Enter key: ");
			scanf("%s", key); 
			sizekey = strlen(key);
			printf("\n");
			for(i=0;i<20;i++)
			{
				printf("*");
			}
			printf(" Encryption ");
			for(j=0;j<20;j++)
			{
				printf("*");
			}
			printf("\n");
			printf("Encryption Phase-1\n");
			ptr = strtok (plaintext," ,.-");
			printf("Plaintext:");
 			 while (ptr != NULL)
  			{
   				//printf ("%s",ptr);
   				strcat(plaintext2,ptr);
    			ptr = strtok (NULL, " ,.-");
    			count++;
  			}
  			 while (plaintext2[t]) 
			{ 
        		chr = plaintext2[t]; 
        		printf("%c", toupper(chr)); 
        		plaintext2[t]=toupper(chr);
        		t++; 
   			} 
  			printf("\n");
			printf("Key: ");
			keycopy(key,plaintext2);
			printf("\n");
			encrpt(plaintext2,key,ctext);
			printf("\n");
			printf("\n");
			printf("Encrption Phase-2\n");
			printf("Inputtext: %s\n",ctext);
			ciphertextsplit(ctext,ciphertext);
			for(j=0;j<52;j++)
			{
				printf("*");
			}
			printf("\n");
			printf("\n");
		}
		else if(flag==2){
						for(i=0;i<20;i++)
			{
				printf("*");
			}
			printf(" Decryption ");
			for(j=0;j<20;j++)
			{
				printf("*");
			}
			printf("\n");
			printf("\n");
			printf("Decryption Phase-2\n");
			printf("Inputtext: %s",ciphertext);
			decryptiontext(ciphertext,key2);
			printf("\n");
			printf("\n");
			printf("Decryption Phase-1\n");
			printf("Inputtext: %s\n",key2);
			printf("key: %s\n",key);
			printf("\n");
			printf("\n");
			founttext(key2,key);
						for(j=0;j<52;j++)
			{
				printf("*");
			}
				printf("\n");
			printf("\n");
		}
		
	}while(flag!=3);
	return 0;
}
