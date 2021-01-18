# Unit Packing

A library for storing unit vectors in a representnation that lends itself to saving space on disk. You can read more on how some of these methods work [here](https://elicdavis.medium.com/a-story-about-information-entropy-and-efficiently-storing-unit-vectors-92b4a68efe67).

## Benchmark

10,000,000 unit vectors where randomly generated and ran through each algorithm available in thie library. If the ability to compress the data/speed is the upmost importance to you, then you should choose the `Coarse24` method. If precision is the upmost importance to you, then you should pick the `Oct24` method.

| Method   | Runtime     | Average Error | File Size After Compression |
|----------|-------------|---------------|-----------------------------|
| Coarse24 | 33.1952358s | 0.003905      | 23,594 KB                   |
| Alg24    | 33.5222204s | 0.000829      | 28,083 KB                   |
| Oct24    | 33.2052488s | 0.000465      | 28,098 KB                   |

## Resources

* [A Survey of Efficient Representations for Independent Unit Vectors](http://jcgt.org/published/0003/02/01/paper.pdf)