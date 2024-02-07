#!/bin/bash
# Rest of your script...

cd $(pwd)/"$1"/"$2"

git remote set-url origin https://github.com/$1/$2.git



gittuf trust init -k ../../keys/root
gittuf trust add-policy-key -k ../../keys/root --policy-key ../../keys/targets.pub
gittuf policy init -k ../../keys/targets



cd ../../