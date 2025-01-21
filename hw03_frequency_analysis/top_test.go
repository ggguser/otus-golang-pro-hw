package hw03frequencyanalysis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}

func TestGetWords(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{input: "a b   c	d", expected: []string{"a", "b", "c", "d"}},
		{input: "a - b", expected: []string{"a", "-", "b"}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			result := getWords(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestGetWordsCount(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string]int
	}{
		{input: []string{"a", "a", "a"}, expected: map[string]int{"a": 3}},
		{input: []string{"a", "b", "c"}, expected: map[string]int{"a": 1, "b": 1, "c": 1}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			result := getWordsCount(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestSortWordsByCount(t *testing.T) {
	tests := []struct {
		input    map[string]int
		expected []string
	}{
		{input: map[string]int{"a": 1, "b": 2, "c": 3}, expected: []string{"c", "b", "a"}},
		{input: map[string]int{"a": 3, "b": 4, "c": 4}, expected: []string{"b", "c", "a"}},
		{input: map[string]int{
			"то":        4,
			"если":      4,
			"а":         6,
			"не":        4,
			"-":         4,
			"и":         6,
			"Кристофер": 4,
			"ты":        5,
			"что":       5,
			"он":        8,
		}, expected: []string{"он", "а", "и", "ты", "что", "-", "Кристофер", "если", "не", "то"}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			result := sortWordsByCount(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
