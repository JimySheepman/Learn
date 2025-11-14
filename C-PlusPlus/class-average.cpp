#include <iostream>
using namespace std;
void readNote(int[], int);
int main()
{
    int class_count, sum = 0;
    int top_point = 0;
    int *note;
    cout << "Enter the class count: ";
    cin >> class_count;
    note = new int[class_count];
    readNote(note, class_count);
    for (int i = 0; i < class_count; i++)
    {
        sum += note[i];
        if (note[i] > top_point)
        {
            top_point = note[i];
        }
    }
    cout << "Class average: " << sum / class_count << endl;
    cout << "Best note: " << top_point << endl;
    delete[] note;

    return 0;
}
void readNote(int class_list[], int count)
{
    cout << "Enter the notes: ";
    for (int i = 0; i < count; i++)
    {
        cin >> class_list[i];
    }
}