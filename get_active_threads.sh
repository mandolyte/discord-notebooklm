#!/bin/sh

BINDIR=$HOME/bin/DiscordChatExporter
CLI=DiscordChatExporter.Cli
EXEC=$BINDIR/$CLI
TOKEN=`cat $HOME/.discord`

echo Start at `date`
$EXEC channels \
    --guild 1124402182171672732 \
    --include-threads active \
    --token $TOKEN \
    > input.txt 2>&1

echo Done at `date`
