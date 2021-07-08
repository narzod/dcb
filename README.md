# dcb - draw checkerboard

dcb is a program to create simple images of checkboard patterns from the command line.

# Usage

## Command-line options (not implemenented yet)

```
dcb - draw checkboard patterns from the command line

USAGE:
    dcb [options] <command> [<arguments...>]

OPTIONS:
    -o, --outfile <FILENAME>     filename to output
    -c1 <HEX COLOR CODE>
    -c2 <HEX COLOR CODE>
    -c3 <HEX COLOR CODE>
    -c3 <HEX COLOR CODE>         specify border color for images which have border width > 0. Border width set with -bw.
    -ct <theme-name>             specify a color pallette
    -t <WxH>                     tile dimensions in pixels
    -m <WxH>                     matrix dimensions in tiles. Tiles are not necessarily squares.
    -n int                       number of frames for animated images
    -d int                       duration per frame for animated images
    -f <FILE FORMAT>             specify SVG or GIF format

```
See: http://jonasjacek.github.io/colors/
  
  
