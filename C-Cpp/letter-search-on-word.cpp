#include <iostream>

using namespace std;

int main()
{
    char word_1[3];
    char word_2[6];
    int res=0;
    cout<<"Enter first word: ";
    cin>>word_1;
    cout<<"Enter second word: ";
    cin>>word_2;
    
    for (int i = 0; i < 3; i++)
    {
        for (int j = 0; j < 6; j++)
        {
            if (word_1[i] == word_2[j])
            {
                res+=1;
                break;
            }
        }
        
    }
    if (res==3)
    {
         cout<< "All letters of the word "<<word_1<<" are found in the word "<<word_2<<endl;
    }
    else
    {
        cout<< "Not all letters of the word "<<word_1<<" are found in the word "<<word_2<<endl;
    }
    
    
    
    
    
    return 0;
}