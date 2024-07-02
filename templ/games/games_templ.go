// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package games

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

type Game struct {
	Name    string `json: name`
	Rating  string `json: rating`
	Comment string `json: comment`
}

func GamesComp(gamesData []Game) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"games\"><h1>Правила игрового аука</h1><p>Я играю в игру 2 часа, и если мне нравится, прохожу до конца. Если вы соблюли все правила, то я не дропаю игру ранее 2 часов прохождения.</p><p>Выкуп аукциона 5000 р</p><h2>Желательно:</h2><ul><li>Сложность:<ul><li>Средняя</li><li>В самый раз</li><li>легко</li></ul></li><li>Возраст:<ul><li>Не старше 10 лет</li></ul></li><li>Язык:<ul><li>Русский</li></ul></li><li>Длительность по мэин сюжету:<ul><li>не более 25 часов</li></ul></li>- не дорожее 3к Рублей</ul><h2>Не желательно:</h2><ul><li>Заказывать игры которые я уже играла \\ дропала</li></ul><h2>Нельзя:</h2><ul><li>Хорроры</li><li>Потенциально страшные игры</li><li>Психологические триллеры</li><li>Игра не доступные для покупки в России</li><li>Стратегии</li><li>игры нарушающие правила стрима \\ твича</li><li>Туканы</li><li>Онлайны</li><li>Соуслайк игры</li><li>Песочницы (Игры у которого нет финального завершения)</li><li>игры которые не для пк</li><li>Стратегии</li><li>Игры по Вахе</li><li>Игры которые были (Есть в списке)</li></ul><p>ps Игра будет дропнута и вам никто не вернет деньги \\ баллы, если вы будите меня бэк ситтить</p><p>pss пренебрегать правилами можно, но и дропнуть игру я смогу в любой момент</p><p>pss Если у вас есть вопросы, лучше спросить их у стримера, чтобы избежать непонимания</p><h1>Бывшие Слоты</h1><table><tr><th>Название</th><th>Оценка</th><th>Комментарий <a href=\"https://twitch.tv/HakaHyHe\">HakaHyHe</a></th></tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, game := range gamesData {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(game.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templ/games/games.templ`, Line: 81, Col: 19}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(game.Rating)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templ/games/games.templ`, Line: 82, Col: 21}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(game.Comment)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templ/games/games.templ`, Line: 83, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</table></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}