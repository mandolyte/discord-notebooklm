#!/bin/sh

BINDIR=$HOME/Downloads/zz
CLI=DiscordChatExporter.Cli
EXEC=$BINDIR/$CLI
TOKEN=`cat $HOME/.discord`

echo Start at `date`
$EXEC channels \
    --guild 1124402182171672732 \
    --include-threads active \
    --token $TOKEN \
    > input.txt

echo Done at `date`
