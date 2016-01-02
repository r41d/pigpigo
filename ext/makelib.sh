#!/bin/bash
if [ ! -d "ext/pigpio" ]; then
	git clone https://github.com/joan2937/pigpio ext/pigpio
fi

pushd ext/pigpio
git pull
popd

make -j5 -C ext/pigpio
#make install
