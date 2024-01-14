#!/bin/bash

for src in *.json
do
  printf "%s\t%d\n" ${src} $(egrep '"id":' ${src} | wc -l )
done
