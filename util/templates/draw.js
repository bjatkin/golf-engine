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
    [224, 60, 40],
    [255, 255, 255],
    [215, 215, 215]],
    [[168, 168, 168],
    [123, 123, 123],
    [52, 52, 52],
    [21, 21, 21]],
    [[13, 32, 48],
    [65, 93, 102],
    [113, 166, 161],
    [189, 255, 202]],
    [[37, 226, 205],
    [10, 152, 172],
    [0, 82, 128],
    [0, 96, 75]],
    [[32, 181, 98],
    [88, 211, 50],
    [19, 157, 8],
    [0, 78, 0]],
    [[23, 40, 8],
    [55, 109, 3],
    [106, 180, 23],
    [140, 214, 18]],
    [[190, 235, 113],
    [238, 255, 169],
    [182, 193, 33],
    [147, 151, 23]],
    [[204, 143, 21],
    [255, 187, 49],
    [255, 231, 55],
    [246, 143, 55]],
    [[173, 78, 26],
    [35, 23, 18],
    [92, 60, 13],
    [174, 108, 55]],
    [[197, 151, 130],
    [226, 215, 181],
    [79, 21, 7],
    [130, 60, 61]],
    [[218, 101, 94],
    [225, 130, 137],
    [245, 183, 132],
    [255, 233, 197]],
    [[255, 130, 206],
    [207, 60, 113],
    [135, 22, 70],
    [163, 40, 179]],
    [[204, 105, 228],
    [213, 156, 252],
    [254, 201, 237],
    [226, 201, 255]],
    [[166, 117, 254],
    [106, 49, 202],
    [90, 25, 145],
    [33, 22, 64]],
    [[61, 52, 165],
    [98, 100, 220],
    [155, 160, 239],
    [152, 220, 255]],
    [[91, 168, 255],
    [10, 137, 255],
    [2, 74, 202],
    [0, 23, 125]],
]