#!/bin/bash

whiptail --msgbox "Bu bir uyari mesaji" 10 50

whiptail --msgbox "Bu bir uyari mesaji" 10 50 --ok-button="Tamam"

#! yes/no

whiptail --title "Anket" --yesno "GNU/Linux'u özgür buluyor musunuz?" \
    --yes-button Evet --no-button Hayır 20 60
if [ $? -eq 0 ]; then
    whiptail --msgbox "Çok haklısınız!" --ok-button Tamam 10 40
else
    whiptail --msgbox "Yanılıyor olabilirsiniz" --ok-button Tamam 10 40
fi

#! inputbox

YENI_HOSTNAME=$(
    whiptail --inputbox "Yeni Hostname Giriniz:" \ 
    8 60 $(hostname) --title "Hostname Değiştirme" --ok-button Tamam \
        --cancel-button İptal 3>&1 1>&2 2>&3
)

exitstat=$?

if [ $exitstat = 0 ] && [ ! -z $YENI_HOSTNAME ]; then
    hostname $YENI_HOSTNAME
else
    echo "işlem iptal edildi"
fi

#! menu

SECIM=$(whiptail --title "Programlama Menüsü" \
    --menu "Bir dil seçin" 17 50 0 \
    "Python" "Guido van Rossum" \
    "C" "Dennis M. Ritchie" \
    "Perl" "Larry Wall" \
    "PHP" "Rasmus Lerdorf" 3>&1 1>&2 2>&3)

if [ $? = 0 ] && [ ! -z $SECIM ]; then
    echo "$SECIM dilini seçtiniz"
else
    echo "Dil seçmediniz"
fi

#! checklist

SECIM=$(whiptail --title "Programlama Menüsü" --checklist \
    "Kullandığınız Dilleri Seçin" 10 60 5 \
    "Python" "Guido van Rossum" ON \
    "C" "Dennis M. Ritchie" OFF \
    "Perl" "Larry Wall" OFF \
    "PHP" "Rasmus Lerdorf" ON 3>&1 1>&2 2>&3)

if [ $? -eq 0 ]; then
    echo "Seçtiğiniz diller:"
    echo "$SECIM"
else
    echo "Dil seçmediniz"
fi

#! radiolist

SECIM=$(whiptail --title "Programlama Menüsü" --radiolist \
    "Kullandığınız Dilleri Seçin" 10 60 5 \
    "Python" "Guido van Rossum" OFF \
    "C" "Dennis M. Ritchie" ON \
    "Perl" "Larry Wall" OFF \
    "PHP" "Rasmus Lerdorf" OFF 3>&1 1>&2 2>&3)

if [ $? -eq 0 ]; then
    echo "Seçtiğiniz diller:"
    echo "$SECIM"
else
    echo "Dil seçmediniz"
fi
