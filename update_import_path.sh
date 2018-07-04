#! /bin/bash
#find ./ -name *.go |xargs  grep -nr github.com/hyperledger 
for i in $(find ./ -name *.go |xargs echo)
    do sed -i 's/github.com\/hyperledger/github.com\/sinochem\-tech/g' $i; echo $i 
done
