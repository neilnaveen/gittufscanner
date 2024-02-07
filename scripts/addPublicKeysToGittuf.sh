#!/bin/bash
# Rest of your script...

cd $(pwd)/"$1"/"$2"

gittuf trust add-policy-key -k ../../keys/root --policy-key ../../"$3".pub

cd ../../