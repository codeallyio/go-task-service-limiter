# Service limiter for free-tier users

## Background

You are the owner of a cloud-processing service. You provide new
users to try your service with the free tier. It is supposed to give
each of them an access to the service with 10 seconds of free
processing time. If the user is using premium mode, they should be
able to omit these limits.


## TODO

### Main goal

Limit every request from free-tier users to maximum 10 seconds.

### Extra goal

Limit all accumulated requests from each free-tier user to maximum 10 seconds.


### Good Luck!
