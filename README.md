# Sublinear Time Complexity Blockchain Verification

Research code for simulating sublinear-complexity blockchain verification (SCV), aimed at large-scale enterprise and distributed systems.

This was my thesis project at the University of Tasmania: *Efficient Blockchain Interoperability: a simulation of an IoT use case*.

Repo: [github.com/lee-junmin/thesis-blockchain](https://github.com/lee-junmin/thesis-blockchain)

## Overview

With normal Simplified Payment Verification (SPV), you walk the chain one block at a time, so verification gets slower as the chain grows. Here, blocks are organised into a tree-like structure using per-block levels. A light client can then verify a block by jumping along higher-level links instead of checking every predecessor.

SCV follows `LevelPrevHash` pointers back toward genesis in roughly logarithmic steps. The benchmarks compare SCV against SPV on verification time and step count, then the same idea is carried into sidechain transfer and IoT network simulations.

## Approach

1. Each block gets a level derived from its hash, plus a `LevelPrevHash` pointing to the previous higher-level block.
2. SPV does a sequential header walk (`O(n)`).
3. SCV skips levels (`O(log n)` expected).
4. Results go out as CSV and get plotted in Python (matplotlib, pandas, numpy).

## Tech stack

- Blockchain and simulations: Go
- Analysis and plots: Python (matplotlib, pandas, numpy, Jupyter)

## Repository layout

```
blockchain/              Core block/chain types (headers, levels, hashing)
sublinearverification/   SCV vs SPV timing and step benchmarks
sidechaintransfer/       Cross-chain / network transfer simulation
iotsim/                  IoT use-case simulation
visualisations/          CSV outputs, notebooks, plot scripts
main.go                  Entry point for running experiments
```

`part1-2/` and `part3/` are snapshots from different stages of the thesis experiments.

## Running

```bash
git clone https://github.com/lee-junmin/thesis-blockchain.git
cd thesis-blockchain
go run .
```

In `main.go`, switch between the experiment entry points (`sublintest`, network tests, `IoTnetworktest`). Exported CSVs land under `visualisations/` and can be plotted with the notebooks or `plot.py`.

## What I measured

- SCV verification time and step count vs chain size (compared to SPV)
- Sidechain transfer behaviour under different network conditions
- An IoT interoperability scenario using the same verification model


