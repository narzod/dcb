# dcb - draw checkerboard

`dcb` is a program to create simple images of checkboard patterns from the command line.

# Usage

## Command-line options

```
dcb - draw checkboard patterns from the command line

USAGE:
    dcb [options] <command> [<arguments...>]

OPTIONS:
    -v   --version           print version info
    -h   --help              print this help
    -o   --outfile           output filename (default out.gif)
    -fc  --fgcolor           tile foreground color (first tile color)
    -bc  --bgcolor           tile background color (border color)
    -bw  --borderwidth       border_width_in_pixels x border_fx_multiplier (AxB)
    -i   --invert            swap colors specied in -fc and -bc
    -t   --tilesize          tile_width x height_in_pixels (AxB)
    -m   --matrixsize        matrix_width x height_in_tiles (AxB)
    -n   --frames            number of frames to include
    -d   --duration          frame duration for animated images in ms
    -f   --format            specify SVG, PNG or GIF format
    -s   --imagesize         image width x height in pixels (overrides -t) (AxB)
    -fx  --borderfx          0=none, 1=alternating, 2=gradient

```
See color info: http://jonasjacek.github.io/colors/ and https://www.astrouw.edu.pl/~jskowron/colors-x11/rgb.html

## Examples

```
dcb -t 32x16 -m 20x10 -bw 1 -fc red -bc green -o test1.gif -i
```

Red and blue chess board. No borders.

```
dcb -t 50x50 -m 8x8 -fc DarkRed -bc SteelBlue -o test2.gif -i colorparings.txt
```

A screenful of red with just one border.
```
dcb -t 600x420 -m 1x1 -fc Black -bc Yellow
```

## Sample input file for sequence of paired colors

```
"Steel Blue"  "Papaya Whip"
"Dark Red"    "Antique White"
"HotPink"     Black
```

```
cat colorpairs.txt | dcb -t 40x40 -m 8x8 -fc DarkRed -bc SteelBlue -o test2.gif -fx=1 -bw=2x3 -i -
```
