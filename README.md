# ConfigSwitcher

### To add a configuration
```bash
 cs config add ${nome-configurazione}
```
Copia la configurazione attuale `config.json` dentro la cartella `.configswitcher`, per poter essere usata dopo con il `nome-configurazione` selezionato.


### To use a configuration
```bash
cs config apply ${nome-configurazione}
```

Copia la configurazione scelta da quelle precedentemente aggiunte, sovrascrivendo il file `config.json`.


### To list all the configurations 

```bash
cs config list
```


### Rember to add `.configswitcher` to your `.gitignore`
