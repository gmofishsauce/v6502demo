**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# 6502 all 256 Opcodes - VisualChips

## 6502 all 256 Opcodes

#### From VisualChips

Starting from [Graham's table](http://www.oxyron.de/html/opcodes02.html), Michael Steil constructed a 3 column table containing the opcode, the mnemomic, the addressing mode and the number of clock cycles.

Then, using the C model 'perfect6502' (available on Github) which uses the same switch-level approach as the [visual6502 simulator](http://visual6502.org/JSSim), he wrote automated tests that feed in all kinds of instructions, and measures cycles, instruction length, memory accesses and register inputs and outputs. It takes about 3 minutes to go through the opcode space.

This is what the output means:

- bytes: opcode is followed by 0s, size is calculated by the BRK address put onto the stack
- cycles: opcode is followed by 0s, on a page boundary, the time is the number of cycles between the opcode fetch and the next fetch.
- (This does not measure the "one more if branch taken/index crosses page boundary" case.)
- AXYSP=>: registers that are used as inputs (S: stack P: status), i.e. behavior (register contents or bus I/O) changes if inputs change
- =>AXYSP=>: registers that are used as outputs, i.e. register changes if inputs (registers, memory) change
- RW: whether this opcode does memory reads and/or writes (instruction fetches are not counted)
- abs, absx, absy, zp, zpx, zpy, izx, izx: addressing mode
- CRASH is printed if the BRK instruction is not detected after 100 half-cycles

This is the output (Graham's list to the left!):

```
00 BRK 7        $00: bytes: 0 cycles: 0 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 01 ORA izx 6    $01: bytes: 2 cycles: 6 A\_\_\_\_=>\_\_\_\_P R\_ izx
 02 *KIL         $02: CRASH
 03 *SLO izx 8   $03: bytes: 2 cycles: 8 A\_\_\_\_=>\_\_\_\_P RW izx
 04 *NOP zp 3    $04: bytes: 2 cycles: 3 \_\_\_\_\_=>\_\_\_\_\_ R\_ zp
 05 ORA zp 3     $05: bytes: 2 cycles: 3 A\_\_\_\_=>A\_\_\_P R\_ zp
 06 ASL zp 5     $06: bytes: 2 cycles: 5 \_\_\_\_\_=>\_\_\_\_P RW zp
 07 *SLO zp 5    $07: bytes: 2 cycles: 5 A\_\_\_\_=>A\_\_\_P RW zp
 08 PHP 3        $08: bytes: 1 cycles: 3 \_\_\_SP=>\_\_\_S\_ \_W 
 09 ORA imm 2    $09: bytes: 2 cycles: 2 \_\_\_\_\_=>A\_\_\_P \_\_ 
 0A ASL 2        $0A: bytes: 1 cycles: 2 A\_\_\_\_=>A\_\_\_P \_\_ 
 0B *ANC imm 2   $0B: bytes: 2 cycles: 2 A\_\_\_\_=>\_\_\_\_P \_\_ 
 0C *NOP abs 4   $0C: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ abs
 0D ORA abs 4    $0D: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ abs
 0E ASL abs 6    $0E: bytes: 3 cycles: 6 \_\_\_\_\_=>\_\_\_\_P RW abs
 0F *SLO abs 6   $0F: bytes: 3 cycles: 6 A\_\_\_\_=>A\_\_\_P RW abs
 10 BPL rel 2*   $10: bytes: 2 cycles: 3 \_\_\_\_P=>\_\_\_\_\_ \_\_ 
 11 ORA izy 5*   $11: bytes: 2 cycles: 5 A\_\_\_\_=>\_\_\_\_P R\_ izy
 12 *KIL         $12: CRASH
 13 *SLO izy 8   $13: bytes: 2 cycles: 8 A\_\_\_\_=>\_\_\_\_P RW izy
 14 *NOP zpx 4   $14: bytes: 2 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ zpx
 15 ORA zpx 4    $15: bytes: 2 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ zpx
 16 ASL zpx 6    $16: bytes: 2 cycles: 6 \_\_\_\_\_=>\_\_\_\_P RW zpx
 17 *SLO zpx 6   $17: bytes: 2 cycles: 6 A\_\_\_\_=>A\_\_\_P RW zpx
 18 CLC 2        $18: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_P \_\_ 
 19 ORA aby 4*   $19: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ absy
 1A *NOP 2       $1A: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 1B *SLO aby 7   $1B: bytes: 3 cycles: 7 A\_\_\_\_=>A\_\_\_P RW absy
 1C *NOP abx 4*  $1C: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ absx
 1D ORA abx 4*   $1D: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ absx
 1E ASL abx 7    $1E: bytes: 3 cycles: 7 \_\_\_\_\_=>\_\_\_\_P RW absx
 1F *SLO abx 7   $1F: bytes: 3 cycles: 7 A\_\_\_\_=>A\_\_\_P RW absx
 20 JSR abs 6    $20: bytes: X cycles: 6 \_\_\_S\_=>\_\_\_S\_ \_W 
 21 AND izx 6    $21: bytes: 2 cycles: 6 \_\_\_\_\_=>A\_\_\_P R\_ izx
 22 *KIL         $22: CRASH
 23 *RLA izx 8   $23: bytes: 2 cycles: 8 \_\_\_\_P=>A\_\_\_P RW izx
 24 BIT zp 3     $24: bytes: 2 cycles: 3 A\_\_\_\_=>\_\_\_\_P R\_ zp
 25 AND zp 3     $25: bytes: 2 cycles: 3 A\_\_\_\_=>A\_\_\_P R\_ zp
 26 ROL zp 5     $26: bytes: 2 cycles: 5 \_\_\_\_P=>\_\_\_\_P RW zp
 27 *RLA zp 5    $27: bytes: 2 cycles: 5 A\_\_\_P=>A\_\_\_P RW zp
 28 PLP 4        $28: bytes: 1 cycles: 4 \_\_\_S\_=>\_\_\_SP \_\_ 
 29 AND imm 2    $29: bytes: 2 cycles: 2 A\_\_\_\_=>A\_\_\_P \_\_ 
 2A ROL 2        $2A: bytes: 1 cycles: 2 A\_\_\_P=>A\_\_\_P \_\_ 
 2B *ANC imm 2   $2B: bytes: 2 cycles: 2 A\_\_\_\_=>\_\_\_\_P \_\_ 
 2C BIT abs 4    $2C: bytes: 3 cycles: 4 A\_\_\_\_=>\_\_\_\_P R\_ abs
 2D AND abs 4    $2D: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ abs
 2E ROL abs 6    $2E: bytes: 3 cycles: 6 \_\_\_\_P=>\_\_\_\_P RW abs
 2F *RLA abs 6   $2F: bytes: 3 cycles: 6 A\_\_\_P=>A\_\_\_P RW abs
 30 BMI rel 2*   $30: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 31 AND izy 5*   $31: bytes: 2 cycles: 5 \_\_\_\_\_=>A\_\_\_P R\_ izy
 32 *KIL         $32: CRASH
 33 *RLA izy 8   $33: bytes: 2 cycles: 8 \_\_\_\_P=>A\_\_\_P RW izy
 34 *NOP zpx 4   $34: bytes: 2 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ zpx
 35 AND zpx 4    $35: bytes: 2 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ zpx
 36 ROL zpx 6    $36: bytes: 2 cycles: 6 \_\_\_\_P=>\_\_\_\_P RW zpx
 37 *RLA zpx 6   $37: bytes: 2 cycles: 6 A\_\_\_P=>A\_\_\_P RW zpx
 38 SEC 2        $38: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_P \_\_ 
 39 AND aby 4*   $39: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ absy
 3A *NOP 2       $3A: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 3B *RLA aby 7   $3B: bytes: 3 cycles: 7 A\_\_\_P=>A\_\_\_P RW absy
 3C *NOP abx 4*  $3C: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ absx
 3D AND abx 4*   $3D: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ absx
 3E ROL abx 7    $3E: bytes: 3 cycles: 7 \_\_\_\_P=>\_\_\_\_P RW absx
 3F *RLA abx 7   $3F: bytes: 3 cycles: 7 A\_\_\_P=>A\_\_\_P RW absx
 40 RTI 6        $40: bytes: X cycles: 6 \_\_\_S\_=>\_\_\_SP \_\_ 
 41 EOR izx 6    $41: bytes: 2 cycles: 6 A\_\_\_\_=>\_\_\_\_P R\_ izx
 42 *KIL         $42: CRASH
 43 *SRE izx 8   $43: bytes: 2 cycles: 8 A\_\_\_\_=>\_\_\_\_P RW izx
 44 *NOP zp 3    $44: bytes: 2 cycles: 3 \_\_\_\_\_=>\_\_\_\_\_ R\_ zp
 45 EOR zp 3     $45: bytes: 2 cycles: 3 A\_\_\_\_=>A\_\_\_P R\_ zp
 46 LSR zp 5     $46: bytes: 2 cycles: 5 \_\_\_\_\_=>\_\_\_\_P RW zp
 47 *SRE zp 5    $47: bytes: 2 cycles: 5 A\_\_\_\_=>A\_\_\_P RW zp
 48 PHA 3        $48: bytes: 1 cycles: 3 A\_\_S\_=>\_\_\_S\_ \_W 
 49 EOR imm 2    $49: bytes: 2 cycles: 2 A\_\_\_\_=>A\_\_\_P \_\_ 
 4A LSR 2        $4A: bytes: 1 cycles: 2 A\_\_\_\_=>A\_\_\_P \_\_ 
 4B *ALR imm 2   $4B: bytes: 2 cycles: 2 A\_\_\_\_=>A\_\_\_P \_\_ 
 4C JMP abs 3    $4C: bytes: X cycles: 3 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 4D EOR abs 4    $4D: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ abs
 4E LSR abs 6    $4E: bytes: 3 cycles: 6 \_\_\_\_\_=>\_\_\_\_P RW abs
 4F *SRE abs 6   $4F: bytes: 3 cycles: 6 A\_\_\_\_=>A\_\_\_P RW abs
 50 BVC rel 2*   $50: bytes: 2 cycles: 3 \_\_\_\_P=>\_\_\_\_\_ \_\_ 
 51 EOR izy 5*   $51: bytes: 2 cycles: 5 A\_\_\_\_=>\_\_\_\_P R\_ izy
 52 *KIL         $52: CRASH
 53 *SRE izy 8   $53: bytes: 2 cycles: 8 A\_\_\_\_=>\_\_\_\_P RW izy
 54 *NOP zpx 4   $54: bytes: 2 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ zpx
 55 EOR zpx 4    $55: bytes: 2 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ zpx
 56 LSR zpx 6    $56: bytes: 2 cycles: 6 \_\_\_\_\_=>\_\_\_\_P RW zpx
 57 *SRE zpx 6   $57: bytes: 2 cycles: 6 A\_\_\_\_=>A\_\_\_P RW zpx
 58 CLI 2        $58: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_P \_\_ 
 59 EOR aby 4*   $59: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ absy
 5A *NOP 2       $5A: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 5B *SRE aby 7   $5B: bytes: 3 cycles: 7 A\_\_\_\_=>A\_\_\_P RW absy
 5C *NOP abx 4*  $5C: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ absx
 5D EOR abx 4*   $5D: bytes: 3 cycles: 4 A\_\_\_\_=>A\_\_\_P R\_ absx
 5E LSR abx 7    $5E: bytes: 3 cycles: 7 \_\_\_\_\_=>\_\_\_\_P RW absx
 5F *SRE abx 7   $5F: bytes: 3 cycles: 7 A\_\_\_\_=>A\_\_\_P RW absx
 60 RTS 6        $60: bytes: X cycles: 6 \_\_\_S\_=>\_\_\_S\_ \_\_ 
 61 ADC izx 6    $61: bytes: 2 cycles: 6 A\_\_\_P=>A\_\_\_P R\_ izx
 62 *KIL         $62: CRASH
 63 *RRA izx 8   $63: bytes: 2 cycles: 8 A\_\_\_P=>A\_\_\_P RW izx
 64 *NOP zp 3    $64: bytes: 2 cycles: 3 \_\_\_\_\_=>\_\_\_\_\_ R\_ zp
 65 ADC zp 3     $65: bytes: 2 cycles: 3 A\_\_\_P=>A\_\_\_P R\_ zp
 66 ROR zp 5     $66: bytes: 2 cycles: 5 \_\_\_\_P=>\_\_\_\_P RW zp
 67 *RRA zp 5    $67: bytes: 2 cycles: 5 A\_\_\_P=>A\_\_\_P RW zp
 68 PLA 4        $68: bytes: 1 cycles: 4 \_\_\_S\_=>A\_\_SP \_\_ 
 69 ADC imm 2    $69: bytes: 2 cycles: 2 A\_\_\_P=>A\_\_\_P \_\_ 
 6A ROR 2        $6A: bytes: 1 cycles: 2 A\_\_\_P=>A\_\_\_P \_\_ 
 6B *ARR imm 2   $6B: bytes: 2 cycles: 2 A\_\_\_P=>A\_\_\_P \_\_ 
 6C JMP ind 5    $6C: bytes: X cycles: 5 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 6D ADC abs 4    $6D: bytes: 3 cycles: 4 A\_\_\_P=>A\_\_\_P R\_ abs
 6E ROR abs 6    $6E: bytes: 3 cycles: 6 \_\_\_\_P=>\_\_\_\_P RW abs
 6F *RRA abs 6   $6F: bytes: 3 cycles: 6 A\_\_\_P=>A\_\_\_P RW abs
 70 BVS rel 2*   $70: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 71 ADC izy 5*   $71: bytes: 2 cycles: 5 A\_\_\_P=>A\_\_\_P R\_ izy
 72 *KIL         $72: CRASH
 73 *RRA izy 8   $73: bytes: 2 cycles: 8 A\_\_\_P=>A\_\_\_P RW izy
 74 *NOP zpx 4   $74: bytes: 2 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ zpx
 75 ADC zpx 4    $75: bytes: 2 cycles: 4 A\_\_\_P=>A\_\_\_P R\_ zpx
 76 ROR zpx 6    $76: bytes: 2 cycles: 6 \_\_\_\_P=>\_\_\_\_P RW zpx
 77 *RRA zpx 6   $77: bytes: 2 cycles: 6 A\_\_\_P=>A\_\_\_P RW zpx
 78 SEI 2        $78: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_P \_\_ 
 79 ADC aby 4*   $79: bytes: 3 cycles: 4 A\_\_\_P=>A\_\_\_P R\_ absy
 7A *NOP 2       $7A: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 7B *RRA aby 7   $7B: bytes: 3 cycles: 7 A\_\_\_P=>A\_\_\_P RW absy
 7C *NOP abx 4*  $7C: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ absx
 7D ADC abx 4*   $7D: bytes: 3 cycles: 4 A\_\_\_P=>A\_\_\_P R\_ absx
 7E ROR abx 7    $7E: bytes: 3 cycles: 7 \_\_\_\_P=>\_\_\_\_P RW absx
 7F *RRA abx 7   $7F: bytes: 3 cycles: 7 A\_\_\_P=>A\_\_\_P RW absx
 80 *NOP imm 2   $80: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 81 STA izx 6    $81: bytes: 2 cycles: 6 A\_\_\_\_=>\_\_\_\_\_ RW izx
 82 *NOP imm 2   $82: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 83 *SAX izx 6   $83: bytes: 2 cycles: 6 \_\_\_\_\_=>\_\_\_\_\_ RW izx
 84 STY zp 3     $84: bytes: 2 cycles: 3 \_\_Y\_\_=>\_\_\_\_\_ \_W zp
 85 STA zp 3     $85: bytes: 2 cycles: 3 A\_\_\_\_=>\_\_\_\_\_ \_W zp
 86 STX zp 3     $86: bytes: 2 cycles: 3 \_X\_\_\_=>\_\_\_\_\_ \_W zp
 87 *SAX zp 3    $87: bytes: 2 cycles: 3 \_\_\_\_\_=>\_\_\_\_\_ \_W zp
 88 DEY 2        $88: bytes: 1 cycles: 2 \_\_Y\_\_=>\_\_Y\_P \_\_ 
 89 *NOP imm 2   $89: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 8A TXA 2        $8A: bytes: 1 cycles: 2 \_X\_\_\_=>A\_\_\_P \_\_ 
 8B *XAA imm 2   $8B: bytes: 2 cycles: 2 \_\_\_\_\_=>A\_\_\_P \_\_ 
 8C STY abs 4    $8C: bytes: 3 cycles: 4 \_\_Y\_\_=>\_\_\_\_\_ \_W abs
 8D STA abs 4    $8D: bytes: 3 cycles: 4 A\_\_\_\_=>\_\_\_\_\_ \_W abs
 8E STX abs 4    $8E: bytes: 3 cycles: 4 \_X\_\_\_=>\_\_\_\_\_ \_W abs
 8F *SAX abs 4   $8F: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ \_W abs
 90 BCC rel 2*   $90: bytes: 2 cycles: 3 \_\_\_\_P=>\_\_\_\_\_ \_\_ 
 91 STA izy 6    $91: bytes: 2 cycles: 6 A\_\_\_\_=>\_\_\_\_\_ RW izy
 92 *KIL         $92: CRASH
 93 *AHX izy 6   $93: bytes: 2 cycles: 6 \_\_\_\_\_=>\_\_\_\_\_ RW izy
 94 STY zpx 4    $94: bytes: 2 cycles: 4 \_\_Y\_\_=>\_\_\_\_\_ RW zpx
 95 STA zpx 4    $95: bytes: 2 cycles: 4 A\_\_\_\_=>\_\_\_\_\_ RW zpx
 96 STX zpy 4    $96: bytes: 2 cycles: 4 \_X\_\_\_=>\_\_\_\_\_ RW zpy
 97 *SAX zpy 4   $97: bytes: 2 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ RW zpy
 98 TYA 2        $98: bytes: 1 cycles: 2 \_\_Y\_\_=>A\_\_\_P \_\_ 
 99 STA aby 5    $99: bytes: 3 cycles: 5 A\_\_\_\_=>\_\_\_\_\_ RW absy
 9A TXS 2        $9A: bytes: X cycles: 2 \_X\_\_\_=>\_\_\_S\_ \_\_ 
 9B *TAS aby 5   $9B: bytes: X cycles: 5 \_\_Y\_\_=>\_\_\_S\_ \_W 
 9C *SHY abx 5   $9C: bytes: 3 cycles: 5 \_\_Y\_\_=>\_\_\_\_\_ RW absx
 9D STA abx 5    $9D: bytes: 3 cycles: 5 A\_\_\_\_=>\_\_\_\_\_ RW absx
 9E *SHX aby 5   $9E: bytes: 3 cycles: 5 \_X\_\_\_=>\_\_\_\_\_ RW absy
 9F *AHX aby 5   $9F: bytes: 3 cycles: 5 \_\_\_\_\_=>\_\_\_\_\_ RW absy
 A0 LDY imm 2    $A0: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_Y\_P \_\_ 
 A1 LDA izx 6    $A1: bytes: 2 cycles: 6 \_\_\_\_\_=>A\_\_\_P R\_ izx
 A2 LDX imm 2    $A2: bytes: 2 cycles: 2 \_\_\_\_\_=>\_X\_\_P \_\_ 
 A3 *LAX izx 6   $A3: bytes: 2 cycles: 6 \_\_\_\_\_=>AX\_\_P R\_ izx
 A4 LDY zp 3     $A4: bytes: 2 cycles: 3 \_\_\_\_\_=>\_\_Y\_P R\_ zp
 A5 LDA zp 3     $A5: bytes: 2 cycles: 3 \_\_\_\_\_=>A\_\_\_P R\_ zp
 A6 LDX zp 3     $A6: bytes: 2 cycles: 3 \_\_\_\_\_=>\_X\_\_P R\_ zp
 A7 *LAX zp 3    $A7: bytes: 2 cycles: 3 \_\_\_\_\_=>AX\_\_P R\_ zp
 A8 TAY 2        $A8: bytes: 1 cycles: 2 A\_\_\_\_=>\_\_Y\_P \_\_ 
 A9 LDA imm 2    $A9: bytes: 2 cycles: 2 \_\_\_\_\_=>A\_\_\_P \_\_ 
 AA TAX 2        $AA: bytes: 1 cycles: 2 A\_\_\_\_=>\_X\_\_P \_\_ 
 AB *LAX imm 2   $AB: bytes: 2 cycles: 2 A\_\_\_\_=>AX\_\_P \_\_ 
 AC LDY abs 4    $AC: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_Y\_P R\_ abs
 AD LDA abs 4    $AD: bytes: 3 cycles: 4 \_\_\_\_\_=>A\_\_\_P R\_ abs
 AE LDX abs 4    $AE: bytes: 3 cycles: 4 \_\_\_\_\_=>\_X\_\_P R\_ abs
 AF *LAX abs 4   $AF: bytes: 3 cycles: 4 \_\_\_\_\_=>AX\_\_P R\_ abs
 B0 BCS rel 2*   $B0: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 B1 LDA izy 5*   $B1: bytes: 2 cycles: 5 \_\_\_\_\_=>A\_\_\_P R\_ izy
 B2 *KIL         $B2: CRASH
 B3 *LAX izy 5*  $B3: bytes: 2 cycles: 5 \_\_\_\_\_=>AX\_\_P R\_ izy
 B4 LDY zpx 4    $B4: bytes: 2 cycles: 4 \_\_\_\_\_=>\_\_Y\_P R\_ zpx
 B5 LDA zpx 4    $B5: bytes: 2 cycles: 4 \_\_\_\_\_=>A\_\_\_P R\_ zpx
 B6 LDX zpy 4    $B6: bytes: 2 cycles: 4 \_\_\_\_\_=>\_X\_\_P R\_ zpy
 B7 *LAX zpy 4   $B7: bytes: 2 cycles: 4 \_\_\_\_\_=>AX\_\_P R\_ zpy
 B8 CLV 2        $B8: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_P \_\_ 
 B9 LDA aby 4*   $B9: bytes: 3 cycles: 4 \_\_\_\_\_=>A\_\_\_P R\_ absy
 BA TSX 2        $BA: bytes: 1 cycles: 2 \_\_\_S\_=>\_X\_\_P \_\_ 
 BB *LAS aby 4*  $BB: bytes: 3 cycles: 4 \_\_\_S\_=>AX\_SP R\_ absy
 BC LDY abx 4*   $BC: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_Y\_P R\_ absx
 BD LDA abx 4*   $BD: bytes: 3 cycles: 4 \_\_\_\_\_=>A\_\_\_P R\_ absx
 BE LDX aby 4*   $BE: bytes: 3 cycles: 4 \_\_\_\_\_=>\_X\_\_P R\_ absy
 BF *LAX aby 4*  $BF: bytes: 3 cycles: 4 \_\_\_\_\_=>AX\_\_P R\_ absy
 C0 CPY imm 2    $C0: bytes: 2 cycles: 2 \_\_Y\_\_=>\_\_\_\_P \_\_ 
 C1 CMP izx 6    $C1: bytes: 2 cycles: 6 A\_\_\_\_=>\_\_\_\_P R\_ izx
 C2 *NOP imm 2   $C2: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 C3 *DCP izx 8   $C3: bytes: 2 cycles: 8 A\_\_\_\_=>\_\_\_\_P RW izx
 C4 CPY zp 3     $C4: bytes: 2 cycles: 3 \_\_Y\_\_=>\_\_\_\_P R\_ zp
 C5 CMP zp 3     $C5: bytes: 2 cycles: 3 A\_\_\_\_=>\_\_\_\_P R\_ zp
 C6 DEC zp 5     $C6: bytes: 2 cycles: 5 \_\_\_\_\_=>\_\_\_\_P RW zp
 C7 *DCP zp 5    $C7: bytes: 2 cycles: 5 A\_\_\_\_=>\_\_\_\_P RW zp
 C8 INY 2        $C8: bytes: 1 cycles: 2 \_\_Y\_\_=>\_\_Y\_P \_\_ 
 C9 CMP imm 2    $C9: bytes: 2 cycles: 2 A\_\_\_\_=>\_\_\_\_P \_\_ 
 CA DEX 2        $CA: bytes: 1 cycles: 2 \_X\_\_\_=>\_X\_\_P \_\_ 
 CB *AXS imm 2   $CB: bytes: 2 cycles: 2 \_\_\_\_\_=>\_X\_\_P \_\_ 
 CC CPY abs 4    $CC: bytes: 3 cycles: 4 \_\_Y\_\_=>\_\_\_\_P R\_ abs
 CD CMP abs 4    $CD: bytes: 3 cycles: 4 A\_\_\_\_=>\_\_\_\_P R\_ abs
 CE DEC abs 6    $CE: bytes: 3 cycles: 6 \_\_\_\_\_=>\_\_\_\_P RW abs
 CF *DCP abs 6   $CF: bytes: 3 cycles: 6 A\_\_\_\_=>\_\_\_\_P RW abs
 D0 BNE rel 2*   $D0: bytes: 2 cycles: 3 \_\_\_\_P=>\_\_\_\_\_ \_\_ 
 D1 CMP izy 5*   $D1: bytes: 2 cycles: 5 A\_\_\_\_=>\_\_\_\_P R\_ izy
 D2 *KIL         $D2: CRASH
 D3 *DCP izy 8   $D3: bytes: 2 cycles: 8 A\_\_\_\_=>\_\_\_\_P RW izy
 D4 *NOP zpx 4   $D4: bytes: 2 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ zpx
 D5 CMP zpx 4    $D5: bytes: 2 cycles: 4 A\_\_\_\_=>\_\_\_\_P R\_ zpx
 D6 DEC zpx 6    $D6: bytes: 2 cycles: 6 \_\_\_\_\_=>\_\_\_\_P RW zpx
 D7 *DCP zpx 6   $D7: bytes: 2 cycles: 6 A\_\_\_\_=>\_\_\_\_P RW zpx
 D8 CLD 2        $D8: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_P \_\_ 
 D9 CMP aby 4*   $D9: bytes: 3 cycles: 4 A\_\_\_\_=>\_\_\_\_P R\_ absy
 DA *NOP 2       $DA: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 DB *DCP aby 7   $DB: bytes: 3 cycles: 7 A\_\_\_\_=>\_\_\_\_P RW absy
 DC *NOP abx 4*  $DC: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ absx
 DD CMP abx 4*   $DD: bytes: 3 cycles: 4 A\_\_\_\_=>\_\_\_\_P R\_ absx
 DE DEC abx 7    $DE: bytes: 3 cycles: 7 \_\_\_\_\_=>\_\_\_\_P RW absx
 DF *DCP abx 7   $DF: bytes: 3 cycles: 7 A\_\_\_\_=>\_\_\_\_P RW absx
 E0 CPX imm 2    $E0: bytes: 2 cycles: 2 \_X\_\_\_=>\_\_\_\_P \_\_ 
 E1 SBC izx 6    $E1: bytes: 2 cycles: 6 A\_\_\_P=>A\_\_\_P R\_ izx
 E2 *NOP imm 2   $E2: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 E3 *ISC izx 8   $E3: bytes: 2 cycles: 8 A\_\_\_P=>A\_\_\_P RW izx
 E4 CPX zp 3     $E4: bytes: 2 cycles: 3 \_X\_\_\_=>\_\_\_\_P R\_ zp
 E5 SBC zp 3     $E5: bytes: 2 cycles: 3 A\_\_\_P=>A\_\_\_P R\_ zp
 E6 INC zp 5     $E6: bytes: 2 cycles: 5 \_\_\_\_\_=>\_\_\_\_P RW zp
 E7 *ISC zp 5    $E7: bytes: 2 cycles: 5 A\_\_\_P=>A\_\_\_P RW zp
 E8 INX 2        $E8: bytes: 1 cycles: 2 \_X\_\_\_=>\_X\_\_P \_\_ 
 E9 SBC imm 2    $E9: bytes: 2 cycles: 2 A\_\_\_P=>A\_\_\_P \_\_ 
 EA NOP 2        $EA: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 EB *SBC imm 2   $EB: bytes: 2 cycles: 2 A\_\_\_P=>A\_\_\_P \_\_ 
 EC CPX abs 4    $EC: bytes: 3 cycles: 4 \_X\_\_\_=>\_\_\_\_P R\_ abs
 ED SBC abs 4    $ED: bytes: 3 cycles: 4 A\_\_\_P=>A\_\_\_P R\_ abs
 EE INC abs 6    $EE: bytes: 3 cycles: 6 \_\_\_\_\_=>\_\_\_\_P RW abs
 EF *ISC abs 6   $EF: bytes: 3 cycles: 6 A\_\_\_P=>A\_\_\_P RW abs
 F0 BEQ rel 2*   $F0: bytes: 2 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 F1 SBC izy 5*   $F1: bytes: 2 cycles: 5 A\_\_\_P=>A\_\_\_P R\_ izy
 F2 *KIL         $F2: CRASH
 F3 *ISC izy 8   $F3: bytes: 2 cycles: 8 A\_\_\_P=>A\_\_\_P RW izy
 F4 *NOP zpx 4   $F4: bytes: 2 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ zpx
 F5 SBC zpx 4    $F5: bytes: 2 cycles: 4 A\_\_\_P=>A\_\_\_P R\_ zpx
 F6 INC zpx 6    $F6: bytes: 2 cycles: 6 \_\_\_\_\_=>\_\_\_\_P RW zpx
 F7 *ISC zpx 6   $F7: bytes: 2 cycles: 6 A\_\_\_P=>A\_\_\_P RW zpx
 F8 SED 2        $F8: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_P \_\_ 
 F9 SBC aby 4*   $F9: bytes: 3 cycles: 4 A\_\_\_P=>A\_\_\_P R\_ absy
 FA *NOP 2       $FA: bytes: 1 cycles: 2 \_\_\_\_\_=>\_\_\_\_\_ \_\_ 
 FB *ISC aby 7   $FB: bytes: 3 cycles: 7 A\_\_\_P=>A\_\_\_P RW absy
 FC *NOP abx 4*  $FC: bytes: 3 cycles: 4 \_\_\_\_\_=>\_\_\_\_\_ R\_ absx
 FD SBC abx 4*   $FD: bytes: 3 cycles: 4 A\_\_\_P=>A\_\_\_P R\_ absx
 FE INC abx 7    $FE: bytes: 3 cycles: 7 \_\_\_\_\_=>\_\_\_\_P RW absx
 FF *ISC abx     $FF: bytes: 3 cycles: 7 A\_\_\_P=>A\_\_\_P RW absx
```

Summary:

- The transistor-level emulation seems to be able to successfully emulate illegal opcodes as well, since the outputs for the illegal opcodes look a lot like Graham's list.
- "Unstable" illegal opcodes are probably not caused by transistor ping-pong, but by leftover trash on the external address or data bus, since the transistor calculation loop with the limiter of 100 never goes above 20, even when testing all opcodes.
- The simulator is the perfect tool to understand illegal opcodes... Well, unless you count "understanding the 6502 schematics and what's actually going on".

Note:

- This is still a work in progress.
- Some instructions were checked against the spec while writing the tests, but not everything is verified, including the simulator.
- BRK, JSR, JMP, branches aren't correct.
- The test could be extended to do special case tests on the illegal opcodes. We can look at bus activity to see what's going on, in order to understand what potential inputs are - it is possible that some instructions do extra write cycles that nobody has measured yet.

Retrieved from "[http://visual6502.org/wiki/index.php?title=6502\_all\_256\_Opcodes](index.php-title-6502_all_256_Opcodes.md)"

