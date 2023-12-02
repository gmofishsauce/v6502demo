**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# Incrementers and adders - VisualChips

## Incrementers and adders

#### From VisualChips

Two common circuits are the incrementer and the adder.  An incrementer
takes one input number, and adds 1 to it; an adder takes two input
numbers and adds them together, and possibly adds a single bit as
well, the "carry in".

There are three ways those circuits are typically implemented in small
(NMOS) circuits: bit-serial, alternating polarity carry chain, and
Manchester carry chain.

**Contents**

- [Bit-serial incrementer](#bit-serial-incrementer)
- [Alternating polarity carry chain](#alternating-polarity-carry-chain)
- [Manchester carry chain](#manchester-carry-chain)
- [Resources](#resources)

#### Bit-serial incrementer

The idea of the bit-serial circuit is to handle one bit per clock, the
bit with the lowest numerical value first.  At every step, there will
be a carry in (which is 1 for the first step), and a carry out (which
becomes the carry in for the next step).  There also of course is an
output bit at every step.

The output bit will be XOR (exclusive-or) of the input bit and the
carry in; the carry out will be the logical AND if the input bit and
the carry in.

#### Alternating polarity carry chain

Instead of clocking N times for N bits, you can put N of those one-bit
circuits in series.  This would ideally make it then work in one clock
period instead of in N.  However, things are not so simple.

It takes time for a logic gate to produce an output.  The time it takes
from when the last input becomes valid to when the output of the gate
becomes valid is called the "propagation delay".  When you tie many
gates together, the propagation delay from any input to any output
should be as small as possible.

The critical path is the carry chain: if all bits in the input are 1,
you get a carry out from every bit, but that output doesn't become
valid until some time after that AND gate's input (the previous carry
bit).

AND gates in NMOS are actually two fundamental gates in series (a NAND
and a NOT).  This would make the critical path take 2N steps.  Luckily,
there is a trick.

Instead of computing the carry for every bit, you comput the inverse of
the carry for the carry out of all the even-numbered bits, and the regular
carry for others.  So for the even bits, instead of the AND gate, you
get a NAND gate, which is a single fundamental gate.  For the odd bits,
you get an AND with one of its inputs complemented.  Now, a NOR gate is
the same as an AND with both inputs complemented.  That is easy to do:
just put an inverter on the input bit (it is not on the critical path,
so it won't hurt!)

Now the carry chain is alternatingly a NAND gate and a NOR gate, only
N fundamental gates total.

#### Manchester carry chain

- tired now, will finish it later*
- TODO: pics!*

#### Resources

- [Wikipedia](http://en.wikipedia.org/wiki/Carry_look-ahead_adder) on Carry look-ahead

Retrieved from "[http://visual6502.org/wiki/index.php?title=Incrementers\_and\_adders](index.php-title-Incrementers_and_adders.md)"

