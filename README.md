# GOLF Engine
the GoLang Fantasy Engine. Right now this is just a convient place for the todo list. No fun stuff yet.

### TODO
---
* Change the sprite functions from using ints to using floats
* Add function to load sprite flag data
* Add startup animation when you start the game
* Create and example golang server for using this framework for playing games online
* Build out the readme so there is good documentation for the API
* Add instructions for installing and playing the golf examples
* Test the golf toolkit on a windows machine

### DONE
---
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