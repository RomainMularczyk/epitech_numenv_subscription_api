package mail

import (
	"fmt"
)

func FormatContent(session string, uniqueStr string) string {
	htmlContent := fmt.Sprintf(
		`<!DOCTYPE html>
    <html lang="fr">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Inscription .env{2024}</title>
        <style>
            header {
                margin-left: 5em;
                padding: 20px;
                display: -webkit-box;
                vertical-align: middle;
                -webkit-align-content: center;
            }
			.header-wrapper {
				display: flex;
				flex-direction: row;
				align-items: center;
			}
            .content {
                margin: 5em;
            }
            .bold {
                font-weight: bold;
            }
            .credits-stack {
                display: flex;
                font-size: 1.1em;
                flex-direction: row;
                justify-content: center;
                align-items: center;
            }
            .code {
              background-color: #4169E1;
              color: #F3F2F2;
              border-radius: 5px 5px 5px 5px;
              padding: 2px 4px 2px 4px;
            }
            p {
                text-align: left;
                font-size: 18px;
            }
            footer {
                display: flex;
                flex-direction: row;
                justify-content: space-evenly;
                text-align: center;
                padding: 10px;
                font-size: 12px;
                margin: 5em;
            }
            .btn-box {
                display: flex;
                flex-direction: row;
                justify-content: center;
            }
            .btn {
                display: inline-block;
                padding: 10px 20px;
                background-color: #4169E1;
                color: #ffffff !important;
                text-decoration: none;
                border-radius: 5px;
                font-weight: bold;
            }
            #title {
                font-weight: bold;
                text-transform: uppercase;
                margin-left: 1em;
            }
        </style>
    </head>
    <body>
        <header>
			<div class="header-wrapper">
				<a target="_blank" href="https://envnum.fr">
				  <img width="60px" src="https://www.imghost.net/ib/DCXSkMBLLzEOGC6_1704764808.png" alt="logo"/>
				</a>
				<h1 id="title">Inscription .env{2024}</h1>
			</div>
        </header>
        <main>
            <div class="content">

                <p>Merci pour votre inscription à la session : <b>%s</b></p>
                <br>
                <p>Pour finaliser votre inscription, merci de rejoindre le serveur Discord 
                    <span class="bold">.env{2024}</span> et d'utiliser la commande 
                    <span class="code">/register</span> avec le code de connexion suivant : 
                    <span class="code">%s</span>
                </p>
                <br>
                <p class="btn-box">
                  <a 
                    href="https://discord.gg/ADaCQ7bV9Y" 
                    class="btn"
                  >
                    Rejoindre le serveur
                  </a>
                </p>
                <br>
                <p>
                    Lors de l'événement, vous pourrez ainsi échanger avec les autres
                    participant.e.s et poser vos questions aux intervenant.e.s.
                </p>
                <br>
          <p>
          Une fois enregistré sur le serveur, vous pouvez rejoindre directement une 
          nouvelle session depuis le canal <span class="code">#welcome</span> en utilisant la commande 
          <span class="code">/subscribe</span>
          suivie du nom de l'intervenant de la session <span class="code">/subscribe &lt;intervenant&gt;</span>
          </p>

            </div>
        </main>
		<footer>
			<div class="credits-stack">
				<p>💚 Team .env&#123;2024&#125; </p>
			</div>
		</footer>
    </body>
    </html>`,
		session,
		uniqueStr,
	)

	return htmlContent
}
