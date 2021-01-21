# Unit Packing

A library for storing unit vectors in a representnation that lends itself to saving space on disk. You can read more on how some of these methods work [here](https://elicdavis.medium.com/a-story-about-information-entropy-and-efficiently-storing-unit-vectors-92b4a68efe67).

## Benchmark

10,000,000 unit vectors where randomly generated and ran through each algorithm available in thie library. The basiline method is simply writing out each vector component as a 32bit float. If the ability to compress the data/speed is the upmost importance to you, then you should choose the `Coarse24` method. If precision is the upmost importance to you, then you should pick the `Oct24` method.

| Dataset | Method | Runtime | Average Error | Uncompressed | Compressed | Compression Ratio |
|-|-|-|-|-|-|-|
| 10 million random | Baseline | N/A | N/A | 120000000 b | 107057857 b | 1.1209
| 10 million random | alg24 | 375.0357ms | 0.000829 | 30000000 b | 28991477 b | 1.0348
| 10 million random | coarse24 | 339.0008ms | 0.003905 | 30000000 b | 26228913 b | 1.1438
| 10 million random | oct24 | 526.9663ms | 0.000465 | 30000000 b | 28991849 b | 1.0348
| alligator.obj Flat | Baseline | N/A | N/A | 38496 b | 11590 b | 3.3215
| alligator.obj Flat | alg24 | 0s | 8.948801 | 9624 b | 34 b | 283.0588
| alligator.obj Flat | coarse24 | 0s | 8.948801 | 9624 b | 33 b | 291.6364
| alligator.obj Flat | oct24 | 0s | 8.948801 | 9624 b | 34 b | 283.0588
| alligator.obj Smooth | Baseline | N/A | N/A | 38496 b | 12747 b | 3.0200
| alligator.obj Smooth | alg24 | 0s | 53.164173 | 9624 b | 34 b | 283.0588
| alligator.obj Smooth | coarse24 | 0s | 53.164173 | 9624 b | 33 b | 291.6364
| alligator.obj Smooth | oct24 | 0s | 53.164173 | 9624 b | 34 b | 283.0588
| armadillo.obj Flat | Baseline | N/A | N/A | 599880 b | 449735 b | 1.3339
| armadillo.obj Flat | alg24 | 3.001ms | NaN | 149970 b | 149042 b | 1.0062
| armadillo.obj Flat | coarse24 | 1.0001ms | 0.024385 | 149970 b | 143304 b | 1.0465
| armadillo.obj Flat | oct24 | 2.9999ms | 0.211910 | 149970 b | 149751 b | 1.0015
| armadillo.obj Smooth | Baseline | N/A | N/A | 599880 b | 556259 b | 1.0784
| armadillo.obj Smooth | alg24 | 1.9989ms | NaN | 149970 b | 140681 b | 1.0660
| armadillo.obj Smooth | coarse24 | 2.0001ms | 1.395041 | 149970 b | 82628 b | 1.8150
| armadillo.obj Smooth | oct24 | 3.0001ms | 1.747437 | 149970 b | 150025 b | 0.9996
| beetle-alt.obj Flat | Baseline | N/A | N/A | 238644 b | 195325 b | 1.2218
| beetle-alt.obj Flat | alg24 | 1.0004ms | 0.333434 | 59661 b | 3638 b | 16.3994
| beetle-alt.obj Flat | coarse24 | 999.9µs | 0.002621 | 59661 b | 3250 b | 18.3572
| beetle-alt.obj Flat | oct24 | 1.0003ms | 0.477801 | 59661 b | 59120 b | 1.0092
| beetle-alt.obj Smooth | Baseline | N/A | N/A | 238644 b | 215595 b | 1.1069
| beetle-alt.obj Smooth | alg24 | 999.9µs | 0.333444 | 59661 b | 3610 b | 16.5266
| beetle-alt.obj Smooth | coarse24 | 0s | 0.002637 | 59661 b | 2922 b | 20.4179
| beetle-alt.obj Smooth | oct24 | 1.0024ms | 0.478242 | 59661 b | 59273 b | 1.0065
| beetle.obj Flat | Baseline | N/A | N/A | 13776 b | 11352 b | 1.2135
| beetle.obj Flat | alg24 | 0s | 0.333440 | 3444 b | 866 b | 3.9769
| beetle.obj Flat | coarse24 | 0s | 0.002775 | 3444 b | 475 b | 7.2505
| beetle.obj Flat | oct24 | 0s | 0.472511 | 3444 b | 3454 b | 0.9971
| beetle.obj Smooth | Baseline | N/A | N/A | 13776 b | 12266 b | 1.1231
| beetle.obj Smooth | alg24 | 0s | 0.333280 | 3444 b | 1421 b | 2.4236
| beetle.obj Smooth | coarse24 | 0s | 0.002977 | 3444 b | 481 b | 7.1601
| beetle.obj Smooth | oct24 | 0s | 0.475055 | 3444 b | 3454 b | 0.9971
| cheburashka.obj Flat | Baseline | N/A | N/A | 80028 b | 45063 b | 1.7759
| cheburashka.obj Flat | alg24 | 0s | 0.333537 | 20007 b | 3135 b | 6.3818
| cheburashka.obj Flat | coarse24 | 0s | 0.003753 | 20007 b | 2711 b | 7.3799
| cheburashka.obj Flat | oct24 | 0s | 0.482498 | 20007 b | 14986 b | 1.3350
| cheburashka.obj Smooth | Baseline | N/A | N/A | 80028 b | 74544 b | 1.0736
| cheburashka.obj Smooth | alg24 | 0s | 0.333344 | 20007 b | 5529 b | 3.6186
| cheburashka.obj Smooth | coarse24 | 0s | 0.003770 | 20007 b | 2869 b | 6.9735
| cheburashka.obj Smooth | oct24 | 0s | 0.483169 | 20007 b | 20022 b | 0.9993
| cow.obj Flat | Baseline | N/A | N/A | 34836 b | 24554 b | 1.4188
| cow.obj Flat | alg24 | 0s | 0.325739 | 8709 b | 6753 b | 1.2896
| cow.obj Flat | coarse24 | 0s | 0.003921 | 8709 b | 3975 b | 2.1909
| cow.obj Flat | oct24 | 0s | 0.461505 | 8709 b | 8648 b | 1.0071
| cow.obj Smooth | Baseline | N/A | N/A | 34836 b | 27294 b | 1.2763
| cow.obj Smooth | alg24 | 0s | NaN | 8709 b | 8155 b | 1.0679
| cow.obj Smooth | coarse24 | 0s | 0.005009 | 8709 b | 6636 b | 1.3124
| cow.obj Smooth | oct24 | 0s | 0.381809 | 8709 b | 8719 b | 0.9989
| fandisk.obj Flat | Baseline | N/A | N/A | 77700 b | 51823 b | 1.4993
| fandisk.obj Flat | alg24 | 1.0273ms | 0.331911 | 19425 b | 3812 b | 5.0958
| fandisk.obj Flat | coarse24 | 0s | 0.003145 | 19425 b | 2035 b | 9.5455
| fandisk.obj Flat | oct24 | 0s | 0.386198 | 19425 b | 9628 b | 2.0176
| fandisk.obj Smooth | Baseline | N/A | N/A | 77700 b | 57639 b | 1.3480
| fandisk.obj Smooth | alg24 | 0s | 0.323590 | 19425 b | 7474 b | 2.5990
| fandisk.obj Smooth | coarse24 | 0s | 0.003126 | 19425 b | 4750 b | 4.0895
| fandisk.obj Smooth | oct24 | 1.0008ms | 0.380434 | 19425 b | 10200 b | 1.9044
| happy.obj Flat | Baseline | N/A | N/A | 591012 b | 512429 b | 1.1534
| happy.obj Flat | alg24 | 2ms | 0.333579 | 147753 b | 23443 b | 6.3026
| happy.obj Flat | coarse24 | 2.0003ms | 0.003850 | 147753 b | 20592 b | 7.1753
| happy.obj Flat | oct24 | 2.9995ms | 0.485776 | 147753 b | 147326 b | 1.0029
| happy.obj Smooth | Baseline | N/A | N/A | 591012 b | 551020 b | 1.0726
| happy.obj Smooth | alg24 | 3.0287ms | 0.333577 | 147753 b | 23222 b | 6.3626
| happy.obj Smooth | coarse24 | 2.0067ms | 0.003804 | 147753 b | 20530 b | 7.1969
| happy.obj Smooth | oct24 | 2.9971ms | 0.484598 | 147753 b | 147513 b | 1.0016
| horse.obj Flat | Baseline | N/A | N/A | 581820 b | 538054 b | 1.0813
| horse.obj Flat | alg24 | 1.9997ms | 0.333578 | 145455 b | 8459 b | 17.1953
| horse.obj Flat | coarse24 | 1.9999ms | 0.003885 | 145455 b | 7709 b | 18.8682
| horse.obj Flat | oct24 | 3.9864ms | 0.482122 | 145455 b | 144906 b | 1.0038
| horse.obj Smooth | Baseline | N/A | N/A | 581820 b | 536719 b | 1.0840
| horse.obj Smooth | alg24 | 2.9987ms | 0.333576 | 145455 b | 7387 b | 19.6907
| horse.obj Smooth | coarse24 | 1.9981ms | 0.003901 | 145455 b | 6751 b | 21.5457
| horse.obj Smooth | oct24 | 3.0025ms | 0.481955 | 145455 b | 144926 b | 1.0037
| igea.obj Flat | Baseline | N/A | N/A | 1612140 b | 1321723 b | 1.2197
| igea.obj Flat | alg24 | 5.9993ms | 0.333576 | 403035 b | 14835 b | 27.1678
| igea.obj Flat | coarse24 | 5ms | 0.004039 | 403035 b | 13573 b | 29.6939
| igea.obj Flat | oct24 | 6.9993ms | 0.488138 | 403035 b | 390750 b | 1.0314
| igea.obj Smooth | Baseline | N/A | N/A | 1612140 b | 1480079 b | 1.0892
| igea.obj Smooth | alg24 | 4.9978ms | 0.333577 | 403035 b | 12498 b | 32.2480
| igea.obj Smooth | coarse24 | 4.9986ms | 0.004048 | 403035 b | 11502 b | 35.0404
| igea.obj Smooth | oct24 | 7.9994ms | 0.488471 | 403035 b | 397009 b | 1.0152
| lucy.obj Flat | Baseline | N/A | N/A | 599844 b | 438612 b | 1.3676
| lucy.obj Flat | alg24 | 1.9995ms | NaN | 149961 b | 133406 b | 1.1241
| lucy.obj Flat | coarse24 | 2.001ms | 23.460864 | 149961 b | 30894 b | 4.8540
| lucy.obj Flat | oct24 | 3.0081ms | 23.944388 | 149961 b | 149006 b | 1.0064
| lucy.obj Smooth | Baseline | N/A | N/A | 599844 b | 555756 b | 1.0793
| lucy.obj Smooth | alg24 | 2.003ms | NaN | 149961 b | 140225 b | 1.0694
| lucy.obj Smooth | coarse24 | 2.032ms | 146.305774 | 149961 b | 22802 b | 6.5767
| lucy.obj Smooth | oct24 | 2.9999ms | 146.802959 | 149961 b | 149762 b | 1.0013
| max-planck.obj Flat | Baseline | N/A | N/A | 600924 b | 537339 b | 1.1183
| max-planck.obj Flat | alg24 | 1.9985ms | NaN | 150231 b | 144494 b | 1.0397
| max-planck.obj Flat | coarse24 | 1.9993ms | 1.012277 | 150231 b | 95463 b | 1.5737
| max-planck.obj Flat | oct24 | 2.9987ms | 1.293638 | 150231 b | 149556 b | 1.0045
| max-planck.obj Smooth | Baseline | N/A | N/A | 600924 b | 557515 b | 1.0779
| max-planck.obj Smooth | alg24 | 3.0026ms | NaN | 150231 b | 135498 b | 1.1087
| max-planck.obj Smooth | coarse24 | 996.5µs | 9.599019 | 150231 b | 31956 b | 4.7012
| max-planck.obj Smooth | oct24 | 3.0001ms | 10.061798 | 150231 b | 149816 b | 1.0028
| nefertiti.obj Flat | Baseline | N/A | N/A | 599652 b | 441478 b | 1.3583
| nefertiti.obj Flat | alg24 | 1.9983ms | NaN | 149913 b | 135446 b | 1.1068
| nefertiti.obj Flat | coarse24 | 1.9777ms | 2.388715 | 149913 b | 69452 b | 2.1585
| nefertiti.obj Flat | oct24 | 2.001ms | 2.741791 | 149913 b | 143185 b | 1.0470
| nefertiti.obj Smooth | Baseline | N/A | N/A | 599652 b | 555645 b | 1.0792
| nefertiti.obj Smooth | alg24 | 2ms | NaN | 149913 b | 136830 b | 1.0956
| nefertiti.obj Smooth | coarse24 | 1.999ms | 19.522894 | 149913 b | 25110 b | 5.9703
| nefertiti.obj Smooth | oct24 | 4ms | 20.003873 | 149913 b | 144842 b | 1.0350
| ogre.obj Flat | Baseline | N/A | N/A | 746328 b | 497314 b | 1.5007
| ogre.obj Flat | alg24 | 3.0028ms | 0.329413 | 186582 b | 100155 b | 1.8629
| ogre.obj Flat | coarse24 | 2.032ms | 0.003807 | 186582 b | 49053 b | 3.8037
| ogre.obj Flat | oct24 | 3.0001ms | 0.467128 | 186582 b | 182476 b | 1.0225
| ogre.obj Smooth | Baseline | N/A | N/A | 746328 b | 645004 b | 1.1571
| ogre.obj Smooth | alg24 | 2.9951ms | NaN | 186582 b | 146612 b | 1.2726
| ogre.obj Smooth | coarse24 | 2.0295ms | 0.007687 | 186582 b | 91240 b | 2.0450
| ogre.obj Smooth | oct24 | 3.0015ms | 0.427751 | 186582 b | 185478 b | 1.0060
| rocker-arm.obj Flat | Baseline | N/A | N/A | 120528 b | 95423 b | 1.2631
| rocker-arm.obj Flat | alg24 | 0s | 0.333589 | 30132 b | 5843 b | 5.1569
| rocker-arm.obj Flat | coarse24 | 970.6µs | 0.004202 | 30132 b | 4688 b | 6.4275
| rocker-arm.obj Flat | oct24 | 1.0008ms | 0.468189 | 30132 b | 30147 b | 0.9995
| rocker-arm.obj Smooth | Baseline | N/A | N/A | 120528 b | 112935 b | 1.0672
| rocker-arm.obj Smooth | alg24 | 0s | 0.333500 | 30132 b | 8548 b | 3.5250
| rocker-arm.obj Smooth | coarse24 | 0s | 0.004184 | 30132 b | 4711 b | 6.3961
| rocker-arm.obj Smooth | oct24 | 0s | 0.467355 | 30132 b | 30147 b | 0.9995
| spot.obj Flat | Baseline | N/A | N/A | 35160 b | 21203 b | 1.6583
| spot.obj Flat | alg24 | 0s | 0.333286 | 8790 b | 3104 b | 2.8318
| spot.obj Flat | coarse24 | 0s | 0.003666 | 8790 b | 1064 b | 8.2613
| spot.obj Flat | oct24 | 0s | 0.479092 | 8790 b | 8683 b | 1.0123
| spot.obj Smooth | Baseline | N/A | N/A | 35160 b | 25476 b | 1.3801
| spot.obj Smooth | alg24 | 0s | 0.331851 | 8790 b | 5924 b | 1.4838
| spot.obj Smooth | coarse24 | 0s | 0.003695 | 8790 b | 2177 b | 4.0377
| spot.obj Smooth | oct24 | 0s | 0.475670 | 8790 b | 8800 b | 0.9989
| stanford-bunny.obj Flat | Baseline | N/A | N/A | 431364 b | 323022 b | 1.3354
| stanford-bunny.obj Flat | alg24 | 2.0001ms | 0.333551 | 107841 b | 12080 b | 8.9272
| stanford-bunny.obj Flat | coarse24 | 999.9µs | 0.003542 | 107841 b | 10775 b | 10.0084
| stanford-bunny.obj Flat | oct24 | 1.999ms | 0.477983 | 107841 b | 103613 b | 1.0408
| stanford-bunny.obj Smooth | Baseline | N/A | N/A | 431364 b | 388377 b | 1.1107
| stanford-bunny.obj Smooth | alg24 | 1ms | 0.333550 | 107841 b | 11395 b | 9.4639
| stanford-bunny.obj Smooth | coarse24 | 1.9623ms | 0.003543 | 107841 b | 10212 b | 10.5602
| stanford-bunny.obj Smooth | oct24 | 3.0002ms | 0.477955 | 107841 b | 105499 b | 1.0222
| teapot.obj Flat | Baseline | N/A | N/A | 43728 b | 28696 b | 1.5238
| teapot.obj Flat | alg24 | 0s | 0.331068 | 10932 b | 7175 b | 1.5236
| teapot.obj Flat | coarse24 | 0s | 0.003947 | 10932 b | 2944 b | 3.7133
| teapot.obj Flat | oct24 | 1.0002ms | 0.473524 | 10932 b | 10786 b | 1.0135
| teapot.obj Smooth | Baseline | N/A | N/A | 43728 b | 33816 b | 1.2931
| teapot.obj Smooth | alg24 | 0s | 0.319197 | 10932 b | 9440 b | 1.1581
| teapot.obj Smooth | coarse24 | 0s | 0.003911 | 10932 b | 6010 b | 1.8190
| teapot.obj Smooth | oct24 | 0s | 0.438176 | 10932 b | 10942 b | 0.9991
| woody.obj Flat | Baseline | N/A | N/A | 8328 b | 2557 b | 3.2569
| woody.obj Flat | alg24 | 0s | 35.678197 | 2082 b | 24 b | 86.7500
| woody.obj Flat | coarse24 | 0s | 35.678197 | 2082 b | 24 b | 86.7500
| woody.obj Flat | oct24 | 0s | 35.678197 | 2082 b | 24 b | 86.7500
| woody.obj Smooth | Baseline | N/A | N/A | 8328 b | 2845 b | 2.9272
| woody.obj Smooth | alg24 | 0s | 201.487992 | 2082 b | 24 b | 86.7500
| woody.obj Smooth | coarse24 | 0s | 201.487992 | 2082 b | 24 b | 86.7500
| woody.obj Smooth | oct24 | 0s | 201.487992 | 2082 b | 24 b | 86.7500
| xyzrgb_dragon.obj Flat | Baseline | N/A | N/A | 1500792 b | 1287715 b | 1.1655
| xyzrgb_dragon.obj Flat | alg24 | 6.0356ms | NaN | 375198 b | 359561 b | 1.0435
| xyzrgb_dragon.obj Flat | coarse24 | 4.999ms | 0.004337 | 375198 b | 316011 b | 1.1873
| xyzrgb_dragon.obj Flat | oct24 | 8.0014ms | 0.339865 | 375198 b | 373510 b | 1.0045
| xyzrgb_dragon.obj Smooth | Baseline | N/A | N/A | 1500792 b | 1385010 b | 1.0836
| xyzrgb_dragon.obj Smooth | alg24 | 4.0007ms | NaN | 375198 b | 359452 b | 1.0438
| xyzrgb_dragon.obj Smooth | coarse24 | 4.9668ms | 0.264061 | 375198 b | 296736 b | 1.2644
| xyzrgb_dragon.obj Smooth | oct24 | 6.9995ms | 0.458829 | 375198 b | 373978 b | 1.0033
## Resources

* [A Survey of Efficient Representations for Independent Unit Vectors](http://jcgt.org/published/0003/02/01/paper.pdf)
* [Alec Jacobson's Common 3D Test Models](https://github.com/alecjacobson/common-3d-test-models)