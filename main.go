package main

import (
	"github.com/lee-junmin/thesis-blockchain/iotsim"
	"github.com/lee-junmin/thesis-blockchain/sidechaintransfer"
	"github.com/lee-junmin/thesis-blockchain/sublinearverification"
)

func main() {
	// Switch which experiment to run:
	// sublintest(100, 10000, 100)
	// normalityNetworkTest()
	IoTnetworktest()
}

// Part 1: SCV vs SPV timing and step-count benchmarks
func sublintest(step int, max int, rep int) {
	sublinearverification.ExportCSV("./visualisations/part1/scv-time.csv", sublinearverification.TestScvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/part1/scv-step.csv", sublinearverification.TestScvStep(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/part1/spv-time.csv", sublinearverification.TestSpvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/part1/spv-step.csv", sublinearverification.TestSpvStep(step, max, rep))
}

// Part 2: sidechain transfer under network failure conditions
func normalityNetworkTest() {
	sidechaintransfer.StartSimulationClock()
	sidechaintransfer.NetworkSimulation()
	sidechaintransfer.ExportCSV("./visualisations/part2/stat-test-15-5000.csv", sidechaintransfer.NormalityNetworkFailureTest(15, 5000, 200, 500))
	sidechaintransfer.ExportCSV("./visualisations/part2/stat-test-30-5000.csv", sidechaintransfer.NormalityNetworkFailureTest(30, 5000, 200, 500))
}

// Part 3: IoT interoperability simulation
func IoTnetworktest() {
	iotsim.ReadIot()
	iotsim.StartSimulationClock()
	iotsim.NetworkSimulation()
	iotsim.ExportCSV("./visualisations/part3/stat-test-15-5000.csv", iotsim.NormalityNetworkFailureTest(15, 5000, 180, 500))
}
