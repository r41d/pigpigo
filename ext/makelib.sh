#!/bin/bash
if [ ! -d "ext/wiringPi" ]
then
  git clone git://git.drogon.net/wiringPi ext/wiringPi > /dev/null
fi

pushd ext/wiringPi > /dev/null
git pull > /dev/null
popd > /dev/null

make -j5 -C ext/wiringPi/wiringPi static 

