package mail

import (
  "fmt"
  "testing"
  "github.com/google/uuid"
) 

func TestFormatContent(t *testing.T) {
  uniqueStr, err := uuid.NewRandom()
  if err != nil {
    fmt.Println("Could not generate a UUID.")
  }

  got := FormatContent("First session", uniqueStr.String())
  want := fmt.Sprintf(`<!DOCTYPE PUBLIC “-//W3C//DTD XHTML 1.0 Transitional//EN” “https://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd”>
    <html xmlns="http://www.w3.org/1999/xhtml">
    <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <title></title>
    </head>
    <body>
      <p>Merci pour votre inscription à la session : First session</p>

      <p>Pour finaliser votre inscription et pouvoir poser vos questions et échanger avec les
      autres participant.e.s, merci de rejoindre le serveur Discord Envnum et d'utiliser la commande 
      /register avec le code de connexion suivant : %s.
      </p>

      <table
        width='100%'
        align='center'
        border='0'
        cellspacing='0'
        cellpadding='0'
        style="table-layout: fixed;"
      >
        <tr>
          <td aligh='center'>
            <a 
              href='https://discord.gg/e3C7v8qPa5' 
              style='
                background-color: #4169E1;
                color: #000000;
                font-size: 1.3em;
                padding: 10px;
                border-radius: 10px;
                text-decoration: none;
              '
            >
              Rejoindre le server
            </a>
          </td>
        </tr>
      </table>
    </body>`, uniqueStr)

  // Assert
  if got != want {
    t.Errorf("got %q, wanted %q", got, want)
  }
}
