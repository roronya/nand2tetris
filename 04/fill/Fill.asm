// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed.
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

//for(int h=0;h<256;h++){
//  for(int i=0;i<32;i++) {
//    M[h][i]=32767;
//    M[@SCREEN+32*h+i]
//  }
//}
//

// max_screen_register = 8196
@8196
D=A
@max_screen_register
M=D

(LOOP)
  // i = 0
  @i
  M=0

  @KBD
  D=M
  @FILL_SCREEN_LOOP
  D;JNE
  @EMPTY_SCREEN_LOOP
  D;JMP
  (FILL_SCREEN_LOOP)
    // current_scrreen_register = 16384 + i
    // Memory[current_screen_register] = -1 = 111...11
    @SCREEN
    D=A
    @i
    A=D+M
    M=-1

    // while(max_screen_register -i == 0)
    @i
    M=M+1
    D=M
    @max_screen_register
    D=M-D
    @FILL_SCREEN_LOOP_END
    D;JEQ
    @FILL_SCREEN_LOOP
    0;JMP
  (FILL_SCREEN_LOOP_END)

  (EMPTY_SCREEN_LOOP)
    // current_scrreen_register = 16384 + i
    // Memory[current_screen_register] = -1 = 111...11
    @SCREEN
    D=A
    @i
    A=D+M
    M=0

    // while(max_screen_register -i == 0)
    @i
    M=M+1
    D=M
    @max_screen_register
    D=M-D
    @EMPTY_SCREEN_LOOP_END
    D;JEQ
    @EMPTY_SCREEN_LOOP
    0;JMP
  (EMPTY_SCREEN_LOOP_END)
(END)
@END
0;JMP
