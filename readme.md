# OWLJACK

**OWLJACK** is a **Clickjacking Testing Tool** designed to check if websites are vulnerable to clickjacking attacks by scanning for the absence of the `X-Frame-Options` header.
```
       ________        
     /          \      
    /            \     
   /   OWLJACK    \    
  /________________\   
    \   /\  /\   /     
     ) ( o  o ) (      
    (    (_)    )      
     \          /       
      \________/
```
   OWLJACK - Your Clickjacking Tool



## Features
- Scans multiple URLs from a text file.
- Detects missing `X-Frame-Options` headers (a sign of vulnerability).
- Simple and easy-to-use command-line interface.
- Stylish owl-themed ASCII art banner.

## Requirements
- Go Programming Language (Version 1.16+)

## Installation
1. Clone or download the source code: `git clone https://github.com/your-repo/owljack.git`
2. Navigate to the project directory: `cd owljack`
3. Build the tool: `go build owljack.go`

## Usage
1. Create a text file (`urls.txt`) with URLs to test, one URL per line:
   - Example:
     - `https://example.com`
     - `https://anotherwebsite.com`
2. Run the tool: `go run owljack.go`
3. When prompted, enter the path to the `urls.txt` file.

## Example Output
`Testing URLs for Clickjacking vulnerability...`  
`[VULNERABLE] https://example.com: No X-Frame-Options header found.`  
`[SECURE] https://anotherwebsite.com: X-Frame-Options = SAMEORIGIN`

