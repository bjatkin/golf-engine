# GOLF Engine
the GoLang Fantasy Engine. Right now this is just a convient place for the todo list. No fun stuff yet.

### TODO
---
* Trow a waning when rebuilding on refresh fails
* Make Blood The Game to demo the engine
  * Add well and aditional tree to the sprite sheet
  * Add demon enemy final boss
  * Use pallet swaps to go from day to night
  * Creppy story told though small conversations with the old man at the church
  * Maybe do celest hair on the character (should be bold red color)
  * Make red blood that you collect as enemies die
  * Use ECS (actors system, ui system, dialouge system etc.)
* Better error message when building fails in the golf_toolkit
* Add custom function for printing output in the golf_toolkit so that the server can output data
* Check out Faith and Faith 2 itch.io
* Create github for golfExamples games
* Add function to load sprite flag data
* Startup Animation?
* Add function to save over the map file, sprite file or sprite flag file
  * Make this fault tolerant so if these functions arent implemented on the server stuff dosent blow up
* Add function to save and load cookies to the browser. Limit the number of 'bytes' that can be saved
* Create and example golang server for using this framework for playing games online

### DONE
---
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
* Make BiBi Duck Game to demo the engine
  * BiBi Duck is a platformer that plays like a mario game. run around, collect fethers and find all the eggs
  * Make sure to feature the Go Gopher to demo a diologue like system
  * Collect eggs to get powerups
    * Jump
    * Hover
    * Double Jump (use wings?)
    * Peck (pick up bugs? to feed to mr frog?)
    * Quack (wake up go gopher? or someone else?)
  * 8x8 baby ducks will follow you when you get an egge
  * Use powerups to progress and find new eggs.
  * When you've found all the eggs and power-ups find MaMa duck with all her bibi's
  * Collect fethers as a bonus for extra challenge
  * Or maybe feathers should be health?
  * Frog Obstical, You need to feed him in order to get him to move
* Add interpreted scripting language to make it more aproacable and to prevent golang from being an install requirement