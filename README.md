# Go-Shell

This is our shell project for VCU's operating systems class (CMSC 312)

It is written in the popular new Go Language which features built-in support for concurrent processes.  

## Install
To install to your current directory, run:

    git clone https://github.com/briansimoni/Go-Shell.git
    
or, if you have a `.zip` file:

    unzip -d  go_shell.zip go_shell

## Usage
To use, first `cd` into the `go_shell` directory. Then run:

    $ go run shell.go
  
You should see  something like:

    /path/to/directory/go_shell/ $ 

Type `help` to see a list of built in commands. Or run a program such as:

    /path/to/directory/go_shell/ $ ping -c3 localhost
    PING localhost (127.0.0.1) 56(84) bytes of data.
    64 bytes from localhost (127.0.0.1): icmp_seq=1 ttl=64 time=0.038 ms
    64 bytes from localhost (127.0.0.1): icmp_seq=2 ttl=64 time=0.040 ms
    64 bytes from localhost (127.0.0.1): icmp_seq=3 ttl=64 time=0.035 ms
    --- localhost ping statistics ---
    3 packets transmitted, 3 received, 0% packet loss, time 1998ms
    rtt min/avg/max/mdev = 0.035/0.037/0.040/0.007 ms

To run a list of commands and programs from a batchfile, first put them into a file, one per line and run as:

    go run shell.go batchfile.gosh
  
You can even run several files sequentially. 

    go run shell.go batch1.gosh batch2.gosh
  
  



