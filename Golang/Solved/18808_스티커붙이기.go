package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N, M, K  int
	R, C     int
	Stickers [][][]int
	cover    [41][41]int
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d %d\n", &N, &M, &K)

	for k := 0; k < K; k++ {

		fmt.Fscanf(r, "%d %d\n", &R, &C)

		stick := make([][]int, R)
		for i := 0; i < R; i++ {
			stick[i] = make([]int, C)
			for j := 0; j < C; j++ {
				fmt.Fscanf(r, "%d ", &stick[i][j])
			}
		}

		Stickers = append(Stickers, stick)
	}

	// 0, 0 부터
	// 스티커의 최대 가로와 세로 [][]

	for len(Stickers) > 0 {

		sticker := Stickers[0]
		Stickers = Stickers[1:]

		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {

				// 회전을 다했는데 안 붙으면 다음 칸으로 간다
				// 모든 칸에서 불가능하면 스티커를 버린다.

				stickToLaptop(i, j, sticker)
			}
		}

	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Printf("%d ", cover[i][j])
		}
		fmt.Println()
	}

	answer := 0

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if cover[i][j] == 1 {
				answer++
			}
		}

	}
	fmt.Fprintln(w, answer)
}

func rotate(sticker [][]int) [][]int {
	temp := [41][41]int{}

	r := len(sticker)
	c := len(sticker[0])

	newSticker := make([][]int, c)
	for i := 0; i < c; i++ {
		newSticker[i] = make([]int, r)
	}
	h := r

	tmp := r
	r = c
	c = tmp

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			temp[i][j] = sticker[h-1-j][i]
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			newSticker[i][j] = temp[i][j]
		}
	}

	return newSticker
}

func stickToLaptop(i, j int, sticker [][]int) {

	available := false

	for k := 0; k < 4; k++ {
		if i+len(sticker)+1 > N || j+len(sticker[0])+1 > M {

			newSticker := rotate(sticker)
			sticker = newSticker

		} else {
			for m := i; m < len(sticker); m++ {
				for l := j; l < len(sticker[0]); l++ {
					if cover[m][l] == 1 && sticker[m][l] == 1 {
						goto out
					}

				}
			}
			available = true
		}
	out:
	}

	// 스티커 합격, 붙여넣기

	if available {
		for k := i; k < len(sticker); k++ {
			for l := j; l < len(sticker[0]); l++ {

				cover[k][l] = sticker[k][l]
			}
		}
	}

}
