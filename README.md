# README

## Overview

This project extracts the "active" posts/threads from the Discord server for Google's NotebookLM project. It utilizes the tooling described below in Appendix B. This project consists of scripts use the DiscordChatExporter tool and manipulate the output.

Feel free to ask questions or raise issues at:
https://github.com/mandolyte/discord-notebooklm/issues

## Wishlist

- add a counter to show progress
    
## What's Here

There are folders which are named after the Discord channel. In each are data and scripts used to export the chat log.

## Process 

### Automation

As of 2025-03-15, a Go language script has been written to automate the detailed steps below. It is located in the folder `main`. Run the script `sh runAll.sh` to use the automation.

Example execution will be similar to the below:

```
$ sh runAll.sh
Extractions Started at Sat Mar 15 08:42:30 PM EDT 2025
Get active threads on Discord server
Start at Sat Mar 15 08:42:31 PM EDT 2025
Done at Sat Mar 15 08:42:42 PM EDT 2025

Look for bugs section
Look for feature-requests section
Start Discord server channel extracts
Worker 5 started: runUseCases.sh
Worker 3 started: runGeneral.sh
Worker 1 started: runBugs.sh
Worker 2 started: runFeatureRequests.sh
Worker 4 started: runIntros.sh
Worker 1 finished: runBugs.sh in 20m14.218398936s
... output:

Worker 2 finished: runFeatureRequests.sh in 31m37.858878637s
... output:

Worker 5 finished: runUseCases.sh in 46m40.676076769s
... output:
Working channel use-cases with ID 1124403655819415592
Start Sat Mar 15 08:42:42 PM EDT 2025
Done Sat Mar 15 09:29:22 PM EDT 2025

Worker 4 finished: runIntros.sh in 1h12m23.175219839s
... output:
Working channel intros with ID 1124438716325306419
Start Sat Mar 15 08:42:42 PM EDT 2025
Done Sat Mar 15 09:55:05 PM EDT 2025

Worker 3 finished: runGeneral.sh in 1h19m1.404259296s
... output:
Working channel general with ID 1124402182909857966
Start Sat Mar 15 08:42:42 PM EDT 2025
Done Sat Mar 15 10:01:43 PM EDT 2025

All workers completed.
Extraction took 1h19m12.117518837s to run
Extractions Completed at Sat Mar 15 10:01:43 PM EDT 2025
$ 

```

### Forum Style Channels

Since each post in a forum is a thread, I have to obtain all the thread numbers for the forum and export them individually. 

*For NLM, there are just two forum style channels that I am doing: Bugs and Feature Requests. So the following only applies to these two channels.*

Here are the steps:

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

For regular channels it isn't required to fetch each message individually. The `use-cases` folder is a regular channel. The channel id is required and all the messages will be exported in one command. See script for details.

## Appendix A - List of Server Channels and IDs

Below shows the command and the output.

```
$ ./DiscordChatExporter.Cli channels --guild 1124402182171672732 --token <t>

1286074990407450777 | student-hub
1182376885280329829 | Mods / system-updates
1167885619989590046 | Read-Only / 🏁-start-here
1124403332262412368 | Read-Only / 📕-guidelines
1182376564525113484 | Read-Only / 📣-announcements
1294057633334165596 | Read-Only / 🚥-status
1173364829810085931 | Discussions / 🐞-bugs
1124438716325306419 | Discussions / 👋🏽-intros
1124402182909857966 | Discussions / 💬-general
1124403655819415592 | Discussions / 📇-use-cases
1173362666266443876 | Discussions / 📋-feature-requests
1280639145785425974 | Discussions / educator-ambassadors
$ 
```

The discord server ID ("guild") may easily be found by inspecting the URL while using the web version of discord. The above for example, came from: 

```
https://discord.com/channels/1124402182171672732
```

## Appendix B - Links and References

This project relies on the installation of the Discord Chat Exporter project. Find this at:

The documentation is at:
https://github.com/Tyrrrz/DiscordChatExporter/blob/2.43.3/.docs/Using-the-CLI.md

The documentation also has the method for which to obtain your token needed to execute the API commands.


