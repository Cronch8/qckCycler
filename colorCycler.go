package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	// Check for root permissions
	if os.Geteuid() != 0 {
		fmt.Println("Running the prism colorCylcler for the QcKPrism program requires root permissions.")
		return
	}

    // change this to where you cloned and compiled the qckprism binary
    const executablePath = "./../qckprism/target/debug/qckprism"

	// variables
    var speed float64 = 0.15
	const updateRate float64 = 0.01
	var t float64 = -100000// -100k to double the time before a reset is required (should last for a few days)
    var gradient [4][3]float64

    // parse arguments
    if len(os.Args) == 2 {
        preset, _ := strconv.ParseInt(os.Args[1], 10, 0)
        gradient = selectColor(int(preset))
    } else if len(os.Args) == 3 {
        preset, _ := strconv.ParseInt(os.Args[1], 10, 0)
        speed, _ = strconv.ParseFloat(os.Args[2], 32)
        gradient = selectColor(int(preset))
    } else {
        gradient = selectColor(3)
    }


	for {
		// Calculate RGB based on the cosine gradient
		r := cosineGradient(t, gradient[0])
		g := cosineGradient(t, gradient[1])
		b := cosineGradient(t, gradient[2])

		r_ := cosineGradient(t+math.Pi*4, gradient[0])
		g_ := cosineGradient(t+math.Pi*4, gradient[1])
		b_ := cosineGradient(t+math.Pi*4, gradient[2])

		// Convert to 8-bit RGB values
		r8 := int(math.Max(math.Min(math.Round(r * 255), 255), 0))
		g8 := int(math.Max(math.Min(math.Round(g * 255), 255), 0))
		b8 := int(math.Max(math.Min(math.Round(b * 255), 255), 0))
		r8_ := int(math.Max(math.Min(math.Round(r_ * 255), 255), 0))
		g8_ := int(math.Max(math.Min(math.Round(g_ * 255), 255), 0))
		b8_ := int(math.Max(math.Min(math.Round(b_ * 255), 255), 0))

		// Format as hex strings
		colorA := fmt.Sprintf("%02x%02x%02x", r8, g8, b8)
		colorB := fmt.Sprintf("%02x%02x%02x", r8_, g8_, b8_)

		// Call qckprism with calculated colors
        cmd := exec.Command(executablePath, "-b", colorA, "-a", colorB)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout
        //fmt.Printf("Running command: ./target/debug/qckprism -a %s -b %s\n", colorA, colorB)

        if err := cmd.Run(); err != nil {
            fmt.Println("Error running qckprism:", err)
            return
        }

		// Increment time and loop
		t += updateRate*speed
		if t >= 100000 {
			t = -100000
		}
        thing := updateRate * float64(time.Second) / speed
		time.Sleep(time.Duration(thing))
	}
}

// cosineGradient calculates a cosine-based gradient for a specific time value
func cosineGradient(t float64, gradient [3]float64) float64 {
	return gradient[0] + gradient[1]*math.Cos(2*math.Pi*(t+gradient[2]))
}

func selectColor(n int) [4][3]float64 {
    switch n {
    case 0:
        // i dont remember what this one was
        return [4][3]float64{
            {0.911, 0.863, 0.800},
            {0.042, 0.205, 0.921},
            {0.336, 1.079, 0.706},
            {0.295, 3.962, 2.978},
        }
    case 1:
        // dark green into Venecity
        return [4][3]float64{
            {0.538, 0.358, 0.934},
            {0.428, 0.698, 0.222},
            {1.500, 2.000, 2.000},
            {0.117, 3.523, 2.968},
        }
    case 2:
        // sexy police into jungle
        return [4][3]float64{
            {0.360, 0.252, 0.934},
            {0.250, 0.530, 0.222},
            {0.857, 0.888, 1.418},
            {0.117, 3.523, 2.968},
        }
    case 3:
        // i have given up on naming these
        return [4][3]float64{
            {0.938, 0.328, 0.718},
            {0.659, 0.438, 0.328},
            {0.388, 0.388, 0.296},
            {2.538, 2.478, 0.168},
        }
    default:
        // Venecity - orage-purple into blue-green
        return [4][3]float64{
            {0.538, 0.718, 1.028},
            {0.468, 0.468, -0.422},
            {1.000, 1.000, 1.000},
            {0.000, 0.468, 0.987},
        }
    }
}
