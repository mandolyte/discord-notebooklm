#!/bin/sh

###
### NOTE: this is for a forum style channel
### where each post is a thread.
###

BINDIR=$HOME/bin/DiscordChatExporter
CLI=DiscordChatExporter.Cli
EXEC=$BINDIR/$CLI
TOKEN=`cat $HOME/.discord`
FOLDER=./data
PWD=`pwd`
CHANNEL=`basename $PWD`
DATE=`date +%Y-%m-%d`
OUTPUT=$DATE--$CHANNEL.txt
# echo $OUTPUT
# exit

# echo Working on channel $CHANNEL

rm -rf $FOLDER
mkdir $FOLDER

# strip out the thread numbers from the channel's input
# echo Strip out the thread numbers from input.txt to thread_numbers.txt
cut -c 4-23 input.txt > thread_numbers.txt


# So I have to have a separate file for each thread. 
# echo Start `date`
for line in `cat thread_numbers.txt`
do
    #echo Working on thread $line
    $EXEC export \
        --channel $line \
        --token $TOKEN \
        --output $FOLDER/$line.txt \
        --format PlainText \
        > export.log 2>&1 
done

# combine them all
cd $FOLDER
cat *.txt > $OUTPUT

# to ensure that the input file is fresh, delete it after using
cd ..
rm input.txt thread_numbers.txt

# echo Done `date`
