#include <linux/init.h>
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/string.h>

struct car_struct {
	pid_t 	cid;
	int 	year;
	char 	*brand;
	char 	*model;
	char 	*color;
	int    motor_volume;
	struct 	list_head list;
};

static LIST_HEAD(car_list);

int simple_init(void)
{
	struct car_struct *head;
	struct car_struct *ptr;
	
	printk(KERN_INFO "Loading Car Module\n");

	head = kmalloc(sizeof(*head), GFP_KERNEL);
	head->cid = 1;
	head->year = 1998;
	head->brand = "Ford";
	head->model = "Focus";
	head->color = "dark blue";
	head->motor_volume = 3;
	INIT_LIST_HEAD(&head->list);
	list_add_tail(&head->list, &car_list);

	head = kmalloc(sizeof(*head), GFP_KERNEL);
	head->cid = 2;
	head->year = 1995;
	head->brand = "Renault";
	head->model = "Megane";
	head->color = "blue";
	head->motor_volume = 1;
	INIT_LIST_HEAD(&head->list);
	list_add_tail(&head->list, &car_list);
	   
	head = kmalloc(sizeof(*head), GFP_KERNEL);
	head->cid = 3;
	head->year = 2005;
	head->brand = "Wolkswagen";
	head->model = "Passat";
	head->color = "silver";
	head->motor_volume = 2;
	INIT_LIST_HEAD(&head->list);
	list_add_tail(&head->list, &car_list);

	printk(KERN_INFO "\nAll the cars are:\n");
	list_for_each_entry(ptr, &car_list, list) 
	{
		printk(KERN_INFO "Cars: CID %d, Year %d , Brand %s , Model %s , Color %s , Motor Volume %d \n", ptr->cid, ptr->year, ptr->brand, ptr->model, ptr->color, ptr->motor_volume);
	}
	printk(KERN_INFO "\nCars have been produced before 2000: \n");   
	list_for_each_entry(ptr, &car_list, list) 
	{
		if (ptr->year <= 2000) 
		{ 
			printk(KERN_INFO "Cars: CID %d, Year %d , Brand %s , Model %s , Color %s , Motor Volume %d \n", ptr->cid, ptr->year, ptr->brand, ptr->model, ptr->color, ptr->motor_volume); 
       	}
	}
	printk(KERN_INFO "\nCID of cars that color is silver:\n");   
	list_for_each_entry(ptr, &car_list, list) 
	{
		if (strcmp(ptr->color, "silver") == 0) 
		{  
            printk(KERN_INFO "Car: CID %d \n", ptr->cid);
		}
	}
	printk(KERN_INFO "\nBrand of car that model is Megane: \n");
	list_for_each_entry(ptr, &car_list, list) 
	{
		if (strcmp(ptr->model, "Megane") == 0) 
		{  
            printk(KERN_INFO "Brand of the car that model is Megane is %s \n", ptr->brand);
		}
	}
	printk(KERN_INFO "\nMotor Volume of car having CID 3 is set from 3 to 4: \n");	   
	list_for_each_entry(ptr, &car_list, list) 
	{
		if (ptr->cid == 3) 
		{
			ptr->motor_volume = 4;
			printk(KERN_INFO "Car: CID %d, Year %d , Brand %s , Model %s , Color %s , Motor Volume %d \n", ptr->cid, ptr->year, ptr->brand, ptr->model, ptr->color, ptr->motor_volume); 
       
		}
	}
	return 0;
}

void simple_exit(void) 
{
	struct  car_struct *ptr, *next;
	printk(KERN_INFO "Removing Module\n");
	list_for_each_entry_safe(ptr, next, &car_list, list) 
	{
		printk(KERN_INFO "Removing Car: %d %d  %s  %s  %s  %d \n", ptr->cid, ptr->year, ptr->brand, ptr->model, ptr->color, ptr->motor_volume); 
       	list_del(&ptr->list);
		kfree(ptr);
	}
}
module_init( simple_init );
module_exit( simple_exit );
MODULE_LICENSE("Dual MPL/GPL");
MODULE_DESCRIPTION(" You need to create, insert and remove module named cars into the Linux kernel for this question. You need to create cars.c file and Makefile file. The module is required to keep and use kernel linked-list data structure whose elements hold a car.");
MODULE_AUTHOR("Yusuf Ali KOYUNCU");

