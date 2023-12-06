#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

# Assign the day parameter
day=$1

# Create the folder if it doesn't exist
folder="day_$day"
mkdir -p "$folder"

# Define the URL
url="https://adventofcode.com/2023/day/$day"
url_input="https://adventofcode.com/2023/day/$day/input"

# Define the output file
output_file="$folder/subject.md"


curl --cookie "session=$AOC_SESSION" -o "$output_file" "$url"
# Use curl to fetch the content and save it to the output file
curl --cookie "session=$AOC_SESSION" -o "$folder/input.txt" "$url_input"

echo "Subject for Day $day has been saved to $output_file"
