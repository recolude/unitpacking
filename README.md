# Unit Packing
[![Build Status](https://travis-ci.com/recolude/unitpacking.svg?branch=main)](https://travis-ci.com/recolude/unitpacking) [![Go Report Card](https://goreportcard.com/badge/github.com/recolude/unitpacking)](https://goreportcard.com/report/github.com/recolude/unitpacking) [![Coverage](https://codecov.io/gh/recolude/unitpacking/branch/main/graph/badge.svg)](https://codecov.io/gh/recolude/unitpacking)

A library for storing unit vectors in a representnation that lends itself to saving space on disk. You can read more on how some of these methods work [here](https://elicdavis.medium.com/a-story-about-information-entropy-and-efficiently-storing-unit-vectors-92b4a68efe67).

## Benchmark

Calculating the smooth and flat normals for a bunch of famous 3D model datasets. Also one dataset is just 10 million randomly generated unit vectors.

| Dataset | Method | Runtime | Average Error | Uncompressed | Compressed | Compression Ratio |
|-|-|-|-|-|-|-|
| 10 million random | Baseline | N/A | N/A | 117187 KB | 107717 KB | 1.0879 |
| 10 million random | alg24 | 372.9999ms | 0.000704 | 29296 KB | 29305 KB | <div style="color:red">0.9997</div> |
| 10 million random | coarse24 | 327.9995ms | 0.003937 | 29296 KB | 29234 KB | 1.0021 |
| 10 million random | oct24 | 600.9719ms | 0.000389 | 29296 KB | 29305 KB | <div style="color:red">0.9997</div> |
| 10 million random | oct32 | 603.0073ms | 0.000024 | 39062 KB | 39074 KB | <div style="color:red">0.9997</div> |
| armadillo flat | Baseline | N/A | N/A | 585 KB | 436 KB | 1.3415 |
| armadillo flat | alg24 | 1.9999ms | 0.000799 | 146 KB | 146 KB | 1.0013 |
| armadillo flat | coarse24 | 2.0023ms | 0.003943 | 146 KB | 146 KB | 1.0030 |
| armadillo flat | oct24 | 2.9998ms | 0.000364 | 146 KB | 146 KB | 1.0015 |
| armadillo flat | oct32 | 2.9998ms | 0.000023 | 195 KB | 172 KB | 1.1347 |
| armadillo smooth | Baseline | N/A | N/A | 585 KB | 540 KB | 1.0836 |
| armadillo smooth | alg24 | 1ms | 0.000805 | 146 KB | 146 KB | <div style="color:red">0.9996</div> |
| armadillo smooth | coarse24 | 2.0019ms | 0.003932 | 146 KB | 146 KB | 1.0003 |
| armadillo smooth | oct24 | 2.0001ms | 0.000365 | 146 KB | 146 KB | <div style="color:red">0.9996</div> |
| armadillo smooth | oct32 | 2.9997ms | 0.000023 | 195 KB | 195 KB | <div style="color:red">0.9997</div> |
| beetle-alt flat | Baseline | N/A | N/A | 233 KB | 189 KB | 1.2318 |
| beetle-alt flat | alg24 | 1.0018ms | 0.001217 | 58 KB | 56 KB | 1.0225 |
| beetle-alt flat | coarse24 | 998.7µs | 0.003954 | 58 KB | 54 KB | 1.0611 |
| beetle-alt flat | oct24 | 1.0194ms | 0.000343 | 58 KB | 57 KB | 1.0092 |
| beetle-alt flat | oct32 | 1.0323ms | 0.000021 | 77 KB | 74 KB | 1.0433 |
| beetle-alt smooth | Baseline | N/A | N/A | 233 KB | 212 KB | 1.0965 |
| beetle-alt smooth | alg24 | 1.004ms | 0.000865 | 58 KB | 57 KB | 1.0138 |
| beetle-alt smooth | coarse24 | 1.0034ms | 0.003950 | 58 KB | 56 KB | 1.0402 |
| beetle-alt smooth | oct24 | 1.0086ms | 0.000380 | 58 KB | 57 KB | 1.0128 |
| beetle-alt smooth | oct32 | 1.0006ms | 0.000024 | 77 KB | 76 KB | 1.0101 |
| beetle flat | Baseline | N/A | N/A | 13 KB | 11 KB | 1.1814 |
| beetle flat | alg24 | 0s | 0.001273 | 3 KB | 3 KB | 1.0053 |
| beetle flat | coarse24 | 0s | 0.003769 | 3 KB | 3 KB | 1.0214 |
| beetle flat | oct24 | 0s | 0.000335 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| beetle flat | oct32 | 0s | 0.000021 | 4 KB | 4 KB | 1.0220 |
| beetle smooth | Baseline | N/A | N/A | 13 KB | 12 KB | 1.0984 |
| beetle smooth | alg24 | 0s | 0.000851 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| beetle smooth | coarse24 | 0s | 0.003959 | 3 KB | 3 KB | 1.0106 |
| beetle smooth | oct24 | 0s | 0.000391 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| beetle smooth | oct32 | 0s | 0.000025 | 4 KB | 4 KB | 1.0000 |
| cheburashka flat | Baseline | N/A | N/A | 78 KB | 43 KB | 1.7861 |
| cheburashka flat | alg24 | 0s | 0.000642 | 19 KB | 14 KB | 1.3377 |
| cheburashka flat | coarse24 | 0s | 0.003960 | 19 KB | 14 KB | 1.3455 |
| cheburashka flat | oct24 | 1.0002ms | 0.000348 | 19 KB | 14 KB | 1.3350 |
| cheburashka flat | oct32 | 0s | 0.000022 | 26 KB | 16 KB | 1.5818 |
| cheburashka smooth | Baseline | N/A | N/A | 78 KB | 72 KB | 1.0831 |
| cheburashka smooth | alg24 | 0s | 0.000616 | 19 KB | 19 KB | <div style="color:red">0.9993</div> |
| cheburashka smooth | coarse24 | 1ms | 0.003965 | 19 KB | 19 KB | 1.0033 |
| cheburashka smooth | oct24 | 0s | 0.000362 | 19 KB | 19 KB | <div style="color:red">0.9993</div> |
| cheburashka smooth | oct32 | 999.6µs | 0.000023 | 26 KB | 26 KB | <div style="color:red">0.9994</div> |
| cow flat | Baseline | N/A | N/A | 34 KB | 23 KB | 1.4391 |
| cow flat | alg24 | 0s | 0.000735 | 8 KB | 8 KB | 1.0099 |
| cow flat | coarse24 | 0s | 0.003966 | 8 KB | 8 KB | 1.0188 |
| cow flat | oct24 | 0s | 0.000339 | 8 KB | 8 KB | 1.0071 |
| cow flat | oct32 | 0s | 0.000021 | 11 KB | 10 KB | 1.1076 |
| cow smooth | Baseline | N/A | N/A | 34 KB | 31 KB | 1.0812 |
| cow smooth | alg24 | 0s | 0.000608 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| cow smooth | coarse24 | 0s | 0.003939 | 8 KB | 8 KB | 1.0036 |
| cow smooth | oct24 | 0s | 0.000364 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| cow smooth | oct32 | 0s | 0.000022 | 11 KB | 11 KB | <div style="color:red">0.9991</div> |
| fandisk flat | Baseline | N/A | N/A | 75 KB | 43 KB | 1.7632 |
| fandisk flat | alg24 | 0s | 0.001230 | 18 KB | 7 KB | 2.4786 |
| fandisk flat | coarse24 | 0s | 0.002818 | 18 KB | 6 KB | 2.9539 |
| fandisk flat | oct24 | 0s | 0.000169 | 18 KB | 9 KB | 2.0176 |
| fandisk flat | oct32 | 0s | 0.000010 | 25 KB | 14 KB | 1.7590 |
| fandisk smooth | Baseline | N/A | N/A | 75 KB | 48 KB | 1.5731 |
| fandisk smooth | alg24 | 0s | 0.001147 | 18 KB | 9 KB | 1.9277 |
| fandisk smooth | coarse24 | 0s | 0.004274 | 18 KB | 8 KB | 2.2281 |
| fandisk smooth | oct24 | 984.3µs | 0.000322 | 18 KB | 10 KB | 1.8154 |
| fandisk smooth | oct32 | 997.5µs | 0.000022 | 25 KB | 15 KB | 1.5830 |
| happy flat | Baseline | N/A | N/A | 577 KB | 496 KB | 1.1628 |
| happy flat | alg24 | 2.0047ms | 0.000865 | 144 KB | 143 KB | 1.0029 |
| happy flat | coarse24 | 2.0153ms | 0.003936 | 144 KB | 142 KB | 1.0102 |
| happy flat | oct24 | 2.9987ms | 0.000346 | 144 KB | 143 KB | 1.0029 |
| happy flat | oct32 | 2.9887ms | 0.000022 | 192 KB | 183 KB | 1.0503 |
| happy smooth | Baseline | N/A | N/A | 577 KB | 531 KB | 1.0851 |
| happy smooth | alg24 | 1.9531ms | 0.000711 | 144 KB | 144 KB | 1.0015 |
| happy smooth | coarse24 | 2.0008ms | 0.003931 | 144 KB | 143 KB | 1.0081 |
| happy smooth | oct24 | 3.0002ms | 0.000359 | 144 KB | 143 KB | 1.0027 |
| happy smooth | oct32 | 2.998ms | 0.000022 | 192 KB | 191 KB | 1.0024 |
| horse flat | Baseline | N/A | N/A | 568 KB | 523 KB | 1.0848 |
| horse flat | alg24 | 1.9993ms | 0.001013 | 142 KB | 140 KB | 1.0076 |
| horse flat | coarse24 | 999.3µs | 0.003928 | 142 KB | 138 KB | 1.0266 |
| horse flat | oct24 | 2.0039ms | 0.000350 | 142 KB | 141 KB | 1.0038 |
| horse flat | oct32 | 2ms | 0.000022 | 189 KB | 188 KB | 1.0043 |
| horse smooth | Baseline | N/A | N/A | 568 KB | 523 KB | 1.0863 |
| horse smooth | alg24 | 2.9995ms | 0.000968 | 142 KB | 141 KB | 1.0069 |
| horse smooth | coarse24 | 1.998ms | 0.003948 | 142 KB | 138 KB | 1.0272 |
| horse smooth | oct24 | 2.9997ms | 0.000360 | 142 KB | 141 KB | 1.0054 |
| horse smooth | oct32 | 1.9995ms | 0.000022 | 189 KB | 188 KB | 1.0051 |
| igea flat | Baseline | N/A | N/A | 1574 KB | 1306 KB | 1.2051 |
| igea flat | alg24 | 4.9878ms | 0.000781 | 393 KB | 379 KB | 1.0384 |
| igea flat | coarse24 | 5.0022ms | 0.003924 | 393 KB | 368 KB | 1.0685 |
| igea flat | oct24 | 8.0014ms | 0.000356 | 393 KB | 381 KB | 1.0314 |
| igea flat | oct32 | 9.0006ms | 0.000022 | 524 KB | 485 KB | 1.0815 |
| igea smooth | Baseline | N/A | N/A | 1574 KB | 1445 KB | 1.0892 |
| igea smooth | alg24 | 5.0004ms | 0.000931 | 393 KB | 385 KB | 1.0202 |
| igea smooth | coarse24 | 3.9984ms | 0.003938 | 393 KB | 374 KB | 1.0516 |
| igea smooth | oct24 | 8.9652ms | 0.000350 | 393 KB | 387 KB | 1.0150 |
| igea smooth | oct32 | 7.0012ms | 0.000022 | 524 KB | 520 KB | 1.0087 |
| lucy flat | Baseline | N/A | N/A | 585 KB | 425 KB | 1.3773 |
| lucy flat | alg24 | 1.9975ms | 0.001102 | 146 KB | 145 KB | 1.0067 |
| lucy flat | coarse24 | 2.0018ms | 0.003923 | 146 KB | 144 KB | 1.0106 |
| lucy flat | oct24 | 2.9989ms | 0.000355 | 146 KB | 145 KB | 1.0064 |
| lucy flat | oct32 | 2.9989ms | 0.000022 | 195 KB | 165 KB | 1.1775 |
| lucy smooth | Baseline | N/A | N/A | 585 KB | 539 KB | 1.0868 |
| lucy smooth | alg24 | 2.0408ms | 0.001102 | 146 KB | 145 KB | 1.0033 |
| lucy smooth | coarse24 | 967.9µs | 0.003938 | 146 KB | 145 KB | 1.0075 |
| lucy smooth | oct24 | 2.9992ms | 0.000362 | 146 KB | 146 KB | 1.0029 |
| lucy smooth | oct32 | 2.0002ms | 0.000023 | 195 KB | 194 KB | 1.0026 |
| max-planck flat | Baseline | N/A | N/A | 586 KB | 519 KB | 1.1289 |
| max-planck flat | alg24 | 2ms | 0.000718 | 146 KB | 145 KB | 1.0075 |
| max-planck flat | coarse24 | 1.001ms | 0.003942 | 146 KB | 145 KB | 1.0095 |
| max-planck flat | oct24 | 2.0003ms | 0.000358 | 146 KB | 146 KB | 1.0045 |
| max-planck flat | oct32 | 3.0012ms | 0.000022 | 195 KB | 190 KB | 1.0275 |
| max-planck smooth | Baseline | N/A | N/A | 586 KB | 540 KB | 1.0866 |
| max-planck smooth | alg24 | 2.0002ms | 0.000703 | 146 KB | 145 KB | 1.0062 |
| max-planck smooth | coarse24 | 2.0005ms | 0.003940 | 146 KB | 145 KB | 1.0086 |
| max-planck smooth | oct24 | 2.9989ms | 0.000360 | 146 KB | 146 KB | 1.0032 |
| max-planck smooth | oct32 | 1.999ms | 0.000022 | 195 KB | 195 KB | 1.0025 |
| nefertiti flat | Baseline | N/A | N/A | 585 KB | 424 KB | 1.3788 |
| nefertiti flat | alg24 | 2.0006ms | 0.000848 | 146 KB | 138 KB | 1.0563 |
| nefertiti flat | coarse24 | 1.9998ms | 0.003932 | 146 KB | 135 KB | 1.0818 |
| nefertiti flat | oct24 | 2.9996ms | 0.000353 | 146 KB | 139 KB | 1.0470 |
| nefertiti flat | oct32 | 3.0256ms | 0.000022 | 195 KB | 166 KB | 1.1747 |
| nefertiti smooth | Baseline | N/A | N/A | 585 KB | 534 KB | 1.0964 |
| nefertiti smooth | alg24 | 2.0001ms | 0.000806 | 146 KB | 140 KB | 1.0417 |
| nefertiti smooth | coarse24 | 1.0005ms | 0.003948 | 146 KB | 135 KB | 1.0793 |
| nefertiti smooth | oct24 | 3.001ms | 0.000371 | 146 KB | 141 KB | 1.0339 |
| nefertiti smooth | oct32 | 3.0012ms | 0.000023 | 195 KB | 191 KB | 1.0215 |
| ogre flat | Baseline | N/A | N/A | 728 KB | 521 KB | 1.3967 |
| ogre flat | alg24 | 2.966ms | 0.000786 | 182 KB | 177 KB | 1.0248 |
| ogre flat | coarse24 | 3.0169ms | 0.003829 | 182 KB | 175 KB | 1.0397 |
| ogre flat | oct24 | 3.9997ms | 0.000334 | 182 KB | 178 KB | 1.0225 |
| ogre flat | oct32 | 3.9973ms | 0.000021 | 242 KB | 214 KB | 1.1346 |
| ogre smooth | Baseline | N/A | N/A | 728 KB | 663 KB | 1.0981 |
| ogre smooth | alg24 | 1.9998ms | 0.000838 | 182 KB | 180 KB | 1.0078 |
| ogre smooth | coarse24 | 1.9984ms | 0.003944 | 182 KB | 178 KB | 1.0183 |
| ogre smooth | oct24 | 3.9972ms | 0.000355 | 182 KB | 181 KB | 1.0064 |
| ogre smooth | oct32 | 3.0011ms | 0.000022 | 242 KB | 240 KB | 1.0094 |
| rocker-arm flat | Baseline | N/A | N/A | 117 KB | 92 KB | 1.2751 |
| rocker-arm flat | alg24 | 0s | 0.001186 | 29 KB | 29 KB | 1.0086 |
| rocker-arm flat | coarse24 | 0s | 0.003932 | 29 KB | 28 KB | 1.0248 |
| rocker-arm flat | oct24 | 0s | 0.000331 | 29 KB | 29 KB | <div style="color:red">0.9995</div> |
| rocker-arm flat | oct32 | 998.8µs | 0.000021 | 39 KB | 35 KB | 1.1119 |
| rocker-arm smooth | Baseline | N/A | N/A | 117 KB | 108 KB | 1.0832 |
| rocker-arm smooth | alg24 | 998.5µs | 0.000970 | 29 KB | 29 KB | 1.0037 |
| rocker-arm smooth | coarse24 | 0s | 0.003938 | 29 KB | 29 KB | 1.0125 |
| rocker-arm smooth | oct24 | 1.0006ms | 0.000349 | 29 KB | 29 KB | <div style="color:red">0.9999</div> |
| rocker-arm smooth | oct32 | 0s | 0.000022 | 39 KB | 39 KB | <div style="color:red">0.9995</div> |
| spot flat | Baseline | N/A | N/A | 34 KB | 20 KB | 1.6768 |
| spot flat | alg24 | 0s | 0.000941 | 8 KB | 8 KB | 1.0150 |
| spot flat | coarse24 | 0s | 0.003928 | 8 KB | 8 KB | 1.0301 |
| spot flat | oct24 | 0s | 0.000344 | 8 KB | 8 KB | 1.0123 |
| spot flat | oct32 | 0s | 0.000021 | 11 KB | 10 KB | 1.1361 |
| spot smooth | Baseline | N/A | N/A | 34 KB | 31 KB | 1.0866 |
| spot smooth | alg24 | 0s | 0.000808 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| spot smooth | coarse24 | 0s | 0.003954 | 8 KB | 8 KB | 1.0047 |
| spot smooth | oct24 | 0s | 0.000358 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| spot smooth | oct32 | 0s | 0.000023 | 11 KB | 11 KB | <div style="color:red">0.9991</div> |
| stanford-bunny flat | Baseline | N/A | N/A | 421 KB | 317 KB | 1.3285 |
| stanford-bunny flat | alg24 | 1.0003ms | 0.000913 | 105 KB | 100 KB | 1.0492 |
| stanford-bunny flat | coarse24 | 2.0005ms | 0.003875 | 105 KB | 97 KB | 1.0855 |
| stanford-bunny flat | oct24 | 2.0006ms | 0.000352 | 105 KB | 101 KB | 1.0410 |
| stanford-bunny flat | oct32 | 3.0007ms | 0.000022 | 140 KB | 123 KB | 1.1341 |
| stanford-bunny smooth | Baseline | N/A | N/A | 421 KB | 376 KB | 1.1186 |
| stanford-bunny smooth | alg24 | 1.0028ms | 0.000712 | 105 KB | 102 KB | 1.0228 |
| stanford-bunny smooth | coarse24 | 999.1µs | 0.003903 | 105 KB | 99 KB | 1.0543 |
| stanford-bunny smooth | oct24 | 2.0005ms | 0.000372 | 105 KB | 102 KB | 1.0241 |
| stanford-bunny smooth | oct32 | 2.0008ms | 0.000023 | 140 KB | 137 KB | 1.0245 |
| teapot flat | Baseline | N/A | N/A | 42 KB | 30 KB | 1.4060 |
| teapot flat | alg24 | 0s | 0.001040 | 10 KB | 10 KB | 1.0319 |
| teapot flat | coarse24 | 0s | 0.003842 | 10 KB | 10 KB | 1.0600 |
| teapot flat | oct24 | 0s | 0.000339 | 10 KB | 10 KB | 1.0135 |
| teapot flat | oct32 | 991.5µs | 0.000021 | 14 KB | 13 KB | 1.0668 |
| teapot smooth | Baseline | N/A | N/A | 42 KB | 39 KB | 1.0864 |
| teapot smooth | alg24 | 0s | 0.000972 | 10 KB | 10 KB | 1.0016 |
| teapot smooth | coarse24 | 0s | 0.003953 | 10 KB | 10 KB | 1.0189 |
| teapot smooth | oct24 | 0s | 0.000365 | 10 KB | 10 KB | <div style="color:red">0.9995</div> |
| teapot smooth | oct32 | 0s | 0.000023 | 14 KB | 14 KB | <div style="color:red">0.9995</div> |
| xyzrgb_dragon flat | Baseline | N/A | N/A | 1465 KB | 1250 KB | 1.1722 |
| xyzrgb_dragon flat | alg24 | 5.0017ms | 0.000775 | 366 KB | 364 KB | 1.0049 |
| xyzrgb_dragon flat | coarse24 | 4.0197ms | 0.003935 | 366 KB | 363 KB | 1.0090 |
| xyzrgb_dragon flat | oct24 | 6.999ms | 0.000367 | 366 KB | 364 KB | 1.0045 |
| xyzrgb_dragon flat | oct32 | 6.996ms | 0.000023 | 488 KB | 465 KB | 1.0503 |
| xyzrgb_dragon smooth | Baseline | N/A | N/A | 1465 KB | 1343 KB | 1.0906 |
| xyzrgb_dragon smooth | alg24 | 4.9998ms | 0.000733 | 366 KB | 364 KB | 1.0042 |
| xyzrgb_dragon smooth | coarse24 | 4.0002ms | 0.003942 | 366 KB | 362 KB | 1.0102 |
| xyzrgb_dragon smooth | oct24 | 7.0011ms | 0.000378 | 366 KB | 364 KB | 1.0048 |
| xyzrgb_dragon smooth | oct32 | 6.9976ms | 0.000024 | 488 KB | 486 KB | 1.0046 |

## Resources

* [A Survey of Efficient Representations for Independent Unit Vectors](http://jcgt.org/published/0003/02/01/paper.pdf)
* [Alec Jacobson's Common 3D Test Models](https://github.com/alecjacobson/common-3d-test-models)