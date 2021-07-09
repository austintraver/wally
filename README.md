# Wally

<img height=400 src="https://d33wubrfki0l68.cloudfront.net/7033dde42ac38d4f2d7da49b841d2a06ecc6fbc0/682b0/static/wally-logo-4b9418edd84190e10b0704f8962e76cb.png" alt="Wally Logo">

Flash your [ZSA Keyboard](https://ergodox-ez.com) the EZ way.

# Introduction

This is my first contribution to the Wally project, I actually just started
using my ErgoDox keyboard just the other day. I've since fallen in love with how
easy it is to create new iterations of my keyboard layout. Wally makes it so
easy to modify my keyboard, and hopefully this tool will be able to make it
easier for you as well!

I wanted to give back to this project, since has already benefited me so
greatly.

## Preview

Below, is an interactive demonstration of the command-line interface.

[![asciicast](https://asciinema.org/a/qRAvPLcGWV3ZZ6Bxts6CCMx4C.svg)](https://asciinema.org/a/qRAvPLcGWV3ZZ6Bxts6CCMx4C)

## Installation

```shell
go get github.com/austintraver/wally
```

## Usage

To flash new firmware onto your keyboard, run the command below in your shell.
Make sure to replace `FIRMWARE` with the path to a firmware file on your computer,
such as `ergodox_ez_layout.hex`.

```shell
wally flash FIRMARE
```

## Feedback

I encourage any and all feedback, requests, and modifications. This is a working
product, but by no means does that make it a *final* product 🌟
