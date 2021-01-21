# Unit Packing
[![Build Status](https://travis-ci.com/recolude/unitpacking.svg?branch=main)](https://travis-ci.com/recolude/unitpacking) [![Go Report Card](https://goreportcard.com/badge/github.com/recolude/unitpacking)](https://goreportcard.com/report/github.com/recolude/unitpacking) [![Coverage](https://codecov.io/gh/recolude/unitpacking/branch/main/graph/badge.svg)](https://codecov.io/gh/recolude/unitpacking)

A library for storing unit vectors in a representnation that lends itself to saving space on disk. You can read more on how some of these methods work [here](https://elicdavis.medium.com/a-story-about-information-entropy-and-efficiently-storing-unit-vectors-92b4a68efe67).

## Benchmark

10,000,000 unit vectors where randomly generated and ran through each algorithm available in thie library. The basiline method is simply writing out each vector component as a 32bit float. If the ability to compress the data/speed is the upmost importance to you, then you should choose the `Coarse24` method. If precision is the upmost importance to you, then you should pick the `Oct24` method.

| Dataset | Method | Runtime | Average Error | Uncompressed | Compressed | Compression Ratio |
|-|-|-|-|-|-|-|
| 10 million random | Baseline | N/A | N/A | 120000000 b | 107057857 b | 1.1209 |
| 10 million random | alg24 | 395.0005ms | 0.000829 | 30000000 b | 28991477 b | 1.0348 |   
| 10 million random | coarse24 | 338.9668ms | 0.003905 | 30000000 b | 26228913 b | 1.1438 |
| 10 million random | oct24 | 526.0306ms | 0.000465 | 30000000 b | 28991849 b | 1.0348 |   
| alligator Flat | Baseline | N/A | N/A | 38496 b | 970 b | 39.6866 |
| alligator Flat | alg24 | 0s | 0.000000 | 9624 b | 34 b | 283.0588 |      
| alligator Flat | coarse24 | 0s | 0.000379 | 9624 b | 464 b | 20.7414 |   
| alligator Flat | oct24 | 1.0008ms | 0.000000 | 9624 b | 34 b | 283.0588 |
| alligator Smooth | Baseline | N/A | N/A | 38496 b | 101 b | 381.1485 |   
| alligator Smooth | alg24 | 0s | 0.000000 | 9624 b | 34 b | 283.0588 |    
| alligator Smooth | coarse24 | 0s | 0.000000 | 9624 b | 33 b | 291.6364 | 
| alligator Smooth | oct24 | 0s | 0.000000 | 9624 b | 34 b | 283.0588 |    
| armadillo Flat | Baseline | N/A | N/A | 599880 b | 447175 b | 1.3415 |
| armadillo Flat | alg24 | 1.9997ms | NaN | 149970 b | 149775 b | 1.0013 |       
| armadillo Flat | coarse24 | 1.993ms | 0.014417 | 149970 b | 149705 b | 1.0018 |
| armadillo Flat | oct24 | 2.9997ms | 0.000364 | 149970 b | 149751 b | 1.0015 |  
| armadillo Smooth | Baseline | N/A | N/A | 599880 b | 553715 b | 1.0834 |
| armadillo Smooth | alg24 | 1.9998ms | NaN | 149970 b | 150025 b | 0.9996 |        
| armadillo Smooth | coarse24 | 1.9993ms | 0.015919 | 149970 b | 150025 b | 0.9996 |
| armadillo Smooth | oct24 | 2.9691ms | 0.000364 | 149970 b | 150025 b | 0.9996 |   
| beetle-alt Flat | Baseline | N/A | N/A | 238644 b | 193730 b | 1.2318 |
| beetle-alt Flat | alg24 | 1.0028ms | NaN | 59661 b | 58349 b | 1.0225 |        
| beetle-alt Flat | coarse24 | 1.0006ms | 0.021670 | 59661 b | 56317 b | 1.0594 |
| beetle-alt Flat | oct24 | 999.9µs | 0.000343 | 59661 b | 59120 b | 1.0092 |    
| beetle-alt Smooth | Baseline | N/A | N/A | 238644 b | 213711 b | 1.1167 |
| beetle-alt Smooth | alg24 | 1.0012ms | NaN | 59661 b | 58763 b | 1.0153 |        
| beetle-alt Smooth | coarse24 | 1.0005ms | 0.019755 | 59661 b | 56792 b | 1.0505 |
| beetle-alt Smooth | oct24 | 997.4µs | 0.000340 | 59661 b | 59251 b | 1.0069 |    
| beetle Flat | Baseline | N/A | N/A | 13776 b | 11661 b | 1.1814 |
| beetle Flat | alg24 | 0s | NaN | 3444 b | 3426 b | 1.0053 |
| beetle Flat | coarse24 | 0s | 0.024597 | 3444 b | 3377 b | 1.0198 |
| beetle Flat | oct24 | 0s | 0.000335 | 3444 b | 3454 b | 0.9971 |
| beetle Smooth | Baseline | N/A | N/A | 13776 b | 12588 b | 1.0944 |
| beetle Smooth | alg24 | 0s | NaN | 3444 b | 3454 b | 0.9971 |
| beetle Smooth | coarse24 | 0s | 0.017801 | 3444 b | 3453 b | 0.9974 |
| beetle Smooth | oct24 | 0s | 0.000341 | 3444 b | 3454 b | 0.9971 |
| cheburashka Flat | Baseline | N/A | N/A | 80028 b | 44807 b | 1.7861 |
| cheburashka Flat | alg24 | 0s | NaN | 20007 b | 14956 b | 1.3377 |
| cheburashka Flat | coarse24 | 0s | 0.026707 | 20007 b | 14868 b | 1.3456 |
| cheburashka Flat | oct24 | 0s | 0.000348 | 20007 b | 14986 b | 1.3350 |
| cheburashka Smooth | Baseline | N/A | N/A | 80028 b | 74126 b | 1.0796 |
| cheburashka Smooth | alg24 | 0s | NaN | 20007 b | 20022 b | 0.9993 |
| cheburashka Smooth | coarse24 | 0s | 0.024624 | 20007 b | 20018 b | 0.9995 |
| cheburashka Smooth | oct24 | 0s | 0.000346 | 20007 b | 20022 b | 0.9993 |
| cow Flat | Baseline | N/A | N/A | 34836 b | 24207 b | 1.4391 |
| cow Flat | alg24 | 0s | NaN | 8709 b | 8624 b | 1.0099 |
| cow Flat | coarse24 | 0s | 0.019476 | 8709 b | 8563 b | 1.0171 |
| cow Flat | oct24 | 1.003ms | 0.000339 | 8709 b | 8648 b | 1.0071 |
| cow Smooth | Baseline | N/A | N/A | 34836 b | 27056 b | 1.2876 |
| cow Smooth | alg24 | 0s | NaN | 8709 b | 8719 b | 0.9989 |
| cow Smooth | coarse24 | 0s | 0.024486 | 8709 b | 8719 b | 0.9989 |
| cow Smooth | oct24 | 0s | 0.000343 | 8709 b | 8719 b | 0.9989 |
| fandisk Flat | Baseline | N/A | N/A | 77700 b | 44067 b | 1.7632 |
| fandisk Flat | alg24 | 0s | NaN | 19425 b | 7891 b | 2.4617 |
| fandisk Flat | coarse24 | 1.0341ms | 0.041221 | 19425 b | 6852 b | 2.8349 |
| fandisk Flat | oct24 | 0s | 0.000169 | 19425 b | 9628 b | 2.0176 |
| fandisk Smooth | Baseline | N/A | N/A | 77700 b | 48784 b | 1.5927 |
| fandisk Smooth | alg24 | 998.4µs | NaN | 19425 b | 8413 b | 2.3089 |
| fandisk Smooth | coarse24 | 0s | 0.038454 | 19425 b | 7710 b | 2.5195 |
| fandisk Smooth | oct24 | 0s | 0.000195 | 19425 b | 10021 b | 1.9384 |
| happy Flat | Baseline | N/A | N/A | 591012 b | 508256 b | 1.1628 |
| happy Flat | alg24 | 2.001ms | NaN | 147753 b | 147331 b | 1.0029 |
| happy Flat | coarse24 | 1.9996ms | 0.023530 | 147753 b | 146473 b | 1.0087 |
| happy Flat | oct24 | 2.999ms | 0.000346 | 147753 b | 147326 b | 1.0029 |
| happy Smooth | Baseline | N/A | N/A | 591012 b | 546903 b | 1.0807 |
| happy Smooth | alg24 | 1.9983ms | NaN | 147753 b | 147627 b | 1.0009 |
| happy Smooth | coarse24 | 1.9689ms | 0.022043 | 147753 b | 146975 b | 1.0053 |
| happy Smooth | oct24 | 1.9966ms | 0.000346 | 147753 b | 147568 b | 1.0013 |
| horse Flat | Baseline | N/A | N/A | 581820 b | 536341 b | 1.0848 |
| horse Flat | alg24 | 1.9974ms | NaN | 145455 b | 144356 b | 1.0076 |
| horse Flat | coarse24 | 2.0007ms | 0.019190 | 145455 b | 141899 b | 1.0251 |
| horse Flat | oct24 | 2.0012ms | 0.000350 | 145455 b | 144906 b | 1.0038 |
| horse Smooth | Baseline | N/A | N/A | 581820 b | 536886 b | 1.0837 |
| horse Smooth | alg24 | 2.0017ms | NaN | 145455 b | 144388 b | 1.0074 |
| horse Smooth | coarse24 | 999.8µs | 0.020305 | 145455 b | 141521 b | 1.0278 |
| horse Smooth | oct24 | 1.9713ms | 0.000350 | 145455 b | 144892 b | 1.0039 |
| igea Flat | Baseline | N/A | N/A | 1612140 b | 1337769 b | 1.2051 |
| igea Flat | alg24 | 6.0021ms | NaN | 403035 b | 388113 b | 1.0384 |
| igea Flat | coarse24 | 5.0023ms | 0.015230 | 403035 b | 378152 b | 1.0658 |
| igea Flat | oct24 | 7.0003ms | 0.000356 | 403035 b | 390750 b | 1.0314 |
| igea Smooth | Baseline | N/A | N/A | 1612140 b | 1479875 b | 1.0894 |
| igea Smooth | alg24 | 6.0002ms | NaN | 403035 b | 395358 b | 1.0194 |
| igea Smooth | coarse24 | 3.9988ms | 0.015598 | 403035 b | 383626 b | 1.0506 |
| igea Smooth | oct24 | 7.9677ms | 0.000356 | 403035 b | 397051 b | 1.0151 |
| lucy Flat | Baseline | N/A | N/A | 599844 b | 435526 b | 1.3773 |
| lucy Flat | alg24 | 1.9937ms | NaN | 149961 b | 148957 b | 1.0067 |
| lucy Flat | coarse24 | 2.0037ms | 0.015410 | 149961 b | 148533 b | 1.0096 |
| lucy Flat | oct24 | 2.9998ms | 0.000355 | 149961 b | 149006 b | 1.0064 |
| lucy Smooth | Baseline | N/A | N/A | 599844 b | 553004 b | 1.0847 |
| lucy Smooth | alg24 | 3.0093ms | NaN | 149961 b | 149733 b | 1.0015 |
| lucy Smooth | coarse24 | 2.0006ms | 0.015031 | 149961 b | 149567 b | 1.0026 |
| lucy Smooth | oct24 | 2.9994ms | 0.000358 | 149961 b | 149776 b | 1.0012 |
| max-planck Flat | Baseline | N/A | N/A | 600924 b | 532289 b | 1.1289 |
| max-planck Flat | alg24 | 3.0255ms | NaN | 150231 b | 149110 b | 1.0075 |
| max-planck Flat | coarse24 | 1.0023ms | 0.014836 | 150231 b | 148973 b | 1.0084 |
| max-planck Flat | oct24 | 2.0035ms | 0.000358 | 150231 b | 149556 b | 1.0045 |
| max-planck Smooth | Baseline | N/A | N/A | 600924 b | 553393 b | 1.0859 |
| max-planck Smooth | alg24 | 3.0584ms | NaN | 150231 b | 149380 b | 1.0057 |
| max-planck Smooth | coarse24 | 1.0005ms | 0.014796 | 150231 b | 149526 b | 1.0047 |
| max-planck Smooth | oct24 | 3.0352ms | 0.000358 | 150231 b | 149842 b | 1.0026 |
| nefertiti Flat | Baseline | N/A | N/A | 599652 b | 434903 b | 1.3788 |
| nefertiti Flat | alg24 | 2.0001ms | NaN | 149913 b | 141924 b | 1.0563 |
| nefertiti Flat | coarse24 | 1.9994ms | 0.018882 | 149913 b | 138697 b | 1.0809 |
| nefertiti Flat | oct24 | 2.9984ms | 0.000353 | 149913 b | 143185 b | 1.0470 |
| nefertiti Smooth | Baseline | N/A | N/A | 599652 b | 548356 b | 1.0935 |
| nefertiti Smooth | alg24 | 3.0004ms | NaN | 149913 b | 143606 b | 1.0439 |
| nefertiti Smooth | coarse24 | 1.9643ms | 0.017349 | 149913 b | 137927 b | 1.0869 |
| nefertiti Smooth | oct24 | 2.9997ms | 0.000353 | 149913 b | 145203 b | 1.0324 |
| ogre Flat | Baseline | N/A | N/A | 746328 b | 534361 b | 1.3967 |
| ogre Flat | alg24 | 2.9998ms | NaN | 186582 b | 182212 b | 1.0240 |
| ogre Flat | coarse24 | 2.0017ms | 0.028269 | 186582 b | 179903 b | 1.0371 |
| ogre Flat | oct24 | 4.0018ms | 0.000334 | 186582 b | 182476 b | 1.0225 |
| ogre Smooth | Baseline | N/A | N/A | 746328 b | 660407 b | 1.1301 |
| ogre Smooth | alg24 | 3.0473ms | NaN | 186582 b | 185172 b | 1.0076 |
| ogre Smooth | coarse24 | 2.0051ms | NaN | 186582 b | 182762 b | 1.0209 |
| ogre Smooth | oct24 | 2.9999ms | NaN | 186582 b | 185486 b | 1.0059 |
| rocker-arm Flat | Baseline | N/A | N/A | 120528 b | 94522 b | 1.2751 |
| rocker-arm Flat | alg24 | 0s | NaN | 30132 b | 29876 b | 1.0086 |
| rocker-arm Flat | coarse24 | 0s | 0.044882 | 30132 b | 29430 b | 1.0239 |
| rocker-arm Flat | oct24 | 1.0019ms | 0.000331 | 30132 b | 30147 b | 0.9995 |
| rocker-arm Smooth | Baseline | N/A | N/A | 120528 b | 111839 b | 1.0777 |
| rocker-arm Smooth | alg24 | 1.004ms | NaN | 30132 b | 29970 b | 1.0054 |
| rocker-arm Smooth | coarse24 | 1.0002ms | 0.039855 | 30132 b | 29693 b | 1.0148 |
| rocker-arm Smooth | oct24 | 0s | 0.000339 | 30132 b | 30147 b | 0.9995 |
| spot Flat | Baseline | N/A | N/A | 35160 b | 20969 b | 1.6768 |
| spot Flat | alg24 | 0s | NaN | 8790 b | 8660 b | 1.0150 |
| spot Flat | coarse24 | 0s | 0.022933 | 8790 b | 8546 b | 1.0286 |
| spot Flat | oct24 | 0s | 0.000344 | 8790 b | 8683 b | 1.0123 |
| spot Smooth | Baseline | N/A | N/A | 35160 b | 25319 b | 1.3887 |
| spot Smooth | alg24 | 968.6µs | NaN | 8790 b | 8800 b | 0.9989 |
| spot Smooth | coarse24 | 0s | 0.023825 | 8790 b | 8773 b | 1.0019 |
| spot Smooth | oct24 | 1.0312ms | 0.000343 | 8790 b | 8800 b | 0.9989 |
| stanford-bunny Flat | Baseline | N/A | N/A | 431364 b | 324699 b | 1.3285 |
| stanford-bunny Flat | alg24 | 999.6µs | NaN | 107841 b | 102775 b | 1.0493 |
| stanford-bunny Flat | coarse24 | 2.0015ms | 0.046346 | 107841 b | 99439 b | 1.0845 |
| stanford-bunny Flat | oct24 | 2.9905ms | 0.010661 | 107841 b | 103613 b | 1.0408 |
| stanford-bunny Smooth | Baseline | N/A | N/A | 431364 b | 387114 b | 1.1143 |
| stanford-bunny Smooth | alg24 | 1.0009ms | NaN | 107841 b | 104977 b | 1.0273 |
| stanford-bunny Smooth | coarse24 | 997.5µs | NaN | 107841 b | 102139 b | 1.0558 |
| stanford-bunny Smooth | oct24 | 2.0177ms | NaN | 107841 b | 105609 b | 1.0211 |
| teapot Flat | Baseline | N/A | N/A | 43728 b | 31101 b | 1.4060 |
| teapot Flat | alg24 | 0s | NaN | 10932 b | 10594 b | 1.0319 |
| teapot Flat | coarse24 | 0s | 0.042057 | 10932 b | 10304 b | 1.0609 |
| teapot Flat | oct24 | 1.0007ms | 0.000339 | 10932 b | 10786 b | 1.0135 |
| teapot Smooth | Baseline | N/A | N/A | 43728 b | 33378 b | 1.3101 |
| teapot Smooth | alg24 | 973.6µs | NaN | 10932 b | 10888 b | 1.0040 |
| teapot Smooth | coarse24 | 0s | 0.043316 | 10932 b | 10603 b | 1.0310 |
| teapot Smooth | oct24 | 1.0307ms | 0.000339 | 10932 b | 10942 b | 0.9991 |
| woody Flat | Baseline | N/A | N/A | 8328 b | 251 b | 33.1793 |
| woody Flat | alg24 | 0s | 0.000000 | 2082 b | 24 b | 86.7500 |
| woody Flat | coarse24 | 0s | 0.000424 | 2082 b | 141 b | 14.7660 |
| woody Flat | oct24 | 0s | 0.000000 | 2082 b | 24 b | 86.7500 |
| woody Smooth | Baseline | N/A | N/A | 8328 b | 43 b | 193.6744 |
| woody Smooth | alg24 | 0s | 0.000000 | 2082 b | 24 b | 86.7500 |
| woody Smooth | coarse24 | 963.5µs | 0.000000 | 2082 b | 24 b | 86.7500 |
| woody Smooth | oct24 | 0s | 0.000000 | 2082 b | 24 b | 86.7500 |
| xyzrgb_dragon Flat | Baseline | N/A | N/A | 1500792 b | 1280374 b | 1.1722 |
| xyzrgb_dragon Flat | alg24 | 5.0046ms | NaN | 375198 b | 373379 b | 1.0049 |
| xyzrgb_dragon Flat | coarse24 | 4.001ms | 0.018748 | 375198 b | 372373 b | 1.0076 |
| xyzrgb_dragon Flat | oct24 | 6.9997ms | 0.000367 | 375198 b | 373510 b | 1.0045 |
| xyzrgb_dragon Smooth | Baseline | N/A | N/A | 1500792 b | 1379672 b | 1.0878 |
| xyzrgb_dragon Smooth | alg24 | 4.9993ms | NaN | 375198 b | 373821 b | 1.0037 |
| xyzrgb_dragon Smooth | coarse24 | 4.0009ms | NaN | 375198 b | 372373 b | 1.0076 |
| xyzrgb_dragon Smooth | oct24 | 8ms | NaN | 375198 b | 373972 b | 1.0033 |

## Resources

* [A Survey of Efficient Representations for Independent Unit Vectors](http://jcgt.org/published/0003/02/01/paper.pdf)
* [Alec Jacobson's Common 3D Test Models](https://github.com/alecjacobson/common-3d-test-models)