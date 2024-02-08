#!/bin/bash

# Navigate to the Git repository directory
cd "$(pwd)/$1/$2"

# Loop through each tag
for tag in $(git tag); do
    # Verify if the tag is signed
    git verify-tag "$tag" &> /dev/null
    if [ $? -eq 0 ]; then
        # The tag is signed, add the rule
        gittuf policy add-rule -k ../../keys/targets --rule-name "protect-$tag" --rule-pattern "git:refs/tags/$tag" --authorize-key gpg:E4F22BADDA796FD6FE27006733896060C7BD6180
    fi
done

# Navigate back to the original directory
cd ../../