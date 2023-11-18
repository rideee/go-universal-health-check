# Universal Health Check written in Go

## General idea
The general idea is to create a simple, fully configurable program (using JSON configuration files) that will check the operating status of running applications/processes/disk usage/etc. locally or on other servers using SSH. Each defined "validation task" would run asynchronously to reduce time, and the results would be displayed in FIFO order.