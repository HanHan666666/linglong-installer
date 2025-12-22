# regexp

Package regexp implements regular expression search.

Documentation: [godoc.org/modernc.org/regexp](http://godoc.org/modernc.org/regexp)

This is a DFA experiment. See also [issue 11646](https://github.com/golang/go/issues/11646).
However, this package does not port the C++ re2 DFA, it's an alternative for
evaluating a different option.

The DFA currently supports a rather limited subset of regexps. It will probably
never support compiling all regexps to a DFA, but it may cover more of them in
the future.


All numbers bellow were produced using Go 1.20.3 on linux/amd64, AMD Ryzen 9
3900X 12-Core Processor.

Attempt to compile a DFA is performed only when no one pass engine was
compiled.

**The DFA is compiled for non one-pass regexp in**

* 41,672 out of 80,581 cases for package tests
* 67 out of 164,692 cases for package benchmarks
* 15 out of 16 cases for the benchmarks game

**The DFA is used to match non one-pass regexp in**

* 48.0% of package tests
* 63.7% of package benchmarks
* 91.4% of the benchmarks game regexredux-dna Go program #1

**Results summary**

* package benchmarks performance sec/op geomean: -24.71%
* package benchmarks throughput B/s geomean: +50.24%
* benchmarks game regexredux-dna wall time: -61.92%

# Compatibility

This package is a drop-in replacement of the stdlib regexp package as in Go
1.20.2. All original tests are passing before tagging a release.

# Benchstat results

2023-04-20

    goos: linux
    goarch: amd64
    pkg: regexp
    cpu: AMD Ryzen 9 3900X 12-Core Processor            
                                    │ log-benchstat-stdlib │        log-benchstat-modernc         │
                                    │        sec/op        │    sec/op     vs base                │
    Find-24                                    209.6n ± 2%    226.6n ± 1%   +8.11% (p=0.000 n=20)
    FindAllNoMatches-24                        97.22n ± 4%   105.15n ± 1%   +8.16% (p=0.000 n=20)
    FindString-24                              200.6n ± 2%    206.7n ± 1%   +2.99% (p=0.001 n=20)
    FindSubmatch-24                            461.8n ± 2%    467.8n ± 3%        ~ (p=0.155 n=20)
    FindStringSubmatch-24                      443.5n ± 2%    466.1n ± 2%   +5.11% (p=0.000 n=20)
    Literal-24                                 66.81n ± 1%    59.08n ± 1%  -11.58% (p=0.000 n=20)
    NotLiteral-24                              974.4n ± 3%    919.2n ± 1%   -5.66% (p=0.000 n=20)
    MatchClass-24                              1.304µ ± 2%    1.083µ ± 1%  -16.95% (p=0.000 n=20)
    MatchClass_InRange-24                      1.262µ ± 1%    1.075µ ± 2%  -14.82% (p=0.000 n=20)
    ReplaceAll-24                              1.665µ ± 3%    1.669µ ± 3%        ~ (p=0.888 n=20)
    AnchoredLiteralShortNonMatch-24            49.72n ± 2%    48.39n ± 1%        ~ (p=0.155 n=20)
    AnchoredLiteralLongNonMatch-24             70.30n ± 1%    49.10n ± 8%  -30.15% (p=0.000 n=20)
    AnchoredShortMatch-24                      78.37n ± 1%    70.37n ± 3%  -10.21% (p=0.000 n=20)
    AnchoredLongMatch-24                      210.30n ± 1%    71.24n ± 2%  -66.12% (p=0.000 n=20)
    OnePassShortA-24                           310.9n ± 2%    320.0n ± 2%   +2.91% (p=0.000 n=20)
    NotOnePassShortA-24                        337.6n ± 1%    345.3n ± 1%   +2.30% (p=0.001 n=20)
    OnePassShortB-24                           257.3n ± 1%    259.4n ± 1%        ~ (p=0.180 n=20)
    NotOnePassShortB-24                        251.3n ± 2%    218.6n ± 2%  -13.01% (p=0.000 n=20)
    OnePassLongPrefix-24                       73.58n ± 2%    69.95n ± 2%   -4.93% (p=0.000 n=20)
    OnePassLongNotPrefix-24                    197.2n ± 1%    196.3n ± 1%        ~ (p=0.365 n=20)
    MatchParallelShared-24                     15.14n ± 3%    15.94n ± 3%   +5.25% (p=0.000 n=20)
    MatchParallelCopied-24                     15.27n ± 2%    18.54n ± 4%  +21.38% (p=0.000 n=20)
    QuoteMetaAll-24                            131.8n ± 1%    139.2n ± 1%   +5.69% (p=0.000 n=20)
    QuoteMetaNone-24                           28.03n ± 1%    27.95n ± 2%        ~ (p=0.713 n=20)
    Compile/Onepass-24                         6.661µ ± 1%    6.731µ ± 1%   +1.05% (p=0.012 n=20)
    Compile/Medium-24                          15.03µ ± 1%    15.53µ ± 1%   +3.29% (p=0.000 n=20)
    Compile/Hard-24                            111.4µ ± 1%    114.2µ ± 1%   +2.51% (p=0.000 n=20)
    Match/Easy0/16-24                          3.279n ± 1%    3.276n ± 2%        ~ (p=0.703 n=20)
    Match/Easy0/32-24                          47.80n ± 1%    53.52n ± 1%  +11.97% (p=0.000 n=20)
    Match/Easy0/1K-24                          209.4n ± 2%    187.1n ± 1%  -10.69% (p=0.000 n=20)
    Match/Easy0/32K-24                         3.216µ ± 1%    3.184µ ± 2%   -0.98% (p=0.015 n=20)
    Match/Easy0/1M-24                          201.4µ ± 0%    202.2µ ± 1%        ~ (p=0.565 n=20)
    Match/Easy0/32M-24                         6.830m ± 1%    6.754m ± 2%   -1.11% (p=0.003 n=20)
    Match/Easy0i/16-24                         3.293n ± 0%    3.260n ± 1%   -1.02% (p=0.006 n=20)
    Match/Easy0i/32-24                         634.2n ± 1%    438.6n ± 1%  -30.83% (p=0.000 n=20)
    Match/Easy0i/1K-24                         18.16µ ± 1%    12.25µ ± 2%  -32.55% (p=0.000 n=20)
    Match/Easy0i/32K-24                        948.2µ ± 1%    386.1µ ± 1%  -59.28% (p=0.000 n=20)
    Match/Easy0i/1M-24                         30.64m ± 1%    12.42m ± 1%  -59.46% (p=0.000 n=20)
    Match/Easy0i/32M-24                        981.6m ± 3%    397.7m ± 1%  -59.49% (p=0.000 n=20)
    Match/Easy1/16-24                          3.262n ± 1%    3.297n ± 1%   +1.07% (p=0.007 n=20)
    Match/Easy1/32-24                          46.53n ± 1%    47.34n ± 0%   +1.76% (p=0.001 n=20)
    Match/Easy1/1K-24                          601.8n ± 1%    592.6n ± 1%   -1.54% (p=0.002 n=20)
    Match/Easy1/32K-24                         24.33µ ± 1%    13.37µ ± 1%  -45.03% (p=0.000 n=20)
    Match/Easy1/1M-24                          837.0µ ± 1%    484.7µ ± 1%  -42.09% (p=0.000 n=20)
    Match/Easy1/32M-24                         26.77m ± 1%    15.56m ± 1%  -41.89% (p=0.000 n=20)
    Match/Medium/16-24                         3.284n ± 1%    3.285n ± 1%        ~ (p=0.899 n=20)
    Match/Medium/32-24                         660.1n ± 1%    476.8n ± 1%  -27.77% (p=0.000 n=20)
    Match/Medium/1K-24                         19.29µ ± 0%    16.76µ ± 1%  -13.10% (p=0.000 n=20)
    Match/Medium/32K-24                        940.9µ ± 1%    535.8µ ± 2%  -43.06% (p=0.000 n=20)
    Match/Medium/1M-24                         29.92m ± 1%    17.30m ± 2%  -42.17% (p=0.000 n=20)
    Match/Medium/32M-24                        949.0m ± 1%    559.3m ± 1%  -41.07% (p=0.000 n=20)
    Match/Hard/16-24                           3.286n ± 2%    3.243n ± 1%        ~ (p=0.090 n=20)
    Match/Hard/32-24                           944.5n ± 2%    938.4n ± 3%        ~ (p=0.140 n=20)
    Match/Hard/1K-24                           28.32µ ± 1%    28.37µ ± 2%        ~ (p=0.815 n=20)
    Match/Hard/32K-24                          1.303m ± 1%    1.317m ± 1%        ~ (p=0.052 n=20)
    Match/Hard/1M-24                           42.27m ± 1%    42.29m ± 1%        ~ (p=0.583 n=20)
    Match/Hard/32M-24                           1.332 ± 1%     1.341 ± 1%        ~ (p=0.056 n=20)
    Match/Hard1/16-24                         3070.0n ± 3%    546.1n ± 2%  -82.21% (p=0.000 n=20)
    Match/Hard1/32-24                          5.936µ ± 2%    1.039µ ± 2%  -82.50% (p=0.000 n=20)
    Match/Hard1/1K-24                         187.93µ ± 1%    30.98µ ± 1%  -83.52% (p=0.000 n=20)
    Match/Hard1/32K-24                        6071.0µ ± 1%    979.2µ ± 1%  -83.87% (p=0.000 n=20)
    Match/Hard1/1M-24                         194.59m ± 1%    31.39m ± 2%  -83.87% (p=0.000 n=20)
    Match/Hard1/32M-24                          6.060 ± 0%     1.010 ± 1%  -83.32% (p=0.000 n=20)
    Match_onepass_regex/16-24                  233.9n ± 1%    234.8n ± 1%        ~ (p=0.372 n=20)
    Match_onepass_regex/32-24                  402.5n ± 1%    403.8n ± 1%        ~ (p=0.245 n=20)
    Match_onepass_regex/1K-24                  10.99µ ± 1%    10.95µ ± 1%        ~ (p=0.499 n=20)
    Match_onepass_regex/32K-24                 347.2µ ± 2%    348.6µ ± 1%        ~ (p=0.565 n=20)
    Match_onepass_regex/1M-24                  11.07m ± 1%    11.09m ± 2%        ~ (p=0.883 n=20)
    Match_onepass_regex/32M-24                 352.1m ± 2%    358.8m ± 1%   +1.88% (p=0.003 n=20)
    geomean                                    7.707µ         5.802µ       -24.71%
    
                                    │ log-benchstat-stdlib │          log-benchstat-modernc           │
                                    │         B/op         │     B/op       vs base                   │
    Find-24                                 0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    FindAllNoMatches-24                     0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    FindString-24                           0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    FindSubmatch-24                         48.00 ±   0%       48.00 ±  0%         ~ (p=1.000 n=20) ¹
    FindStringSubmatch-24                   32.00 ±   0%       32.00 ±  0%         ~ (p=1.000 n=20) ¹
    Literal-24                              0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    NotLiteral-24                           0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    MatchClass-24                           0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    MatchClass_InRange-24                   0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    ReplaceAll-24                           96.00 ±   1%       96.00 ±  0%     0.00% (p=0.008 n=20)
    AnchoredLiteralShortNonMatch-24         0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    AnchoredLiteralLongNonMatch-24          0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    AnchoredShortMatch-24                   0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    AnchoredLongMatch-24                    0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    OnePassShortA-24                        0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    NotOnePassShortA-24                     0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    OnePassShortB-24                        0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    NotOnePassShortB-24                     0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    OnePassLongPrefix-24                    0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    OnePassLongNotPrefix-24                 0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    MatchParallelShared-24                  0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    MatchParallelCopied-24                  0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    QuoteMetaAll-24                         64.00 ±   0%       64.00 ±  0%         ~ (p=1.000 n=20) ¹
    QuoteMetaNone-24                        0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Compile/Onepass-24                    3.961Ki ±   0%     3.977Ki ±  0%    +0.39% (p=0.000 n=20)
    Compile/Medium-24                     9.203Ki ±   0%     9.562Ki ±  0%    +3.90% (p=0.000 n=20)
    Compile/Hard-24                       82.77Ki ±   0%     83.13Ki ±  0%    +0.43% (p=0.000 n=20)
    Match/Easy0/16-24                       0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy0/32-24                       0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy0/1K-24                       0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy0/32K-24                      0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy0/1M-24                       1.000 ±    ?       0.000 ±  0%  -100.00% (p=0.000 n=20)
    Match/Easy0/32M-24                      18.00 ± 172%       18.00 ±  0%     0.00% (p=0.002 n=20)
    Match/Easy0i/16-24                      0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy0i/32-24                      0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy0i/1K-24                      0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy0i/32K-24                     2.000 ± 200%       1.000 ±  0%   -50.00% (p=0.000 n=20)
    Match/Easy0i/1M-24                     228.00 ±  61%       33.00 ±  3%   -85.53% (p=0.000 n=20)
    Match/Easy0i/32M-24                   4.188Ki ±  28%     1.033Ki ±  3%   -75.33% (p=0.000 n=20)
    Match/Easy1/16-24                       0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy1/32-24                       0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy1/1K-24                       0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy1/32K-24                      0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Easy1/1M-24                       2.000 ± 200%       1.000 ±  0%   -50.00% (p=0.000 n=20)
    Match/Easy1/32M-24                     132.50 ±  52%       42.00 ±  2%   -68.30% (p=0.000 n=20)
    Match/Medium/16-24                      0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Medium/32-24                      0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Medium/1K-24                      0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Medium/32K-24                     6.000 ±  67%       1.000 ±  0%   -83.33% (p=0.000 n=20)
    Match/Medium/1M-24                     209.00 ±  61%       46.00 ±  2%   -77.99% (p=0.000 n=20)
    Match/Medium/32M-24                   4.188Ki ±  64%     1.551Ki ±  3%   -62.97% (p=0.003 n=20)
    Match/Hard/16-24                        0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Hard/32-24                        0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Hard/1K-24                        0.000 ±   0%       1.000 ±   ?         ? (p=0.022 n=20)
    Match/Hard/32K-24                       3.000 ± 200%       9.000 ± 67%         ~ (p=0.231 n=20)
    Match/Hard/1M-24                        316.0 ±  64%       310.0 ± 62%         ~ (p=0.982 n=20)
    Match/Hard/32M-24                     8.500Ki ±  65%     8.500Ki ± 65%         ~ (p=1.000 n=20)
    Match/Hard1/16-24                       0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Hard1/32-24                       0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match/Hard1/1K-24                       6.000 ±    ?       0.000 ±  0%  -100.00% (p=0.000 n=20)
    Match/Hard1/32K-24                     47.000 ±  66%       2.000 ±  0%   -95.74% (p=0.000 n=20)
    Match/Hard1/1M-24                      564.50 ± 175%       85.00 ±  2%   -84.94% (p=0.000 n=20)
    Match/Hard1/32M-24                    3.008Ki ±   0%     3.102Ki ±  3%         ~ (p=0.960 n=20)
    Match_onepass_regex/16-24               0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match_onepass_regex/32-24               0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match_onepass_regex/1K-24               0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20) ¹
    Match_onepass_regex/32K-24              0.000 ±   0%       0.000 ±  0%         ~ (p=1.000 n=20)
    Match_onepass_regex/1M-24               29.00 ±   3%       30.00 ±  0%    +3.45% (p=0.009 n=20)
    Match_onepass_regex/32M-24            1.033Ki ±   3%     1.033Ki ±  3%         ~ (p=1.000 n=20)
    geomean                                              ²                  ?                       ²
    ¹ all samples are equal
    ² summaries must be >0 to compute geomean
    
                                    │ log-benchstat-stdlib │         log-benchstat-modernc         │
                                    │      allocs/op       │  allocs/op   vs base                  │
    Find-24                                  0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    FindAllNoMatches-24                      0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    FindString-24                            0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    FindSubmatch-24                          1.000 ±  0%     1.000 ±  0%        ~ (p=1.000 n=20) ¹
    FindStringSubmatch-24                    1.000 ±  0%     1.000 ±  0%        ~ (p=1.000 n=20) ¹
    Literal-24                               0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    NotLiteral-24                            0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    MatchClass-24                            0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    MatchClass_InRange-24                    0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    ReplaceAll-24                            5.000 ±  0%     5.000 ±  0%        ~ (p=1.000 n=20) ¹
    AnchoredLiteralShortNonMatch-24          0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    AnchoredLiteralLongNonMatch-24           0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    AnchoredShortMatch-24                    0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    AnchoredLongMatch-24                     0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    OnePassShortA-24                         0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    NotOnePassShortA-24                      0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    OnePassShortB-24                         0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    NotOnePassShortB-24                      0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    OnePassLongPrefix-24                     0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    OnePassLongNotPrefix-24                  0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    MatchParallelShared-24                   0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    MatchParallelCopied-24                   0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    QuoteMetaAll-24                          2.000 ±  0%     2.000 ±  0%        ~ (p=1.000 n=20) ¹
    QuoteMetaNone-24                         0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Compile/Onepass-24                       52.00 ±  0%     52.00 ±  0%        ~ (p=1.000 n=20) ¹
    Compile/Medium-24                        112.0 ±  0%     116.0 ±  0%   +3.57% (p=0.000 n=20)
    Compile/Hard-24                          424.0 ±  0%     428.0 ±  0%   +0.94% (p=0.000 n=20)
    Match/Easy0/16-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0/32-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0/1K-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0/32K-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0/1M-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0/32M-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0i/16-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0i/32-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0i/1K-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0i/32K-24                      0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0i/1M-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy0i/32M-24                      7.000 ± 71%     1.000 ±   ?  -85.71% (p=0.000 n=20)
    Match/Easy1/16-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy1/32-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy1/1K-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy1/32K-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy1/1M-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Easy1/32M-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Medium/16-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Medium/32-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Medium/1K-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Medium/32K-24                      0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Medium/1M-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Medium/32M-24                      7.000 ± 86%     1.000 ±  0%  -85.71% (p=0.000 n=20)
    Match/Hard/16-24                         0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard/32-24                         0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard/1K-24                         0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard/32K-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard/1M-24                         0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard/32M-24                        19.00 ± 89%     19.00 ± 89%        ~ (p=1.000 n=20)
    Match/Hard1/16-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard1/32-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard1/1K-24                        0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard1/32K-24                       0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match/Hard1/1M-24                        0.000 ±   ?     0.000 ±  0%    0.00% (p=0.001 n=20)
    Match/Hard1/32M-24                       2.000 ±  0%     3.000 ± 33%        ~ (p=0.960 n=20)
    Match_onepass_regex/16-24                0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match_onepass_regex/32-24                0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match_onepass_regex/1K-24                0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match_onepass_regex/32K-24               0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match_onepass_regex/1M-24                0.000 ±  0%     0.000 ±  0%        ~ (p=1.000 n=20) ¹
    Match_onepass_regex/32M-24               1.000 ±   ?     1.000 ±   ?        ~ (p=1.000 n=20)
    geomean                                              ²                 -4.87%                ²
    ¹ all samples are equal
    ² summaries must be >0 to compute geomean
    
                               │ log-benchstat-stdlib │         log-benchstat-modernc          │
                               │         B/s          │      B/s       vs base                 │
    QuoteMetaAll-24                     101.33Mi ± 1%    95.89Mi ± 1%    -5.36% (p=0.000 n=20)
    QuoteMetaNone-24                     884.6Mi ± 1%    887.3Mi ± 2%         ~ (p=0.718 n=20)
    Match/Easy0/16-24                    4.544Gi ± 1%    4.548Gi ± 2%         ~ (p=0.703 n=20)
    Match/Easy0/32-24                    638.5Mi ± 1%    570.2Mi ± 1%   -10.69% (p=0.000 n=20)
    Match/Easy0/1K-24                    4.553Gi ± 2%    5.099Gi ± 1%   +12.00% (p=0.000 n=20)
    Match/Easy0/32K-24                   9.491Gi ± 1%    9.585Gi ± 2%    +0.98% (p=0.015 n=20)
    Match/Easy0/1M-24                    4.849Gi ± 0%    4.829Gi ± 1%         ~ (p=0.565 n=20)
    Match/Easy0/32M-24                   4.575Gi ± 1%    4.627Gi ± 2%    +1.12% (p=0.003 n=20)
    Match/Easy0i/16-24                   4.526Gi ± 0%    4.572Gi ± 1%    +1.02% (p=0.005 n=20)
    Match/Easy0i/32-24                   48.12Mi ± 1%    69.58Mi ± 1%   +44.58% (p=0.000 n=20)
    Match/Easy0i/1K-24                   53.78Mi ± 1%    79.73Mi ± 2%   +48.24% (p=0.000 n=20)
    Match/Easy0i/32K-24                  32.96Mi ± 1%    80.94Mi ± 1%  +145.57% (p=0.000 n=20)
    Match/Easy0i/1M-24                   32.64Mi ± 1%    80.51Mi ± 1%  +146.68% (p=0.000 n=20)
    Match/Easy0i/32M-24                  32.60Mi ± 2%    80.47Mi ± 1%  +146.83% (p=0.000 n=20)
    Match/Easy1/16-24                    4.568Gi ± 1%    4.520Gi ± 1%    -1.06% (p=0.007 n=20)
    Match/Easy1/32-24                    656.0Mi ± 1%    644.6Mi ± 0%    -1.74% (p=0.000 n=20)
    Match/Easy1/1K-24                    1.585Gi ± 1%    1.609Gi ± 1%    +1.56% (p=0.002 n=20)
    Match/Easy1/32K-24                   1.255Gi ± 1%    2.282Gi ± 1%   +81.91% (p=0.000 n=20)
    Match/Easy1/1M-24                    1.167Gi ± 1%    2.015Gi ± 1%   +72.68% (p=0.000 n=20)
    Match/Easy1/32M-24                   1.167Gi ± 1%    2.009Gi ± 1%   +72.09% (p=0.000 n=20)
    Match/Medium/16-24                   4.537Gi ± 1%    4.536Gi ± 1%         ~ (p=0.883 n=20)
    Match/Medium/32-24                   46.23Mi ± 1%    64.00Mi ± 1%   +38.44% (p=0.000 n=20)
    Match/Medium/1K-24                   50.64Mi ± 0%    58.27Mi ± 1%   +15.07% (p=0.000 n=20)
    Match/Medium/32K-24                  33.22Mi ± 1%    58.33Mi ± 2%   +75.61% (p=0.000 n=20)
    Match/Medium/1M-24                   33.42Mi ± 1%    57.79Mi ± 2%   +72.95% (p=0.000 n=20)
    Match/Medium/32M-24                  33.72Mi ± 1%    57.22Mi ± 1%   +69.67% (p=0.000 n=20)
    Match/Hard/16-24                     4.534Gi ± 2%    4.595Gi ± 1%         ~ (p=0.091 n=20)
    Match/Hard/32-24                     32.31Mi ± 2%    32.52Mi ± 3%         ~ (p=0.144 n=20)
    Match/Hard/1K-24                     34.48Mi ± 1%    34.42Mi ± 2%         ~ (p=0.774 n=20)
    Match/Hard/32K-24                    23.98Mi ± 1%    23.73Mi ± 1%    -1.05% (p=0.048 n=20)
    Match/Hard/1M-24                     23.66Mi ± 1%    23.64Mi ± 1%         ~ (p=0.569 n=20)
    Match/Hard/32M-24                    24.03Mi ± 1%    23.86Mi ± 1%         ~ (p=0.062 n=20)
    Match/Hard1/16-24                    4.969Mi ± 3%   27.938Mi ± 2%  +462.28% (p=0.000 n=20)
    Match/Hard1/32-24                    5.140Mi ± 2%   29.383Mi ± 2%  +471.61% (p=0.000 n=20)
    Match/Hard1/1K-24                    5.198Mi ± 1%   31.528Mi ± 1%  +506.61% (p=0.000 n=20)
    Match/Hard1/32K-24                   5.145Mi ± 1%   31.915Mi ± 1%  +520.30% (p=0.000 n=20)
    Match/Hard1/1M-24                    5.136Mi ± 1%   31.853Mi ± 2%  +520.24% (p=0.000 n=20)
    Match/Hard1/32M-24                   5.283Mi ± 1%   31.672Mi ± 1%  +499.46% (p=0.000 n=20)
    Match_onepass_regex/16-24            65.22Mi ± 1%    65.00Mi ± 1%         ~ (p=0.383 n=20)
    Match_onepass_regex/32-24            75.84Mi ± 1%    75.57Mi ± 1%         ~ (p=0.239 n=20)
    Match_onepass_regex/1K-24            88.87Mi ± 1%    89.20Mi ± 1%         ~ (p=0.490 n=20)
    Match_onepass_regex/32K-24           90.00Mi ± 2%    89.65Mi ± 1%         ~ (p=0.570 n=20)
    Match_onepass_regex/1M-24            90.31Mi ± 1%    90.17Mi ± 2%         ~ (p=0.899 n=20)
    Match_onepass_regex/32M-24           90.88Mi ± 2%    89.20Mi ± 1%    -1.85% (p=0.003 n=20)
    geomean                              146.3Mi         219.8Mi        +50.24%

# Benchmarks game regexredux-dna results

See [https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/regexredux-go-1.html](https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/regexredux-go-1.html)

2023-04-20

    goos: linux
    goarch: amd64
    pkg: modernc.org/regexp
    cpu: AMD Ryzen 9 3900X 12-Core Processor            
    BenchmarkGameStdlib
    BenchmarkGameStdlib-24     	       1	18544006464 ns/op	1166947000 B/op	   21919 allocs/op
    BenchmarkGameModernc
    BenchmarkGameModernc-24    	       1	7061945499 ns/op	1168672840 B/op	   67698 allocs/op
