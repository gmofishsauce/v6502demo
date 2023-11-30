**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# 6507 Decode ROM - VisualChips

## 6507 Decode ROM

#### From VisualChips
(Redirected from [6507 Decode PLA](index.php-title-6507_Decode_PLA~redirect-no))

The Decode ROM[1](#endnote-1) is a 130x21 bits ROM in the 6502 that is used to decode the instruction and to control various units of the CPU.

Some basic (maybe partially incorrect) information: [http://www.pagetable.com/?p=39](http://www.pagetable.com/?p=39)

This is a verified correct transcription of the ROM, taken from the Atari 6507 diagram sheets (this ROM differs from the one in the NMOS 6502, as simulated by the [[visual6502 simulator](http://visual6502.org/JSSim)]):

```
100XX1XX 3 X STY
XXX100XX 1 3 T3INDYA
XXX110XX 1 2 T2ABSY
1100XXXX 3 0 T0CPYINY
100110XX 3 0 T0TYAA
1X0010XX 3 0 T0DEYINY
000000XX 3 5 T5INT
10XXXXXX 2 X LDXSDX
XXX1X1XX X 2 T2ANYX
XXX000XX 1 2 T2XIND
100010XX 2 0 T0TXAA
110010XX 2 0 T0DEX
1110XXXX 3 0 T0CPXINX
100110XX 2 0 T0TXS
100XXXXX 2 X SDX
101XXXXX 2 0 T0TALDTSX
110010XX 2 1 T1DEX
111010XX 3 1 T1INX
101110XX 2 0 T0TSX
1X0010XX 3 1 T1DEYINY
101XX1XX 3 0 T0LDY1
1010XXXX 3 0 T0LDY2TAY
0XX0X0XX 3 2 CCC
001000XX 3 0 T0JSR
0X0010XX 3 0 T0PSHASHP
011000XX 3 4 T4RTS
0X1010XX 3 3 T3PLAPLPA
010000XX 3 5 T5RTI
011XXXXX 2 X RORRORA
001000XX 3 2 T2JSR
01X011XX 3 X JMPA
XXXXXXXX X 2 T2
XXX011XX X 2 T2EXT
01X000XX 3 X RTIRTS
XXX000XX 1 4 T4XIND
XXXXXXXX X 0 T0A
XXXX0XXX X 2 T2NANYABS
010000XX 3 4 T4RTIA
00X000XX 3 4 T4JSRINT
0XX0XXXX 3 3 NAME1:T3\_RTI\_RTS\_JSR\_JMP\_INT\_PULA\_PUPL
XXX100XX 1 3 T3INDYB
XXX000XX 1 3 T3XIND
XXX100XX 1 4 T4INDYA
XXX100XX 1 2 T2INDY
XXX11XXX X 3 T3ABSXYA
0X1010XX 3 X PULAPULP
111XXXXX 2 X INC
010XXXXX 1 0 T0EOR
110XXXXX 1 0 T0CMP
11X0XXXX 3 0 NAME2:T0\_CPX\_CPY\_INX\_INY
X11XXXXX 1 0 T0ADCSBC
111XXXXX 1 0 T0SBC
001XXXXX 2 X ROLROLA
01X011XX 3 3 T3JMP
000XXXXX 1 0 T0ORA
00XXXXXX 2 X NAME8:ROL\_ROLA\_ASL\_ASLA
100110XX 3 0 T0TYAB
100010XX 2 0 T0TXAB
X11XXXXX 1 1 T1ADCSBCA
0XXXXXXX 1 1 NAME7:T1\_AND\_EOR\_OR\_ADC
0XX010XX 2 1 NAME4:T1\_ASLA\_ROLA\_LSRA
011010XX 3 0 T0PULA
XXX11XXX X 4 T4ABSXYA
XXX100XX 1 5 T5INDY
101XXXXX 1 0 T0LDA
XXXXXXXX 1 0 T0G1
001XXXXX 1 0 T0AND
0010X1XX 3 0 T0BITA
0XX010XX 2 0 NAME6:T0\_ASLA\_ROLA\_LSRA
101010XX 2 0 T0TAX
101010XX 3 0 T0TAY
01X010XX 2 0 T0LSRA
01XXXXXX 2 X LSRLSRA
001000XX 3 5 T5JSRA
XXX100XX 3 2 T2BR
000000XX 3 2 T2INT
001000XX 3 3 T3JSR
XXXX01XX X 2 T2ANYZP
XXXX00XX 1 2 T2ANYIND
XXXXXXXX X 4 T4
XXXXXXXX X 3 T3
0X0000XX 3 0 T0RTIINT
01X011XX 3 0 T0JMP
0XX0X0XX 3 2 NAME3:T2\_RTI\_RTS\_JSR\_INT\_PULA\_PUPLP\_PSHA\_PSHP
011000XX 3 5 T5RTS
XXXX1XXX X 2 T2ANYABS
100XXXXX 1 X STA
010010XX 3 2 T2PSHA
XXX100XX 3 0 T0BR
0XX010XX 3 X PSHPULA
XXX000XX 1 5 T5XIND
XXXX1XXX X 3 T3ANYABS
XXX100XX 1 4 T4INDYB
XXX11XXX X 3 T3ABSXYB
0X0000XX 3 X RTIINT
001000XX 3 X JSR
01X011XX 3 X JMPB
11X00XXX 3 1 T1CPX2CY2
00X010XX 2 1 T1ASLARLA
11X011XX 3 1 T1CPX1CY1
110XXXXX 1 1 T1CMP
X11XXXXX 1 1 T1ADCSBCB
00XXXXXX 2 X NAME5:ROL\_ROLA\_ASL\_ASLA
X1XXXXXX 2 X LSRRADCIC
0010X1XX 3 1 T1BIT
000010XX 3 2 T2PSHP
000000XX 3 4 T4INT
100XXXXX X X STASTYSTX
XXX11XXX X 4 T4ABSXYB
XXXX00XX 1 5 T5ANYIND
XXX001XX X 2 T2ZP
XXX011XX X 3 T3ABS
XXX101XX X 3 T3ZPX
0X0010XX 3 2 T2PSHASHP
01X000XX 3 5 T5RTIRTS
001000XX 3 5 T5JSRB
01X011XX 3 5 T4JMP
010011XX 3 2 T2JMPABS
0X1010XX 3 3 T3PLAPLPB
XXX100XX 3 3 T3BR
0010X1XX 3 0 T0BITB
010000XX 3 4 T4RTIB
001010XX 3 0 T0PULP
0XX010XX 3 X PSHPULB
101110XX 3 X CLV
00X110XX 3 0 T0CLCSEC
01X110XX 3 0 T0CLISEI
11X110XX 3 0 T0CLDSED
0XXXXXXX X X NI7P
X0XXXXXX X X NI6P
```

The format is:

```
76543210 G T NAME
```

- The first column represents the 8 bits of the IR. 1 means the bit has to be 1 for the line to fire, 0 means the bit has to be 0, X is a don't care. Note that the lower two bits are always XX - the decode ROM doesn't actually check these, but check a cooked version of these bits instead.
- The second column is the "G" input, it must match 1, 2, 3 or it's a don't care. G is derived from the lower two bits of the IR:

```
G1 = IR0
G2 = IR1
G3 = !IR0 & !IR1
```

- The third column is "T", which is the clock cycle in which the line fires (0..5).

Some observations:

1. There are 15 duplicates in the decode ROM:

```
$ for i in `sort pla.txt | cut -c -12 | uniq -c | sort -n | grep "^   2" | cut -c 6-17 | sed -e "s/ /./g"`; do grep $i pla.txt; done
```

We assume this has been done because they had no way of routing the output of some line where they wanted, so they put the same line at a different location again.

2. As an example, ADC # is 2 cycles, but there are lines that match T=[2..4]. In practice, these will never fire; they are meant for other instructions that have a similar encoding and do have T>2.

3. About G, and how it explains many illegal opcodes:
Orlando and I reverse engineered this by dumping operation lists with decode.rb and filtering which Gs made sense. The funny thing here is that this leads to the table:

```
00 -> G3
01 -> G1
10 -> G2
11 -> G1/2
```

11 is the don't care case, there are no opcodes XXXXXX11 that are documented. So in order to simplify the G encoding, 11 has both G1 and G2 turned on, so all G=1 and G=2 lines fire. And this explains A LOT of things, like how LDA (0xAD, G=1) and LDX (0xAE, G=2) become LAX (0xAF, G=1 and G=2):

```
LDA T=0
XXXXXXXX X 0 T0A
101XXXXX 1 0 T0LDA
XXXXXXXX 1 0 T0G1
X0XXXXXX X X NI6P

LDX T=0
10XXXXXX 2 X LDXSDX
101XXXXX 2 0 T0TALDTSX
XXXXXXXX X 0 T0A
X0XXXXXX X X NI6P

LAX T0
10XXXXXX 2 X LDXSDX
101XXXXX 2 0 T0TALDTSX
XXXXXXXX X 0 T0A
101XXXXX 1 0 T0LDA
XXXXXXXX 1 0 T0G1
X0XXXXXX X X NI6P
```

...which is pretty much LDA and LDX joined!

If you look at [http://www.oxyron.de/html/opcodes02.html](http://www.oxyron.de/html/opcodes02.html) , you can see that columns 3, 7, B and F are illegal; this is the G=1+G=3 case. These columns basically execute all operations of the two preceding columns at the same time. Note that (as far as I checked) the cycle count (number in the table) is always the MAX() of the two opcodes it consists of.

This was known:

[http://www.viceteam.org/plain/64doc.txt](http://www.viceteam.org/plain/64doc.txt)

Other undocumented instructions usually cause two preceding opcodes being executed.

But now we have more of a clue what's happening there...

Column 2 is KIL, column 3 has mostly *8* cycle instructions, which is weird. There are no regular 8 cycles ones. Not sure this is accurate in the docs. I see a correlation between the KIL and the 8 cycles. My theory is that KIL overflows the cycle counter. Not sure why column 3 doesn't inherit that feature.

The MAX() property and the KIL/Column 3 thing might explain how the cycle counter gets reset to 0... what triggers that. Orlando is also looking into this, comparing neighbor opcodes that terminate differently.

Here is the ruby program that accepts an opcode at the command line and prints the sequence of clocks:

```
#! /usr/bin/env ruby

if ($*.length < 1)
    print "usage: #{$0} <value>\n"
    exit
end

opc = eval($*[0])
b0 = (opc & 1) != 0
b1 = (opc & 2) != 0
gmatch = Array.new
gmatch[1] = b0
gmatch[2] = b1
gmatch[3] = !b0 && !b1

bin = ("%!b(MISSING)" %!o(MISSING)pc)

input = Array.new

File.open('pla.txt').each\_line do |s|
  next if s =~ /^#.*/ # skip lines starting with '#'
  input += [ s.chop.split(/ /) ]
end

6.times do |time|
  print "T=#{time}\n"
  input.each do |ni, g, t, name|
    print "#{ni} #{g} #{t} #{name}\n" if (bin.match(ni.gsub(/X/, ".")) && (t == "X" || time == t.to\_i) && (g == "X" || gmatch[g.to\_i]))
  end
  print "\n"
end
```

It also needs to read a file 'pla.txt' which has a tabulation [as found here](index.php-title-6502_all_256_Opcodes).

### Notes
[^1](#ref-1) The "Decode ROM" is named as a ROM in [Hanson's block diagram](index.php-title-Hanson~27s_Block_Diagram), although it has wordline inputs and no address decoder. It is sometimes described as a PLA although it also lacks an AND plane. It is a structured layout of NOR gates with many common inputs, as compared to the unstructured gates found in the central decode logic, sometimes known as the random logic (meaning not structured).Retrieved from "[http://visual6502.org/wiki/index.php?title=6507\_Decode\_ROM](index.php-title-6507_Decode_ROM)"

