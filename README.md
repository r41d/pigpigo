
 gpigo - a raspberry pi (2) gpio go wrapper
============================================

This encapsulates the wiringPi C library (see [https://wiringPi.com]).


**Note: gpigo is in early dev stage - Use on your own risk - Do _NOT_ use in production environments!**
**Note: gpigo is intended to run on a raspberry pi or a raspberry pi 2. (Cross)building on other architectures may lead to unpredictable results.**

Build:
1. go get the repo:

	go get -u -d code.bitsetter.de/tk/gpigo

   Don't be confused about the repo url. This will redirect to this github repo.

2. fetch and build dependencies:

	go generate

   This fetches and builds (but not installs) wiringPi.

3. Install gpigo

	go install

   You're ready to go ;)

