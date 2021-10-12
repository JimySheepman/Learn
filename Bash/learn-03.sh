#!/usr/bin/sh

# ! case statement
if [ $# -ne 3 ]; then
    echo "Enter the 3 args!!"
    exit 1
fi
num1=$1 num2=$2
operator=$3

case $operator in
"+")
    echo "Result: $(expr $num1 + $num2)"
    ;;
"-")
    echo "Result: $(expr $num1 - $num2)"
    ;;
"*")
    echo "Result: $(expr $num1 * $num2)"
    ;;
"/")
    echo "Result: $(expr $num1 / $num2)"
    ;;
"%")
    echo "Result: $(expr $num1 % $num2)"
    ;;
*)
    echo "Error..!"
    ;;
esac

# ! array

os=("ubuntu","windows",'manjaro',"FreeBSD")
os[4]='mac' # cahnge index

unset os[0] # delet index

echo "${os[@]}"  # print all item
echo "${os[0]}"  # print select indet
echo "${!os[@]}" # print array index
echo "${#os[@]}" # print array len

# ! while

n=1

while [ $n -le 10 ]; do
    echo "$n"
    # n=$(expr $n + 1)
    # n=$(( n + 1 )) (( n++ ))
    sleep 1
done

#-----------ikinci örnek------------

while (($n <= 10)); do
    echo $n
    ((n++))
done

#-----------üçüncü örnek------------
# "read" komutu ile döngü kurma
while read line; do
    echo $line
done <test3.sh

#-----------dörtüncü örnek------------

cat test3.sh | while read line; do
    echo $line
done

# ! Until loop

n=1
until [ $n -gt 10 ]; do
    echo $n
    ((n++))
done

# ! For loop

for myVar in 1 2 3 4 5; do
    echo $myVar
done

#----------------------------

for myVar in {1..10..2}; do
    echo $myVar
done

#----------------------------

for ((i = 0; i < 5; i++)); do
    echo $i
done

#! /bin/bash

for cmd in ls pwd date; do
    echo "--------------$cmd--------------"
    $cmd
done

for item in *; do
    if [ -d "$item" ]; then
        echo "$item"
    fi
done
