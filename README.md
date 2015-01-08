4 bits to do to compile and run go code:

1) set your directory structure like so:
  <basedir> go\
	    ]_\bin	// into where binary is installed
	    ]_\pkg	// (dont know what this is for)
	    ]_\src	// where your source packages go
	       ]_\<packagename>\<packagename>.go

2) set GOPATH to <basedir>\go

3) cd to <basedir>\go\src\<packagename>, and type

 PROMPT> go install

4) if the install is successful, execute the binary from anywhere, eg:

 PROMPT> <basedir>\go\bin\<packagename>


for example:
1. dir structure:
widst001:go_big_or_go_home wid$ ls -R phrasebook/
go

phrasebook//go:
bin	pkg	src

phrasebook//go/bin:
hello

phrasebook//go/pkg:

phrasebook//go/src:
hello		interface

phrasebook//go/src/hello:
hello.go

phrasebook//go/src/interface:
interface.go
widst001:go_big_or_go_home wid$ 


2. set GOPATH (mac terminal)
widst001:go_big_or_go_home wid$ export GOPATH=/Users/wid/workspace/git/GITHUB/go_big_or_go_home/phrasebook/go
widst001:go_big_or_go_home wid$ echo $GOPATH
/Users/wid/workspace/git/GITHUB/go_big_or_go_home/phrasebook/go


3. compile
widst001:go_big_or_go_home wid$ cd phrasebook/go/src/hello/

widst001:hello wid$ go install

4. execute

widst001:hello wid$ ../../bin/hello 
hello... framingham?
widst001:hello wid$ 


more here:
https://groups.google.com/forum/#!topic/golang-nuts/J5NP0Y6nnYE
