![Go](https://img.shields.io/badge/go-1.20+-blue?logo=go) ![License](https://img.shields.io/badge/license-Apache-blue)

## Description

The repository hosts the implementation of SeerEVM, an advanced transaction execution engine for EVM-compatible blockchains.

SeerEVM incorporates fine-grained branch prediction to fully exploit pre-execution effectiveness. Seer predicts state-related branches using a two-level prediction approach, reducing inconsistent execution paths more efficiently than executing all possible branches. To enable effective reuse of pre-execution results, Seer employs checkpoint-based fast-path execution, enhancing transaction execution for both successful and unsuccessful predictions.

## Precondition

**Testbeds:** 

- Amazon EC2 c7i.8xlagre instance (32 vCPUs, 64 GB memory), running Ubuntu 22.04 LTS

**Necessary software environment:**

- Installation of go (go version >= 1.20)

Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:

```shell script
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
```

Add /usr/local/go/bin to the PATH environment variable.

```shell script
export PATH=$PATH:/usr/local/go/bin
```

- Install Go-related modules/packages

```shell script
# Under the current directory `SeerEVM`
$ go mod tidy
```

**Dataset:**

Due to the size of the dataset used in our paper exceeding 100 GB, we have trimmed the dataset to facilitate artifact evaluation. We provide a tiny-scale of Ethereum datasets (approximately 370 MB), spanning 1000 blocks (including state data):

- Block height ranging from 14,650,000 ~ 14,651,000, shown in `./ethereumdata_1465_small`

## Usage

Here, we show how to run SeerEVM to produce several critical experimental results. All the test scripts are presented in the directory `./experiments`. Note that due to the different idle state of the machine's resources and the tiny scale of Ethereum datasets—which results in minimal disk I/O operations—it is normal for the reproducible results to have deviations with those reported in the paper.

#### 1.1 Prediction accuracy

Enter the directory `./experiments` and run the script `prediction_height.sh`:

```shell script
$ cd experiments
$ ./prediction_height.sh
```

This script run would output the prediction accuracy across 1,000 blocks shown in `prediction_height.txt`. This experimental result presents the prediction accuracy of complex contract transactions (C-TXs) and normal contract transactions (N-TXs) at every 200-block offset.

#### 1.2 Pre-execution latency

In the current directory, run the script `pre_execution_lat_large.sh`:

```shell script
$ ./pre_execution_lat_large.sh
```

This script run would output the pre-execution latency under different number of transactions shown in `preExecution_large_ratio.txt`. For each transaction scale, three different disorder rates (0, 0.4, 0.8) are used. Under each disorder rate, the pre-execution latency of both C-TXs and N-TXs is presented, along with the total latency of pre-executing all transactions.

#### 1.3 Speedup over single transaction execution

In the current directory, run the script `speedup_tx.sh`:

```shell script
$ ./speedup_tx.sh
```

This script run would output the speedup distribution across smart contract transactions shown in `speedup_perTx_full.txt`. This experimental result presents the average speedups on C-TXs and N-TXs, as well as the average speedup for all transactions. 

#### 1.4 Concurrent execution performance

In the current directory, run the script `con_execution_abort.sh`:

```shell script
$ ./con_execution_abort.sh
```

This script run would output the abort rate under varying number of concurrent transactions shown in `concurrent_abort.txt`. This experimental result presents the average abort rates for C-TXs and N-TXs.

Then, run the script `con_execution_speedup.sh`:

```shell script
$ ./con_execution_speedup.sh
```

This script run would output the speedup over serial execution by using different threads shown in `concurrent_speedup.txt`. This experimental result presents the latency and throughput of SeerEVM's concurrent execution, alongside those of the serial execution for comparison. Note that, compared to the full-node storage size simulated in the paper, the dataset used here is relatively small. As a result, the reduction in disk I/O overhead achieved by SeerEVM's pre-execution over serial execution is less significant, leading to a lower speedup than reported in the paper.

## Citation

If you are interested in our work and intend to use SeerEVM in your research, please consider citing our paper. Below is the BibTex format for citation:

```bibtex
@article{zhang2024seer,
  title={Seer: Accelerating Blockchain Transaction Execution by Fine-Grained Branch Prediction},
  author={Zhang, Shijie and Cheng, Ru and Liu, Xinpeng and Xiao, Jiang and Jin, Hai and Li, Bo},
  journal={Proceedings of the VLDB Endowment},
  volume={18},
  number={3},
  pages={822--835},
  year={2024},
  publisher={VLDB Endowment}
}
```

