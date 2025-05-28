package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreGraphics -framework Foundation
#include <CoreGraphics/CoreGraphics.h>
#include <stdio.h>

void moveMouseTo(int x, int y) {
    CGPoint point = CGPointMake(x, y);
    CGEventRef event = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, point, kCGMouseButtonLeft);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void getMousePosition(int *x, int *y) {
    CGEventRef event = CGEventCreate(NULL);
    CGPoint point = CGEventGetLocation(event);
    *x = (int)point.x;
    *y = (int)point.y;
    CFRelease(event);
}
*/
import "C"

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func moveMouse() {
	// Get current mouse position
	var currentX, currentY C.int
	C.getMousePosition(&currentX, &currentY)

	// Move mouse a small random amount and then back
	offsetX := rand.Intn(21) - 10 // Random number between -10 and 10
	offsetY := rand.Intn(21) - 10 // Random number between -10 and 10

	// Move mouse slightly away from current position
	C.moveMouseTo(currentX+C.int(offsetX), currentY+C.int(offsetY))
	time.Sleep(100 * time.Millisecond) // 0.1 second duration

	// Return to original position
	C.moveMouseTo(currentX, currentY)
	time.Sleep(100 * time.Millisecond) // 0.1 second duration

	fmt.Printf("Mouse moved at %s\n", time.Now().Format("15:04:05"))
}

func main() {
	fmt.Println("Mouse jiggler started. Press Ctrl+C to exit.")

	// Set up signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Create a ticker for 3 minutes (180 seconds)
	ticker := time.NewTicker(3 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			moveMouse()
		case <-c:
			fmt.Println("\nMouse jiggler stopped.")
			return
		}
	}
}
