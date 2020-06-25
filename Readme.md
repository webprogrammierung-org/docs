# Installieren

## Python 3 und PIP 3

### Ubuntu 

```

sudo apt-get install python3  python3-pip


```

### Windows 

* [Python 3](https://www.python.org/)

## Sphinx Dokumente

```

pip3 install -U Sphinx

```

## Sphinx Theme

###  Installieren Sphinx-rtd-theme

Dieses Theme wird auf PyPI verteilt und kann mit pip3 installiert werden:


```
pip3 install sphinx-rtd-theme

```

Um das Layout in Ihrem Sphinx-Projekt zu verwenden, müssen Sie das Folgende zu Ihrer conf.py-Datei hinzufügen:

```
import sphinx_rtd_theme

extensions = [
    ...
    "sphinx_rtd_theme",
]

html_theme = "sphinx_rtd_theme"

```

### Installieren Mardown 

* Installieren Sie den Markdown-Parser recommonmark:

```
pip install --upgrade recommonmark

```

* Recommonmark zur Liste der eingerichteten Erweiterungen hinzufügen:

```
extensions = ['recommonmark']

```

* Wenn Sie Markdown-Dateien mit anderen Erweiterungen als .md verwenden möchten, passen Sie die Variable source_suffix an. Im folgenden Beispiel wird Sphinx so konfiguriert, dass alle Dateien mit den Erweiterungen .md und .txt als Markdown geparst werden:

```
source_suffix = {
    '.rst': 'restructuredtext',
    '.txt': 'markdown',
    '.md': 'markdown',
}

```

### Go get 

```

go get github.com/webprogrammierung-org/docs

```
