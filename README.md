
# Google Translate CLI Tool

This is a command-line interface (CLI) tool for Google Translate, built using Go. This tool allows users to translate text from one language to another directly from the terminal.
Getting Started

You need to have Go installed on your machine to run this tool.

Clone the repository to your local machine:


```bash
git clone https://github.com/Bostigger/google-translate-cli.git
```


Navigate to the cloned directory:

```bash
cd google-translate-cli
```


Usage

You can run the tool using the go run command followed by the parameters for source language, target language, and the text to translate. The parameters are as follows:

    -s: Source Language (example is "en").
    -t: Target Language (example is "fr").
    -st: Text to translate.




Example command:

```bash
go run main.go -s en -t fr -st "hello peter how are you doing"
```

This will translate the English text "hello peter how are you doing" into French and print the translated text to the terminal. In this case, the output will be:

Output:

```bash
bonjour Pierre comment vas tu
```


Contributing

If you find any bugs or would like to contribute to this project, please feel free to open an issue or a pull request.


