# CyberSpace-go

Is a DGF frontend for the terminal.


## Feture tracker


* Board editor
    * Load Board
        * From marshalled data
        * From Text file
    * Load defaults
      * White 5.5 komi
      * x handicap
    * Place stones
    * Scroll back/forth
* Raw ASCII display
* Raw UTF-16 (nerdfont) display
* Mouse input
* Quarter zoom


### ASCII

The most basic configuration uses only ASCII characters to print a board.
```
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . O . . . . 
. . . . . . . . . . . . . . @ O O . . 
. . . + . . . . . + . . . . . @ @ . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . + . . . . . + . . . . . + . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . + . . . . . + . . . . . + . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
. . . . . . . . . . . . . . . . . . . 
```
The following ASCII is used:

Intersections are periosds (.) with spaces between.

Starpoints: +

Stone A : @

Stone B : O (capital letter o)

## Raw nerdfont display
As a simple but aesthetically pleasing startign point, the following board display is used

󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟  󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟    󰧟 󰧟
󰧟 󰧟 󰧟 󰧞 󰧟 󰧟 󰧟 󰧟 󰧟 󰧞 󰧟 󰧟 󰧟 󰧟 󰧟   󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧞 󰧟 󰧟 󰧟 󰧟 󰧟 󰧞 󰧟 󰧟 󰧟 󰧟 󰧟 󰧞 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧞 󰧟 󰧟 󰧟 󰧟 󰧟 󰧞 󰧟 󰧟 󰧟 󰧟 󰧟 󰧞 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟
󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟 󰧟

The following utf-16 characters are used to compose a raw text board:
intersections : 󰧟 \udb82\udddf
start points : 󰧞 \udb82\uddde
stone A (black on a dark them, white on a light theme) :  \uebb4
stone B :  \uebb5


As characters have greater hieght than width, spaces are place between each of the intersections and/or stones.
