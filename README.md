<p align="center">
  <img width="100" src="https://github.com/hungvminh/vscode-portable/blob/master/res/papp.png">
</p>

<p align="center">
  <a href="https://github.com/hungvminh/vscode-portable/releases"><img src="https://img.shields.io/github/release/hungvminh/vscode-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/hungvminh/vscode-portable/releases"><img src="https://img.shields.io/github/downloads/hungvminh/vscode-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://github.com/hungvminh/vscode-portable/actions?workflow=build"><img src="https://img.shields.io/github/actions/workflow/status/hungvminh/vscode-portable/build.yml?label=build&logo=github&style=flat-square" alt="Build Status"></a>
</p>

# VSCode Portable - Custom Build

A portable version of Visual Studio Code that can run from any location without installation.

## Features

- 🚀 Completely portable - no installation required
- 💾 All data stored in portable directory
- 🔧 Easy to use - just extract and run
- 🆕 Latest VSCode version (1.101) with all new AI features
- 🔒 Secure and isolated from system VSCode installations

## Download

Download the latest version from the [Releases](https://github.com/hungvminh/vscode-portable/releases) page.

## Usage

1. Download the latest release
2. Extract the ZIP file to your preferred location
3. Run `vscode-portable.exe`
4. Your settings and extensions will be stored in the `data` folder

## Command Line Access

Use `code.cmd` for command-line access to VSCode features:

```bash
code.cmd myfile.txt
code.cmd --list-extensions
code.cmd --install-extension ms-python.python