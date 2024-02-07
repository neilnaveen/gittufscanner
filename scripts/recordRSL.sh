#!/bin/bash
# Rest of your script...

# Navigate to the Git repository directory
cd $(pwd)/"$1"/"$2"
# Loop through each tag
for tag in $(git tag); do
    gittuf rsl record "$tag"
done


cd ../../