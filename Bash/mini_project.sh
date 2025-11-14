630 630 630 #! /bin/bash
args_arr1=("$@")
echo "The shell name is $BASH"
echo "The shell version is $BASH_VERSION"
echo "Our home directory is $HOME "
echo "Current working directory is $PWD"
echo "----------------------------------"
echo "Number of arguments is $#"
echo "----------------------------------"
echo "array1 size is ${#args_arr1[@]}"
echo "array1 is [ ${args_arr1[@]} ] "
echo "----------------------------------"
args_arr2=("${args_arr1[@]:4}")
echo "array2 size is ${#args_arr2[@]}"
echo "array2 is [ ${args_arr2[@]} ] "
echo "Please choose an index of arry2 to remove an element"
read index_number
unset args_arr2[index_number]
echo "array2 after removing element at index  $index_number is [ ${args_arr2[@]} ] "
echo "after removing element at index $index_number array2 size is ${#args_arr2[@]} "
echo "----------------------------------"
for cmd in "${args_arr2[@]}"; do
    echo "---------------$cmd---------------"
    $cmd
done
echo "----------------------------------"
num1=${args_arr1[0]}
opt=${args_arr1[1]}
num2=${args_arr1[2]}
fact=1
for ((i = 2; i <= num1; i++)); do
    fact=$((fact * i))
done
echo "Factorial of $first_number is $fact "
echo "----------------------------------"
if [[ $opt == "+" ]]; then
    echo "The result of $fact $opt $num2 is $((fact + num2))"
elif [[ $opt == "-" ]]; then
    echo "The result of $fact $opt $num2 is $((fact - num2))"
elif [[ $opt == "*" ]]; then
    echo "The result of $fact $opt $num2 is $((fact * num2))"
elif [[ $opt == "/" ]]; then
    echo "The result of $fact $opt $num2 is $((fact / num2))"
elif [[ $opt == "%" ]]; then
    echo "The result of $fact $opt $num2 is $((fact % num2))"
else
    echo "illegal sing"
fi
echo "----------------------------------"
if [ -e ${args_arr1[3]} ]; then
    echo "${args_arr1[3]} file found"
else
    echo "${args_arr1[3]} not found"
fi
echo "----------------------------------"
