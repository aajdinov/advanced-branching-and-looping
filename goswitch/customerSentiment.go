package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

type rating float32

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	initial       rating = 5.0
	negative      rating = -0.6
	extraNegatice rating = -1.2
)

type result struct {
	feedbackDate     string
	feedbackTotal    int
	feedbackPositive int
	feedbackNegative int
	feedbackNeutral  int
}

var customerSentiment []rating

func main() {
	f, err := os.Open("customerSentiment.csv")
	if err != nil {
		exitProgram(err.Error())
	}

	defer f.Close()

	r := bufio.NewReader(f)
	str, err := r.ReadString('\n')
	if err != nil {
		exitProgram(err.Error())
	}

	var feedback result
	feedback.feedbackDate = str
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if len(str) > 10 {
			feedback.feedbackTotal++

			var customerRating rating
			customerRating = initial

			text := strings.Split(str, " ")
			for _, word := range text {
				switch s := strings.Trim(strings.ToLower(word), " ,.!;:!?\t\n\r"); s {
				case "pleasure", "impressed", "wonderful", "fantastic", "splendid", "great":
					customerRating += extraPositive
				case "good", "help", "helpful", "nice", "friendly", "happy", "satisfied", "excellent", "thank", "thanks":
					customerRating += positive
				case "not helpful", "sad", "angry", "improve", "annoy":
					customerRating += negative
				case "terrible", "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
					customerRating += extraNegatice
				}
			}

			switch {
			case customerRating > 8.0:
				feedback.feedbackPositive++
			case customerRating >= 4.0:
				feedback.feedbackNeutral++
			case customerRating < 4.0:
				feedback.feedbackNegative++
			}

			customerSentiment = append(customerSentiment, customerRating)
		}
	}
	feedbackTable(feedback, customerSentiment)
}

func exitProgram(s string) {
	log.Fatal("Exiting the program with: ", s)
}

func feedbackTable(f result, c []rating) {
	writer := tabwriter.NewWriter(os.Stdout, 1, 4, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintf(writer, f.feedbackDate)
	fmt.Fprintf(writer, "Total feedback:\t%v\n", f.feedbackTotal)
	fmt.Fprintf(writer, "Positive feedback:\t%v\n", f.feedbackPositive)
	fmt.Fprintf(writer, "Neutral feedback:\t%v\n", f.feedbackNeutral)
	fmt.Fprintf(writer, "Negative feedback:\t%v\n", f.feedbackNegative)
	fmt.Fprintf(writer, "Customer sentiment:\t%v\n", c)
	writer.Flush()
}
