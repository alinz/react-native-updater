#!/bin/bash

usage()
{
cat << EOF
usage: $0 options

React-Native-Updater Build tool

OPTIONS:
   -u      Relative Updater App Folder
   -a      Relative Main App Folder
   -o      Absolute Output Folder
EOF
}

UPDATER=
MAIN=
OUTPUT_IOS=$(pwd)/ios
while getopts ":u:a:ios:" OPTION
do
     case $OPTION in
         u)
             UPDATER=$(pwd)/$OPTARG
             ;;
         a)
             MAIN=$(pwd)/$OPTARG
             ;;
         ios)
             OUTPUT_IOS=$OPTARG
             ;;
         *)
             usage
             exit
             ;;
     esac
done

cd $UPDATER; react-native bundle --out $OUTPUT_IOS/update.jsbundle
cd $MAIN; react-native bundle --out $OUTPUT_IOS/main.jsbundle
