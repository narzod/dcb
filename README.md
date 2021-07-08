# dcb - draw checkerboard

dcb is a program to create simple images of checkboard patterns from the command line.

# Usage

## Command-line options (not implemenented yet)

```
dcb - draw checkboard patterns from the command line

USAGE:
    dcb [options] <command> [<arguments...>]

OPTIONS:
    -v, --version
    -h, --help
    -o, --outfile <FILENAME>     output filename
    -ct, --theme                 specify a thee-color pallette by name
    -bw, --borderwidth           border witdh for tile
    -c1, --color1 <COLOR NAME>   color override for first color in pallette
    -c2, --color2 <COLOR NAME>   color override for second color in pallette
    -c3, --color3 <COLOR NAME>, --bordercolor
                                 color override for third color in pallette which specifies border color.
                                 Used only for images which have border width > 0. Border width set with -bw.
    -i, --invert                 swap first two colors in pallette
    -ts, --tilesize              tile width and height in pixels
    -ms, --matrixsize            matrix width and height in tiles. Tiles are not necessarily squares.
    -n, --frames                 number of frames used for animated images
    -d, --duration               duration per frame for animated images
    -f, --format <FILE FORMAT>   specify SVG or GIF format
    -s, --imagesize              image width and height in pixels. Overrides tilesize option.

```
See: http://jonasjacek.github.io/colors/
  
  
