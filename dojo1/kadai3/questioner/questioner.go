package questioner

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// Quiz ...
type Quiz struct {
	limitTime time.Duration
	quiz      []string
	count     int
}

// New ...
func New(time time.Duration, quiz []string) *Quiz {
	q := Quiz{
		limitTime: time,
		quiz:      quiz,
		count:     0,
	}

	return &q
}

// 与えられた英単語の選択肢からランダムに選んで返す
func selectQuiz(q []string) string {
	// 乱数をさせる
	i := rand.Intn(len(q))

	// 受け取ったListの中からランダムな選択肢を返す
	return q[i]
}

func (q *Quiz) judge(userAnswer, correctAnswer string) {
	if userAnswer == correctAnswer {
		q.count++
		fmt.Println("正解です")
	} else {
		fmt.Println("不正解です")
	}
}

func getUserInput(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()

	return ch
}

// Start ...
func (q *Quiz) Start() {
	fmt.Println("タイピングゲームを開始します。表示される英単語をできるだけ多く入力してください")
	fmt.Println("制限時間は", q.limitTime, "です。スタート！")
	finish := time.After(q.limitTime)
	ch := getUserInput(os.Stdin)
	for {
		select {
		case <-finish:
			fmt.Println("タイピング終了です。")
			fmt.Println("あなたの成績は", q.count, "です")
			return
		default:
			correctAnswer := selectQuiz(q.quiz)
			fmt.Println(correctAnswer)
			if input, ok := <-ch; ok {
				q.judge(input, correctAnswer)
			} else {
				os.Exit(1)
			}
		}
	}
}
