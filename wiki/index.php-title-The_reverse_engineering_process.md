**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# The reverse engineering process - VisualChips

## The reverse engineering process

#### From VisualChips
**Contents**

- [Overview](#overview)
- [Microphotography](#microphotography)
- [De-layering](#de-layering)
- [Resources](#resources)

### Overview

To help explain which state each of our projects is at, here's a description of the steps we follow:

- Get a chip, usually just one of a particular kind but sometimes more
- Depackage the chip
- Chips with a metal lid or a ceramic sandwich [package](http://en.wikipedia.org/wiki/Dual_in-line_package) are preferable since these have no plastic in contact with the die.
- Chips packaged in plastic must be treated with very hot, very nasty acids which we do at a local laboratory with proper equipment
- Photograph the exposed surface of the chip through a microscope
- Many separate photographs must be taken to cover the surface at high enough resolution
- Stitch the photographs into a single large image
- Alignment data is used to correct individual photographs for optical distortions
- Usually, de-layer the chip to reveal hidden or obscured lower features
- Photograph and stitch each layer image
- Align all layer images to each other
- Create polygon models of each part of the chip based on the aligned images
- Convert the polygon data into a description we can simulate
- Investigate the behaviour of the chip by simulation
- Investigate the layout and logic design
- Write up our results on this wiki

### Microphotography

Based on our own work and advice from several professionals in the field

- A 20x objective is great, while 100x is overkill and difficult to work with
- 10x is sometimes adequate for chips with 4 um to 6 um feature sizes, but its better to shoot at higher magnification and downsample the result.
- Useful whole-chip images are typically 6000 to 10000 pixels on a side
- Use an X-Y table to ensure no rotation between the successive images
- A position readout is not needed, and position information from the microscope is not used to stitch images
- Try to get the chip dead level so its entire surface is in the focal plane
- A tip-tilt stage with micrometer drive is essential for this, unless you are very patient
- Use a manual fixed exposure, zoom, and white balance for all images
- Microscopes with a variable zoom are not helpful and could waste a lot of your time later on
- Save images in RAW format if possible at the highest quality
- Aim for at least 200 pixels of overlap between adjacent images

### De-layering

Stripping away individual layers of a chip to reveal the parts and features below can be one of the most difficult and even hazardous procedures owing to the chemicals involved and their byproducts.

- Some labs may use repeated mechanical or chemical-mechanical polishing and photography to image successive layers
- This is more common for modern devices, especially those that have been planarized during manufacture
- It may be riskier and costlier for the older chips we study which have only a single metal layer and whos surfaces are very irregular
- Plasma etching and various chemicals can be used to remove all the material of a particular layer at once

### Resources

Labs:

- [Raw Science](https://www.rawscience.co.uk/reverse-enginering/decapsulation.aspx) a lab in the UK who deprocessed and photographed the Spectrum ULA
- [3g forensics](http://www.3gforensics.co.uk/content.php/203) a lab in the UK who deprocessed the Tube ULA
- [EAG](https://www.eag.com/services/engineering/failure-analysis/) formerly MEFAS, a failure analysis lab in Irvine California, mentioned in [this posting](http://www.atariage.com/forums/topic/136706-internal-antic-and-gtia-schematics/page__view__findpost__p__1651531?s=de4cd5a79909d3bcb06b0384e3039745) by Henry of reactivemicro.com on AtariAge forums

Papers and websites:

- [[1]](http://visual6502.org/downloads.html) Visual6502's PDFs relating to Greg James' presentation at SIGGRAPH 2010
- [Degate](http://www.degate.org/), GPL software to recover netlist from layout, especially of cell-based designs
- [Reverse-Engineering a Cryptographic RFID Tag](http://www.usenix.org/events/sec08/tech/nohl.html) Usenix paper by Nohl, Evans, Starbug and Plötz
- [Reverse-engineering the HP-35](http://www.pmonta.com/calculators/hp-35/) website by Peter Monta
- [The Decapping Project](http://guru.mameworld.info/decap/index.html) website on ROM dumping for MAME
- [Silicon Pr0n](http://siliconpr0n.wikispaces.com/) "A Reverse Engineering Wiki"

Mailing lists, blogs and forum postings:

- [Reversing the Tube ULA (destructively)](http://lists.cloud9.co.uk/pipermail/bbc-micro/2010-October/009437.html) post and thread on the BBC-Micro mailing list. Also found [here](http://mdfs.net/Archive/BBCMicro/2010/10/29/182154.htm)
- [post](http://lists.cloud9.co.uk/pipermail/bbc-micro/2010-October/009443.html) containing Christian Sattler's advice on photography
- [The Decapping Project WIP Page: A Blog About Decapping For MAME](http://decap.mameworld.info/)

See also our [Educational Resources](index.php-title-Educational_Resources.md) page

Retrieved from "[http://visual6502.org/wiki/index.php?title=The\_reverse\_engineering\_process](index.php-title-The_reverse_engineering_process.md)"

