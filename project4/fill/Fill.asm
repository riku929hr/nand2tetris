// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

//// Replace this comment with your code.
// infinite loop
(LOOP)
    // i=0
    @i
    M=0
    // input keyboard
    @KBD
    D=M
    // if KBD != 0, then jump to BLACK
    @BLACK
    D;JGT

    (WHITE)
        // if i>=8192, then jump to LOOP
        @i
        D=M
        @8192
        D=D-A
        @LOOP
        D;JGE

        // R[i+SCREEN] = 0
        @i 
        D=M // read i
        @SCREEN
        A=D+A // i+SCREEN
        M=0

        // i++
        @i
        M=M+1
    @WHITE
    0;JMP
            
    (BLACK)
        // if i>=8192, then jump to LOOP
        @i
        D=M
        @8192
        D=D-A
        @LOOP
        D;JGE

        // R[i+SCREEN] = -1
        @i 
        D=M // read i
        @SCREEN
        A=D+A // i+SCREEN
        M=-1

        // i++
        @i
        M=M+1

    // back to BLACK loop
    @BLACK
    0;JMP

@LOOP
0;JMP

(END)
@END
0;JMP
