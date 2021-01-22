# Unit Packing
[![Build Status](https://travis-ci.com/recolude/unitpacking.svg?branch=main)](https://travis-ci.com/recolude/unitpacking) [![Go Report Card](https://goreportcard.com/badge/github.com/recolude/unitpacking)](https://goreportcard.com/report/github.com/recolude/unitpacking) [![Coverage](https://codecov.io/gh/recolude/unitpacking/branch/main/graph/badge.svg)](https://codecov.io/gh/recolude/unitpacking)

A library for storing unit vectors in a representnation that lends itself to saving space on disk. You can read more on how some of these methods work [here](https://elicdavis.medium.com/a-story-about-information-entropy-and-efficiently-storing-unit-vectors-92b4a68efe67).

## Benchmark

Calculating the smooth and flat normals for a bunch of famous datasets. Also one dataset is just 10 million randomly generated unit vectors.

| Dataset | Method | Runtime | Average Error | Uncompressed | Compressed | Compression Ratio |
|-|-|-|-|-|-|-|
| 10 million random | Baseline | N/A | N/A | 117187 KB | 107717 KB | 1.0879 |
| 10 million random | alg24 | 388.9914ms | 0.000704 | 29296 KB | 29305 KB | <div style="color:red">0.9997</div> |
| 10 million random | coarse24 | 334.0091ms | 0.008057 | 29296 KB | 29274 KB | 1.0008 |
| 10 million random | oct24 | 661.999ms | 0.000389 | 29296 KB | 29305 KB | <div style="color:red">0.9997</div> |
| alligator Flat | Baseline | N/A | N/A | 37 KB | 0 KB | 39.6866 |
| alligator Flat | alg24 | 1.0341ms | 0.000000 | 9 KB | 0 KB | 283.0588 |
| alligator Flat | coarse24 | 0s | 0.000379 | 9 KB | 0 KB | 20.7414 |
| alligator Flat | oct24 | 0s | 0.000000 | 9 KB | 0 KB | 283.0588 |
| alligator Smooth | Baseline | N/A | N/A | 37 KB | 1 KB | 23.3734 |
| alligator Smooth | alg24 | 0s | 0.000310 | 9 KB | 1 KB | 7.4547 |
| alligator Smooth | coarse24 | 0s | 0.005084 | 9 KB | 1 KB | 7.6200 |
| alligator Smooth | oct24 | 0s | 0.000331 | 9 KB | 1 KB | 7.4837 |
| armadillo Flat | Baseline | N/A | N/A | 585 KB | 436 KB | 1.3415 |
| armadillo Flat | alg24 | 2.0004ms | 0.000799 | 146 KB | 146 KB | 1.0013 |
| armadillo Flat | coarse24 | 996.5µs | 0.014417 | 146 KB | 146 KB | 1.0018 |
| armadillo Flat | oct24 | 3.0012ms | 0.000364 | 146 KB | 146 KB | 1.0015 |
| armadillo Smooth | Baseline | N/A | N/A | 585 KB | 540 KB | 1.0836 |
| armadillo Smooth | alg24 | 2.0011ms | 0.000805 | 146 KB | 146 KB | <div style="color:red">0.9996</div> |
| armadillo Smooth | coarse24 | 999.8µs | 0.010859 | 146 KB | 146 KB | <div style="color:red">0.9998</div> |
| armadillo Smooth | oct24 | 3.0002ms | 0.000365 | 146 KB | 146 KB | <div style="color:red">0.9996</div> |
| beetle-alt Flat | Baseline | N/A | N/A | 233 KB | 189 KB | 1.2318 |
| beetle-alt Flat | alg24 | 0s | 0.001217 | 58 KB | 56 KB | 1.0225 |
| beetle-alt Flat | coarse24 | 999.7µs | 0.021670 | 58 KB | 54 KB | 1.0594 |
| beetle-alt Flat | oct24 | 999.3µs | 0.000343 | 58 KB | 57 KB | 1.0092 |
| beetle-alt Smooth | Baseline | N/A | N/A | 233 KB | 212 KB | 1.0965 |
| beetle-alt Smooth | alg24 | 968.5µs | 0.000865 | 58 KB | 57 KB | 1.0138 |
| beetle-alt Smooth | coarse24 | 1.0029ms | 0.003912 | 58 KB | 56 KB | 1.0398 |
| beetle-alt Smooth | oct24 | 1ms | 0.000380 | 58 KB | 57 KB | 1.0128 |
| beetle Flat | Baseline | N/A | N/A | 13 KB | 11 KB | 1.1814 |
| beetle Flat | alg24 | 0s | 0.001283 | 3 KB | 3 KB | 1.0053 |
| beetle Flat | coarse24 | 0s | 0.024597 | 3 KB | 3 KB | 1.0198 |
| beetle Flat | oct24 | 0s | 0.000335 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| beetle Smooth | Baseline | N/A | N/A | 13 KB | 12 KB | 1.0984 |
| beetle Smooth | alg24 | 0s | 0.000851 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| beetle Smooth | coarse24 | 0s | 0.004469 | 3 KB | 3 KB | 1.0088 |
| beetle Smooth | oct24 | 0s | 0.000391 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| cheburashka Flat | Baseline | N/A | N/A | 78 KB | 43 KB | 1.7861 |
| cheburashka Flat | alg24 | 998.1µs | 0.000642 | 19 KB | 14 KB | 1.3377 |
| cheburashka Flat | coarse24 | 1.0003ms | 0.026707 | 19 KB | 14 KB | 1.3456 |
| cheburashka Flat | oct24 | 1.0001ms | 0.000348 | 19 KB | 14 KB | 1.3350 |
| cheburashka Smooth | Baseline | N/A | N/A | 78 KB | 72 KB | 1.0831 |
| cheburashka Smooth | alg24 | 995.5µs | 0.000616 | 19 KB | 19 KB | <div style="color:red">0.9993</div> |
| cheburashka Smooth | coarse24 | 0s | 0.014946 | 19 KB | 19 KB | 1.0025 |
| cheburashka Smooth | oct24 | 0s | 0.000362 | 19 KB | 19 KB | <div style="color:red">0.9993</div> |
| cow Flat | Baseline | N/A | N/A | 34 KB | 23 KB | 1.4391 |
| cow Flat | alg24 | 988.5µs | 0.000735 | 8 KB | 8 KB | 1.0099 |
| cow Flat | coarse24 | 0s | 0.019476 | 8 KB | 8 KB | 1.0171 |
| cow Flat | oct24 | 0s | 0.000339 | 8 KB | 8 KB | 1.0071 |
| cow Smooth | Baseline | N/A | N/A | 34 KB | 31 KB | 1.0812 |
| cow Smooth | alg24 | 0s | 0.000608 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| cow Smooth | coarse24 | 0s | 0.013519 | 8 KB | 8 KB | 1.0020 |
| cow Smooth | oct24 | 0s | 0.000364 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| fandisk Flat | Baseline | N/A | N/A | 75 KB | 43 KB | 1.7632 |
| fandisk Flat | alg24 | 1.0093ms | 0.001277 | 18 KB | 7 KB | 2.4617 |
| fandisk Flat | coarse24 | 1.001ms | 0.041221 | 18 KB | 6 KB | 2.8349 |
| fandisk Flat | oct24 | 999.1µs | 0.000169 | 18 KB | 9 KB | 2.0176 |
| fandisk Smooth | Baseline | N/A | N/A | 75 KB | 48 KB | 1.5731 |
| fandisk Smooth | alg24 | 0s | 0.001147 | 18 KB | 9 KB | 1.9277 |
| fandisk Smooth | coarse24 | 0s | 0.004418 | 18 KB | 8 KB | 2.2553 |
| fandisk Smooth | oct24 | 999.9µs | 0.000322 | 18 KB | 10 KB | 1.8154 |
| happy Flat | Baseline | N/A | N/A | 577 KB | 496 KB | 1.1628 |
| happy Flat | alg24 | 2.0228ms | 0.000865 | 144 KB | 143 KB | 1.0029 |
| happy Flat | coarse24 | 1ms | 0.023530 | 144 KB | 143 KB | 1.0087 |
| happy Flat | oct24 | 2.008ms | 0.000346 | 144 KB | 143 KB | 1.0029 |
| happy Smooth | Baseline | N/A | N/A | 577 KB | 531 KB | 1.0851 |
| happy Smooth | alg24 | 1.9871ms | 0.000711 | 144 KB | 144 KB | 1.0015 |
| happy Smooth | coarse24 | 1.9998ms | 0.010774 | 144 KB | 143 KB | 1.0070 |
| happy Smooth | oct24 | 2.9961ms | 0.000359 | 144 KB | 143 KB | 1.0027 |
| horse Flat | Baseline | N/A | N/A | 568 KB | 523 KB | 1.0848 |
| horse Flat | alg24 | 1.9999ms | 0.001013 | 142 KB | 140 KB | 1.0076 |
| horse Flat | coarse24 | 2ms | 0.019190 | 142 KB | 138 KB | 1.0251 |
| horse Flat | oct24 | 2.9887ms | 0.000350 | 142 KB | 141 KB | 1.0038 |
| horse Smooth | Baseline | N/A | N/A | 568 KB | 523 KB | 1.0863 |
| horse Smooth | alg24 | 3.0045ms | 0.000968 | 142 KB | 141 KB | 1.0069 |
| horse Smooth | coarse24 | 2.0007ms | 0.010709 | 142 KB | 138 KB | 1.0253 |
| horse Smooth | oct24 | 1.9996ms | 0.000360 | 142 KB | 141 KB | 1.0054 |
| igea Flat | Baseline | N/A | N/A | 1574 KB | 1306 KB | 1.2051 |
| igea Flat | alg24 | 5.9969ms | 0.000781 | 393 KB | 379 KB | 1.0384 |
| igea Flat | coarse24 | 4.9941ms | 0.015230 | 393 KB | 369 KB | 1.0658 |
| igea Flat | oct24 | 6.9925ms | 0.000356 | 393 KB | 381 KB | 1.0314 |
| igea Smooth | Baseline | N/A | N/A | 1574 KB | 1445 KB | 1.0892 |
| igea Smooth | alg24 | 4.9971ms | 0.000931 | 393 KB | 385 KB | 1.0202 |
| igea Smooth | coarse24 | 5.9996ms | 0.011622 | 393 KB | 374 KB | 1.0504 |
| igea Smooth | oct24 | 7.989ms | 0.000350 | 393 KB | 387 KB | 1.0150 |
| lucy Flat | Baseline | N/A | N/A | 585 KB | 425 KB | 1.3773 |
| lucy Flat | alg24 | 2.0027ms | 0.001102 | 146 KB | 145 KB | 1.0067 |
| lucy Flat | coarse24 | 2.0005ms | 0.015410 | 146 KB | 145 KB | 1.0096 |
| lucy Flat | oct24 | 2.9983ms | 0.000355 | 146 KB | 145 KB | 1.0064 |
| lucy Smooth | Baseline | N/A | N/A | 585 KB | 539 KB | 1.0868 |
| lucy Smooth | alg24 | 2.0001ms | 0.001102 | 146 KB | 145 KB | 1.0033 |
| lucy Smooth | coarse24 | 1.9862ms | 0.008282 | 146 KB | 145 KB | 1.0063 |
| lucy Smooth | oct24 | 2.9983ms | 0.000362 | 146 KB | 146 KB | 1.0029 |
| max-planck Flat | Baseline | N/A | N/A | 586 KB | 519 KB | 1.1289 |
| max-planck Flat | alg24 | 2.0013ms | 0.000718 | 146 KB | 145 KB | 1.0075 |
| max-planck Flat | coarse24 | 1.9999ms | 0.014836 | 146 KB | 145 KB | 1.0084 |
| max-planck Flat | oct24 | 2.998ms | 0.000358 | 146 KB | 146 KB | 1.0045 |
| max-planck Smooth | Baseline | N/A | N/A | 586 KB | 540 KB | 1.0866 |
| max-planck Smooth | alg24 | 2.9988ms | 0.000703 | 146 KB | 145 KB | 1.0062 |
| max-planck Smooth | coarse24 | 1.0001ms | 0.013936 | 146 KB | 145 KB | 1.0072 |
| max-planck Smooth | oct24 | 2.9988ms | 0.000360 | 146 KB | 146 KB | 1.0032 |
| nefertiti Flat | Baseline | N/A | N/A | 585 KB | 424 KB | 1.3788 |
| nefertiti Flat | alg24 | 1.9995ms | 0.000848 | 146 KB | 138 KB | 1.0563 |
| nefertiti Flat | coarse24 | 1.9998ms | 0.018882 | 146 KB | 135 KB | 1.0809 |
| nefertiti Flat | oct24 | 2.9993ms | 0.000353 | 146 KB | 139 KB | 1.0470 |
| nefertiti Smooth | Baseline | N/A | N/A | 585 KB | 534 KB | 1.0964 |
| nefertiti Smooth | alg24 | 1.9999ms | 0.000806 | 146 KB | 140 KB | 1.0417 |
| nefertiti Smooth | coarse24 | 2.0364ms | 0.016187 | 146 KB | 135 KB | 1.0774 |
| nefertiti Smooth | oct24 | 3ms | 0.000371 | 146 KB | 141 KB | 1.0339 |
| ogre Flat | Baseline | N/A | N/A | 728 KB | 521 KB | 1.3967 |
| ogre Flat | alg24 | 3.003ms | 0.000824 | 182 KB | 177 KB | 1.0240 |
| ogre Flat | coarse24 | 2.024ms | 0.028269 | 182 KB | 175 KB | 1.0371 |
| ogre Flat | oct24 | 4ms | 0.000334 | 182 KB | 178 KB | 1.0225 |
| ogre Smooth | Baseline | N/A | N/A | 728 KB | 663 KB | 1.0981 |
| ogre Smooth | alg24 | 2.001ms | 0.000838 | 182 KB | 180 KB | 1.0078 |
| ogre Smooth | coarse24 | 1.9961ms | 0.010169 | 182 KB | 179 KB | 1.0171 |
| ogre Smooth | oct24 | 3.9642ms | 0.000355 | 182 KB | 181 KB | 1.0064 |
| rocker-arm Flat | Baseline | N/A | N/A | 117 KB | 92 KB | 1.2751 |
| rocker-arm Flat | alg24 | 0s | 0.001190 | 29 KB | 29 KB | 1.0086 |
| rocker-arm Flat | coarse24 | 0s | 0.044882 | 29 KB | 28 KB | 1.0239 |
| rocker-arm Flat | oct24 | 1.0004ms | 0.000331 | 29 KB | 29 KB | <div style="color:red">0.9995</div> |
| rocker-arm Smooth | Baseline | N/A | N/A | 117 KB | 108 KB | 1.0832 |
| rocker-arm Smooth | alg24 | 0s | 0.000970 | 29 KB | 29 KB | 1.0037 |
| rocker-arm Smooth | coarse24 | 999.3µs | 0.014634 | 29 KB | 29 KB | 1.0109 |
| rocker-arm Smooth | oct24 | 1.0058ms | 0.000349 | 29 KB | 29 KB | <div style="color:red">0.9999</div> |
| spot Flat | Baseline | N/A | N/A | 34 KB | 20 KB | 1.6768 |
| spot Flat | alg24 | 0s | 0.000941 | 8 KB | 8 KB | 1.0150 |
| spot Flat | coarse24 | 0s | 0.022933 | 8 KB | 8 KB | 1.0286 |
| spot Flat | oct24 | 1.0009ms | 0.000344 | 8 KB | 8 KB | 1.0123 |
| spot Smooth | Baseline | N/A | N/A | 34 KB | 31 KB | 1.0866 |
| spot Smooth | alg24 | 0s | 0.000808 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| spot Smooth | coarse24 | 0s | 0.009306 | 8 KB | 8 KB | 1.0031 |
| spot Smooth | oct24 | 0s | 0.000358 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| stanford-bunny Flat | Baseline | N/A | N/A | 421 KB | 317 KB | 1.3285 |
| stanford-bunny Flat | alg24 | 1.032ms | 0.011212 | 105 KB | 100 KB | 1.0511 |
| stanford-bunny Flat | coarse24 | 997.9µs | 0.046346 | 105 KB | 97 KB | 1.0847 |
| stanford-bunny Flat | oct24 | 1.0007ms | 0.013426 | 105 KB | 101 KB | 1.0410 |
| stanford-bunny Smooth | Baseline | N/A | N/A | 421 KB | 376 KB | 1.1186 |
| stanford-bunny Smooth | alg24 | 996.7µs | 0.000712 | 105 KB | 102 KB | 1.0228 |
| stanford-bunny Smooth | coarse24 | 999.8µs | 0.008954 | 105 KB | 100 KB | 1.0526 |
| stanford-bunny Smooth | oct24 | 1.9991ms | 0.000372 | 105 KB | 102 KB | 1.0241 |
| teapot Flat | Baseline | N/A | N/A | 42 KB | 30 KB | 1.4060 |
| teapot Flat | alg24 | 0s | 0.001040 | 10 KB | 10 KB | 1.0319 |
| teapot Flat | coarse24 | 0s | 0.042057 | 10 KB | 10 KB | 1.0609 |
| teapot Flat | oct24 | 0s | 0.000339 | 10 KB | 10 KB | 1.0135 |
| teapot Smooth | Baseline | N/A | N/A | 42 KB | 39 KB | 1.0864 |
| teapot Smooth | alg24 | 0s | 0.000972 | 10 KB | 10 KB | 1.0016 |
| teapot Smooth | coarse24 | 0s | 0.006317 | 10 KB | 10 KB | 1.0161 |
| teapot Smooth | oct24 | 980.1µs | 0.000365 | 10 KB | 10 KB | <div style="color:red">0.9995</div> |
| woody Flat | Baseline | N/A | N/A | 8 KB | 0 KB | 33.1793 |
| woody Flat | alg24 | 0s | 0.000000 | 2 KB | 0 KB | 86.7500 |
| woody Flat | coarse24 | 0s | 0.000424 | 2 KB | 0 KB | 14.7660 |
| woody Flat | oct24 | 0s | 0.000000 | 2 KB | 0 KB | 86.7500 |
| woody Smooth | Baseline | N/A | N/A | 8 KB | 0 KB | 17.5696 |
| woody Smooth | alg24 | 0s | 0.000307 | 2 KB | 0 KB | 6.0000 |
| woody Smooth | coarse24 | 0s | 0.004974 | 2 KB | 0 KB | 6.3865 |
| woody Smooth | oct24 | 0s | 0.000326 | 2 KB | 0 KB | 6.0523 |
| xyzrgb_dragon Flat | Baseline | N/A | N/A | 1465 KB | 1250 KB | 1.1722 |
| xyzrgb_dragon Flat | alg24 | 5.0008ms | 0.000775 | 366 KB | 364 KB | 1.0049 |
| xyzrgb_dragon Flat | coarse24 | 3.0004ms | 0.018748 | 366 KB | 363 KB | 1.0076 |
| xyzrgb_dragon Flat | oct24 | 7.0002ms | 0.000367 | 366 KB | 364 KB | 1.0045 |
| xyzrgb_dragon Smooth | Baseline | N/A | N/A | 1465 KB | 1343 KB | 1.0906 |
| xyzrgb_dragon Smooth | alg24 | 5.0082ms | 0.000733 | 366 KB | 364 KB | 1.0042 |
| xyzrgb_dragon Smooth | coarse24 | 3.9702ms | 0.008780 | 366 KB | 363 KB | 1.0091 |
| xyzrgb_dragon Smooth | oct24 | 7ms | 0.000378 | 366 KB | 364 KB | 1.0048 |

## Resources

* [A Survey of Efficient Representations for Independent Unit Vectors](http://jcgt.org/published/0003/02/01/paper.pdf)
* [Alec Jacobson's Common 3D Test Models](https://github.com/alecjacobson/common-3d-test-models)