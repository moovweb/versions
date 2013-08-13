# Versions #

Search for versioned filenames by version patterns.

## Viable Patterns ##

    Less Than :         "< 0.4.5"
    Less or Equals :    "<= 0.4.5"
    Equals :            "= 0.4.5" or "0.4.5"
    Pessimistic :       "~> 0.3.45"
    Greater or Equals : ">= 0.4.5"
    Greater Than :      "> 0.4.5"


# Go Library Interface #

    // Returns the newest matching file
    versions.FindByName("/home/garfield/.rvm/gems/ruby-1.9.2-p180/", "nokogiri")

    // Returns all matching files
    versions.FindByNameAndVersion("/home/garfield/.rvm/gems/ruby-1.9.2-p180/", "nokogiri", "~> 1.4.2")


# Command Line Interface #

    # Returns the newest matching file
    version /home/garfield/.rvm/gems/ruby-1.9.2-p180/ nokogiri

    # Returns all matching files
    version /home/garfield/.rvm/gems/ruby-1.9.2-p180/ nokogiri "~> 1.4.2"


# License #

Copyright (C) 2012 Sean Jezewski / Moovweb

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
