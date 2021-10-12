#include<stdio.h>
#define PI 3.14
int main(){
 char character;
 double circle_radius=0;
 int rectangle_width=0;
 int rectangle_length=0;
 int square_side=0;
 double circle_area=0;
 double circle_circumference=0;
 int rectangle_area=0;
 int rectangle_circumference=0;
 int square_area=0;
 int square_circumference=0;
 
 printf("Enter the shape(C/c, R/r, S/s) :");
 scanf("%c",&character);
 if(character=='C'||character=='c')
 	{
 		printf("Enter the raius of circle:");
 		scanf("%lf",&circle_radius);
 		circle_area=PI*(circle_radius*circle_radius);
 		circle_circumference=PI*2*circle_radius;
 		printf("Circle's area is %.2lf and circumference is %.2lf",circle_area,circle_circumference);
 	}
 else if(character=='R'||character=='r')
 	{
 		printf("Enter the width and length of rectangle:");
 		scanf("%d%d",&rectangle_width,&rectangle_length);
 		if(rectangle_width==rectangle_length)
		 {
 			printf("Error. It is not a rectangle.");
		 }
		 else
		 {
		 	rectangle_area=rectangle_width*rectangle_length;
		 	rectangle_circumference=2*(rectangle_width+rectangle_length);
		 	printf("Ractangle's area is %d and circumference is %d",rectangle_area,rectangle_circumference);
		 }
	}
 else if(character=='S' ||character=='s')
 	{
 		printf("Enter the side of square:");
 		scanf("%d",&square_side);
 		square_area=square_side*square_side;
 		square_circumference=4*square_side;
 		printf("Square's area is %d and circumference is %d",square_area,square_circumference);
	}
 else
 	{
 	printf("Eror! It is not a corret shape.");
	}
 return 0;
 }
 
