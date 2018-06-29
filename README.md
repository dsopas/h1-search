h1-search
===========

We created this tool to fill out needs to gather information on most common
issues on particular [HackerOne][3] bounty programs. The tool will connect to H1
and retrieve all the public disclosed reports and display them in a local webserver. 
You can search for specific attacks and click to see the report. 
h1-search was developed by David Sopas [@dsopas][1] and Paulo Silva [@pauloasilva_com][2].

## Build

```
$ go build h1-search.go hacktivity.go

## Usage

```
$ ./h1-search localtapiola
```

### Examples

```shell
$ ./h1-search localtapiola
2018/06/29 11:27:18 ::: Loading h1-search.go...
2018/06/29 11:27:18 ::: by @dsopas and @pauloasilva_com
2018/06/29 11:27:18 :::::::::::::::::::::::::::::::::::.
2018/06/29 11:27:21 ::: Creating template file...
2018/06/29 11:27:21 ::: Getting results from HackerOne...
2018/06/29 11:27:23 ::: Getting page 1 of 4...
2018/06/29 11:27:25 ::: Getting page 2 of 4...
2018/06/29 11:27:26 ::: Getting page 3 of 4...
2018/06/29 11:27:28 ::: Getting page 4 of 4...
2018/06/29 11:27:28 ::: File is ready...
2018/06/29 11:27:28 ::: Click to view the results: http://localhost:3000/
```

![h1-search screenshot](screenshot_demo.png)

## License

Copyright (c) 2018 David Sopas

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

[1]: https://www.twitter.com/dsopas
[2]: https://www.twitter.com/pauloasilva_com
[3]: https://www.hackerone.com