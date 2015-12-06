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

To run a list of commands and programs from a batchfile, first put them into a file, one per line and run as:

    go run shell.go batchfile.gosh
  
You can even run several files sequentially. 

    go run shell.go batch1.gosh batch2.gosh
  
  



