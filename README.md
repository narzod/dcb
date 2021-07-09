# dcb - draw checkerboard

dcb is a program to create simple images of checkboard patterns from the command line.

# Usage

## Command-line options

```
dcb - draw checkboard patterns from the command line

USAGE:
    dcb [options] <command> [<arguments...>]

OPTIONS:
    -v, --version
    -h, --help
    -o, --outfile                output filename                                        -- default out.gif
    -fc, --fgcolor               tile foreground color (first tile color)               -- default white
    -bc, --bgcolor               tile background color (border color)                   -- default black
    -bw, --borderwidth           [border width in pixels] x [border fx multiplier]      -- default 1x1
    -i, --invert                 swap colors specied in -fc and -bc
    -t, --tilesize               tile width x height in pixels                          -- default 80x80
    -m, --matrixsize             matrix width x height in tiles                         -- default 8x8
    -n, --frames                 number of frames to include                            -- default 1
    -d, --duration               frame duration for animated images in ms               -- default 100                           
    -f, --format                 specify SVG or GIF format                              -- default GIF
    -s, --imagesize              image width x height in pixels (overrides -t)
    -fx, --borderfx              0=none, 1=gradient, 2=alternating

```
See: http://jonasjacek.github.io/colors/

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
SteelBlue  SeaFoam
DarkRed    Yellow
HotPink    White
```

```
cat colorpairs.txt | dcb -t 40x40 -m 8x8 -fc DarkRed -bc SteelBlue -o test2.gif -fx=1 -bw=2x3 -i -
```
