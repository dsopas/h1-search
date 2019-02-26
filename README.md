h1-search
===========

We created this tool to fill out the need of gathering information on most common
issues on particular [HackerOne][3] bounty programs. h1-search will connect to H1
and retrieve all the public disclosed reports on that specific program and display them in a local webserver. 
Beware that __H1 has rate limit on GET requests__ so don't abuse it too much.

The tool provides you the possibility of searching for specific attacks and direct link to the report. 
h1-search was developed by David Sopas [@dsopas][1] and Paulo Silva [@pauloasilva_com][2].

## Build

```
$ go build h1-search.go hacktivity.go
```

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

[1]: https://www.twitter.com/dsopas
[2]: https://www.twitter.com/pauloasilva_com
[3]: https://www.hackerone.com
