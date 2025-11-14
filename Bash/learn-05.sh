#!/bin/bash
# learn-05.sh <indirilecek paket adı>

# Debian Yukleme Fonksiyonu
debian_yukle () {
    # program zaten yuklu mu?
    dpkg -l | grep $1 > /dev/null
    if [ $? -eq 1 ]; then
         echo "Yukleme basliyor..."
         apt-get install $1
    else
        echo "$1 zaten yuklu"
    fi

    return 0
}

redhat_yukle () {
    # program zaten yuklu mu?
    rpm -qa | grep $1 > /dev/null
    if [ $? -eq 1 ]; then
        echo "Yukleme basliyor..."
        yum install $1
    else
        echo "$1 zaten yuklu"
    fi

    return 0
}

# Root yetkiniz var mi?
if [ $(id -u) -ne 0 ]; then
    echo "root yetkisi ile calistirilmali"
    exit 1
fi

# Parametre sayisi kontrol ediliyor
if [ $# -lt 1 ]; then
    echo "Parametre vermediniz"
    echo "Kullanım: yukle.sh program-adi"
    exit 1
fi

# Eger sistem Debian ise
if [ -f /etc/debian_version ]; then
    debian_yukle $1

# Eger sistem Red Hat ise
elif [ -f /etc/redhat-release ]; then
    redhat_yukle $1

# Tanimadigimiz sistem
else
    echo "Tanimadigimiz bir sistem?"
    exit 1
fi

exit 0