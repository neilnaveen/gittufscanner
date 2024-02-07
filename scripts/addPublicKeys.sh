#!/bin/bash

REPO_PATH="$(pwd)/$1/$2"
KEYS_DIR="$(pwd)/keys"

# Check if the KEYS_DIR exists
if [ ! -d "$KEYS_DIR" ]; then
    echo "Keys directory not found: $KEYS_DIR" >&2
    exit 1
fi

cd "$REPO_PATH"

# Getting public keys from the tags
for tag in $(git tag); do
    # Extract the key id used to sign the tag
    key_id=$(git tag -v "$tag" 2>&1 | grep 'using RSA key' | awk '{print $NF}')

    if [ -n "$key_id" ]; then
        # Export the public key
        gpg --export --armor "$key_id" > temp_key.pub

        # Check if key export was successful
        if [ $? -eq 0 ]; then
            # Generate a unique filename for the key
            key_filename=$(echo "$key_id" | awk '{print substr($0, length($0) - 7)}')
            key_filepath="$KEYS_DIR/${key_filename}.pub"

            # Check if the key file already exists
            if [ ! -f "$key_filepath" ]; then
                # Move the key to the keys directory
                mv temp_key.pub "$key_filepath"
                echo "$key_filepath"
            else
                rm temp_key.pub
            fi
        else
            rm temp_key.pub
        fi
    fi
done

cd ../../