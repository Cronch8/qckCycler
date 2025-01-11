# QcK Prism Color Cycler

A simple Go project to create changing RGB colors for SteelSeries QcK Prism mousepads using cosine-based gradients ([gradient idea came from here, use this to make new gradients](https://github.com/thi-ng/cgg?tab=readme-ov-file)).

There are a few predifined color cycles, but feel free to modify them for yourself using the tool linked above.

## Prerequisites
1. **Rust and Cargo**: Required to compile the [qckprism](https://github.com/zapp88/qckprism/tree/master) binary.
2. **SteelSeries QcK Prism Mousepad**: Why would you use this if you dont have the mouspad itself?

The Go Programming Language is Required if you want to compile the color cycler yourself.

## usage
1. Clone this repository
2. Build the `qckprism` tool
   Ensure the compiled binary is located at `./target/debug/qckprism`.
3. change the path at the start of the program to point to the compiled executable of qckprism
4. Compile the Go program (if you want)
5. dont complain how janky it is to do it this way is because i made this tool for myself and i didnt really think about user frendliness
6. Run the color cycler with root permissions, then it can pass those perms onto the qckprism tool.
```bash
sudo ./colorCycler [gradient number]
```
This utility itself doesnt use root perms, just uses it to run the qckprism with root each time it updates the colours. I know an udev rule can be used (as is mentioned onthe qckprism github page) but i couldnt get it to work so i just used sudo

### Gradient Options
- **Default**: Orange-purple into blue-green.
- `0`: Soft neutral gradient.
- `1`: Dark green transitioning into orange-blue
- `2`: Vibrant "sexy police" transitioning into green-ish jungle.

for example
```bash
sudo ./colorCycler 1
```
If no gradient ID is provided, the default gradient is used.

(Yes half of this readme was generated with ai, i'm lazy alright)
