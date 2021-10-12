#!/usr/bin/sh

# ! run comment
echo "Shell name $BASH"
echo "Shell version $BASH_VERSION"
echo "Main directory $HOME"
echo "Working directory $PWD"

# ! output
name="merlins"
echo $name
echo "username = $name"

# ! input,
echo "Enter the multiple  names: "
read name_2 name_3 name_4
echo "usernames list: $name $name_2 $name_3 $name_4"

# ! input options
read -p 'username: ' user_name
read -sp 'password: ' user_pass
echo "Enter the name: ";read -a names
echo "Names: ${names[0]},${names[1]},${names[2]}"
echo "Username is $user_name"
echo "Password is $user_pass"

# ! example
echo $1 $2 $3 # give a args in terminal
echo "comment = $0"
args=("$@") # create list
echo ${args[1]} ${args[2]} # select index list and print screen
echo $@ # print all list
echo $# # count of list