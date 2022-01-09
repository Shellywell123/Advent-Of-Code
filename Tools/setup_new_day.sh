#!/bin/bash
# small script to help with the setup of each AOC puzzle

# user input year

echo "What year are you working on: (input year e.g 2021)"
read y

# user input day num x/25
echo "What number day do you want to setup: (input no e.g 5)"
read x

url=https://adventofcode.com/${y}/day/${x}/input

#if one char add 0 to front to maintain file order
len=`expr length "${x}"`
if [[ ${len}=1 ]];
then 
num="0${x}"
else
    num=${x}
fi

# user input file ext 
echo "What Language will you be using: (input file ext e.g python=py C++=cpp)"
read ext

# create dir
mkdir Day${num}

# create files
touch Day${num}/Part_1.${ext} Day${num}/Part_2.${ext} Day${num}/inputs.txt Day${num}/tests.txt

# populate inputs.txt (currently requires login)
curl -o Day${num}/inputs.txt ${url}