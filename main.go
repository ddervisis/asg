package main

import (
	"bufio"
	"os"
	"os/user"
	"path/filepath"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(p string, initialList *[]string) {
	if _, err := os.Stat(p); err == nil {
		fileHandle, err := os.Open(p)
		check(err)

		fileScanner := bufio.NewScanner((fileHandle))
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			var currentText = fileScanner.Text()
			if !slices.Contains(*initialList, currentText) {
				*initialList = append(*initialList, currentText)
			}
		}
	}
}

func createDir(p string) {
	err := os.MkdirAll(p, os.ModeDir)
	check(err)
}

func main() {
	var initialCars []string
	initialCars = append(initialCars,
		"alpine_a110_gt4",
		"amr_v12_vantage_gt3",
		"amr_v8_vantage_gt3",
		"amr_v8_vantage_gt4",
		"audi_r8_gt4",
		"audi_r8_lms",
		"audi_r8_lms_evo",
		"audi_r8_lms_evo_ii",
		"audi_r8_lms_gt2",
		"bentley_continental_gt3_2016",
		"bentley_continental_gt3_2018",
		"bmw_m2_cs_racing",
		"bmw_m4_gt3",
		"bmw_m4_gt4",
		"bmw_m6_gt3",
		"chevrolet_camaro_gt4r",
		"ferrari_296_gt3",
		"ferrari_488_challenge_evo",
		"ferrari_488_gt3",
		"ferrari_488_gt3_evo",
		"ginetta_g55_gt4",
		"honda_nsx_gt3",
		"honda_nsx_gt3_evo",
		"jaguar_g3",
		"ktm_xbow_gt2",
		"ktm_xbow_gt4",
		"lamborghini_gallardo_rex",
		"lamborghini_huracan_gt3",
		"lamborghini_huracan_gt3_evo",
		"lamborghini_huracan_gt3_evo2",
		"lamborghini_huracan_st",
		"lamborghini_huracan_st_evo2",
		"lexus_rc_f_gt3",
		"maserati_mc20_gt2",
		"maserati_mc_gt4",
		"mclaren_570s_gt4",
		"mclaren_650s_gt3",
		"mclaren_720s_gt3",
		"mclaren_720s_gt3_evo",
		"mercedes_amg_gt2",
		"mercedes_amg_gt3",
		"mercedes_amg_gt3_evo",
		"mercedes_amg_gt4",
		"nissan_gt_r_gt3_2017",
		"nissan_gt_r_gt3_2018",
		"porsche_718_cayman_gt4_mr",
		"porsche_935",
		"porsche_991ii_gt3_cup",
		"porsche_991ii_gt3_r",
		"porsche_991_gt2_rs_mr",
		"porsche_991_gt3_r",
		"porsche_992_gt3_cup",
		"porsche_992_gt3_r",
	)

	var initialTracks []string
	initialTracks = append(initialTracks,
		"Barcelona",
		"brands_hatch",
		"cota",
		"donington",
		"Hungaroring",
		"Imola",
		"indianapolis",
		"Kyalami",
		"Laguna_Seca",
		"misano",
		"monza",
		"mount_panorama",
		"nurburgring",
		"oulton_park",
		"Paul_Ricard",
		"red_bull_ring",
		"Silverstone",
		"snetterton",
		"Spa",
		"Suzuka",
		"Valencia",
		"watkins_glen",
		"Zandvoort",
		"Zolder",
	)

	readFile(filepath.Join("cars.txt"), &initialCars)
	readFile(filepath.Join("tracks.txt"), &initialTracks)

	u, _ := user.Current()
	var basePath = filepath.Join(u.HomeDir, "Documents", "Assetto Corsa Competizione", "Setups")

	for _, car := range initialCars {
		createDir(filepath.Join(basePath, car))
		for _, track := range initialTracks {
			createDir(filepath.Join(basePath, car, track))
		}
	}

}
