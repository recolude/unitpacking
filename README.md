# Unit Packing

A library for storing unit vectors in a representnation that lends itself to saving space on disk. You can read more on how some of these methods work [here](https://elicdavis.medium.com/a-story-about-information-entropy-and-efficiently-storing-unit-vectors-92b4a68efe67).

## Benchmark

10,000,000 unit vectors where randomly generated and ran through each algorithm available in thie library. The basiline method is simply writing out each vector component as a 32bit float. If the ability to compress the data/speed is the upmost importance to you, then you should choose the `Coarse24` method. If precision is the upmost importance to you, then you should pick the `Oct24` method.

| Dataset | Method | Runtime | Average Error | Uncompressed | Compressed | Compression Ratio |
|-|-|-|-|-|-|-|
| 10 Million Random | Baseline | N/A         | N/A           | 117,188 KB   | 100,217 KB | 1.1693            |
| 10 million random | alg24 | 911.9958ms | 0.0008 | 30000000 | 28991477 | 1.0348
| 10 million random | coarse24 | 731.9991ms | 0.0039 | 30000000 | 26228913 | 1.1438
| 10 million random | oct24 | 1.0780012s | 0.0005 | 30000000 | 28991849 | 1.0348

## Resources

* [A Survey of Efficient Representations for Independent Unit Vectors](http://jcgt.org/published/0003/02/01/paper.pdf)
* [Alec Jacobson's Common 3D Test Models](https://github.com/alecjacobson/common-3d-test-models)