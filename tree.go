package main

import (
	"github.com/zoh/decision-tree/tree"
)

func main() {

	t := tree.DecisionTree{CategoryAttr: "person", IgnoredAttribute: []string{}}

	trainigSet := tree.TrainingSet{
		tree.TrainingItem{"person": "Homer", "hairLength": 0, "weight": 250, "age": 36, "sex": "male"},
		tree.TrainingItem{"person": "Marge", "hairLength": 10, "weight": 150, "age": 34, "sex": "female"},
		tree.TrainingItem{"person": "Bart", "hairLength": 2, "weight": 90, "age": 10, "sex": "male"},
		tree.TrainingItem{"person": "Lisa", "hairLength": 6, "weight": 78, "age": 8, "sex": "female"},
		tree.TrainingItem{"person": "Maggie", "hairLength": 4, "weight": 20, "age": 1, "sex": "female"},
		tree.TrainingItem{"person": "Abe", "hairLength": 1, "weight": 170, "age": 70, "sex": "male"},
		tree.TrainingItem{"person": "Selma", "hairLength": 8, "weight": 160, "age": 41, "sex": "female"},
		tree.TrainingItem{"person": "Otto", "hairLength": 10, "weight": 180, "age": 38, "sex": "male"},
		tree.TrainingItem{"person": "Krusty", "hairLength": 6, "weight": 200, "age": 45, "sex": "male"}}

	tree.TrainingTree(&t, trainigSet)
	item := tree.TrainingItem{
		"person": "Comic guy", "hairLength": 8, "weight": 290, "age": 38}

	t.Predict(item)
	t.SaveToHtml("e:/decisionTree/simpsons.html")

}
