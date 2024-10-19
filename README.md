# README

## IDEAS

- add a counter to show progress
    
## What's Here

There are folders which are named after the Discord channel. In each are data and scripts used to export the chat log.

## Process 

### Forum Style Channels

Since each post in a forum is a thread, I have to obtain all the thread numbers for the forum and export them individually. Here are the steps:

First, use the `get_active_threads.sh` script to obtain the active posts from all the channels. The output for this will be `input.txt`. 

Sample output:

```
$ ./get_active_threads.sh 
Start at Sat Oct 19 09:33:24 AM BST 2024
Done at Sat Oct 19 09:33:47 AM BST 2024
```

Second, open the file and locate the forum channel of interest and copy/paste those lines into the channel subfolder with the same name, namely, "input.txt"

Third, run the `export.sh` script to loop thru the thread number file, exporting each chat.

The script creates a log file named `export.log` and will concatenate all the exported chats into one text file named `Active_${CHANNEL}.txt`.

### Regular Channels

TBS
