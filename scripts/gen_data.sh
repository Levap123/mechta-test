#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 <number_of_items>"
    exit 1
fi

num_items=$1
output_file="data.txt"

echo "" > $output_file

for ((i=0; i<num_items; i++))
do
    a=$((RANDOM % 21 - 10))
    b=$((RANDOM % 21 - 10))
    if [[ $i -eq 0 ]]; then
        echo -n "[{\"a\": $a, \"b\": $b}" >> $output_file
    else
        echo -n ",{\"a\": $a, \"b\": $b}" >> $output_file
    fi
done

echo "]" >> $output_file
