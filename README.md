# Ascii Art Justify

## Description

This project is an extension of a previous project going by a similar name. The initial project, ascii-art, was a project that was used to print text data in an artistic format. In it, there were a few files we referred to as banner files that contained graphic representations of printable ascii characters. When the user needed to print their data on the console, the program would instead print the graphic representation of the same. 

This project is an extension by the virtue that it offers the user the ability to use an optional flag when invoking the program. This flag is then used to decide how the text would appear in the console. The new concepts introduced is alignment. The user now has the option to request the text to be aligned in a particular way.


## Features
- Converts strings into ASCII art representations.
- Supports multiple banner formats including `shadow`, `standard`, and `thinkertoy`.
- Support for aligning the graphic representation on the terminal


#### Limitations
- The program is designed to work with printable ASCII characters only.
- Characters outside the range of space (` `) to tilde (`~`) are not supported and may not render correctly.

#### Prerequisites
- Go programming language
- ASCII banner files in the correct format

#### Banner files [BANNER]

The project also allows the user to include the banner file name that they need to be used to display the art. Currently we only accept three banner files which are: 

  + [`shadow`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
  + [`standard`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
  + [`thinkertoy`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)



#### Alignment [OPTION]

 These include and limited to:

**NOTE: The example are not to be taken literary, but as a rough estimate of how the text should appear**

+ left
+ right
+ center
+ justify

A random example is the alignment: center

```bash

                                                   _              _   _          
                                                  | |            | | | |         
                                                  | |__     ___  | | | |   ___   
                                                  |  _ \   / _ \ | | | |  / _ \  
                                                  | | | | |  __/ | | | | | (_) | 
                                                  |_| |_|  \___| |_| |_|  \___/  
                                                                                 
                                                                                 

```


## Usage

The format has to be (replace the square bracket with the specific requirement): 

go run . [OPTION] [STRING] [BANNER]

**OPTION** allows you to specify the textual alignment

**STRING** is the actual string whose representation will be printed

**BANNER** is also an optional argument to switch between the banner files to be used

**NOTE:** While the OPTION and BANNER are totally optional, in order to get the artistic representation, you will need to provide the STRING argument: 

go run . [STRING]

The following is an example of how to do so:

##### installation

```bash
git clone https://learn.zone01kisumu.ke/git/bnyatoro/ascii-art-justify.git
cd ascii-art-justify
clear
```
##### run the program

```bash
go run . --align=left "Hello"
```

Additionally, to switch between banner files, you can provide an additional argument which has to be the last argument

```bash
go run . --align=center "Hello" "shadow"
```


## Contributions

#### contibutors

- [sfana](https://learn.zone01kisumu.ke/git/shfana)
- [bnyatoro](https://learn.zone01kisumu.ke/git/bnyatoro)
- [anoduor](https://learn.zone01kisumu.ke/git/anoduor)

#### To contribute

Go to the repository at
[ascii-art-justify](https://learn.zone01kisumu.ke/git/bnyatoro/ascii-art-justify) and fork the repository. Clone your fork locally and make any changes you need to be made. After completing all git processes and pusing to your fork, you can issue a pull request and we assure you that your contributions wiil be taken into consideration.


