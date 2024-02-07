#!/bin/bash

# Navigate to the Git repository directory
cd $(pwd)/"$1"/"$2"

# Loop through each tag
for tag in $(git tag); do
   gittuf policy add-rule -k ../../keys/targets --rule-name "protect-$tag" --rule-pattern git:refs/tags/$tag --authorize-key gpg:E4F22BADDA796FD6FE27006733896060C7BD6180
done


cd ../../