# Go-Shell

This is our shell project for VCU's operating systems class (CMSC 312)

It is written in the popular new Go Language. 

## Install
todo



-----------------------------------------------------------------
For group members, we can deploy to VCU's server with something like sftp.
Assuming you use linux or MacOS

From a terminal,

sftp -P 20023 username@servername.vcu.edu

you can upload the entire repository by simply 
$put -r local_folder

Some useful commands in sftp:
lls // list contents in local directory
ls // list contents in remote directory
lcd // change directory locally
cd // change directory remotely
and so on...

The server is running CentOS, so if you are on Mac or Windows, your executables created with '$go install' will probably not work unless you change your enviornment variables locally to compile for linux.
