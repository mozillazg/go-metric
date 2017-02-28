# go-metric

[![Build Status](https://travis-ci.org/mozillazg/go-metric.svg?branch=master)](https://travis-ci.org/mozillazg/go-metric)
[![Go Report Card](https://goreportcard.com/badge/github.com/mozillazg/go-metric)](https://goreportcard.com/report/github.com/mozillazg/go-metric)

## install

`go get -u github.com/mozillazg/go-metric`


## CLI

install:

`go get -u github.com/mozillazg/go-metric/cmd/parse_metrics`

or download from [release](https://github.com/mozillazg/go-metric/releases).


```
$ cat example/metrics.txt |parse_metrics
#Time	CPU(%)(T: 16)	RAM(%)(T: 31.4 GB)	Disk(%)(T: 916.6 GB)
2017-02-27 20:02:23	1.29	0.60	0.03
2017-02-27 20:02:38	404.97	5.89	0.08
2017-02-27 20:02:53	4.05	9.28	0.08
2017-02-27 20:03:08	3.18	13.48	0.08
2017-02-27 20:03:23	2.38	17.49	0.08
```
