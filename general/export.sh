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
CHANNELID=1124402182909857966
DATE=`date +%Y-%m-%d`
OUTPUT=$DATE--$CHANNEL.txt

echo Working channel $CHANNEL with ID $CHANNELID

rm -rf $FOLDER
mkdir $FOLDER

echo Start `date`

$EXEC export \
    --channel $CHANNELID \
    --token $TOKEN \
    --output $FOLDER/$OUTPUT \
    --format PlainText \
    > export.log 2>&1 

echo Done `date`
