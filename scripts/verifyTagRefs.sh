#!/bin/bash

# Navigate to the Git repository directory
cd "$(pwd)/$1/$2"

for tag in $(git tag); do
  # Verify if the tag is signed
  git verify-tag "$tag" &> /dev/null
    if [ $? -eq 0 ]; then
       gittuf verify-tag $tag
    fi
done

cd ../../
