# Exercice Go + HTMX
Projet démontrant l'utilisation de la librairie front-end `HTMX` avec le langage de programmation `Go`. Le tout connecté à une base de données `SQLite`.

## Librairies et langages
* **Front-end**
  * [HTMX](https://htmx.org/)
  * [Tailwind CSS](https://tailwindcss.com/)
* **Back-end**
  * [Go](https://go.dev/)
    * [Gin Gonic (Framework Web)](https://gin-gonic.com/)
    * [GORM (ORM pour Go)](https://gorm.io/index.html)
* **Base de données**
  * [SQLite](https://www.sqlite.org/index.html)

## Mise en place
1. Clôner ce répertoire, utiliser la branche `main`.
2. Effectuer les commandes suivantes dans le terminal au document racine:
    * `npm install`
    * `npm run build-css`
    * `npm run dev`

Si vous recevez un message d'erreur concernant la variable CGO_ENABLED qui a une valeur `0`, effectuer la commande suivante et refaire la commande `npm run dev`:
```bash
go env -w CGO_ENABLED=1
```

### Windows
Il est possible d'avoir une erreur concernant GCC à cause de l'implémentation de l'ORM `Gorm`. Pour régler le problème, télécharger et installer le compilateur GCC ici: https://jmeubank.github.io/tdm-gcc/

Fermer les terminaux et les réouvrir pour compléter leur installation.