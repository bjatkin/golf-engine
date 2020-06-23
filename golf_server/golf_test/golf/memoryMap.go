package golf

// Addresses
// ScreenBuff: 0x0000 - 0x3601
//  Col Buff: 0x0000 - 0x23FF
const screenColBuffBase = 0x0000

// 	Pal Buff: 0x2400 - 0x35FF
const screenPalBuffBase = 0x2400

//  Pal Set: 0x3600
const screenPalSet = 0x3600

// BG Color: 0x3601 - high 3 bits
const bgColor = 0x3601

// CameraX: 0x3602-0x3603
const cameraX = 0x3602

// CameraY: 0x3604-0x3605
const cameraY = 0x3604

// Frames: 0x3606-0x3608
const frames = 0x3606

// ClipX: 0x3609
const clipX = 0x3609

// ClipY: 0x360A
const clipY = 0x360A

// ClipW: 0x360B
const clipW = 0x360B

// ClipH: 0x360C
const clipH = 0x360C

// Mouse:
//  X: 0x360D
const mouseX = 0x360D

//	Y: 0x360E
const mouseY = 0x360E

//	Left Click: 0x360F
//	Middle Click: 0x360F
//	Right Click: 0x360F
//	Mouse Style: 0x360F
const mouseBase = 0x360F

// Keyboard: 0x3610-0x3646
const keyBase = 0x3610

// InternalSpriteSheet: 0x3647-0x3F47 [0x0900]
const internalSpriteColBase = 0x3647
const internalSpritePalBase = 0x3C47

// SpriteSheet: 0x3F48-0x6F48 [0x3000]
const spriteColBase = 0x3F48
const spritePalBase = 0x5F48

// ActiveSpriteBuff: 0x6F49-0x6F4B
const activeSpriteColBuff = 0x6F49
const activeSpritePalBuff = 0x6F4B

// MapData (256x64 / 512 8x8): 0x6F4D - 0xB74D
const mapBase = 0xB74D //Start from the high memory so the map grows down

// SpriteFlags (512 8x8): 0xB74E - B94E
const spriteFlags = 0xB74E

//My Guess Is that i'll need about 10-20k bytes for music

/* TODO:
[x] Fix fantasyfong.png
[] Add a textF function? for formated text?
[] Add startup animation? How to display logo?
[] Add function to load map data (Make sure to load it in reverse)
[] Interleave pixel data the same way the map is done to make the map sprite swap easier
	- col0[00112233] col1[44556677] pal[01234567]
[] Fix sprite importing
[] how should the code be structured? so that it just works?
[] Add in a penalty for the load sprites and load map (60-120 frames?) to prevent new users from
	messing with multiple sprite sheets
[] when you download the github you should be able to __ go run [golf_term] __ in any directory in order to
	set up a project or projects in that directory
[] golf_term should have a map editor, sprite editor, sprite flag editor and be able to save those to files
[] Add func to load sprite flags? Or can this be part of loading the sprites? (maybe just a 512 byte header???)
*/
