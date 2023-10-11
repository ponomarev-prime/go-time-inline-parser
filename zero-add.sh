#!/bin/bash

input_string="$1"

if [ ${#input_string} -ge 19 ]; then
    echo "Input string is already 19 characters or longer."
else
    while [ ${#input_string} -lt 19 ]; do
        input_string="${input_string}0"
    done
    echo "$input_string"
fi