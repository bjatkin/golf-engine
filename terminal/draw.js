// Get the canvas and context
var canvas = document.getElementById("mycanvas");
var context = canvas.getContext("2d");
var width = 320;
var height = 288;
var imagedata = context.createImageData(width, height);

var pixelData = new Uint8Array(((width * height)/2)+1);

// Prevent the context menue on the canvas so left click works
canvas.addEventListener('contextmenu', e=> e.preventDefault());

// Create the image
function drawScreen() {
    // Set the pallets from the first byte of data
    palletMasks = [
        0b00000111,
        0b01110000,
    ]
    a = palletList[pixelData[0] & 0b00000111]
    b = palletList[(pixelData[0] & 0b01110000) / 16]
    for (i = 0; i < 8; i++) {
        pallet[i] = a[i]
        pallet[i+8] = b[i]
    }
 
    //CONVERT 4-BIT DATA FROM WASM
    for (i = 1; i < pixelData.length; i++) {
        // First Pixel
        index = (i-1)*8;
        p = 0b00001111 & pixelData[i];
        imagedata.data[index] = pallet[p][0];
        imagedata.data[index+1] = pallet[p][1];
        imagedata.data[index+2] = pallet[p][2];
        imagedata.data[index+3] = 255;

        // Second Pixel
        index += 4;
        p = (0b11110000 & pixelData[i])/16;
        imagedata.data[index] = pallet[p][0];
        imagedata.data[index+1] = pallet[p][1];
        imagedata.data[index+2] = pallet[p][2];
        imagedata.data[index+3] = 255;

        lastIndex = index
    }

    // Draw the image data to the canvas
    context.putImageData(imagedata, 0, 0);
}

// create pallets https://coolors.co/

//white and red colors
pallet = [ 
    //Pallet 1
    [0, 0, 0],
    [36, 36, 36],
    [73, 73, 73],
    [108, 108, 108],
    [144, 144, 144],
    [180, 180, 180],
    [216, 216, 216],
    [255, 255, 255],
    //Pallet 2
    [15, 0, 0],
    [36, 10, 10],
    [73, 20, 20],
    [108, 30, 30],
    [144, 40, 40],
    [180, 50, 50],
    [216, 60, 60],
    [255, 70, 70],
]

white = [
    [0, 0, 0],
    [36, 36, 36],
    [73, 73, 73],
    [108, 108, 108],
    [144, 144, 144],
    [180, 180, 180],
    [216, 216, 216],
    [255, 255, 255],
]

red = [
    [15, 0, 0],
    [36, 10, 10],
    [73, 20, 20],
    [108, 30, 30],
    [144, 40, 40],
    [180, 50, 50],
    [216, 60, 60],
    [255, 70, 70],
]

green = [
    [0, 15, 0],
    [10, 36, 10],
    [20, 73, 20],
    [30, 108, 30],
    [40, 144, 40],
    [50, 180, 50],
    [60, 216, 60],
    [70, 255, 70],
]

darkGreen = [
    [0, 15, 40],
    [10, 36, 50],
    [20, 73, 60],
    [30, 108, 70],
    [40, 144, 80],
    [50, 180, 90],
    [60, 216, 100],
    [70, 255, 110],
]

blue = [
    [0, 0, 15],
    [10, 10, 36],
    [20, 20, 73],
    [30, 30, 108],
    [40, 40, 144],
    [50, 50, 180],
    [60, 60, 216],
    [70, 70, 255],
]

darkBlue = [
    [40, 0, 15],
    [50, 10, 36],
    [60, 20, 73],
    [70, 30, 108],
    [80, 40, 144],
    [90, 50, 180],
    [100, 60, 216],
    [110, 70, 255],
]

yellow = [
    [15, 15, 0],
    [36, 36, 10],
    [73, 73, 20],
    [108, 108, 30],
    [144, 144, 40],
    [180, 180, 50],
    [216, 216, 60],
    [255, 255, 70],
]

purple = [
    [15, 0, 15],
    [36, 10, 36],
    [73, 20, 73],
    [108, 30, 108],
    [144, 40, 144],
    [180, 50, 180],
    [216, 60, 216],
    [255, 70, 255]
]

palletList = [
    white,
    red,
    green,
    darkGreen,
    blue,
    darkBlue,
    yellow,
    purple,
]
