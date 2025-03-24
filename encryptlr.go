package encryptLR

import "fmt"

func decodeLR(encoded string) string {
	const numDigits = 10
	n := len(encoded)
	dp := make([][]string, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]string, numDigits)
	}

	for currDigit := 0; currDigit < numDigits; currDigit++ {
		dp[0][currDigit] = fmt.Sprintf("%d", currDigit)
	}

	for i := 1; i <= n; i++ {
		for currDigit := 0; currDigit < numDigits; currDigit++ {
			for prevDigit := 0; prevDigit < numDigits; prevDigit++ {
				if isValidTransition(rune(encoded[i-1]), prevDigit, currDigit) && dp[i-1][prevDigit] != "" {
					updateSequence(dp, i, currDigit, dp[i-1][prevDigit], currDigit)
				}
			}
		}
	}

	bestSequence := ""
	for _, seq := range dp[n] {
		if seq != "" && (bestSequence == "" || sum(seq) < sum(bestSequence)) {
			bestSequence = seq
		}
	}

	return bestSequence
}

func isValidTransition(char rune, prev int, curr int) bool {
	switch char {
	case 'L':
		return prev > curr
	case 'R':
		return prev < curr
	case '=':
		return prev == curr
	default:
		return false
	}
}

func updateSequence(dp [][]string, i, currDigit int, prevSequence string, currInt int) {
	candidateSequence := prevSequence + fmt.Sprintf("%d", currInt)

	if dp[i][currDigit] == "" || sum(candidateSequence) < sum(dp[i][currDigit]) {
		dp[i][currDigit] = candidateSequence
	}
}

func sum(s string) int {
	total := 0

	for _, ch := range s {
		total += int(ch - '0')
	}

	return total
}
