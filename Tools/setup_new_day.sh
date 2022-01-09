#!/bin/bash
# small script to help with the setup of each AOC puzzle

# user input day num x/25
echo "What number day do you want to setup: (recommend 0X format if less than 10)"
read x

# user input file ext 
echo "What Language will you be using: (input file ext) python=py C++=cpp: "
read ext

# create dir
mkdir Day${x}

#create files
touch Day${x}/Part_1.${ext} Day${x}/Part_2.${ext} Day${x}/inputs.txt Day${x}/tests.txt