#include <iostream>
#define CLASS_COUNT 3 // student numbers
using namespace std;

class Student
{
private:
    int no,midterm,final;
public:
    int getNo(){
        return no;
    }
    int getMidterm() {
        return midterm;
    }
    int getFinal(){
        return final;
    }
    int setNo(int _no){
        no = _no;
    }
    int setMidterm(int _midterm){
        midterm=_midterm;
    }
    int setFinal(int _final){
        final=_final;
    }
};
int main(){
    Student class_student[CLASS_COUNT];
    int no,midterm,final;
    float average;
    for (int i = 0; i < CLASS_COUNT; i++)
    {
        cout<<"Enter the stutent no, midterm and final: ";
        cin>>no>>midterm>>final;
        Student student;
        student.setNo(no);
        student.setMidterm(midterm);
        student.setFinal(final);
        class_student[i]=student;
    }
    for (int i = 0; i < CLASS_COUNT; i++)
    {
        average=0.4*class_student[i].getMidterm()+0.6*class_student[i].getFinal();
        cout<<"The average of student number "<<class_student[i].getNo()<<" is "<<average<<endl;
    }
    
    return 0;
}