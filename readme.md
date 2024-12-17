## game koding

Das hier ist ein Lernprojekt um einen Einblick zu bekommen wie Spiele programmiert werden.

## Einrichtung
- Wir setzen VS Code ein, was installiert sein sollte. Hier kann man das runterladen: https://code.visualstudio.com/
- Programmiert wird in Go, das bekommt man hier: https://go.dev/ 
    - Es muss möglich sein in VS Code ein Terminal zu öffnen (im Menü oben) und dort `go version` einzugeben und eine Antwort zu bekommen wie "go1.23.4"
    - In vs code kann man in der linken Leiste unter Extensions nach "go" suchen. Die sollte auch installiert sein, dann versteht VS Code unseren Go code und zeigt Fehler and etc. 
- Wir setzen die ebitengine ein, die für 2D Spiele gedacht ist:
    - Hier die Website mit Beispielen und Doku: https://ebitengine.org/
    - Hier die Sources dazu: https://github.com/hajimehoshi/ebiten
- Bisher wurden die unglaublichen Grafiken mit Photopea (einer kostenlosen Photoshop Alternative) gemacht: https://www.photopea.com/

## Mit git arbeiten
git ist ein Werkzeug, mit dem Software-Entwickler gemeinsam von ihren Rechnern aus an einem Projekt arbeiten können. Das machen wir hier auch direkt. Dazu muss auf dem Rechner git installiert sein. Das bekommt man hier:
https://git-scm.com/downloads

Wenn git installiert ist, findet VS Code das und bietet alle wichtigen Funktionen von git direkt an.

## Laufen lassen per git clone
Wenn man VS Code startet oder ein neues Fenster öffnet (kann man im Menü machen -> neues Fenster) gibt es dort in der leeren Fläche einen Link "Projekt aus git clonen" oder so ähnlich. Da kann man drauf klicken, dann wird man gefragt nach einem Link zum repository. Spoiler: da gehört diese URL rein: `https://github.com/martinfriedel/gamekoden.git`, aber bekommen tut man diese, indem man auf github bei  https://github.com/martinfriedel/gamekoden auf den grünen "<> Code" button klickt, dann dort in dem "HTTPS" Reiter diese URL rauskopiert. 

VS Code sollte nun nach einem Ordner suchen, wo man denn die Daten von github auf die Festplatte speichern möchte. Am bestenen einen Ordner machen wo Programmierprojekte reingehören und den auswählen. Es wird dann ein Unterordner "gamekoden" angelegt und da werden die Dateien abgelegt. 

## Alternativ: Laufen lassen per Download
Du kannst diesen Code von Github herunterladen, indem du auf den grünen Button "<> Code" klickst, dann auf "download ZIP". 

Dann packst du das ZIP aus und öffnest den Ordner mit VS Code. Dann ein Terminal öffnen und `go run .` eingeben. Das Spiel sollte in einem Fenster starten. 



