# GOLF Engine
the GoLang Fantasy Engine. Right now this is just a convient place for the todo list. No fun stuff yet.

### TODO
---
* Make the map tool respect SprOpts, or make a separete MapOpts for pallet swaping
* Make tools for importings maps.
  * Import CSV as map
  * Import image as map
* Add function to load map data (reverse order)
* Add function to load sprite flag data
* Determine final codebase structure (seperate repo for demo games that are linked in the readme)
* Startup Animation?

### DONE
---
* Create csv format for the flags (Probably just 1's and 0's) [x]
* Import sprite flag csv (Should be pretty easy) [x]
* Create the Golf ToolKit [x]
* Add about and help commands [x]
* Add clear and !! command to the go toolkit [x]

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
  * Make sure to features the Go Gopher to demo a diologue like system
* Make Blood The Game to demo the engine
  * Black and white 1 bit diablo like game
  * Use pallet swaps to go from day to night
  * Creppy story told though small conversations with the old man at the church
  * Fight demony thingies
  * Maybe do celest hair on the character (should be bold red color)
  * Make red blood that you collect as enemies die
* Add interpreted scripting language to make it more aproacable and to prevent golang from being an install requirement

### To Consider
---
* should the map size be configurable?
  * the top 2 bytes could be used to configure the width and height (1-256)
  * 128x128 might be a bit too narrow for some games
  * 256x64 might be a bit too short for others
  * additionally that would give more flexibility memorywise since you can load the map into the sprites
  * however it could add a lot of unwanted complexity
  * how woud this be implemented?

### Fix
* oo for the filled in circle is a bad letter sequence. Picke a sequence that is less common in regular words