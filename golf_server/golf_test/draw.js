// Get the canvas and context
var canvas = document.getElementById("golfcanvas");
var context = canvas.getContext("2d");
var imagedata = context.createImageData(768, 768);

// Create the shared video buffer
var screenCol = (192/4)*192;
var screenPal = (192/8)*192;
var setPal = 1;
var screenBuff = new Uint8Array(setPal+screenCol+screenPal);

// Prevent the context menue on the canvas so left click works
canvas.addEventListener('contextmenu', e=> e.preventDefault());

// Engine Pallets
var pallets = [
    // Pallet 0
    [[0, 0, 0], //black
    [84, 84, 84], // dark gray
    [168, 168, 168], // light gray
    [255, 255, 255]], // white
    // Pallet 1
    [[54, 36, 26], // Brown
    [119, 97, 65], // Cream
    [181, 165, 125], // Light Cream
    [218, 207, 190]], // Pale
    // Pallet 2
    [[75, 6, 16], // Maroon
    [128, 8, 19], // Dark Red
    [194, 10 , 23], // Red
    [243, 13, 30]], // Light Red
    // Pallet 3
    [[245, 61, 74], // Dark Pink
    [248, 110, 120], // Pink
    [250, 180, 158], // Light Pink
    [251, 214, 189]], //Salmon
    // Pallet 4
    [[29, 44, 29], // Dark Green
    [79, 124, 69], // Mid Green
    [154, 185, 121], // Light Green
    [214, 227, 175]], // Pale Green 2
    // Pallet 5
    [[51, 77, 44], // Tween Green
    [116, 168, 88], // Nuther Green
    [205, 222, 170], // Green Light
    [232, 237, 191]], // Pale Green
    // Pallet 6
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 7
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 8
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 9
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 10
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 11
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 12
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 13
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 14
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 15
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
    // Pallet 16
    [[0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255], // REPLACE THIS
    [0, 0, 255]], // REPLACE THIS
];

var fPallet = [0,1,2,3,4,5,6,7]; // Place Holder Data
// Create the image
function drawScreen() {
    // Set the pallets from the first byte of data
    a = pallets[(screenBuff[screenCol+screenPal] & 0b11110000) / 16];
    b = pallets[(screenBuff[screenCol+screenPal] & 0b00001111)];
    for (let i = 0; i < 4; i++) {
        fPallet[i] = a[i];
        fPallet[i+4] = b[i];
    }

    // Fill the Screen Buffer
    c =[0,0,0,0,0,0,0,0];
    p =[0,0,0,0,0,0,0,0];
    for (let i=0; i < screenPal; i++) {
        c[0] = 0b00000011 & screenBuff[i*2];
        p[0] = 0b00000001 & screenBuff[screenCol+i];
        c[1] = (0b00001100 & screenBuff[i*2]) / 4;
        p[1] = (0b00000010 & screenBuff[screenCol+i]) / 2;
        c[2] = (0b00110000 & screenBuff[i*2]) / 16;
        p[2] = (0b00000100 & screenBuff[screenCol+i]) / 4;
        c[3] = (0b11000000 & screenBuff[i*2]) / 64;
        p[3] = (0b00001000 & screenBuff[screenCol+i]) / 8;
        c[4] = (0b00000011 & screenBuff[i*2+1]);
        p[4] = (0b00010000 & screenBuff[screenCol+i]) / 16;
        c[5] = (0b00001100 & screenBuff[i*2+1]) / 4;
        p[5] = (0b00100000 & screenBuff[screenCol+i]) / 32;
        c[6] = (0b00110000 & screenBuff[i*2+1]) / 16;
        p[6] = (0b01000000 & screenBuff[screenCol+i]) / 64;
        c[7] = (0b11000000 & screenBuff[i*2+1]) / 64;
        p[7] = (0b10000000 & screenBuff[screenCol+i]) / 128;
        
        let x = i*8 % 192
        let y = Math.floor(i / 24)
        for(let j=0; j < 8; j++) {
            let col = fPallet[c[j]+(p[j]*4)]
            drawHiResPixel(x+j, y, col)
        }
    }
 
    // Draw the image data to the canvas
    context.putImageData(imagedata, 0, 0);
}

function drawHiResPixel(x, y, rgb) {
    s = x*4 + y*192*4*4
    for(let i = s; i < s+4; i++) {
        for(let j = 0; j < 4; j++) {
            imagedata.data[i*4+j*192*4*4] = rgb[0]
            imagedata.data[i*4+1+j*192*4*4] = rgb[1]
            imagedata.data[i*4+2+j*192*4*4] = rgb[2]
            imagedata.data[i*4+3+j*192*4*4] = 255
        }
    }
}