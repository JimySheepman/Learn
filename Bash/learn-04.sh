#!/usr/bin/sh

select name in ali ayse fatma salih; do
    echo "$name secildi"
done

#----------------ikici örnek-----------

select name in ali ayse fatma salih; do
    case $name in
    ali)
        echo "ali secildi."
        ;;
    *)
        echo "Hata: lütfen 1..4 arasında bir sayı girin"
        ;;

    esac
done

# ! function

function hello() {
    echo "Hello"
}

quit() {
    exit 0
}

function print() {
    echo $1 # birinci argumanı yazar
}

# quit
print Hello
hello
