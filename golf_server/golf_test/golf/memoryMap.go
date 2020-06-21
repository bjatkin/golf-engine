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

// ActiveSpriteBuff: 0x6F49-0x6FB
const activeSpriteColBuff = 0x6F49
const activeSpritePalBuff = 0x6F4B

//My Guess Is that i'll need about 10k bytes for music
