
# ASCII Art

## Overview
The ASCII Art is a Go program designed to convert strings into graphical representations using ASCII characters. It supports optional color formatting using the `--color` flag. The program can handle input with numbers, letters, spaces, special characters, and newline (`\n`) characters.

## Installation
Clone the repository and navigate to the project directory:

```bash
git clone https://learn.zone01kisumu.ke/git/bnyatoro/ascii-art
cd ascii-art
```

Run the program as shown in the usage section.

## Features
- Converts strings into ASCII art representations.
- Each ASCII character has a height of 8 lines.
- Characters are separated by a newline character.
- Supports multiple banner formats including `shadow`, `standard`, and `thinkertoy`.
- Allows output customization using the `--color` flag.


## Limitations
- The program is designed to work with printable ASCII characters only.
- Characters outside the range of space (` `) to tilde (`~`) are not supported and may not render correctly.

## Prerequisites
- Go programming language
- ASCII banner files in the correct format

## Usage
To use the ASCII Art Generator, compile and run the program with a string argument. Here are some examples:

### Basic Usage
```bash
go run . "Hello" | cat -e
```
```console
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
```                              
```bash
go run . "\n" | cat -e
```
```console
$
```
```bash
go run . "HeLlO" | cat -e
```
```console
 _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $
 ```                                     

### Using the `--color` Flag
You can specify a color using the `--color` flag. Supported colors include `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, and `white`.

#### Color the Entire Word
```bash
go run . --color=red "Hello"
```
```diff

- _              _   _          
-| |            | | | |         
-| |__     ___  | | | |   ___   
-|  _ \   / _ \ | | | |  / _ \  
-| | | | |  __/ | | | | | (_) | 
-|_| |_|  \___| |_| |_|  \___/  
                       
```

#### Color Specific Characters
```bash
go run . --color=green "e" "Hello"
```
```console
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
```

### Switching Banner Files
You can switch between different banner files by including the name of the file after the input string. If a name is not provided, the default is `standard.txt`.

Supported banner files:
- `standard.txt`
- `thinkertoy.txt`
- `shadow.txt`

**With Extension Included**
```bash
go run . "Hello There" "standard.txt" | cat -e
```
```console
 _    _          _   _               _______   _                           $
| |  | |        | | | |             |__   __| | |                          $
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| $
                                                                           $
                                                                           $
```

**Without Extension**
```bash
go run . "Hello There" "standard" | cat -e
```
```console
 _    _          _   _               _______   _                           $
| |  | |        | | | |             |__   __| | |                          $
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| $
                                                                           $
                                                                           $
```

### Non-Existent Banner Files
If a banner file is provided but does not exist, you will be informed if the word is relevant. However, empty spaces and new lines will still print even without the banner file's presence.

### Special Characters
While the program attempts to support as many special characters as possible, the following are supported:
- `\n` (Newline)
- `\t` (Tab)
- `\v` (Vertical Tab)
- `\b` (Backspace)

Unsupported special characters:
- `\a` (Alert)
- `\r` (Carriage Return)
- `\f` (Form Feed)

### Testing
You can run the tests provided like this while in the root folder of the project:

```bash
$ cd asciiart
$ go test
```

### Contributing
Contributions to the ASCII Art Generator are welcome. Please ensure that your code adheres to Go's best practices and includes unit tests for new features.
```

### Summary of Changes:
1. Added the `--color` flag feature to the `README`.
2. Updated examples to include `--color` usage.
3. Enhanced the `Usage` section with banner files and special characters.