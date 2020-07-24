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

var fPallet = [0,1,2,3,4,5,6,7]; // Place Holder Data
// Create the image
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

function drawScreen() {
    // Set the pallets from the first byte of data
    a = pallets[(screenBuff[screenCol+screenPal] & 0b11110000) / 16];
    b = pallets[(screenBuff[screenCol+screenPal] & 0b00001111)];
    for (let i = 0; i < 4; i++) {
        fPallet[i] = a[i];
        fPallet[i+4] = b[i];
    }

    // Fill the Screen Buffer with pixels
    for (let i=0; i < 192*192; i++) {
        let index = Math.floor(Math.floor(i/4) / 2 * 3)
        let pIndex = index + (2-index%3)
        let shift = (i % 4) * 2
        let pShift = i % 8

        let c = (0b00000011 << shift  & screenBuff[index]) >> shift
        let p = (0b00000001 << pShift & screenBuff[pIndex]) >> pShift
        let color = fPallet[c + p*4]
        x = Math.floor(i % 192)
        y = Math.floor(i / 192)

        drawHiResPixel(x, y, color)
    }
 
    // Draw the image data to the canvas
    context.putImageData(imagedata, 0, 0);
}

let pallets = [
    [[0, 0, 0],
    [64, 64, 64],
    [96, 96, 96],
    [128, 128, 128]],

    [[144, 144, 144],
    [160, 160, 160],
    [192, 192, 192],
    [255, 255, 255]],

    [[21, 3, 7],
    [48, 11, 14],
    [74, 18, 21],
    [126, 33, 35]],

    [[178, 48, 49],
    [230, 63, 63],
    [236, 126, 124],
    [242, 189, 184]],

    [[22, 6, 2],
    [47, 25, 7],
    [72, 44, 11],
    [121, 81, 19]],

    [[220, 155, 35],
    [228, 170, 58],
    [235, 185, 81],
    [250, 215, 126]],

    [[2, 24, 12],
    [50, 76, 45],
    [74, 102, 61],
    [97, 128, 77]],

    [[124, 153, 100],
    [151, 178, 122],
    [178, 203, 144],
    [205, 227, 166]],

    [[7, 13, 17],
    [33, 66, 72],
    [58, 119, 126],
    [83, 172, 180]],

    [[108, 225, 234],
    [128, 229, 232],
    [148, 232, 230],
    [188, 238, 226]],

    [[7, 1, 26],
    [14, 18, 55],
    [21, 34, 83],
    [28, 50, 111]],

    [[35, 66, 139],
    [47, 85, 165],
    [59, 104, 191],
    [82, 141, 242]],

    [[0, 0, 0],
    [197, 17, 17],
    [20, 58, 133],
    [255, 255, 255]],

    [[181, 98, 11],
    [241, 131, 14],
    [243, 152, 56],
    [246, 178, 105]],

    [[246, 189, 96],
    [247, 237, 226],
    [245, 202, 195],
    [132, 165, 157]],

    [[21, 34, 83],
    [70, 135, 143],
    [148, 227, 68],
    [226, 243, 228]],
]