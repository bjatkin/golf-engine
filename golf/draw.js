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

let pallets = [[[0, 0, 0], [96, 96, 96], [144, 144, 144], [192, 192, 192]],[[64, 64, 64], [128, 128, 128], [160, 160, 160], [255, 255, 255]],[[21, 3, 7], [74, 18, 21], [174, 48, 49], [236, 126, 124]],[[48, 11, 14], [126, 33, 35], [230, 63, 63], [242, 189, 184]],[[22, 6, 2], [72, 44, 11], [220, 155, 35], [235, 185, 81]],[[47, 25, 7], [121, 81, 19], [228, 170, 58], [250, 215, 126]],[[25, 21, 0], [155, 90, 16], [210, 118, 20], [242, 178, 123]],[[2, 24, 12], [74, 102, 61], [124, 153, 100], [178, 203, 144]],[[50, 76, 45], [97, 128, 77], [151, 178, 122], [205, 227, 166]],[[7, 13, 17], [58, 119, 126], [108, 225, 234], [148, 232, 230]],[[33, 66, 72], [83, 172, 180], [128, 229, 232], [188, 238, 226]],[[7, 1, 26], [21, 34, 83], [35, 66, 165], [59, 104, 191]],[[14, 18, 55], [28, 50, 111], [47, 85, 165], [82, 141, 242]],[[21, 34, 83], [70, 135, 143], [148, 227, 68], [226, 243, 228]],[[0, 48, 59], [255, 119, 119], [255, 206, 150], [241, 242, 218]],[[0, 0, 0], [197, 17, 17], [20, 58, 133], [255, 255, 255]]]