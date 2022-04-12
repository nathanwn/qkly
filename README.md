<div align="center">

  <h1>qkly</h1>
  <h5>CLI automation tool for competitive programming, aiming to make things happen a bit more quickly (hence the name) in an organised manner.</h5>

[![Golang](https://img.shields.io/badge/Go-grey.svg?style=for-the-badge&logo=go)](http://www.lua.org)
![Work In Progress](https://img.shields.io/badge/Work%20In%20Progress-yellow?style=for-the-badge)

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/nathan-wien/qkly/Test?style=flat-square)](https://github.com/nathan-wien/qkly/actions?query=workflow%3ATest)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.17-61CFDD.svg?style=flat-square)

</div>

## :book: Contents

- [:book: Contents](#book-contents)
- [:rocket: Features](#rocket-features)
  - [:arrow_forward: `qkly fetch`](#arrow_forward-qkly-fetch)
- [:package: Installation](#package-installation)
  - [:hammer: Manual installation](#hammer-manual-installation)
- [:heavy_check_mark: Supported Online Judges](#heavy_check_mark-supported-online-judges)
- [:memo: Roadmap](#memo-roadmap)

## :rocket: Features

This tool saves you from tedious repetitive actions when solving problems, such as manually copying and pasting sample test cases, copying your template files, etc. . In addition, it helps organising all the problems that you solve into an organised workspace structure for future reference.

With `qkly`, your workspace structure would look similar to the following:

```
.
├── qkly.yaml
├── solutions
│   ├── atcoder
│   │   └── abc246
│   │       └── a
│   │           ├── 1.in.txt
│   │           ├── 1.out.txt
│   │           ├── 2.in.txt
│   │           ├── 2.out.txt
│   │           ├── main.cpp
│   │           └── Makefile
│   └── codeforces
│       └── 1665
│           └── a
│               ├── 1.in.txt
│               ├── 1.out.txt
│               ├── main.cpp
│               └── Makefile
└── templates
    └── default
        ├── main.cpp
        └── Makefile
```

- `qkly.yaml` is the configuration file for a `qkly` workspace. At the moment, the only thing that you can set is the port that `qkly fetch` listens on. An empty `qkly.yaml` file in the current directory is suffice to run `qkly fetch`.
- `solutions` is the directory where all your solutions go. All tasks are organised nicely into a hierarchy of judges, contests, and problem ids.
- `templates/default` is the directory where all your template files go. These files will be automatically copied to each task directory on fetching with `qkly fetch`.

The current version of the tool offers the following features:

### :arrow_forward: `qkly fetch`

This command, used in combination with the browser plugin [Competitive Companion](https://github.com/jmerle/competitive-companion) (available on Firefox and Chrome), initiates a local http server to listen for POST requests from Competitive Companion to get sample test cases, create solution directories and copy all of your template files into these directories.

Note that you could only run `qkly fetch` in the root directory of your workspace.

The general use case is as follows:

- The user runs `qkly fetch`. `qkly` starts waiting for Competitive Companion to fetch information of one or more tasks.
- The user navigates to a problem or contest on an online judge website and clicks on the Competitive Companion plus icon in the browser's toolbar. When this happen, Competitive Companion sends one or more POST requests to `qkly`, each contains information of a task with all of its sample test cases.

`qkly` will then:

- Create a designated directory for each task/problem based on the online judge, the id of the contest and the id of the problem.
- Fetch sample test cases into the task directory.
- Copy template files into the task directory automatically.

## :package: Installation

Since the project is still in its very early version, no binary is released just yet.

### :hammer: Manual installation

- Requirement: latest version of [go](https://go.dev/doc/install) (at least `1.17`)

```sh
$ git clone https://github.com/nathan-wien/qkly.git
$ cd qkly
$ go install ./cmd/qkly
```

## :heavy_check_mark: Supported Online Judges

At the moment, this tool supports the following online judges:

| Judge                                     | Judge ID     |
| ----------------------------------------- | ------------ |
| [AtCoder](https://atcoder.jp)             | `atcoder`    |
| [CodeForces](https://codeforces.com)      | `codeforces` |
| [CSES](https://www.cses.fi/problemset)    | `cses`       |
| [HackerRank](https://www.hackerrank.com/) | `hackerrank` |

## :memo: Roadmap

T.B.D.
