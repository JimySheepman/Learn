#!/usr/bin/sh

# ! if else parameter
# -eq : equals
# -what : not equal
# -gt : big
# -ge : greater than or equal to
# -lt : less
# -le : less than or equal to
# < : is less (in double parenthesis)
# <= : less than or equal (in double parenthesis)
# > : greater (in double parenthesis)
# >= : greater than or equal to (in double parenthesis)

# ! example 1
count=10
if [ $count -eq 9 ]; then
    echo "false"
fi
if (($count > 9)); then
    echo "count bigger then 9"
fi

# ! example 2
word="abc"

if [ $word == "abc" ]; then
    echo "word is equel 'abc'"
fi

# ! example 3
word=a

if [[ $word < "b" ]]; then
    echo "Condition is true"
else
    echo "Condition is false"
fi

# ! example 4

word=a

if [[ $word == "b" ]]; then
    echo "word = 'b'"
elif [[ $word == "a" ]]; then
    echo "word = 'a'"
else
    echo "word != a or b"
fi

# ! example 5
if [ $# -ne 3 ]; then
    echo "Enter the 3 args!!"
    exit 1
fi

num1=$1
num2=$2
operator=$3

if [[ $operator == "+" ]]; then
    echo "Result: $((num1 + num2))"
elif [[ $operator == "-" ]]; then
    echo "Result: $((num1 - num2))"
elif [[ $operator == "*" ]]; then
    echo "Result: $((num1 * num2))"
elif [[ $operator == "/" ]]; then
    echo "Result: $((num1 / num2))"
elif [[ $operator == "%" ]]; then
    echo "Result: $((num1 % num2))"
else
    echo "Erroor!!!"
fi

# ! File handling

# ! example 1
echo -e "Enter the file name: \c"
read fileName

if [ -e $fileName ]; then
    echo "$fileName is founded!"
else
    echo "$fileName is not founded..!"
fi

# ! example 2
echo -e "Enter the file name: \c"
read fileName

if [ -f $fileName ]; then
    if [ -w $fileName ]; then
        echo "write a text. Exit prees CTRL-D."
        cat >>$fileName
    else
        echo "Permission denied..!"
    fi
else
    echo "$fileName does not exist."
fi
