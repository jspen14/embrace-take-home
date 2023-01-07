########## INSTRUCTIONS ##########
RUN PROGRAM:
go run cmd/main.go data/[file_name].csv

RUN TEST SUITE:
go test ./...

########## EXPLANATION ##########
My philosophy in building this app was to make it as easy as possible to enhance. So, I used a few interfaces. If 
at some point I wanted to change the data retriever to read from a different file format or maybe a database, I would
simply need to make another implementation of the interface. 

As quoted in the prompt, the "complete data set is much larger than the sample." To ensure the program can run, I iteratively
read over the CSV so I wasn't loading the whole thing at once. Additionally, I applied a compression technique (only storing one 
session per user, per day) that was quite effective. Specifically, it resulted in a 50X compression. There are 5,598,398 data points
in the data file, and I was able to reduce that to 129,442.


########## PROMPT ##########
INSTRUCTIONS

You can find the data file to be processed at https://drive.google.com/drive/folders/1E9U-G7ySNo5jGxT8AtnQGsGofsVHsC5D?usp=sharing. 

You may use the language of your choice, and use of any of the language’s standard features or libraries is 
encouraged. If you have any questions about your choice of language, please ask before starting this exercise.


PROBLEM

In this exercise, we will be analyzing an application's user retention as we compute the number of users that user the 
application for a number of consecutive days.  The provided CSV file documents user activity for the application. Each
line in the file indicates the time of activity and the user associated with the activity. Determine how addictive the 
application is by computing how many consecutive days users that started using the application on a given day ended up 
using it for. In this exercise we will consider the date range 1/1/2021 to 1/14/2021 and you can assume that nobody used
the application on 12/31/2020 or 1/15/2021. 

INPUT AND OUTPUT

Your program should take one command line argument which is the path to the CSV file. 

Your program should print to stdout and should consist of one line of output for each day in the date 
range, with each line having the format. This should go to stdout.

DAY,USERS_PLAYING_1_DAY,USERS_PLAYING_2_DAYS,.....,USERS_PLAYING_14_DAYS

where DAY is an integer in the range from 1 to 14 and USERS_PLAYING_N_DAY(S) is a count of the people who started 
playing that day and then played for N consecutive days in a row. If they only played the day they started, and did not 
return the following day, this should be considered one day of consecutive usage. If a user returns to the application
after a gap in usage, they should be counted the same way as a user who is seen for the first time. The sample data provided covers the time range
1/1/2021 to 1/14/2021. The time in seconds from start of epoch to 1/1/2021 00:00:00 is 1609459200.

Input File Format
* CSV
* a timestamp and user ID pair on each line
* each line ends with a '\n'
* the timestamp is field 1 and the user ID is field 2
* timestamp is 32-bit int, indicating the number of seconds from start of
epoch (1/1/1970 00:00) relative to UTC
* user ID is 64-bit int
* the data in the file is sorted in ascending order of timestamps



Some things that may simplify the implementation
* You do not need to handle malformed input data. You can assume that
all the data in the input is consistent with the format specified above.
* You do not need to handle invalid paths passed as the arg to the
program

Provide the following once you have completed the coding exercise:
* A README with:
  * Any instructions needed to build (if necessary) and run the
application.
  * Any requirements or constraints that must be met to build and/or
run the program (e.g. platform, tool versions, language versions)
* The source code for your solution

EXAMPLE

In the following example, we track the activity of 5 users. 
```
1609459200,1
1609459260,1
1609459200,2
1609459200,3
1609459200,4
1609545600,1
1609545600,3
1609632000,1
1609632000,2
1609632000,3
1609718400,1
1609718400,2
1609804800,1
1609804800,5
```

The previous input file should print the following output to stdout:

```
1,2,0,1,0,1,0,0,0,0,0,0,0,0,0
2,0,0,0,0,0,0,0,0,0,0,0,0,0,0
3,0,1,0,0,0,0,0,0,0,0,0,0,0,0
4,0,0,0,0,0,0,0,0,0,0,0,0,0,0
5,1,0,0,0,0,0,0,0,0,0,0,0,0,0
6,0,0,0,0,0,0,0,0,0,0,0,0,0,0
7,0,0,0,0,0,0,0,0,0,0,0,0,0,0
8,0,0,0,0,0,0,0,0,0,0,0,0,0,0
9,0,0,0,0,0,0,0,0,0,0,0,0,0,0
10,0,0,0,0,0,0,0,0,0,0,0,0,0,0
11,0,0,0,0,0,0,0,0,0,0,0,0,0,0
12,0,0,0,0,0,0,0,0,0,0,0,0,0,0
13,0,0,0,0,0,0,0,0,0,0,0,0,0,0
14,0,0,0,0,0,0,0,0,0,0,0,0,0,0
```

Four users began streaks on day 1: Users 2 and 4 had a one day streak, User 3 had a 
three day streak, and User 1 had a five day streak. That means that the first line 
of the output should be

1,2,0,1,0,1,0,0,0,0,0,0,0,0,0

No users began a streak on day 2, so the second line in the output should be

2,0,0,0,0,0,0,0,0,0,0,0,0,0,0

User 3 started a two day streak on day 3, and they were the only user that started
a streak for that day, yielding the following line:

3,0,1,0,0,0,0,0,0,0,0,0,0,0,0

No users began a streak on day 4, so the fourth line is 

4,0,0,0,0,0,0,0,0,0,0,0,0,0,0

User 5 began a single day streak on day 5, so the fifth line is 

5,1,0,0,0,0,0,0,0,0,0,0,0,0,0

There was no further activity afterwards, so the rest of the output has zeroes.
