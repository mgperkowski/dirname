# dirname
A simple library to get the relative directory path similar to __dirname in Node.js

##Installation

```bash
go get github.com/mgperkowski/dirname
```

##Features

Offers two methods to return either the directory that the executable is stored in or the directory the process was called from.

GetExecutableDirname returns the working directory when using "go run ..."

##Example Usage

```go

package main

import(
    "github.com/mgperkowski/dirname"
)

func main{
    executableDirname, err := dirname.GetExecutableDirname() //This gets the directory that the executable is located in -- useful for interacting with files located in the executable directory

    workingDirname, err := dirname.GetWorkingDirname() //Returns the directory the process was called from -- useful for interacting with files in the current working directory

    ...
}

```

