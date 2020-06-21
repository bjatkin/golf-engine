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

// CameraX: 0x3603-0x3604
const cameraX = 0x3603

// CameraY: 0x3605-0x3606
const cameraY = 0x3605

// Frames: 0x3607-0x3609
const frames = 0x3607

// ClipX: 0x360A
const clipX = 0x360A

// ClipY: 0x360B
const clipY = 0x360B

// ClipW: 0x360C
const clipW = 0x360C

// ClipH: 0x360D
const clipH = 0x360D

// Mouse:
//  X: 0x360E
const mouseX = 0x360E

//	Y: 0x360F
const mouseY = 0x360F

//	Left Click: 0x3610
//	Middle Click: 0x3610
//	Right Click: 0x3610
//	Mouse Style: 0x3610
const mouseBase = 0x3610

// Keyboard: 0x3611-0x3647
const keyBase = 0x3611

// InternalSpriteSheet: 0x3648-0x3F48 [0x0900]
const internalSpriteColBase = 0x3648
const internalSpritePalBase = 0x3C48

// SpriteSheet: 0x3F49-0x6F49 [0x3000]
// these need to be declared as vars so that we can
// swap to the internal sprite sheet for text/ mouse/ logo drawing
var spriteColBase = 0x3F49
var spritePalBase = 0x5F49
