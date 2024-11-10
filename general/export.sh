#!/bin/sh

###
### NOTE: this is for a forum style channel
### where each post is a thread.
###

BINDIR=$HOME/Downloads/zz
CLI=DiscordChatExporter.Cli
EXEC=$BINDIR/$CLI
TOKEN=`cat $HOME/.discord`
FOLDER=./data
PWD=`pwd`
CHANNEL=`basename $PWD`
CHANNELID=1124402182909857966

echo Working channel $CHANNEL with ID $CHANNELID

rm -rf $FOLDER
mkdir $FOLDER

echo Start `date`

$EXEC export \
    --channel $CHANNELID \
    --token $TOKEN \
    --output $FOLDER/$CHANNEL.txt \
    --format PlainText \
    > export.log 2>&1 

echo Done `date`
