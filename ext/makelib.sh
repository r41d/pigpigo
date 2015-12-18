#!/bin/bash
if [ ! -d "ext/wiringPi" ]
then
  git clone git://git.drogon.net/wiringPi ext/wiringPi
fi

pushd ext/wiringPi
git pull
popd

make -j5 -C ext/wiringPi/wiringPi static
cp ext/wiringPi/wiringPi/*.a ext/

