# Pen-testing Directory Structure

This code will create a file structure that can be used to organize data aquired during a pen-test of a target.
I made this because I wanted to stay organized, and I found myself irritated by the constant manual necessity of creating these directory structure for every target.

**START-NOTE** I will be looking to make this code read from likely a YAML/TOML file to create whatever directory structure the user wants. Right now it is hard-coded to be what is represented below. **END-NOTE**

## Tool Functionality:

- Generates a directory structure that can be used to keep track of information
- Will allow for the '-h or -help' flag to be provided for more information
- This tool has three optional arguments:
  - '-name OR --name' | name the starting directory (e.g., targets name) | if no arg passed, the starting directory will be named 'defaultProject'
  - '-path OR --path' | name the path for where to create the directory structure (e.g., /home/projects) | if no arg is passed, the current directory will be used
  - '-branch OR --branch' | choose whether to create just one of the second-level branches (i.e., IPT, EPT) | if no arg is passed, both will be created

## Tool Requirements:

- To use the default functionality of this tool, no additional libraries or modules are needed
- This code does either need to be pulled down and compiled with GO compiler, or the binary can be taken from the GitHub repo

## Quick Notes:

- This was designed in its current state to work with a Linux OS. However if desired it can be altered to fit a Windows OS.

## TREE STRUCTURE

```
  Acme Company
    ├── EPT (external pen-test)
    │   ├── evidence
    │   │   ├── credentials
    │   │   ├── data
    │   │   └── screenshots
    │   ├── logs
    │   ├── scans
    │   ├── scope
    │   └── tools
    └── IPT (internal pen-test)
        ├── evidence
        │   ├── credentials
        │   ├── data
        │   └── screenshots
        ├── logs
        ├── scans
        ├── scope
        └── tools

```
