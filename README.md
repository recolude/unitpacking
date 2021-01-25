# Unit Packing
[![Build Status](https://travis-ci.com/recolude/unitpacking.svg?branch=main)](https://travis-ci.com/recolude/unitpacking) [![Go Report Card](https://goreportcard.com/badge/github.com/recolude/unitpacking)](https://goreportcard.com/report/github.com/recolude/unitpacking) [![Coverage](https://codecov.io/gh/recolude/unitpacking/branch/main/graph/badge.svg)](https://codecov.io/gh/recolude/unitpacking)

A library for storing unit vectors in a representnation that lends itself to saving space on disk. You can read more on how some of these methods work [here](https://elicdavis.medium.com/a-story-about-information-entropy-and-efficiently-storing-unit-vectors-92b4a68efe67).

## Benchmark

Calculating the smooth and flat normals for a bunch of famous 3D model datasets. Also one dataset is just 10 million randomly generated unit vectors.

| Dataset | Method | Runtime | Average Error | Uncompressed | Compressed | Compression Ratio |
|-|-|-|-|-|-|-|
| 10 million random | Baseline | N/A | N/A | 117187 KB | 107717 KB | 1.0879 |
| 10 million random | alg24 | 380.9725ms | 0.000704 | 29296 KB | 29305 KB | <div style="color:red">0.9997</div> |
| 10 million random | coarse24 | 329.9998ms | 0.008057 | 29296 KB | 29274 KB | 1.0008 |
| 10 million random | oct24 | 599.9993ms | 0.000389 | 29296 KB | 29305 KB | <div style="color:red">0.9997</div> |
| armadillo flat | Baseline | N/A | N/A | 585 KB | 436 KB | 1.3415 |
| armadillo flat | alg24 | 1.9657ms | 0.000799 | 146 KB | 146 KB | 1.0013 |
| armadillo flat | coarse24 | 1.9986ms | 0.014417 | 146 KB | 146 KB | 1.0018 |
| armadillo flat | oct24 | 1.9995ms | 0.000364 | 146 KB | 146 KB | 1.0015 |
| armadillo smooth | Baseline | N/A | N/A | 585 KB | 540 KB | 1.0836 |
| armadillo smooth | alg24 | 2.0019ms | 0.000805 | 146 KB | 146 KB | <div style="color:red">0.9996</div> |
| armadillo smooth | coarse24 | 1.9995ms | 0.010859 | 146 KB | 146 KB | <div style="color:red">0.9998</div> |
| armadillo smooth | oct24 | 3.033ms | 0.000365 | 146 KB | 146 KB | <div style="color:red">0.9996</div> |
| beetle-alt flat | Baseline | N/A | N/A | 233 KB | 189 KB | 1.2318 |
| beetle-alt flat | alg24 | 1.0042ms | 0.001217 | 58 KB | 56 KB | 1.0225 |
| beetle-alt flat | coarse24 | 0s | 0.021670 | 58 KB | 54 KB | 1.0594 |
| beetle-alt flat | oct24 | 1.0001ms | 0.000343 | 58 KB | 57 KB | 1.0092 |
| beetle-alt smooth | Baseline | N/A | N/A | 233 KB | 212 KB | 1.0965 |
| beetle-alt smooth | alg24 | 0s | 0.000865 | 58 KB | 57 KB | 1.0138 |
| beetle-alt smooth | coarse24 | 999.7µs | 0.003912 | 58 KB | 56 KB | 1.0398 |
| beetle-alt smooth | oct24 | 2.0007ms | 0.000380 | 58 KB | 57 KB | 1.0128 |
| beetle flat | Baseline | N/A | N/A | 13 KB | 11 KB | 1.1814 |
| beetle flat | alg24 | 0s | 0.001283 | 3 KB | 3 KB | 1.0053 |
| beetle flat | coarse24 | 0s | 0.024597 | 3 KB | 3 KB | 1.0198 |
| beetle flat | oct24 | 0s | 0.000335 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| beetle smooth | Baseline | N/A | N/A | 13 KB | 12 KB | 1.0984 |
| beetle smooth | alg24 | 0s | 0.000851 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| beetle smooth | coarse24 | 0s | 0.004469 | 3 KB | 3 KB | 1.0088 |
| beetle smooth | oct24 | 0s | 0.000391 | 3 KB | 3 KB | <div style="color:red">0.9971</div> |
| cheburashka flat | Baseline | N/A | N/A | 78 KB | 43 KB | 1.7861 |
| cheburashka flat | alg24 | 0s | 0.000642 | 19 KB | 14 KB | 1.3377 |
| cheburashka flat | coarse24 | 0s | 0.026707 | 19 KB | 14 KB | 1.3456 |
| cheburashka flat | oct24 | 0s | 0.000348 | 19 KB | 14 KB | 1.3350 |
| cheburashka smooth | Baseline | N/A | N/A | 78 KB | 72 KB | 1.0831 |
| cheburashka smooth | alg24 | 0s | 0.000616 | 19 KB | 19 KB | <div style="color:red">0.9993</div> |
| cheburashka smooth | coarse24 | 0s | 0.014946 | 19 KB | 19 KB | 1.0025 |
| cheburashka smooth | oct24 | 0s | 0.000362 | 19 KB | 19 KB | <div style="color:red">0.9993</div> |
| cow flat | Baseline | N/A | N/A | 34 KB | 23 KB | 1.4391 |
| cow flat | alg24 | 0s | 0.000735 | 8 KB | 8 KB | 1.0099 |
| cow flat | coarse24 | 1ms | 0.019476 | 8 KB | 8 KB | 1.0171 |
| cow flat | oct24 | 0s | 0.000339 | 8 KB | 8 KB | 1.0071 |
| cow smooth | Baseline | N/A | N/A | 34 KB | 31 KB | 1.0812 |
| cow smooth | alg24 | 998.4µs | 0.000608 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| cow smooth | coarse24 | 0s | 0.013519 | 8 KB | 8 KB | 1.0020 |
| cow smooth | oct24 | 0s | 0.000364 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| fandisk flat | Baseline | N/A | N/A | 75 KB | 43 KB | 1.7632 |
| fandisk flat | alg24 | 0s | 0.001277 | 18 KB | 7 KB | 2.4617 |
| fandisk flat | coarse24 | 0s | 0.041221 | 18 KB | 6 KB | 2.8349 |
| fandisk flat | oct24 | 1.0213ms | 0.000169 | 18 KB | 9 KB | 2.0176 |
| fandisk smooth | Baseline | N/A | N/A | 75 KB | 48 KB | 1.5731 |
| fandisk smooth | alg24 | 0s | 0.001147 | 18 KB | 9 KB | 1.9277 |
| fandisk smooth | coarse24 | 0s | 0.004418 | 18 KB | 8 KB | 2.2553 |
| fandisk smooth | oct24 | 0s | 0.000322 | 18 KB | 10 KB | 1.8154 |
| happy flat | Baseline | N/A | N/A | 577 KB | 496 KB | 1.1628 |
| happy flat | alg24 | 1.0246ms | 0.000865 | 144 KB | 143 KB | 1.0029 |
| happy flat | coarse24 | 2.0312ms | 0.023530 | 144 KB | 143 KB | 1.0087 |
| happy flat | oct24 | 2.9992ms | 0.000346 | 144 KB | 143 KB | 1.0029 |
| happy smooth | Baseline | N/A | N/A | 577 KB | 531 KB | 1.0851 |
| happy smooth | alg24 | 2.9991ms | 0.000711 | 144 KB | 144 KB | 1.0015 |
| happy smooth | coarse24 | 3.0024ms | 0.010774 | 144 KB | 143 KB | 1.0070 |
| happy smooth | oct24 | 4.0332ms | 0.000359 | 144 KB | 143 KB | 1.0027 |
| horse flat | Baseline | N/A | N/A | 568 KB | 523 KB | 1.0848 |
| horse flat | alg24 | 1.9912ms | 0.001013 | 142 KB | 140 KB | 1.0076 |
| horse flat | coarse24 | 999.8µs | 0.019190 | 142 KB | 138 KB | 1.0251 |
| horse flat | oct24 | 2.9995ms | 0.000350 | 142 KB | 141 KB | 1.0038 |
| horse smooth | Baseline | N/A | N/A | 568 KB | 523 KB | 1.0863 |
| horse smooth | alg24 | 2.0007ms | 0.000968 | 142 KB | 141 KB | 1.0069 |
| horse smooth | coarse24 | 1.9802ms | 0.010709 | 142 KB | 138 KB | 1.0253 |
| horse smooth | oct24 | 2.999ms | 0.000360 | 142 KB | 141 KB | 1.0054 |
| igea flat | Baseline | N/A | N/A | 1574 KB | 1306 KB | 1.2051 |
| igea flat | alg24 | 4.9976ms | 0.000781 | 393 KB | 379 KB | 1.0384 |
| igea flat | coarse24 | 4.0006ms | 0.015230 | 393 KB | 369 KB | 1.0658 |
| igea flat | oct24 | 7ms | 0.000356 | 393 KB | 381 KB | 1.0314 |
| igea smooth | Baseline | N/A | N/A | 1574 KB | 1445 KB | 1.0892 |
| igea smooth | alg24 | 4.999ms | 0.000931 | 393 KB | 385 KB | 1.0202 |
| igea smooth | coarse24 | 4.0013ms | 0.011622 | 393 KB | 374 KB | 1.0504 |
| igea smooth | oct24 | 7.0047ms | 0.000350 | 393 KB | 387 KB | 1.0150 |
| lucy flat | Baseline | N/A | N/A | 585 KB | 425 KB | 1.3773 |
| lucy flat | alg24 | 2ms | 0.001102 | 146 KB | 145 KB | 1.0067 |
| lucy flat | coarse24 | 1.9984ms | 0.015410 | 146 KB | 145 KB | 1.0096 |
| lucy flat | oct24 | 2.9989ms | 0.000355 | 146 KB | 145 KB | 1.0064 |
| lucy smooth | Baseline | N/A | N/A | 585 KB | 539 KB | 1.0868 |
| lucy smooth | alg24 | 2.0009ms | 0.001102 | 146 KB | 145 KB | 1.0033 |
| lucy smooth | coarse24 | 1.9672ms | 0.008282 | 146 KB | 145 KB | 1.0063 |
| lucy smooth | oct24 | 3.0007ms | 0.000362 | 146 KB | 146 KB | 1.0029 |
| max-planck flat | Baseline | N/A | N/A | 586 KB | 519 KB | 1.1289 |
| max-planck flat | alg24 | 2.0051ms | 0.000718 | 146 KB | 145 KB | 1.0075 |
| max-planck flat | coarse24 | 1.0047ms | 0.014836 | 146 KB | 145 KB | 1.0084 |
| max-planck flat | oct24 | 1.9989ms | 0.000358 | 146 KB | 146 KB | 1.0045 |
| max-planck smooth | Baseline | N/A | N/A | 586 KB | 540 KB | 1.0866 |
| max-planck smooth | alg24 | 1.9998ms | 0.000703 | 146 KB | 145 KB | 1.0062 |
| max-planck smooth | coarse24 | 1.0002ms | 0.013936 | 146 KB | 145 KB | 1.0072 |
| max-planck smooth | oct24 | 2.9969ms | 0.000360 | 146 KB | 146 KB | 1.0032 |
| nefertiti flat | Baseline | N/A | N/A | 585 KB | 424 KB | 1.3788 |
| nefertiti flat | alg24 | 2.0002ms | 0.000848 | 146 KB | 138 KB | 1.0563 |
| nefertiti flat | coarse24 | 1.004ms | 0.018882 | 146 KB | 135 KB | 1.0809 |
| nefertiti flat | oct24 | 2.999ms | 0.000353 | 146 KB | 139 KB | 1.0470 |
| nefertiti smooth | Baseline | N/A | N/A | 585 KB | 534 KB | 1.0964 |
| nefertiti smooth | alg24 | 1.9983ms | 0.000806 | 146 KB | 140 KB | 1.0417 |
| nefertiti smooth | coarse24 | 999.2µs | 0.016187 | 146 KB | 135 KB | 1.0774 |
| nefertiti smooth | oct24 | 2.0129ms | 0.000371 | 146 KB | 141 KB | 1.0339 |
| ogre flat | Baseline | N/A | N/A | 728 KB | 521 KB | 1.3967 |
| ogre flat | alg24 | 1.9992ms | 0.000824 | 182 KB | 177 KB | 1.0240 |
| ogre flat | coarse24 | 1.9709ms | 0.028269 | 182 KB | 175 KB | 1.0371 |
| ogre flat | oct24 | 3.0004ms | 0.000334 | 182 KB | 178 KB | 1.0225 |
| ogre smooth | Baseline | N/A | N/A | 728 KB | 663 KB | 1.0981 |
| ogre smooth | alg24 | 1.9996ms | 0.000838 | 182 KB | 180 KB | 1.0078 |
| ogre smooth | coarse24 | 1.9993ms | 0.010169 | 182 KB | 179 KB | 1.0171 |
| ogre smooth | oct24 | 2.9999ms | 0.000355 | 182 KB | 181 KB | 1.0064 |
| rocker-arm flat | Baseline | N/A | N/A | 117 KB | 92 KB | 1.2751 |
| rocker-arm flat | alg24 | 999.4µs | 0.001190 | 29 KB | 29 KB | 1.0086 |
| rocker-arm flat | coarse24 | 0s | 0.044882 | 29 KB | 28 KB | 1.0239 |
| rocker-arm flat | oct24 | 1.0003ms | 0.000331 | 29 KB | 29 KB | <div style="color:red">0.9995</div> |
| rocker-arm smooth | Baseline | N/A | N/A | 117 KB | 108 KB | 1.0832 |
| rocker-arm smooth | alg24 | 1.0005ms | 0.000970 | 29 KB | 29 KB | 1.0037 |
| rocker-arm smooth | coarse24 | 0s | 0.014634 | 29 KB | 29 KB | 1.0109 |
| rocker-arm smooth | oct24 | 997.6µs | 0.000349 | 29 KB | 29 KB | <div style="color:red">0.9999</div> |
| spot flat | Baseline | N/A | N/A | 34 KB | 20 KB | 1.6768 |
| spot flat | alg24 | 0s | 0.000941 | 8 KB | 8 KB | 1.0150 |
| spot flat | coarse24 | 0s | 0.022933 | 8 KB | 8 KB | 1.0286 |
| spot flat | oct24 | 999.3µs | 0.000344 | 8 KB | 8 KB | 1.0123 |
| spot smooth | Baseline | N/A | N/A | 34 KB | 31 KB | 1.0866 |
| spot smooth | alg24 | 0s | 0.000808 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| spot smooth | coarse24 | 0s | 0.009306 | 8 KB | 8 KB | 1.0031 |
| spot smooth | oct24 | 0s | 0.000358 | 8 KB | 8 KB | <div style="color:red">0.9989</div> |
| stanford-bunny flat | Baseline | N/A | N/A | 421 KB | 317 KB | 1.3285 |
| stanford-bunny flat | alg24 | 999.1µs | 0.011212 | 105 KB | 100 KB | 1.0511 |
| stanford-bunny flat | coarse24 | 1.0004ms | 0.046346 | 105 KB | 97 KB | 1.0847 |
| stanford-bunny flat | oct24 | 2.0003ms | 0.013426 | 105 KB | 101 KB | 1.0410 |
| stanford-bunny smooth | Baseline | N/A | N/A | 421 KB | 376 KB | 1.1186 |
| stanford-bunny smooth | alg24 | 2.0103ms | 0.000712 | 105 KB | 102 KB | 1.0228 |
| stanford-bunny smooth | coarse24 | 1.0001ms | 0.008954 | 105 KB | 100 KB | 1.0526 |
| stanford-bunny smooth | oct24 | 2.0003ms | 0.000372 | 105 KB | 102 KB | 1.0241 |
| teapot flat | Baseline | N/A | N/A | 42 KB | 30 KB | 1.4060 |
| teapot flat | alg24 | 0s | 0.001040 | 10 KB | 10 KB | 1.0319 |
| teapot flat | coarse24 | 0s | 0.042057 | 10 KB | 10 KB | 1.0609 |
| teapot flat | oct24 | 0s | 0.000339 | 10 KB | 10 KB | 1.0135 |
| teapot smooth | Baseline | N/A | N/A | 42 KB | 39 KB | 1.0864 |
| teapot smooth | alg24 | 0s | 0.000972 | 10 KB | 10 KB | 1.0016 |
| teapot smooth | coarse24 | 0s | 0.006317 | 10 KB | 10 KB | 1.0161 |
| teapot smooth | oct24 | 0s | 0.000365 | 10 KB | 10 KB | <div style="color:red">0.9995</div> |
| xyzrgb_dragon flat | Baseline | N/A | N/A | 1465 KB | 1250 KB | 1.1722 |
| xyzrgb_dragon flat | alg24 | 6.0266ms | 0.000775 | 366 KB | 364 KB | 1.0049 |
| xyzrgb_dragon flat | coarse24 | 4.9876ms | 0.018748 | 366 KB | 363 KB | 1.0076 |
| xyzrgb_dragon flat | oct24 | 8.0027ms | 0.000367 | 366 KB | 364 KB | 1.0045 |
| xyzrgb_dragon smooth | Baseline | N/A | N/A | 1465 KB | 1343 KB | 1.0906 |
| xyzrgb_dragon smooth | alg24 | 4.0009ms | 0.000733 | 366 KB | 364 KB | 1.0042 |
| xyzrgb_dragon smooth | coarse24 | 4.0008ms | 0.008780 | 366 KB | 363 KB | 1.0091 |
| xyzrgb_dragon smooth | oct24 | 6.9902ms | 0.000378 | 366 KB | 364 KB | 1.0048 |


## Resources

* [A Survey of Efficient Representations for Independent Unit Vectors](http://jcgt.org/published/0003/02/01/paper.pdf)
* [Alec Jacobson's Common 3D Test Models](https://github.com/alecjacobson/common-3d-test-models)