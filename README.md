# GoLF Engine
the GoLang Fantasy Engine (GoLF Engine) is a retro game engine. It draws inspiration from fantasy console projects 
like [pico-8](https://www.lexaloffle.com/pico-8.php), [tic-80](https://tic.computer/), and [pyxle](https://pypi.org/project/pyxel/). 
Like those projects it is designed to be a retro feeling game creation/ playing tool. Unlike those
projects GoLF is more minimal in scope and only provides an api and a small set of tools to help you create your games. Tools like
an image editor, and code editor are not built in. Despite this creating games in GoLF is still easy and should still maintain
the retro game feel.

# Getting Started
Installing the golf engine is simple. You can install the engine and required tools by running the following commands.
from your terminal run 

`go get github.com/bjatkin/golf-engine/golf`

ignore any errors you get about syscall/js.

then run

`go get github.com/bjatkin/golf-engine/util`

then create a directory for your game. for example.

`mkdir hello world`

navigte into this direcotry and run the golf_toolkit binary, located in the golf engine utils folder

`./Users/[your user name]/go/src/github.com/bjatkin/golf-engine/util/golf_toolkit`

This will start the golf toolkit which you'll need to help you develop golf engine games. Now you can init your new project.

`init <project name>`

This will create all the nessisary files and folders for you to develop your game. You can edit the assets/spritesheet.png file
to add sprites to your game. you can also edit assets/map.png to create a map for your game. You can start writing your golf code
in the main.go file.

# Developing your game

# Releasing your game

# GoLF toolkit commands
  * about: Displays some simple information about the golf toolkit and why it exsists.
  * exit: Quits the golf toolkit. Will stop the development server if it's running.
  * map: Takes a map file location, a sprite file location, and an output file location. the map file should be a png file filled with 8x8 sprites from the specified sprite sheet. The result is saved to the output file.
  * csvmap: Takes a csv map file location and an output file. This file should have a list of sprite indexes which corresponds to the index of 8x8 sprites on the sprite sheet. The resut is save to the output file.
  * sprite: Takes a sprite file location and an output file location. Converts the sprite sheet into golf data and saves it to the output file location. Sprite sheets must only use 2 pallets from the 16 GoLF pallets
  * flag: Takes a flage file location and an output file location. Contains a list of flags that corespond to the sprite sheet. Each flag should be 8 characters long and consist of 1's(flag is set) and 0's(flag is not set).
  This file need not contain all flags for all 512 sprites.
  * startserver: starts a development server. This will automatically only your default browsers to localhost:8080 where you can play your game. Each time you reload your game your project will be rebuilt.
  * stopserver: stops the deveopment server if it's running. Otherwise does nothing.
  * build: builds the current project, creating a new WASM game file.
  * init: init's a new game in the current directory. Creates the following files.
    * assets
      * map.png - an empty my file.
      * spritesheet.png - an empyt sprite sheet.
    * web
      * index.html - simple html file with a canvas to run your game.
      * wasm_exec.js - the golang wasm glue file.
    * main.go - some boiler plate code to get you started.
    * golf_config - used by the golf_toolkit to compile your project.
    * build.sh - build file.
  * config: takes a golf_config property name and prints the current value. valid config property names are listed below.
    * name - your project name.
    * spriteFile - the sprite file to be converted when build is run.
    * mapFile - the map file to be converted when build is run.
    * flagFile - the flag file to be converted when build is run.
    * outputSpriteFile - the go file to write the converted sprite data to.
    * outputMapFile - the go file to write the converted map data to.
    * outputFlagFile - the go file to write the converted flag data to.
  * setconfig - takes a golf_config property name and a new value. the new value is assigned to that value in the golf_config file.
  * clear - clears the terminal screen.
  * help - displays all the golf toolkit commands.
  * !! - re-run the last executed command.

# Specs
  * 192 x 192 screen size
  * 64 total colors split into 16 pallets with 4 colors each
  * 8 on screen colors consisting of any 2 of the 16 predefined pallets
  * one 256 x 128 sprite sheet for a total of 512 8x8 sprites
  * one 128 x 128 tile map
  * 60 FPS
  * Mouse support with 3 different cursor styles
  * Keyboard support

# The GoLF color pallet
golf uses a pallet of 64 colors split into 16 four color pallets. You can mix and match these pallets however you want but only 2 can be used at a time

![Color Pallet](https://github.com/bjatkin/golf-engine/blob/master/images/golf_color_pallet.png)

### HEX
Pallet | Color 1 | Color 2 | Color 3 | Color 4
-------|---------|---------|---------|---------
0 | #000000 | #606060 | #909090 | #c0c0c0
1 | #404040 | #808080 | #a0a0a0 | #ffffff
2 | #150307 | #4a1215 | #ae3031 | #ec7e7c
3 | #300b0e | #7e2123 | #363f3f | #f2bdb8
4 | #160602 | #482c06 | #dc9b23 | #ebb951
5 | #2f1907 | #795013 | #e4aa3a | #fad77e
6 | #191500 | #9b5a10 | #d27614 | #f2b27b
7 | #02180c | #4a663d | #7c9964 | #bccb90
8 | #324c2d | #61804d | #97b27a | #cde3a6
9 | #070d11 | #3a777e | #6ce1ea | #94e8e6
10 | #214248 | #53acb4 | #80e5e8 | #bceee2
11 | #07011a | #152253 | #2342a5 | #3b68bf
12 | #0e1237 | #1c326f | #2f55a5 | #528df2
13 | #152253 | #46878f | #94e344 | #e2f3e4
14 | #00303b | #ff7777 | #ffce96 | #f1f2da
15 | #000000 | #c51111 | #143a85 | #ffffff

### RGB
Pallet | Color 1 | Color 2 | Color 3 | Color 4
-------|---------|---------|---------|---------
0 | (0, 0, 0) | (96, 96, 96) | (144, 144, 144) | (192, 192, 192)
1 | (64, 64, 64) | (128, 128, 128) | (160, 160, 160) | (255, 255, 255)
2 | (21, 3, 7) | (74, 18, 21) | (174, 48, 49) | (236, 126, 124)
3 | (48, 11, 14) | (126, 33, 35) | (230, 63, 63) | (242, 189, 184)
4 | (22, 6, 2) | (72, 44, 11) | (220, 155, 35) | (235, 185, 81)
5 | (47, 25, 7) | (121, 81, 19) | (228, 170, 58) | (250, 215, 126)
6 | (25, 21, 0) | (155, 90, 16) | (210, 118, 20) | (242, 178, 123)
7 | (2, 24, 12) | (74, 102, 61) | (124, 153, 100) | (178, 203, 144)
8 | (50, 76, 45) | (97, 128, 77) | (151, 178, 122) | (205, 227, 166)
9 | (7, 13, 17) | (58, 119, 126) | (108, 225, 234) | (148, 232, 230)
10 | (33, 66, 72) | (83, 172, 180) | (128, 229, 232) | (188, 238, 226)
11 | (7, 1, 26) | (21, 34, 83) | (35, 66, 165) | (59, 104, 191)
12 | (14, 18, 55) | (28, 50, 111) | (47, 85, 165) | (82, 141, 242)
13 | (21, 34, 83) | (70, 135, 143) | (148, 227, 68) | (226, 243, 228)
14 | (0, 48, 59) | (255, 119, 119) | (255, 206, 150) | (241, 242, 218)
15 | (0, 0, 0) | (197, 17, 17) | (20, 58, 133) | (255, 255, 255)

# Example Games
There are a few example projects located at golfExamples which you can download, play, and review to better understand how the golf engine works.
You play these games in one of two ways. You can download the WASM files here. Then run the golf_toolkit and run ‘play <wasmfile.wasm>’
This will start a local server and allow you to play the game on localhost:8080. You can also clone the github repo. 
Once cloned you can enter the directory for one of the examples (e.g. bibiDuck) and then run the golf_toolkit. 
From there simple run ‘startserver’ the local development server will be started and you’ll be able to play the game at localhost:8080. 
Cloning the github project is the preferred method as this allows you to see the code as well as modify it to better understand how the API works.

# Releasing My Game
Making your GoLF games available for others to play is easy. You can distribute your game in one of two ways.
First if your friends have the golf_toolkit installed you can simply send them the WASM file generated by go build.
This will have all the necessary data so that your friends can play the game using the ‘play <wasmfile.wasm>’ command.
The second option is to export your game as a ‘native’ application. The way this works is by producing an executable 
which will start a localhost server when run eliminating the need for the golf_toolkit. This file can be produced 
by using the exportnative command in the golf toolkit. This is ideal for distributing your game to a wide audience. 
The final method is making your game playable from the web. This is the preferred option as it’s what the GoLF engine was designed for. 
You can run exportweb from the golf_toolkit to produce a dist directory. This directory will contain a basic index.html file, 
a wasm file with your game, a js file with the glue code required to run golang web assembly and a small go file with an example server. 
You can replace/ modify the index.html file in any way you want as long as you don’t modify the canvas element or any of the script tags.
Then, either using the default go server or a server of your choice, simply server the index.html file, the wasm file and the js file to
a user when they visit the appropriate page. Setting up a server is beyond the scope of this documentation but there is a large amount of
excellent material on how to do this already. Visit here or here for some material on the matter.


# GoLF API

### GoLF structs

**golf.SOp:** This structure is a list of options that can be passed to a sprite to change how it is drawn.
  * FH: flip the sprite horizontally.
  * FV: flip the sprite vertically.
  * TCol: set the sprites tranparency color.
  * PFrom & PTo: change the sprites pallet. colors number n in PFrom is converted to color number n in PTo.
  * W: width of the sprite in tiles to read from the spritesheet. (e.g. W: 2 is 16 pixels in width).
  * H: height of the sprite in tiles to read from the spritesheet. (e.g. H: 2 is 16 pixels tall).
  * SW: the amount to scale the width of the sprite. default value is 1 or no scaling.
  * SH: the amount to scale the height of the sprite. default value is 1 or no scaling.
  * Fixed: if this is set to true then the sprite ignores the camera x & y when draing. Useful for UI.

**golf.TOp:** this structure is a list of options that can be passed to text functions to change how text is drawn.
  * Col: The color to draw the text.
	* Fixed: If this is set to true the text ignores the camera.
	* SW: the amount to scale the width of the text.
  * SH: the amount to scale the height of the text.

### Golf Types

**golf.Col:** a golf color. there are 8 colors ranging from Col0 to Col7. The first 4 colors map to pallet A and the last
four map to pallet B.

**golf.Pal:** a golf pallet. There are 16 available pallets (Pal0 to Pal15). These can be used to give you game a unique feel/ look.

### The GoLF Engine

**NewEngine(updateFunc func(), draw func()):** creates a golf engine instance and returns a pointer to the engine. 
The golf engine is the main object used to perform most of the golf functions

**engine.Run():** starts the game engine running. Once this is run the update function will be called 60 times a second and the draw function will be called 60 times a second.

**engine.Frames():** returns the number of frames that have passed since the game engine was started. This count includes
the startup animation frames. The startup animation is 254 frames meaning the first frame that the update/ draw function 
will be called is frame 255.

**engine.DrawMouse(style int):** sets the draw style for the mouse indictor.
  * 0 = no mouse cursor is drawn
  * 1 = a mouse arrow is drawn
  * 2 = a hand cursor is drawn
  * 3 = a cross cursoe is drawn

**engine.Cls(col golf.Col):** fills the screen with the col color.

**engine.Camera(x, y int):** Changes the X, Y coordinates of the camera. This value is then subtracted from the
X, Y coordinates of all future drawing calls. This is useful for moving the panning the screen around.

**engine.Clip(x, y, w, h int):** clips all future draw functions with upper left corner at point (x, y) and with w and heigh h.

**engine.RClip():** resets the screen clipping so that no screen pixels are clipped.

**engine.PalA(pallet golf.Pal):** sets the first pallet.

**engine.PalB(pallet golf.Pal):** sets the second pallet.

**engine.PalGet():** returns the first and second pallets that are currently set.

### Shapes

**engine.Pset(x, y float64, col golf.Col):** Sets the pixel on the screen at point (x, y) to the color col.

**engine.Pget(x, y float64):** Gets the color currently set at screen pixel (x, y).

**engine.Rect(x, y, w, h float64, col golf.Col):** Draw an empty rectangle outline with the specified draw color.

**engine.RectFill(x, y, w, h float64, col golf.Col):** Draw a filled rectangle with the specified draw color.

**engine.Line(x1, y1, x2, y2 float64, col golf.Col):** Draw a line from point (x1, y1) to (x2, y2). The line is drawn with
the specified color.

**engine.Circ(xc, yx, r float64, col golf.Col):** Draw a circle outline with center at point (xc, yc) with radius r.
The outline is drawn with the specified color.

**engine.CircFill(xc, yc, r float64, col golf.Col):** Draw a filled circle with center at point (xc, yc) with radius r.
The circle is drawn with the specified color

### Controlls

**engine.Btn(key golf.Key):** returns true if the given key is being held on this frame.

**engine.Btnp(key golf.Key):** returns true if the given key was first pressed on this frame.

**engine.Btnr(key golf.Key):** returns true if the given key was released on this frame.

**engine.Mbtn(key golf.MouseBtn):** returns true if the given mouse key is being held on this frame.

**engine.Mbtnp(key golf.MouseBtn):** returns true if the given mouse key was first pressed on this frame.

**engine.Mbtnr(key golf.MouseBtn):** returns true if the given mouse was released ont his frame.

### Map

**engine.LoadMap(mapData [0x4800]byte):** load the map data into memory.

**engine.Map(mx, my, mw, mh int, dx, dy float64, opts ...SOp):** Draws the map data onto the screen witht he left coordinate 
at screen point dx, dy. mx and my are the map coordinates in tiles and dw and mh are the map size in tiles. opts are optional and change how each individual map tile is drawn.

**engine.Mset(x, y, t int):** sets the map tile to sprite number t at the map coordinate (x, y)

**engine.Mget(x, y int):** returns the sprite index of the tile a the map coordinate (x, y)

### Sprites

**engine.LoadSprs(sheet [0x3000]byte):** load the sprite sheet data into memory.

**engine.LoadFlags(flags [0x200]byte):** load the sprite flags into memory. Each sprite in the sprite sheet has 1 byte 
(or 8 flags) associated with is that can be set and then later checked. The meaning of each of these flags is 
totally up to the needs of the progarmmer.

**engine.Spr(n int, x, y float64, opts ...SOp):** draw sprite number n at screen position x, y. opts are optional and change
how the sprite is drawn on screen. the sprite sheet is broken up into 8x8 areas that are then number from the top left 
to the bottom right. Usually the first 8x8 sprite is not used as this sprite is drawn as a transparent tile when used on the map screen.

**engine.SSpr(sx, sy, sw, sh int, dx, dy float64, opts ...SOp):** a more general version of the spr function. It draws a sprite from an abitrary spot on the sprite sheet with abitrary size to the screen. sx and sy are the pixel coordiantes of upper left corner of the sprite on the sprite sheet. sw and sh are the sprites withd and height respectivly. dx and dy are the screen coordinates that the sprite is drawn to. opts is optional and changes how the sprite is drawn on screen.

**engine.Fget(n, f int):** returns flag number f associated with sprite number n.

**engine.Fset(n, f int, s bool):** sets the flag number f for sprite n to the same value as s.

**engine.FgetByte(n int):** returns the full byte assocated with sprite number n.

**engine.FsetByte(n int, b byte):** sets the full byte assocated with sprite number n to the value of b.

### Text

**engine.Text(x, y float64, text string, opts ...TOp):** draws the text on screen at point (x, y), all text is converted to the golf
engines internal font which is all upper case. There are also several sequences that are converted in to golf emojis. escaped sequences are listed bellow. opts are optional and modify how the text is drawn.
  * (<) left button
  * (>) right button
  * (^) up button
  * (v) down button
  * (x) x button
  * (o) o button
  * (l) l shoulder button
  * (r) r shoulder button
  * (+) + button
  * (-) - button
  * :) smily face
  * :( frowny face
  * x( angry face
  * :| meh face
  * =[ boxy face
  * |^ up arrow
  * |v down arrow
  * <- left arrow
  * -> right arrow
  * $$ pound symbol
  * @@ small black dot
  * <| speaker symbole
  * <3 white heart
  * <4 black heart
  * +1 plus one symbole
  * -1 minus one symbole
  * ~~ the pi symbole
  * () tall black dot
  * [] dark square
  * :; dither pattern
  * ** start symbole
  note: if you need to draw one of these patterns without it being drawn as an emoji you can use the '^' symbole to escape.
  the pattern. (e.g. ^** will be drawn as two asterix characters rather than a star)

**engine.TextL(text string, opts ...TOp):** draws text in the upper left hand corner of the screen. Each time TextL called a new
line is added.

**engine.TextR(text string, opts ...TOp):** draws text in the upper right hand corner of the screen. Each time TextR is called
a new line is added.

### Cart Data

**engine.Dset(name string, data []byte):** stores persistent data to a users browser as a cookie. only 1024 bytes or less can be stored and the name must be alpa numeric.
the name is used to save the data so it can be retrived later. Keep in mind this name should be unique or it may get overwritten by other games.

**engine.Dget(name string):** retrievs data stored with Dset. In addition to returning the data it returns a bool which is true if the saved data was successfully found.

# The GoLF memory map
Another goal of golf is to be a 'hackable' engine. To achieve this golf uses virtual ram (stored in engine.RAM). This virtual ram
stores sprite data, map data, the screen buffer and much more. Bellow is a list of all the important memory addresses in the 
virtual ram. (you can also view memroy addresses by looking at the memoryMap.go file)

  * Screen Buffer: 0x0000 - 0x3600, this data is coppied to the screen once per frame
  * Screen Pallet: 0x3600, the two screen pallets, top 4 bits are pallet 1 and bottom 4 bits are pallet 2
  * Start Screen Length: 0x3601, the number of frames to play the startup animation. If you set this to 0 you can skip the startup animtion.
    If you choose to do this please credit the project some other way in your game.
  * CameraX: 0x3602-0x3603, the 16 bit x coordinate of the camera.
  * CameraY: 0x3604-0x3605, the 16 bit y coordinate of the camera.
  * Frames: 0x3606-0x3608, the 24 bit number that counts the frames since the game engine was started.
  * ClipX: 0x3609, the x coordinate of the clipping rect.
  * ClipY: 0x360A, the Y coordinate of the clipping rect.
  * ClipW: 0x360B, the width of the clipping rect.
  * ClipH: 0x360C, the height of the clipping rect.
  * MouseX: 0x360D, the x coordinate of the mouse.
  * MouseY: 0x360E, the y coordinate of the mouse.
  * Left Click: 0x360F, the click state of the left mouse button (00 - unclicked, 01 - click started, 10 - click ended, 11 - pressed)
  * Middle Click: 0x360F, the click state of the middle mouse button (00 - unclicked, 01 - click started, 10 - click ended, 11 - pressed
  * Right Click: 0x360F, the click state of the right mouse button (00 - unclicked, 01 - click started, 10 - click ended, 11 - pressed
  * Mouse Style: 0x360F, the draw style of the mouse (00 - mouse currsor is not drawn, 01 - arrow, 10 - hand cursor, 11 - cross cursor)
  * Keyboad Key Stat: 0x3601 - 0x3646, the pressed state of all the keys on the keyboard (00 - unpressed, 01 - press started, 10 - press ended, 11 - pressed). Keys are indexed in this array based on their js keyCode - 9 (backspace keycode)
  * Internal Sprite Sheet: 0x3647 - 0x3F47, sprite data for the golf font, emojies, logo and mouse sprites.
  * Sprite Sheet: 0x3F48 - 0x6F48, the data for the user sprite sheet, this data is stored in the compressed format described below.
  * Active Sprite Buff: 0x6F49 - 0x6F4A, 16 bit address that points to the memory location that will be used by the sprite functions. You can use this to swap to the internal sprite sheet or reindex sprites on the sprite sheet.
  * Map Data: 0x6F4B - 0xB74B, the map data, this data is stored in the compressed format described below.
  * Sprite Flag Data:  0xB74C - 0xB94C, the sprite flag data, each sprite gets one byte of data which is 8 flags.

# Data Packing
Part of the goal for this project was to make the console feel somewhat ‘retro’ no just in it’s visual style but also internally.
I decided to do this by creating simulated memory which is the engine.RAM array.
Using this memory efficiently was a goal since early on in development.
Because of this some of the internal representations of data can get a little difficult to understand in order to aid 
in this I’ve included the following sections of GoLF pixel data and map tile data. 
While this may seem complicated it is totally possible to use the engine without understanding these
concepts so unless you're interested in how the internals of GoLF works feel free to skip these sections.

# GoLF Pixel Data
GoLF supports up to 8 colors on screen at a time. Unfortunately 8 colors only fills up 3 bits which makes packing the grafix data efficiently rather difficult.
In order resolve this the golf engine splits a pixels color from its pallet. This means that grafix data is represented as follows in GoLF simulated memory.
2 bytes with 4 pixel intensities each followed by one byte with 8 color pallets, one for each of the previous 8 pixel intensities.
You can learn more about his by looking at the pget and pset functions in the golf engine.

![Sprite Data](https://github.com/bjatkin/golf-engine/blob/master/images/sprite_data.png)

# GoLF Map Tile Data
The golf map supports indexing 512 8x8 sprites from the sprite sheet. This means a map tile value is a maximum of 9 bits. 
Like with pixel data this odd sizing makes it a little tricky to use memory efficiently. 
In order to deal with this I pack pixel data using the following method.
The sprite sheet is broken up into two halves each with 256 sprites.
The top half is the low half (low memory) and the bottom half is the high half (high memory). 
I then idex each tile with an 8 bit integer 0-255 and store which half of the sprite sheet it belongs to separately.
This data is then packed into RAM as follows. 8 bytes with a 0-255 index for each tile, followed by 1 byte 
with 8 bits to indicate whether the previous 8 tiles are in high or low memory.

![Map Tile Data](https://github.com/bjatkin/golf-engine/blob/master/images/map_data.png)

# GoLF Graphics Memory Layout (Map and SpriteSheet)
Additionally it’s worth noting that the GoLF sprite data and GoLF map data are placed next to each other in memory and grow in opposite directions.
Sprites grow from low memory to high memory and the map grows from high memory to low memory.
This allows for the sprite sheet or the map to expand beyond the default sizes if needed.
In the case the the spritesheet ‘overgrows’ the map keep in mind that the map will only index the first
512 tiles and that the sprite flags will only apply to the first 512 sprites as well.

![Map Vs Sprite Data](https://github.com/bjatkin/golf-engine/blob/master/images/sprite_vs_map.png)

### TODO
---
* Add instructions for installing and playing the golf examples
* Test the golf toolkit on a windows machine

### TODO long term
---
* Make it more fantasy console like
  * add a golf terminal that runs in the browser
  * build with golf engine so it has the same feel
  * sprite editor in the golf terminal
  * map editor and viewer
  * sprite flag editor 
* Sound? (I still have 20k memory for this)
* Add interpreted scripting language to make it more aproacable and to prevent golang from being an install requirement
* Let text use multips options with {} syntac to start and end option sections
* add vertical and horizontal flip to the text functions