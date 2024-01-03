#include <iostream>
#include <limits>

int main()
{
    std::cout<< "Here is some text.\n";
    std::cout << "Hello, world!"<< std::endl;

    std::cin.clear();
    std::cin.ignore(std::numeric_limits<std::streamsize>::max(),'\n');
    std::cin.get();

    return 0;
}