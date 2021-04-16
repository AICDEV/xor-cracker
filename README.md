# xor-cracker

A simple, but effective, xor breaker written in golang. The breaker uses frequency analysis compined with hamming distance to computed the most probable key that was used to xor a message.

## Usage
For a better compatibility, the encrypted messages must be encoded in hexadecimal.

### Encryption example
On the following page you can find a predefined recipe to XOR encode a message and to encode the output in Hex. [Create and encrypted test message](https://cyberchef.nullteilerfrei.de/#recipe=XOR(%7B'option':'UTF8','string':'golang'%7D,'Standard',false)To_Hex('None',0)&input=V2hlbiBHbyBtb2R1bGVzIHdlcmUgcmVsZWFzZWQgdHdvIHllYXJzIGFnbywgdGhleSBjb21wbGV0ZWx5IGNoYW5nZWQgdGhlIGxhbmRzY2FwZSBvZiBHbyBkZXZlbG9wZXIgdG9vbGluZy4gVG9vbHMgbGlrZSBnb2ltcG9ydHMgYW5kIGdvZGVmIHByZXZpb3VzbHkgZGVwZW5kZWQgb24gdGhlIGZhY3QgdGhhdCBjb2RlIHdhcyBzdG9yZWQgaW4geW91ciAkR09QQVRILiBXaGVuIHRoZSBHbyB0ZWFtIGJlZ2FuIHJld3JpdGluZyB0aGVzZSB0b29scyB0byB3b3JrIHdpdGggbW9kdWxlcywgd2UgaW1tZWRpYXRlbHkgcmVhbGl6ZWQgdGhhdCB3ZSBuZWVkZWQgYSBtb3JlIHN5c3RlbWF0aWMgYXBwcm9hY2ggdG8gYnJpZGdlIHRoZSBnYXAuCgpBcyBhIHJlc3VsdCwgd2UgYmVnYW4gd29ya2luZyBvbiBhIHNpbmdsZSBHbyBsYW5ndWFnZSBzZXJ2ZXIsIGdvcGxzLCB3aGljaCBwcm92aWRlcyBJREUgZmVhdHVyZXMsIHN1Y2ggYXMgYXV0b2NvbXBsZXRpb24sIGZvcm1hdHRpbmcsIGFuZCBkaWFnbm9zdGljcyB0byBhbnkgY29tcGF0aWJsZSBlZGl0b3IgZnJvbnRlbmQuIFRoaXMgcGVyc2lzdGVudCBhbmQgdW5pZmllZCBzZXJ2ZXIgaXMgYSBmdW5kYW1lbnRhbCBzaGlmdCBmcm9tIHRoZSBlYXJsaWVyIGNvbGxlY3Rpb25zIG9mIGNvbW1hbmQtbGluZSB0b29scy4KCkluIGFkZGl0aW9uIHRvIHdvcmtpbmcgb24gZ29wbHMsIHdlIHNvdWdodCBvdGhlciB3YXlzIG9mIGNyZWF0aW5nIGEgc3RhYmxlIGVjb3N5c3RlbSBvZiBlZGl0b3IgdG9vbGluZy4gTGFzdCB5ZWFyLCB0aGUgR28gdGVhbSB0b29rIHJlc3BvbnNpYmlsaXR5IGZvciB0aGUgR28gZXh0ZW5zaW9uIGZvciBWUyBDb2RlLiBBcyBwYXJ0IG9mIHRoaXMgd29yaywgd2Ugc21vb3RoZWQgdGhlIGV4dGVuc2lvbuKAmXMgaW50ZWdyYXRpb24gd2l0aCB0aGUgbGFuZ3VhZ2Ugc2VydmVy4oCUYXV0b21hdGluZyBnb3BscyB1cGRhdGVzLCByZWFycmFuZ2luZyBhbmQgY2xhcmlmeWluZyBnb3BscyBzZXR0aW5ncywgaW1wcm92aW5nIHRoZSB0cm91Ymxlc2hvb3Rpbmcgd29ya2Zsb3csIGFuZCBzb2xpY2l0aW5nIGZlZWRiYWNrIHRocm91Z2ggYSBzdXJ2ZXkuIFdl4oCZdmUgYWxzbyBjb250aW51ZWQgdG8gZm9zdGVyIGEgY29tbXVuaXR5IG9mIGFjdGl2ZSB1c2VycyBhbmQgY29udHJpYnV0b3JzIHdobyBoYXZlIGhlbHBlZCB1cyBpbXByb3ZlIHRoZSBzdGFiaWxpdHksIHBlcmZvcm1hbmNlLCBhbmQgdXNlciBleHBlcmllbmNlIG9mIHRoZSBHbyBleHRlbnNpb24u)

### Crack Encrypted Message
Get your encrypted message in hexadecimal encoding. Param "-m" is your encrypted message and "-l" is the maximum key length the programm should test for.
```sh
go run main.go -m 3007090f4e20084f010e0a120b0a1f411902150a4c130b0b020e1f040a47131803411702061d1f410f0008434c1506021e4f0f0e03170b0a1804021e470c04000000020b4c15060247030d0f0a14040e1c044e08014f2b0e4e030219090d0117021d4c1501080b06020640473300030d1d470b0607044e00080601110115131c4c000003470803050b01471f1e04180e081a1f0d1747030a1c040003020b4c0e0047130709410806041b4c150606134f0f0e0a0247180d124e1413001e040a470e014c180112154f48262137263b244f4e300f0a02411a0f024f2b0e4e13020e01410c02000e02411c02101d05150709004f18090b14024f180e010b144f180e4e10081d0741190e13074c0c0103120309124247100a4c08030a020b05001a020b164c130b060b0616040a4713070d154e10024f02040b03020b4c004e0a081d09411d1e141b090c0f130e0c4c001e1715000d02064713004c031c0e030809411a0f024f0b001e496d652d124e06471d09121b0b13434c160b47050a0b00004710001e0a0709004f030f4e06471c050f090b024f2b0e4e0b06010b140f00024f1f041c11021d4041090817031f4d4e100f060f094e1715001a080a02144f25252b47010a0d151b15021c40411d1204074c001d47061a180e0d080a1f00041a0e08014041080815020d151a0e090840410f09034f08080f0009001f150704144f180e4e0609164c02010a170e18080c0b024f09050713081d4c071c08091b090f0a49473b04081d47170a1e120714130a02154e06090b4c14000e010609054e14021d1a041c470e1c4c004e01120108000302091b0d0d4e140f060a154e01150001411a0f024f09001c0b0e0a1e410d080b0309021a0e08011f410101470c030c0306090b410d0709024f180e010b1441666b2709470e080507130e0002411a0847180313050e09084c0e004700001c0d1d4b471809411d08120804154e08130709134e1006161f410101470c1e040f130e010b410f47141b0d030202470a0f0e1d1e141b090c4e08014f09050713081d4c1501080b06020640472b0e1f154e1e020e1e4d4e130f0a4c260147130a0d0c4e13080007411c02141f030f1d0e050600081a1e470903134e130f0a4c2601470217180400140e0002410808154f3a324e24080b094f4e26144f1c001c1347000a411a0f0e1c4c1601150c434c160b471402030e1a0f020b4c150602470a14150b091406030f8ce7fe1c4c08001302081e001a0e08014c1607130f4f18090b470b0e02061b06000a4c120b15110a1e83eef3061a180e0306130602064e00081f00124e12170b0d150b144b4f1e040f15150e02060709004f0d0f0a4704030d1307011e0602064e00081f00124e14021b1808000014434c08031715001a080000471b04044e131500190302021407030e1a0e09084c1601150c09000e194b470e02054e140803050207130e010b410802020b0e000d0c471b0413011200074c004e14121d1a04174947380983eefe110a4c000214084f0f0e00130e0119040a4713004c070114130a1e410f470400010c1b090e1b15410101470e0f150711024f19120b15144f0d0f0a47040002151c0e051a180e1c144718040e4e0f0619094106020b1f09054e12144f050c1e15081909411a0f024f1f150f050e030515174b471f0913080815020d0f0d024b4f0d0f0a47121c09134e021f1f09130702090c09410101471b04044e20084f09191a02091c050e0049 -l 10
```

### Example program output
![Example Output](https://github.com/AICDEV/xor-cracker/blob/master/xor_cracker_example.png)

## Tech
The project is implemented in [golang](https://golang.org/). All calculations are based on:
* [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance)
* [Frequency Analysis](https://en.wikipedia.org/wiki/Frequency_analysis)

### How to train new languages
There are three reference texts in english language in the project. But you can simply put documents (.txt files) in any language in the data folder. Just make sure that there is enough text and that you do not mix languages !!
