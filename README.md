# scg
For scg testing interview

# Getting Started
```
docker build -t scg . 
docker run --rm -p 5000:5000 -it scg
```
## Description
### Question 1 : Find x,y,z
> `/find-xyz`

Let's find relationship about this sequence number 

<img width="559" alt="Screen Shot 2563-04-07 at 21 57 04" src="https://user-images.githubusercontent.com/45779140/78686341-1e71ac00-791d-11ea-8eeb-f3cb4ae82483.png">

and find the stable range of number in 2nd level .
I see this relationship and then go to code for finding number in 2nd level and x,y,z value

### Question 2 : Find b,c
> `/find-bc`

A+B = 23
A + C = -21
Question provide value of A = 21
Let's move A to minus B and C .

### Question 3 : Short direction with Google API
I didn't do anything because GoogleAPI needs to pay for using it.

### Question 4 : Line Messaging API

> `/callback`

Line messaging API gives the time of user sending a message then I timed it in for loop . If system can't find the answer of user's question in 10 sec , System'll decide to answer like ' can't find any answer' to user.
