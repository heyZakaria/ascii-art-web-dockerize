# Description
 Ascii Art Web Stylize is a web GUI convert text to Ascii Art
# Authors
- Zakaria ABDELALI
- Ayoub RABYA
- Ayyoub MAHFOUD
# Usage: how to run
in the Terminal go to the project Directory and write ``` go run main.go ``` and follow the link http://localhost:8080 that appears in the terminal
# Implementation details: algorithm

#### 1. Setup the server using the http Methods HandleFunc and ListenAndServe:
 *We use ListenAndServe to run the server and HandleFunc to handle two paths "/" using HomeHandler and "/ascii-art" using AsciiArtHandler, additionaly we create a struct to hold our data (Banner, Text, Ascii-Art, CharsNotAllowed, Error)*
in details:
- **HomeHandler function:**
    * checking the path if it's not match "/"
    * checking the method if it's not match "GET"
    * parsing the html file (index.html)
    * if we have banner and text we pass them to the AsciiGenerator we put the return value in the struct
    * we check if there is any not allowed characters so we add them in the Error field 
    * Execute the template
    * Clean data from struct
- **AsciiArtHandler function:**
    * Checking the method if it's not match "POST"
    * we assign data to text and banner
    * we check if the banner and text
    * we redirect to the root path
