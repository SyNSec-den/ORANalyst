package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type color int

const (
	Uncolored color = iota
	Red
	Green
)

var (
	trackedDir    string
	baseDir       string
	sourceDir     string
	destDir       string
	outputDir     string
	lastRunDir    string
	cumulativeDir string
)

var (
	fileCoverInfoMap map[string]*[]coverInfo
	coverList        []float64
	lastRunList      []float64
)

type coverInfo struct {
	coverIdx     int
	startLineNum int // The line number of the start of the cover block in the original file
	endLineNum   int // The line number of the end of the cover block in the original file
}

func fileServer() {
	fs := http.FileServer(http.Dir(outputDir))
	http.Handle("/", fs)

	log.Println("Serving files on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupCover() {
	createDir(destDir)
	createDir(lastRunDir)
	createDir(cumulativeDir)
	fileCoverInfoMap = make(map[string]*[]coverInfo)
	coverList = make([]float64, 0)
	lastRunList = make([]float64, 0)
	blackLast, greenLast, redLast := 0, 0, 0
	blackCummulative, greenCummulative, redCummulative := 0, 0, 0

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			panic(fmt.Sprintf("error getting relative path %q: %v\n", path, err))
		}
		// if !strings.HasPrefix(relPath, "gopath") {
		// 	return nil
		// }
		destPath := filepath.Join(destDir, relPath)
		lastRunPath := filepath.Join(lastRunDir, relPath)
		cumulativePath := filepath.Join(cumulativeDir, relPath)

		if info.IsDir() {
			err = os.MkdirAll(destPath, info.Mode())
			if err != nil {
				panic(fmt.Sprintf("error creating directory %q: %v\n", destPath, err))
			}
			err = os.MkdirAll(lastRunPath, info.Mode())
			if err != nil {
				panic(fmt.Sprintf("error creating directory %q: %v\n", lastRunPath, err))
			}
			err = os.MkdirAll(cumulativePath, info.Mode())
			if err != nil {
				panic(fmt.Sprintf("error creating directory %q: %v\n", cumulativePath, err))
			}
			return nil
		} else {
			fmt.Println("Path:", relPath) // Print the relative path to the directory
			if !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}

			var srcData []byte
			instrumentData, err := ioutil.ReadFile(path) // Read the file content
			if err != nil {
				panic(fmt.Sprintf("error reading file %q: %v\n", path, err))
			}
			destFile, err := os.Create(destPath)
			if err != nil {
				panic(fmt.Sprintf("error creating file %q: %v\n", destPath, err))
			}
			originalFile := isInstrumented(string(instrumentData))
			var srcFile *os.File
			if originalFile == "" {
				srcFile, err = os.Open(path)
				if err != nil {
					return err
				}

				srcData, err = ioutil.ReadFile(path)
				if err != nil {
					panic(fmt.Sprintf("error reading file %q: %v\n", path, err))
				}
			} else {
				srcFile, err = os.Open(originalFile)
				if err != nil {
					return err
				}
				srcData, err = ioutil.ReadFile(originalFile)
				if err != nil {
					panic(fmt.Sprintf("error reading file %q: %v\n", path, err))
				}
			}
			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				panic(fmt.Sprintf("error copying file %q: %v\n", destPath, err))
			}

			instrumentLines := strings.Split(string(instrumentData), "\n")
			originalLines := strings.Split(string(srcData), "\n")
			coverInfoList, _ := constructCoverInfo(instrumentLines, originalLines)
			fileCoverInfoMap[relPath] = &coverInfoList

			emptyCover := make([]byte, CoverSize)

			// lastRun info
			colorInfo := assignOutputColor(srcData, coverInfoList, emptyCover)
			htmlString, b, g, r := convertToColoredHTML(srcData, colorInfo)
			if strings.HasSuffix(lastRunPath, ".go") {
				lastRunPath = strings.TrimSuffix(lastRunPath, ".go") + ".html"
			} else {
				panic(fmt.Sprintf("error file name %q: %v\n", lastRunPath, err))
			}
			err = ioutil.WriteFile(lastRunPath, []byte(htmlString), 0644)
			if err != nil {
				panic(fmt.Sprintf("error writing file %q: %v\n", lastRunPath, err))
			}
			if strings.HasPrefix(relPath, trackedDir) {
				blackLast += b
				greenLast += g
				redLast += r
			}

			// cumulative info
			colorInfo = assignOutputColor(srcData, coverInfoList, emptyCover)
			htmlString, b, g, r = convertToColoredHTML(srcData, colorInfo)
			if strings.HasSuffix(cumulativePath, ".go") {
				cumulativePath = strings.TrimSuffix(cumulativePath, ".go") + ".html"
			} else {
				panic(fmt.Sprintf("error file name %q: %v\n", cumulativePath, err))
			}
			err = ioutil.WriteFile(cumulativePath, []byte(htmlString), 0644)
			if err != nil {
				panic(fmt.Sprintf("error writing file %q: %v\n", cumulativePath, err))
			}
			if strings.HasPrefix(relPath, trackedDir) {
				blackCummulative += b
				greenCummulative += g
				redCummulative += r
			}

			destFile.Close()
			srcFile.Close()

		}
		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("error walking the path %q: %v\n", sourceDir, err))
	}

	htmlContent := generateSummary(blackLast, greenLast, redLast, blackCummulative, greenCummulative, redCummulative)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}
	filePath := filepath.Join(outputDir, "summary.html")
	if err := ioutil.WriteFile(filePath, []byte(htmlContent), 0644); err != nil {
		panic(fmt.Sprintf("Error writing to file: %v\n", err))
	}
}

func displayCover(coverTab []byte) {
	blackLast, greenLast, redLast := 0, 0, 0
	blackCummulative, greenCummulative, redCummulative := 0, 0, 0

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			panic(fmt.Sprintf("error getting relative path %q: %v\n", path, err))
		}
		// if !strings.HasPrefix(relPath, "gopath") {
		// 	return nil
		// }
		destPath := filepath.Join(destDir, relPath)
		lastRunPath := filepath.Join(lastRunDir, relPath)
		cumulativePath := filepath.Join(cumulativeDir, relPath)

		if !info.IsDir() {
			if !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}
			srcData, err := ioutil.ReadFile(destPath)
			if err != nil {
				panic(fmt.Sprintf("error reading file %q: %v\n", path, err))
			}

			coverInfoList := *fileCoverInfoMap[relPath]
			// lastRun info
			colorInfo := assignOutputColor(srcData, coverInfoList, coverTab)
			htmlString, b, g, r := convertToColoredHTML(srcData, colorInfo)
			if strings.HasSuffix(lastRunPath, ".go") {
				lastRunPath = strings.TrimSuffix(lastRunPath, ".go") + ".html"
			} else {
				panic(fmt.Sprintf("error file name %q: %v\n", lastRunPath, err))
			}
			err = ioutil.WriteFile(lastRunPath, []byte(htmlString), 0644)
			if err != nil {
				panic(fmt.Sprintf("error writing file %q: %v\n", lastRunPath, err))
			}
			if strings.HasPrefix(relPath, trackedDir) {
				blackLast += b
				greenLast += g
				redLast += r
			}

			// cumulative info
			colorInfo = assignOutputColor(srcData, coverInfoList, CumulativeCover)
			htmlString, b, g, r = convertToColoredHTML(srcData, colorInfo)
			if strings.HasSuffix(cumulativePath, ".go") {
				cumulativePath = strings.TrimSuffix(cumulativePath, ".go") + ".html"
			} else {
				panic(fmt.Sprintf("error file name %q: %v\n", cumulativePath, err))
			}
			err = ioutil.WriteFile(cumulativePath, []byte(htmlString), 0644)
			if err != nil {
				panic(fmt.Sprintf("error writing file %q: %v\n", cumulativePath, err))
			}
			if strings.HasPrefix(relPath, trackedDir) {
				blackCummulative += b
				greenCummulative += g
				redCummulative += r
			}

		}
		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("error walking the path %q: %v\n", sourceDir, err))
	}

	htmlContent := generateSummary(blackLast, greenLast, redLast, blackCummulative, greenCummulative, redCummulative)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}
	filePath := filepath.Join(outputDir, "summary.html")
	if err := ioutil.WriteFile(filePath, []byte(htmlContent), 0644); err != nil {
		panic(fmt.Sprintf("Error writing to file: %v\n", err))
	}

}

func generateSummary(blackLast, greenLast, redLast, blackCummulative, greenCummulative, redCummulative int) string {
	lastRunCovered := 100.00
	if greenLast+redLast != 0 {
		lastRunCovered = float64(greenLast) / float64(greenLast+redLast) * 100
	}
	cummulativeCovered := 100.00
	if greenCummulative+redCummulative != 0 {
		cummulativeCovered = float64(greenCummulative) / float64(greenCummulative+redCummulative) * 100
	}

	lastRunList = append(lastRunList, lastRunCovered)
	coverList = append(coverList, cummulativeCovered)

	formatList := func(list []float64) string {
		strArr := make([]string, len(list))
		for i, v := range list {
			strArr[i] = strconv.FormatFloat(v, 'f', 2, 64)
		}

		// Join the string array with commas
		result := strings.Join(strArr, ", ")
		return fmt.Sprintf("[%s]", result)
	}

	html := fmt.Sprintf(`
	<html>
	<head>
		<title>Statistics</title>
		<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
		<style>
			body { font-family: Arial, sans-serif; }
			h2 { color: #333; }
		</style>
	</head>
	<body>
		<h2>Statistics for last run:</h2>
		<p>Covered lines: %d</p>
		<p>Uncovered lines: %d</p>
		<p>Untracked lines: %d</p>
		<p>Percentage of covered lines: %.2f%%</p>

		<h2>Cummulative statistic:</h2>
		<p>Covered lines: %d</p>
		<p>Uncovered lines: %d</p>
		<p>Untracked lines: %d</p>
		<p>Percentage of covered lines: %.2f%%</p>

		<h2>Graphs</h2>
		<canvas id="lastRunChart" width="800" height="200"></canvas>
		<canvas id="cumulativeChart" width="800" height="200"></canvas>

		<script>
			var ctx1 = document.getElementById('lastRunChart').getContext('2d');
			var lastRunChart = new Chart(ctx1, {
				type: 'line',
				data: {
					labels: Array.from({length: %d}, (_, i) => i + 1),
					datasets: [{
						label: 'Last Run',
						data: %v,
						borderColor: 'green',
						fill: false,
					}]
				}
			});

			var ctx2 = document.getElementById('cumulativeChart').getContext('2d');
			var cumulativeChart = new Chart(ctx2, {
				type: 'line',
				data: {
					labels: Array.from({length: %d}, (_, i) => i + 1),
					datasets: [{
						label: 'Cumulative',
						data: %v,
						borderColor: 'blue',
						fill: false,
					}]
				}
			});
		</script>

	</body>
	</html>
	`, greenLast, redLast, blackLast, lastRunCovered,
		greenCummulative, redCummulative, blackCummulative, cummulativeCovered,
		len(lastRunList), formatList(lastRunList), len(coverList), formatList(coverList))

	return html
}

func colorToHTML(c color) string {
	switch c {
	case Red:
		return "red"
	case Green:
		return "green"
	default:
		return "black"
	}
}

func convertToColoredHTML(originalFile []byte, colorInfo []color) (html string,
	blackCount, greenCount, redCount int) {
	lines := strings.Split(string(originalFile), "\n")

	var result strings.Builder
	blackCount, greenCount, redCount = 0, 0, 0
	for i, line := range lines {
		// Preserve tabs and spaces for HTML
		line = strings.ReplaceAll(line, "\t", " ") // or use " " for narrower spacing
		line = strings.ReplaceAll(line, " ", " ")

		// Handle empty lines
		if len(strings.TrimSpace(line)) == 0 {
			line = " "
		}

		if i < len(colorInfo) { // Make sure there's a color defined for this line
			if colorInfo[i] == Uncolored {
				blackCount++
			} else if colorInfo[i] == Red {
				redCount++
			} else if colorInfo[i] == Green {
				greenCount++
			}
			color := colorToHTML(colorInfo[i])
			result.WriteString(fmt.Sprintf(`<div style="white-space: pre; color: %s">%s</div>`, color, line))
		} else {
			result.WriteString(fmt.Sprintf("<div style=\"white-space: pre; color: black\">%s</div>", line))
			blackCount++
		}
	}
	// Calculate percentage
	percentageGreen := 100.0
	if greenCount+redCount != 0 {
		percentageGreen = (float64(greenCount) / float64(greenCount+redCount)) * 100
	}

	// Append the summary to the HTML
	result.WriteString(fmt.Sprintf(`
		<div style="margin-top: 20px; border-top: 1px solid #ccc; padding-top: 10px; color: black">
			<p><strong>Summary:</strong></p>
			<p>Covered lines: %d</p>
			<p>Uncovered lines: %d</p>
			<p>Other lines: %d</p>
			<p>Percentage of covered lines): %.2f%%</p>
		</div>
	`, greenCount, redCount, blackCount, percentageGreen))
	return result.String(), blackCount, greenCount, redCount
}

func assignOutputColor(originalFile []byte, coverInfoList []coverInfo, coverTab []byte) []color {
	lineNum := len(strings.Split(string(originalFile), "\n"))
	colorInfo := make([]color, lineNum)
	for _, coverInfo := range coverInfoList {
		if coverInfo.coverIdx >= 0 && coverInfo.coverIdx < CoverSize {
			if coverTab[coverInfo.coverIdx] > 0 {
				for i := coverInfo.startLineNum - 1; i < coverInfo.endLineNum; i++ {
					colorInfo[i] = Green
				}
			} else {
				for i := coverInfo.startLineNum - 1; i < coverInfo.endLineNum; i++ {
					colorInfo[i] = Red
				}
			}
		}
	}
	return colorInfo
}

// []coverInfo contains the information of the cover blocks in the instrumented file,
// ordered by their appearance of the cover start in the source file
// map[int]int maps the cover index to the index in the []coverInfo array
func constructCoverInfo(srcLines []string, originalLines []string) ([]coverInfo, map[int]int) {
	coverInfoList := make([]coverInfo, 0)
	coverIdxMap := make(map[int]int)
	for i, line := range srcLines {
		if isCoverStart(line) != -1 {
			coverIdx := isCoverStart(line)
			if i == 0 {
				panic("Error parsing line comment: first line cannot be cover start")
			}
			_, lineNum := parseSrcLineComment(srcLines[i-1])
			if lineNum == -1 {
				// panic("Error parsing line comment: cannot find line number for cover start")
				continue
			}
			coverInfo := coverInfo{coverIdx: coverIdx, startLineNum: lineNum}
			coverInfoList = append(coverInfoList, coverInfo)
			coverIdxMap[coverIdx] = len(coverInfoList) - 1
		} else if isCoverEnd(line) != -1 {

			coverIdx := isCoverEnd(line)
			if _, ok := coverIdxMap[coverIdx]; !ok {
				// panic("Error parsing line comment: cannot find cover start for cover end")
				continue
			}
			_, lineNum := parseSrcLineComment(srcLines[i-1])
			if lineNum == -1 {
				coverInfoList[coverIdxMap[coverIdx]].endLineNum = coverInfoList[coverIdxMap[coverIdx]].startLineNum
				// panic(fmt.Sprintf("Error parsing line comment: cannot find line number for cover start, line: %v, srcLine: %v\n", line, srcLines[i-1]))
				continue
			}
			for lineNum < len(originalLines) {
				if len(strings.TrimSpace(originalLines[lineNum])) == 0 {
					lineNum++
					continue
				}
				if strings.TrimSpace(originalLines[lineNum]) == "}" {
					lineNum++
					break
				}
				break
			}
			coverInfoList[coverIdxMap[coverIdx]].endLineNum = lineNum
		}
	}
	return coverInfoList, coverIdxMap
}

// returns the original go file path and line number
func parseSrcLineComment(line string) (string, int) {
	line = strings.TrimSpace(line)
	pattern := `//line (.+\.go):(\d+)`
	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(line)

	if len(match) > 2 {
		// Convert string to integer
		intValue, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return "", -1
		}
		return match[1], intValue
	} else {
		return "", -1
	}
}

func isInstrumented(fileContent string) (originalFile string) {
	if !strings.Contains(fileContent, "_go_fuzz_dep_.CoverTab[") {
		return ""
	}

	fileToken := strings.Split(fileContent, "\n")
	for _, line := range fileToken {
		if file, lineNum := parseSrcLineComment(line); file != "" && lineNum != -1 {
			return file
		}
		// line = strings.TrimSpace(line)
		// if strings.HasPrefix(line, "//line") {
		// 	lineToken := strings.Split(line, " ")
		// 	if len(lineToken) >= 2 {
		// 		pathToken := strings.Split(lineToken[len(lineToken)-1], ":")
		// 		if len(pathToken) >= 2 {
		// 			_, err := strconv.Atoi(pathToken[len(pathToken)-1])
		// 			if err != nil {
		// 				continue
		// 			}
		// 			lineNumLen := len(pathToken[len(pathToken)-1])
		// 			originalFile = strings.TrimSpace(line[6 : len(line)-lineNumLen-1])
		// 			return originalFile
		// 		}
		// 	}
		// }
	}
	return ""
}

func isCoverStart(line string) int {
	line = strings.TrimSpace(line)
	pattern := `_go_fuzz_dep_\.CoverTab\[(\d+)\]\+\+`
	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(line)

	if len(match) > 0 {
		intValue, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return -1
		}
		return intValue
	} else {
		return -1
	}
}

func isCoverEnd(line string) int {
	line = strings.TrimSpace(line)
	// Match _ = "end of CoverTab[i]" where i is a positive integer, and capture the value of i
	pattern := `// = "end of CoverTab\[(\d+)\]"`
	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(line)

	if len(match) > 1 {
		// Convert string to integer
		intValue, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return -1
		}
		return intValue
	} else {
		pattern2 := `// _ = "end of CoverTab\[(\d+)\]"`
		re2 := regexp.MustCompile(pattern2)

		match2 := re2.FindStringSubmatch(line)

		if len(match2) > 1 {
			// Convert string to integer
			intValue2, err := strconv.Atoi(match2[1])
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return -1
			}
			return intValue2
		} else {
			return -1
		}
	}
}

func createDir(dest string) {
	// Check if destDir exists
	_, err := os.Stat(dest)
	if err == nil {
		// Directory exists, remove it
		err = os.RemoveAll(dest)
		if err != nil {
			panic(err)
		}
	} else if os.IsNotExist(err) {
		// Directory does not exist, do nothing
	} else {
		// For other types of errors
		panic(err)
	}

	// Make new destination directory
	err = os.MkdirAll(dest, 0755)
	if err != nil {
		panic(err)
	}
}
