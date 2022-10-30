## GoStoragePing
---

* This is a tool to ping **filesystems**.

What it basically does, is:

* Creates a file.
* Writes "ping" into the file.
* Measures how much it took.
* Reads the "ping" from the file.
* Measures how much it took.
* Deletes the file once you press CTRL+C.

What it lacks for now is a way to avoid OS cache, but I found it to measure pretty well when the system is under load. My guess it's that as im rewriting the file every time, the cache has to be refreshed, so every time is like the first time.  
I had options to run write or read separately, but because this cache issue, read was giving useless results.  

**Is not** a tool to measure it **precisely**, but instead, to detect when the system is under load, and how much the latency difference is compared to it's normal state.  

**IT READS AND WRITE CONTINOUSLY, SO DO NOT LEAVE IT RUNNING ON FLASH STORAGE FOR LONG PERIODS, IT COULD KILL YOUR SSD.** Although, the writen information is very small, you have been warned.  

I wrote this as a way to measure latency on our samba servers at work.
