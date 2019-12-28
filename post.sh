#!/bin/sh

while true; do 
    NOW=$( date "+%Y-%m-%dT:%H:%M:%S+01:00")
    R=$(( ( RANDOM % 2 )  + 1 ))
    if [ $R == 1 ]
        then echo 'time="$(NOW)" level=info msg="job successfully posted" http status=200 id=4867' >> /tmp/job-postings.log
    fi
    if [ $R == 2 ]
        then echo 'time="$(NOW)" level=error msg="job failed" http status=400 id=4867' >> /tmp/job-postings.log
    fi
    sleep 2
done