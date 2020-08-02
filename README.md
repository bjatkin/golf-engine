# GoLF Engine
the GoLang Fantasy Engine (GoLF Engine) is a fantasy engine. It draws inspiration from projects fantasy console projects 
like pico-8 and tic-80. Like those projects it is designed to be a retro feeling game creation/ playing tool. Unlike those
projects GoLF is more minimal in scope and only provides an api and a small set of tools to create youre game. Tools like
and image editor, and code editor are not built in. fairly restrictive game creation tool focused on creating
a more managble game creation enviroment the a fully featured game engine would offer. This project differs from those
projects in a few importan ways however.

# Specs
  * 192 x 192 screen size
  * 64 total colors split into 16 pallets with 4 colors each
  * 8 on screen colors consisting of any 2 of the 16 predefined pallets
  * one 256 x 128 sprite sheet for a total of 512 8x8 sprites
  * one 128 x 128 tile map
  * 60 FPS
  * Mouse support with 3 different cursor styles
  * Keyboard support

# Getting Started
The golf engine is just a go package. It can be installed by running `go get github.com/bjatkin/golf-engine/golf`.
There is also a toolkit which you will need to complie you game as well as for importing the sprite sheet and map file.
you can install this toolkit by running `go get github.com/bjatkin/golf-engine/util`. 
Once both these are installed you can use the golf toolkit to start a new project. Simply create a directory for you new game.
Then, open that directory in termnal and run the golf_toolkit program.
The golf_toolkit program is located in the golf-engine/util directory so you can run `./Users/[your user name]/go/src/github.com/bjatkin/golf-engine/util/golf_toolkit` to start it.
Note that this is for MacOSX users and Windows/ Linux users will have to change the path so it correctly points to the golf-engine/util directory.
Once you've started the golf_toolkit you can run `init <project name>` command where project name is the name of your game.
This will create all the nessisary files for you to start building your first game

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

# GoLF RAM
The golf engine object uses a RAM array to store crucial data about the current state of the engine. 
For example the screen buffer and sprite data. This makes the engine itself very hackable and exposes additional functionality that the API does not cover.
the memoryMap.go file lists all the critical locations in memory and this fill is reiterated below in the API section of the documentation.
You are encouraged to dig around here and mess with the values to get a better understanding of how the engine works.

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

# GoLF Map Tile Data
The golf map supports indexing 512 8x8 sprites from the sprite sheet. This means a map tile value is a maximum of 9 bits. 
Like with pixel data this odd sizing makes it a little tricky to use memory efficiently. 
In order to deal with this I pack pixel data using the following method.
The sprite sheet is broken up into two halves each with 256 sprites.
The top half is the low half (low memory) and the bottom half is the high half (high memory). 
I then idex each tile with an 8 bit integer 0-255 and store which half of the sprite sheet it belongs to separately.
This data is then packed into RAM as follows. 8 bytes with a 0-255 index for each tile, followed by 1 byte 
with 8 bits to indicate whether the previous 8 tiles are in high or low memory.

# GoLF Graphics Memory Layout (Map and SpriteSheet)
Additionally it’s worth noting that the GoLF sprite data and GoLF map data are placed next to each other in memory and grow in opposite directions.
Sprites grow from low memory to high memory and the map grows from high memory to low memory.
This allows for the sprite sheet or the map to expand beyond the default sizes if needed.
In the case the the spritesheet ‘overgrows’ the map keep in mind that the map will only index the first
512 tiles and that the sprite flags will only apply to the first 512 sprites as well.

# GoLF and WASM
WASM is an exciting new technology and brings a lot of new power and many new programing languages to the web. GoLF was designed
from the groud up to work with WASM. 

# GoLF API

NewEngine(updateFunc func(), draw func()): creates a golf engine instance and returns a pointer to the engine. The golf engine is
the main object used to perform most of the golf functions

engine.Run(): starts the game engine running. Once this is run the update function will be called 60 times a second and the draw function will be called 60 times a second.

engine.Frames(): returns the number of frames that have passed since the game engine was started. This count includes
the startup animation frames. The startup animation is 254 frames meaning the first frame that the update/ draw function 
will be called is frame 255.

engine.DrawMouse(style int): sets the draw style for the mouse indictor.
  * 0 = no mouse cursor is drawn
  * 1 = a mouse arrow is drawn
  * 2 = a hand cursor is drawn
  * 3 = a cross cursoe is drawn

engine.Cls(col golf.Col): fills the screen with the col color.

engine.Camera(x, y int): Changes the X, Y coordinates of the camera. This value is then subtracted from the
X, Y coordinates of all future drawing calls. This is useful for moving the panning the screen around.

engine.Rect(x, y, w, h float64, col golf.Col): Draw an empty rectangle outline with the specified draw color.

engine.RectFill(x, y, w, h float64, col golf.Col): Draw a filled rectangle with the specified draw color.

engine.Line(x1, y1, x2, y2 float64, col golf.Col): Draw a line from point (x1, y1) to (x2, y2). The line is drawn with
the specified color.

engine.Circ(xc, yx, r float64, col golf.Col): Draw a circle outline with center at point (xc, yc) with radius r.
The outline is drawn with the specified color.

engine.CircFill(xc, yc, r float64, col golf.Col): Draw a filled circle with center at point (xc, yc) with radius r.
The circle is drawn with the specified color

engine.Clip(x, y, w, h int): clips all future draw functions with upper left corner at point (x, y) and with w and heigh h.

engine.RClip(): resets the screen clipping so that no screen pixels are clipped.

engine.Pset(x, y float64, col golf.Col): Sets the pixel on the screen at point (x, y) to the color col.

engine.Pget(x, y float64): Gets the color currently set at screen pixel (x, y).

engine.PalA(pallet golf.Pal): sets the first pallet.

engine.PalB(pallet golf.Pal): sets the second pallet.

engine.PalGet(): returns the first and second pallets that are currently set.

engine.Btn(key golf.Key): returns true if the given key is being held on this frame.

engine.Btnp(key golf.Key): returns true if the given key was first pressed on this frame.

engine.Btnr(key golf.Key): returns true if the given key was released on this frame.

engine.Mbtn(key golf.MouseBtn): returns true if the given mouse key is being held on this frame.

engine.Mbtnp(key golf.MouseBtn): returns true if the given mouse key was first pressed on this frame.

engine.Mbtnr(key golf.MouseBtn): returns true if the given mouse was released ont his frame.

engine.LoadMap(mapData [0x4800]byte): load the map data into memory.

engine.Map(mx, my, mw, mh int, dx, dy float64, opts ...SOp): Draws the map data onto the screen witht he left coordinate 
at screen point dx, dy. mx and my are the map coordinates in tiles and dw and mh are the map size in tiles. opts are optional and change how each individual map tile is drawn.

engine.Mset(x, y, t int): sets the map tile to sprite number t at the map coordinate (x, y)

engine.Mget(x, y int): returns the sprite index of the tile a the map coordinate (x, y)

engine.LoadSprs(sheet [0x3000]byte): load the sprite sheet data into memory.

engine.LoadFlags(flags [0x200]byte): load the sprite flags into memory. Each sprite in the sprite sheet has 1 byte 
(or 8 flags) associated with is that can be set and then later checked. The meaning of each of these flags is 
totally up to the needs of the progarmmer.

engine.Spr(n int, x, y float64, opts ...SOp): draw sprite number n at screen position x, y. opts are optional and change
how the sprite is drawn on screen. the sprite sheet is broken up into 8x8 areas that are then number from the top left 
to the bottom right. Usually the first 8x8 sprite is not used as this sprite is drawn as a transparent tile when used on the map screen.

engine.SSpr(sx, sy, sw, sh int, dx, dy float64, opts ...SOp): a more general version of the spr function. It draws a sprite from an
abitrary spot on the sprite sheet with abitrary size to the screen. sx and sy are the pixel coordiantes of upper left corner of the
sprite on the sprite sheet. sw and sh are the sprites withd and height respectivly. dx and dy are the screen coordinates that
the sprite is drawn to. opts is optional and changes how the sprite is drawn on screen.

engine.Fget(n, f int): returns flag number f associated with sprite number n.

engine.Fset(n, f int, s bool): sets the flag number f for sprite n to the same value as s.

engine.FgetByte(n int): returns the full byte assocated with sprite number n.

engine.FsetByte(n int, b byte): sets the full byte assocated with sprite number n to the value of b.

golf.SOp: this structure represents a list of options that can be passed to a sprite to change how it is drawn.
  * FH: flip the sprite horizontally.
  * FV: flip the sprite vertically.
  * TCol: set the sprites tranparency color.
  * PFrom & PTo: change the sprites pallet. colors number n in PFrom is converted to color number n in PTo.
  * W: width of the sprite in tiles to read from the spritesheet. (e.g. W: 2 is 16 pixels in width).
  * H: height of the sprite in tiles to read from the spritesheet. (e.g. H: 2 is 16 pixels tall).
  * SW: the amount to scale the width of the sprite. default value is 1 or no scaling.
  * SH: the amount to scale the height of the sprite. default value is 1 or no scaling.
  * Fixed: if this is set to true then the sprite ignores the camera x & y when draing. Useful for UI.

engine.Text(x, y float64, text string, opts ...TOp): draws the text on screen at point (x, y), all text is converted to the golf
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

engine.TextL(text string, opts ...TOp): draws text in the upper left hand corner of the screen. Each time TextL called a new
line is added.

engine.TextR(text string, opts ...TOp): draws text in the upper right hand corner of the screen. Each time TextR is called
a new line is added.

golf.Col: a golf color. there are 8 colors ranging from Col0 to Col7. The first 4 colors map to pallet A and the last
four map to pallet B.

golf.Pal: a golf pallet. There are 16 available pallets (Pal0 to Pal15). These can be used to give you game a unique feel/ look.

# The GoLF memory map
the golf engine is designed to be hackable as was as to maintin a retro feel which developing. In order to achieve this the
golf engine uses a block of memory that can be accessed using the RAM member variable of a golf engine instance. The engine.RAM
contains all the data for the sprite memory, the map memory, the screen buffer and even the keyboard state. If there is anything
that the API does not expose you can probable read or write that data from the engine memory. In order to help you with this 
you may with to look at the memoryMap.go file. This contains a list of all the memory addresses used by golf to run the engine.

# The GoLF color pallet
part of golf's goal as a engine is to all the creation of games with a unique and recognizable style. Visual style is an important
aspect of this and so the engine uses a unique method for drawing colors on screen. Each pixel can only be one of 4 colors but it
can also belong to either of 2 pallets. This means that you can effectively draw up to 8 colors on screen at a time. There is no
limit to how often you can change these pallets either allowing you to create uinque and interesting games with creative graphics.
This system may seem restrictive but that is by design. The golf engine firmly believes that restrictions foster creativity and
also help developers finish projects which are major goals of this engine. The 16 available 4 color pallets are shown bellow 
(in order from top Pal0 to bottom Pal15). Play around and find unique and interesting pallet combinations!

# Developing your game

# Releasing your game

### TODO
---
* Build out the readme so there is good documentation for the API
* Change Cls to take a Col so we dont need to store a bg color (Update the README)
* Create an example golang server for using this framework for playing games online
* Add instructions for installing and playing the golf examples
* Test the golf toolkit on a windows machine

### DONE
---
* Fix mouse transparency [x]
* Clean up startup animation code [x]
* Use all 8 colors on the loading animation to make the fading nicer [x]
* Fix text alpha [x]
* Show error on server automatic build in golf toolkit [x]
* Test new map code [x]
* chang the way the map code works so that it uses the new color atlas and works with the new sprite import code [x]
* give X, Y coords on unknown color in image [x]
* Change build code to use hex instead of bin to make smaller generated files [x]
* Test the new sprite importing code [x]
* Inject draw.js to make golf require less dependancies [x]
* Default to the black and gray pallets [x]
* Fix color pallet swapping caused by the startup anim (allow the user to set the pallets) [x]
* Template generation code should create arrays rather than slices [x]
* Add last 4 color pallets (Black White Red Blue pallet) [x]
* Add scale width and scale height opts to the text drawing [x]
* Add startup animation when you start the game [x]
* Fix the color pallet [x]
* Add function to load sprite flag data [x]
* Change the sprite functions from using ints to using floats [x]
* Change the SprOpts and TextOpts to be Sop and Top to make the API more terse [x]
* Add a save cart data function (save to a browser cookie) [x]
* Read cart data function (read from the browsers cookies) [x]
* Add readme to go examples [x]
* Create github for golfExamples games [x]
* Check out Faith and Faith 2 itch.io [x]
* Better error message when building fails in the golf_toolkit [x]
* Throw a waning when rebuilding on refresh fails [x]
* Add in command to modify the config file [x]
* Fix the gaps that start to form in the map file as you scale up (only a problem with scale w/h) [x]
* Make the map tool respect SprOpts, or make a separete MapOpts for pallet swaping [x]
* Make the build tool work with csv map files [x]
* Make golf_config not a hidden file [x]
* Use \n chars to seperate golf config files rather than commas. this will make it easier to read and use [x]
* restructure code. Delete bad old code. Make the repo nice and clean [x]
* Add function to load map data (reverse order) [x]
* Make tools for importings maps. [x]
  * Import CSV as map [x]
  * Import image as map [x]
* restructure init so that starting code is nicer (e.g. js folder, assets folder etc.) [x]
* oo for the filled in circle is a bad letter sequence. Picke a sequence that is less common in regular words [x]
* Create csv format for the flags (Probably just 1's and 0's) [x]
* Import sprite flag csv (Should be pretty easy) [x]
* Create the Golf ToolKit [x]
* Add about and help commands [x]
* Add clear and !! command to the go toolkit [x]
* Fix Filled Circles so the look the same as hollo ones [x]
* Code automatically recompiles on browsers refresh [x]
* Start and Stop server with golf_toolkit (nonblocking) [x]
* Add build and init to golf_toolkit [x]

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