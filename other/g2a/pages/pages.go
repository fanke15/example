package pages

const (
	Page1 = `{"type":"page","title":"404","body":[{"type":"markdown","value":"# 🚫 Oops,  找不到对应的页面配置\n[👉 点击我返回页面列表](/)"}],"regions":["body"]}`

	Page2 = `

{
    "type": "page",
    "data": {
        "name": "amis",
        "age": 1
    },
    "body": [
        {
            "type": "tpl",
            "tpl": "my name is ${name}"
        },
        {
            "type": "tpl",
            "tpl": "I am ${age} years old"
        }
    ]
}

`
)

var (
	Page3 = []byte(`{
  "type": "page",
  "title": "用户登录",
  "body": [
    {
      "type": "html",
      "html": "<div style='display: flex; justify-content: center; align-items: center; margin: 96px 0px 8px;'><img src='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAFgAAABYBAMAAACDuy0HAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAkUExURQuQ9A+M7hCM7xCM7hCM7hCM7w+M7hCM7hCM7xCM7w+M7hCM7p2RGX4AAAALdFJOUwHxFtl5wi6YRq9emx6XmgAAA0VJREFUSMeNV89rGkEU3m6sVXNKK6TEiweLhJ56C/HikqN3EbwEihRv7RZKIRchBHq2EChesvgXiC5mmX+u+2Pem/fezLo7Bw/6OTPv+773YzzPsVp/97+9mqs1Ukrd1MTOVbamdbD+gyrW9xrgQGOP/Wrso8bG99XYpcbuX6qxQ43dhdXYTk+D1zUJVpTkaFaGbW40tgvfLErP8IOeIC2PYOwEbyVpjfzP+9UJgpG0jo4gsdW5BtJA5eZcyRBgDZQgDS1iOxAJvpEWsWlHgq9ktHoRSvB6SNoCIoCoV5aDkbQhnr+U28D1dmNGcLpSsZ/4Bb/CNqEzAthq6iTCR4vkR/tzcu5EUg8RJDqCNkTU95pSp0ifFAMBAE4z+J3QaSEtQk3+lus0lBZBXrNozww4DaFh5RWEdMiiPSei7v+MDMHcuEW0GXg34k4wFllyuTNw3OHoS2lc0DYHewyd3AvZsYp8zsFGR+Kw1heqtNk5PbEn7YTGvUUvZzsnNJZQEow++GbAWrtpaekdxgTsRfRIu4pcFOAjbnaQeUUtwsHNf31ZesekDsaFgrKeoUVYtAVY1jMkDS0S5WJpIyWsgzxIggs7xeA62psiaREtAeysdqb6fVKieoK4ibHQzOOsIcFom4SYYs2uYedgrt21Ev0348nk4EZRsKmCY7OZlYPoCggfS8VwJplR2slZn1Il7QMPvTN+w4sd3d1ZhWe4Mwm567TIc16LYqupzBzN49bjYFPz1padstMKcLsvG6Em0P9B86oAn3etFrvifOaya7AxYkCLjPgrgNNYhY8v+9h34FJt9HMoB47kTjaldi7Km/TzwnaYKFz+BMEmBxs9huV5ddRtIpE1wBp/suIPYsh8s8cfpV69ljxxW1qm3xu+ZEU0ZRqiTmloKGet3b1IPj86sypzkMnBCU+MQPaHzkikOvkR/WXKCgYgZSfXOpQNcCo8NVAJQafOUY1/O3IPdgPHeSyv3OPl2BrKHE+FJ0lgIGWnK+I/bq35j60JnbqQ4JJnBRaxKzfBJQ+E2UC5qCwRoecmWJRNnoPd0++rBcUmVW+xR1LpV5XPwUBZWp5Ab8TMWOdZ+lzvEduZ18eme//88Mv1/X8KQ6zq3Tt8/QAAAABJRU5ErkJggg==' alt='logo' style='margin-right: 8px; width: 48px;'><span style='font-size: 32px; font-weight: bold;'>G2a Admin</span></div><div style='width: 100%; text-align: center; color: rgba(0, 0, 0, 0.45); margin-bottom: 40px;'>Amis是一个低代码前端框架，可以减少页面开发工作量，极大提升效率<br/></div>",
      "id": "u:ab5a58fa4273"
    },
    {
      "type": "grid",
      "valign": "middle",
      "align": "center",
      "columns": [
        {
          "lg": 3,
          "md": 3,
          "sm": 3,
          "valign": "middle",
          "body": [
            {
              "type": "form",
              "mode": "horizontal",
              "horizontal": {
                "left": 3,
                "right": 9
              },
              "title": "",
              "submitText": "登录",
              "body": [
                {
                  "type": "input-text",
                  "name": "mobile",
                  "label": "手机号",
                  "required": true,
                  "id": "u:de62abf78425",
                  "validations": {
                    "isTelNumber": true
                  },
                  "value": "110"
                },
                {
                  "type": "input-password",
                  "name": "password",
                  "label": "密码",
                  "value": "admin",
                  "required": true,
                  "maxLength": 128,
                  "id": "u:ff7f98c09a31"
                },
                {
                  "type": "checkbox",
                  "name": "rememberMe",
                  "label": "记住登录",
                  "id": "u:359093b19d9d"
                }
              ],
              "actions": [
                {
                  "type": "button-toolbar",
                  "buttons": [
                    {
                      "type": "button",
                      "actionType": "link",
                      "label": "注册",
                      "link": "",
                      "id": "u:3d6c526bff91"
                    },
                    {
                      "type": "button",
                      "actionType": "submit",
                      "label": "登录",
                      "level": "primary",
                      "id": "u:792559ff0e7c"
                    }
                  ],
                  "id": "u:9a29ae396a5e"
                }
              ],
              "actionsClassName": "no-border m-none p-none",
              "wrapWithPanel": true,
              "panelClassName": "",
              "api": {
                "url": "",
                "method": "POST"
              },
              "silentPolling": false,
              "redirect": "",
              "id": "u:1d400b717192"
            }
          ],
          "id": "u:7ec8898f5956"
        }
      ],
      "id": "u:edbe4e87fd14"
    }
  ],
  "id": "u:9c2ae9613794"
}`)
)
