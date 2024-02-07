#!/bin/bash
# Rest of your script...

# Navigate to the Git repository directory
cd "$(pwd)/$1/$2"

for tag in $(git tag); do
  gittuf verify-tag $tag
done

cd ../../
